package main

import (
	"fmt"
	"github.com/KyberNetwork/tradelogs/v2/internal/server"
	"log"
	"os"

	libapp "github.com/KyberNetwork/tradelogs/v2/pkg/app"
	storageTypes "github.com/KyberNetwork/tradelogs/v2/pkg/storage/tradelogs/types"
	zxotcStorage "github.com/KyberNetwork/tradelogs/v2/pkg/storage/tradelogs/zxotc"
	"github.com/KyberNetwork/tradinglib/pkg/dbutil"
	"github.com/jmoiron/sqlx"
	"github.com/urfave/cli"
	"go.uber.org/zap"
)

func main() {
	app := libapp.NewApp()
	app.Name = "trade logs crawler service"
	app.Action = run
	app.Flags = append(app.Flags, libapp.PostgresSQLFlags("tradelogs_v2")...)
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
	l.Infow("Starting trade logs server")

	db, err := initDB(c)
	l.Infow("init db successfully")
	if err != nil {
		return fmt.Errorf("cannot init DB: %w", err)
	}

	// trade log storages
	storage := []storageTypes.Storage{
		zxotcStorage.New(l, db),
	}

	s := server.NewTradeLogs(l, storage, c.String(libapp.HTTPTradeLogsServerFlag.Name))
	return s.Run()
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