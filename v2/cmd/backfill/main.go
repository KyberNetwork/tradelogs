package main

import (
	"fmt"
	"log"
	"os"

	"github.com/KyberNetwork/go-binance/v2"
	"github.com/KyberNetwork/tradelogs/v2/internal/worker"
	libapp "github.com/KyberNetwork/tradelogs/v2/pkg/app"
	"github.com/KyberNetwork/tradelogs/v2/pkg/handler"
	"github.com/KyberNetwork/tradelogs/v2/pkg/kafka"
	"github.com/KyberNetwork/tradelogs/v2/pkg/parser"
	"github.com/KyberNetwork/tradelogs/v2/pkg/parser/zxotc"
	"github.com/KyberNetwork/tradelogs/v2/pkg/pricefiller"
	"github.com/KyberNetwork/tradelogs/v2/pkg/rpcnode"
	"github.com/KyberNetwork/tradelogs/v2/pkg/storage/backfill"
	"github.com/KyberNetwork/tradelogs/v2/pkg/storage/tradelogs"
	storageTypes "github.com/KyberNetwork/tradelogs/v2/pkg/storage/tradelogs/types"
	zxotcStorage "github.com/KyberNetwork/tradelogs/v2/pkg/storage/tradelogs/zxotc"
	"github.com/KyberNetwork/tradinglib/pkg/dbutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jmoiron/sqlx"
	"github.com/urfave/cli"
	"go.uber.org/zap"
)

func main() {
	app := libapp.NewApp()
	app.Name = "trade log parser service"
	app.Action = run

	app.Flags = append(app.Flags, libapp.PriceFillerFlags()...)
	app.Flags = append(app.Flags, libapp.RPCNodeFlags()...)
	app.Flags = append(app.Flags, libapp.PostgresSQLFlags("tradelogs_v2")...)
	app.Flags = append(app.Flags, libapp.KafkaFlag()...)

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
	l.Infow("Starting backfill service")

	db, err := initDB(c)
	if err != nil {
		return fmt.Errorf("cannot init DB: %w", err)
	}

	// trade log manager
	storages := []storageTypes.Storage{
		zxotcStorage.New(l, db),
	}
	manager := tradelogs.NewManager(l, storages)

	// price filler
	binanceClient := binance.NewClient(c.String(libapp.BinanceAPIKeyFlag.Name), c.String(libapp.BinanceSecretKeyFlag.Name))
	_, err = pricefiller.NewPriceFiller(l, binanceClient, storages)
	if err != nil {
		return fmt.Errorf("cannot init price filler: %w", err)
	}

	// backfill storage
	s := backfill.New(l, db)

	// rpc node to query trace call
	rpcURL := c.StringSlice(libapp.RPCUrlFlagName)
	if len(rpcURL) == 0 {
		return fmt.Errorf("rpc url is empty")
	}

	ethClients := make([]*ethclient.Client, len(rpcURL))
	for i, url := range rpcURL {
		client, err := ethclient.Dial(url)
		if err != nil {
			return fmt.Errorf("cannot dial eth client: %w", err)
		}
		ethClients[i] = client
	}
	rpcNode := rpcnode.NewClient(l, ethClients...)

	parsers := []parser.Parser{
		//kyberswap.MustNewParser(),
		zxotc.MustNewParser(),
		//paraswap.MustNewParser(),
		//kyberswaprfq.MustNewParser(),
		//hashflowv3.MustNewParser(),
		//oneinchv6.MustNewParser(traceCalls),
		//uniswapxv1.MustNewParser(traceCalls),
		//uniswapx.MustNewParser(traceCalls),
		//bebop.MustNewParser(traceCalls),
		//zxrfqv3.MustNewParserWithDeployer(traceCalls, ethClient, common.HexToAddress(parser.Deployer0xV3)),
	}

	// kafka broadcast topic
	broadcastTopic := c.String(libapp.KafkaBroadcastTopic.Name)
	err = kafka.ValidateTopicName(broadcastTopic)
	if err != nil {
		return fmt.Errorf("invalid kafka topic: %w", err)
	}

	// kafka publisher for broadcasting trade logs
	kafkaPublisher, err := kafka.NewPublisher(libapp.KafkaConfigFromFlags(c))
	if err != nil {
		return fmt.Errorf("cannot create kafka publisher: %w", err)
	}

	// trade log handler
	tradeLogHandler := handler.NewTradeLogHandler(l, rpcNode, manager, parsers, broadcastTopic, kafkaPublisher)

	// parse log worker
	w := worker.NewBackFiller(tradeLogHandler, s, l, rpcNode)

	return w.Run()
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
