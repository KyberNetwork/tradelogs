package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/KyberNetwork/tradelogs/pkg/dune"
	"github.com/KyberNetwork/tradelogs/pkg/parser"
	"github.com/KyberNetwork/tradelogs/pkg/parser/oneinch"
	"github.com/KyberNetwork/tradelogs/pkg/parser/oneinchv6"
	"github.com/KyberNetwork/tradelogs/pkg/rpcnode"
	"github.com/KyberNetwork/tradelogs/pkg/tracecall"

	libapp "github.com/KyberNetwork/tradelogs/internal/app"
	"github.com/KyberNetwork/tradelogs/internal/bigquery"
	"github.com/KyberNetwork/tradelogs/internal/dbutil"
	backfill "github.com/KyberNetwork/tradelogs/internal/server/backfill"
	tradelogs "github.com/KyberNetwork/tradelogs/internal/server/tradelogs"
	"github.com/KyberNetwork/tradelogs/internal/worker"
	"github.com/KyberNetwork/tradelogs/pkg/evmlistenerclient"
	"github.com/KyberNetwork/tradelogs/pkg/parser/hashflow"
	hashflowv3 "github.com/KyberNetwork/tradelogs/pkg/parser/hashflow_v3"
	"github.com/KyberNetwork/tradelogs/pkg/parser/kyberswap"
	kyberswaprfq "github.com/KyberNetwork/tradelogs/pkg/parser/kyberswap_rfq"
	"github.com/KyberNetwork/tradelogs/pkg/parser/native"
	"github.com/KyberNetwork/tradelogs/pkg/parser/paraswap"
	"github.com/KyberNetwork/tradelogs/pkg/parser/tokenlon"
	"github.com/KyberNetwork/tradelogs/pkg/parser/zxotc"
	"github.com/KyberNetwork/tradelogs/pkg/parser/zxrfq"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/urfave/cli"
	"go.uber.org/zap"
)

func main() {
	app := libapp.NewApp()
	app.Name = "trade logs crawler service"
	app.Action = run
	app.Flags = append(app.Flags, libapp.PostgresSQLFlags("tradelogs")...)
	app.Flags = append(app.Flags, libapp.RedisFlags()...)
	app.Flags = append(app.Flags, libapp.EvmListenerFlags()...)
	app.Flags = append(app.Flags, libapp.HTTPServerFlags()...)
	app.Flags = append(app.Flags, libapp.BigqueryFlags()...)
	app.Flags = append(app.Flags, libapp.RPCNodeFlags()...)

	if err := app.Run(os.Args); err != nil {
		log.Panic(err)
	}
}

func run(c *cli.Context) error {
	logger, _, flush, err := libapp.NewLogger(c)
	if err != nil {
		return fmt.Errorf("new logger: %w", err)
	}

	defer flush()

	zap.ReplaceGlobals(logger)
	l := logger.Sugar()
	l.Infow("app starting ..")

	db, err := initDB(c)
	if err != nil {
		l.Panicw("cannot init DB", "err", err)
	}
	redis, err := initRedis(c)
	if err != nil {
		l.Panicw("cannot init redis", "err", err)
	}
	s := storage.New(l, db)
	listener := evmlistenerclient.New(l, libapp.EvmlistenerConfigFromCli(c), redis)
	err = listener.Init(context.Background())
	if err != nil {
		l.Errorw("Error while init listener service")
		return err
	}
	httpClient := &http.Client{
		Timeout: time.Second * 30,
		Transport: &http.Transport{
			IdleConnTimeout:       time.Second * 120,
			ResponseHeaderTimeout: time.Second * 30,
		},
	}
	rpcClient, err := rpcnode.NewClient(httpClient, c.String(libapp.RPCUrlFlagName))
	if err != nil {
		panic(err)
	}
	traceCalls := tracecall.NewCache(rpcClient)

	parsers := []parser.Parser{kyberswap.MustNewParser(),
		zxotc.MustNewParser(),
		zxrfq.MustNewParser(),
		tokenlon.MustNewParser(),
		paraswap.MustNewParser(),
		hashflow.MustNewParser(),
		native.MustNewParser(),
		kyberswaprfq.MustNewParser(),
		hashflowv3.MustNewParser(),
		oneinch.MustNewParser(traceCalls),
		oneinchv6.MustNewParser(traceCalls),
	}

	tradeLogChan := make(chan storage.TradeLog, 1000)
	w, err := worker.New(l, s, listener, tradeLogChan, parsers...)
	if err != nil {
		l.Errorw("Error while init worker")
		return err
	}
	bigQueryWorker, err := bigquery.NewWorker(libapp.BigqueryProjectIDFFromCli(c), s, parsers)
	if err != nil {
		l.Errorw("error when init big query worker", "error", err)
	}

	httpBackfill := backfill.New(c.String(libapp.HTTPBackfillServerFlag.Name), bigQueryWorker,
		backfill.NewDuneWoker(dune.NewClient(c.String(libapp.DuneURLFlag.Name), c.String(libapp.DuneKeyFlag.Name), httpClient), s),
		parsers)
	go func() {
		if err := httpBackfill.Run(); err != nil {
			panic(err)
		}
	}()

	bc := tradelogs.NewBroadcaster(tradeLogChan)
	go bc.BroadcastLog()
	httpTradelogs := tradelogs.New(l, s, c.String(libapp.HTTPServerFlag.Name), bc)
	go func() {
		if err := httpTradelogs.Run(); err != nil {
			panic(err)
		}
	}()

	return w.Run(context.Background())
}

func initDB(c *cli.Context) (*sqlx.DB, error) {
	db, err := libapp.NewDB(map[string]interface{}{
		"host":     c.String(libapp.PostgresHost.Name),
		"port":     c.Int(libapp.PostgresPort.Name),
		"user":     c.String(libapp.PostgresUser.Name),
		"password": c.String(libapp.PostgresPassword.Name),
		"dbname":   c.String(libapp.PostgresDatabase.Name),
		"sslmode":  "disable",
	})
	if err != nil {
		return nil, err
	}

	_, err = dbutil.RunMigrationUp(db.DB, c.String(libapp.PostgresMigrationPath.Name),
		c.String(libapp.PostgresDatabase.Name))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func initRedis(c *cli.Context) (redis.UniversalClient, error) {
	redis := redis.NewUniversalClient(libapp.RedisOptionFromFlags(c))
	if _, err := redis.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}
	return redis, nil
}
