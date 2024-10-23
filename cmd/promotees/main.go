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

	w := promotionworker.New(l, s, listener, parsers)

	apiServer := server.New(s, c.String(libapp.HTTPPromoteeServerFlag.Name))
	go func() {
		if err := apiServer.Run(); err != nil {
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
