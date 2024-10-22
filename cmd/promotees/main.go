package main

import (
	"context"
	"fmt"
	"log"
	"os"

	libapp "github.com/KyberNetwork/tradelogs/internal/app"
	"github.com/KyberNetwork/tradelogs/internal/dbutil"
	"github.com/KyberNetwork/tradelogs/internal/promotionworker"
	"github.com/KyberNetwork/tradelogs/pkg/evmlistenerclient"

	server "github.com/KyberNetwork/tradelogs/internal/server/promotees"
	"github.com/KyberNetwork/tradelogs/pkg/promotionparser"
	"github.com/KyberNetwork/tradelogs/pkg/promotionparser/oneinchv2"
	"github.com/KyberNetwork/tradelogs/pkg/promotionstorage"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/urfave/cli"
	"go.uber.org/zap"
)

func main() {
	app := libapp.NewApp()
	app.Name = "Promotees fetcher service"
	app.Action = run
	app.Flags = append(app.Flags, libapp.PostgresSQLFlags("tradelogs")...)
	app.Flags = append(app.Flags, libapp.RedisFlags()...)
	app.Flags = append(app.Flags, libapp.EvmListenerFlags()...)
	app.Flags = append(app.Flags, libapp.HTTPServerFlags()...)

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
	s := promotionstorage.New(l, db)
	listener := evmlistenerclient.New(l, libapp.EvmlistenerConfigFromCli(c), redis)
	err = listener.Init(context.Background())
	if err != nil {
		l.Errorw("Error while init listener service")
		return err
	}

	parsers := []promotionparser.Parser{oneinchv2.MustNewParser()}

	w, err := promotionworker.New(l, s, listener, parsers)
	if err != nil {
		l.Errorw("Error while init worker")
		return err
	}

	apiServer := server.New(s, c.String(libapp.HTTPPromoteeServerFlag.Name))
	go func() {
		if err := apiServer.Run(); err != nil {
			panic(err)
		}
	}()

	// err = listener.SendLog(context.Background(), "{\"revertedBlocks\":null,\"newBlocks\":[{\"number\":20389423,\"hash\":\"0xfbce2d84f1822fabc605e5c7492f201e18d4b9f309c7224093f76af17971de58\",\"timestamp\":1721981219,\"parentHash\":\"0x9e8b6b70a88f6813968733450aed3298782017ac96562c49f7fd2a2123a995cc\",\"reorgedHash\":\"\",\"logs\":[{\"address\":\"0xF55684BC536487394B423e70567413faB8e45E26\",\"topics\":[\"0xb863cf86b291171e4b0332ea12b59af030f98a2c74a6d51effaf1109ae4c7f1e\"],\"data\":\"0x000000000000000000000000a8be6b2afe6e060985675675615c2108a66135c80000000000000000000000000000000000000000000000000000000000000001000000000000000000000000b6613cc55866e282638006455390207c1d485be9\",\"blockNumber\":\"0x1371E2F\",\"transactionHash\":\"0x8a695a7b9264fef4c70b78df9b787989559f7a63916511faee79a667fb602163\",\"transactionIndex\":\"0x152\",\"blockHash\":\"0xfbce2d84f1822fabc605e5c7492f201e18d4b9f309c7224093f76af17971de58\",\"logIndex\":\"0x152\",\"removed\":false}]}]}")
	// if err != nil {
	// 	l.Errorw("Error sending log", "error", err)
	// }

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
