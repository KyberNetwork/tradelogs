package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	libapp "github.com/KyberNetwork/tradelogs/v2/pkg/app"
	"github.com/KyberNetwork/tradelogs/v2/pkg/mtm"
	pricefiller "github.com/KyberNetwork/tradelogs/v2/pkg/price_filler"
	dashboardStorage "github.com/KyberNetwork/tradelogs/v2/pkg/storage/dashboard"
	bebopStorage "github.com/KyberNetwork/tradelogs/v2/pkg/storage/tradelogs/bebop"
	cowProtocolStorage "github.com/KyberNetwork/tradelogs/v2/pkg/storage/tradelogs/cow_protocol"
	hashflowv3Storage "github.com/KyberNetwork/tradelogs/v2/pkg/storage/tradelogs/hashflow_v3"
	kyberswapStorage "github.com/KyberNetwork/tradelogs/v2/pkg/storage/tradelogs/kyberswap"
	kyberswaprfqStorage "github.com/KyberNetwork/tradelogs/v2/pkg/storage/tradelogs/kyberswap_rfq"
	oneinchv6Storage "github.com/KyberNetwork/tradelogs/v2/pkg/storage/tradelogs/oneinch_v6"
	pancakeswapStorage "github.com/KyberNetwork/tradelogs/v2/pkg/storage/tradelogs/pancakeswap"
	paraswapStorage "github.com/KyberNetwork/tradelogs/v2/pkg/storage/tradelogs/paraswap"
	storageTypes "github.com/KyberNetwork/tradelogs/v2/pkg/storage/tradelogs/types"
	uniswapxStorage "github.com/KyberNetwork/tradelogs/v2/pkg/storage/tradelogs/uniswapx"
	zxotcStorage "github.com/KyberNetwork/tradelogs/v2/pkg/storage/tradelogs/zxotc"
	zxrfqv3Storage "github.com/KyberNetwork/tradelogs/v2/pkg/storage/tradelogs/zxrfqv3"
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
	app.Flags = append(app.Flags, libapp.PriceFillerFlags()...)

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
	l.Infow("Starting price filler service")

	db, err := initDB(c)
	l.Infow("init db successfully")
	if err != nil {
		return fmt.Errorf("cannot init DB: %w", err)
	}

	// trade log storages
	s := []storageTypes.Storage{
		kyberswapStorage.New(l, db),
		zxotcStorage.New(l, db),
		paraswapStorage.New(l, db),
		kyberswaprfqStorage.New(l, db),
		hashflowv3Storage.New(l, db),
		oneinchv6Storage.New(l, db),
		uniswapxStorage.New(l, db),
		bebopStorage.New(l, db),
		zxrfqv3Storage.New(l, db),
		pancakeswapStorage.New(l, db),
		cowProtocolStorage.New(l, db),
	}
	httpClient := &http.Client{}
	mtmClient := mtm.NewMtmClient(c.String(libapp.MarkToMarketURLFlag.Name), httpClient)
	dashboardStorage := dashboardStorage.New(l, db)
	priceFiller, err := pricefiller.NewPriceFiller(l, s, mtmClient, dashboardStorage)
	if err != nil {
		l.Errorw("Error while init price filler")
		return err
	}
	fillPriceInterval := c.Duration(libapp.FillPriceTimeIntervalFlag.Name)
	if fillPriceInterval <= 0 {
		l.Errorw("invalid value of fill-price-time-interval-flag")
		return fmt.Errorf("invalid value of fill-price-time-interval-flag")
	}
	priceFiller.Run(fillPriceInterval)
	return nil
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
