package main

import (
	"fmt"
	libapp "github.com/KyberNetwork/tradelogs/v2/internal/app"
	"github.com/KyberNetwork/tradelogs/v2/internal/dbutil"
	"github.com/KyberNetwork/tradelogs/v2/internal/worker"
	"github.com/KyberNetwork/tradelogs/v2/pkg/evmlistenerclient"
	"github.com/KyberNetwork/tradelogs/v2/pkg/handler"
	"github.com/KyberNetwork/tradelogs/v2/pkg/kafka"
	"github.com/KyberNetwork/tradelogs/v2/pkg/parser"
	"github.com/KyberNetwork/tradelogs/v2/pkg/parser/zxotc"
	"github.com/KyberNetwork/tradelogs/v2/pkg/rpcnode"
	"github.com/KyberNetwork/tradelogs/v2/pkg/storage"
	storageTypes "github.com/KyberNetwork/tradelogs/v2/pkg/storage/types"
	zxotcStorage "github.com/KyberNetwork/tradelogs/v2/pkg/storage/zxotc"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jmoiron/sqlx"
	"github.com/urfave/cli"
	"go.uber.org/zap"
	"log"
	"os"
)

func main() {
	app := libapp.NewApp()
	app.Name = "trade log parser service"
	app.Action = run

	app.Flags = append(app.Flags, libapp.PostgresSQLFlags("tradelogs_v2")...)
	app.Flags = append(app.Flags, libapp.RedisFlags()...)
	app.Flags = append(app.Flags, libapp.EvmListenerFlags()...)
	app.Flags = append(app.Flags, libapp.HTTPServerFlags()...)
	app.Flags = append(app.Flags, libapp.RPCNodeFlags()...)
	app.Flags = append(app.Flags, libapp.KafkaFlag()...)
	//app.Flags = append(app.Flags, pricefiller.PriceFillerFlags()...)

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
	l.Infow("Starting log parser service")

	db, err := initDB(c)
	if err != nil {
		l.Panicw("cannot init DB", "err", err)
	}

	s := storage.NewManager(l, []storageTypes.Storage{
		zxotcStorage.New(l, db),
	})

	// rpc node to query trace call
	rpcURL := c.StringSlice(libapp.RPCUrlFlagName)
	if len(rpcURL) == 0 {
		return fmt.Errorf("rpc url is empty")
	}

	ethClients := make([]*ethclient.Client, len(rpcURL))
	for i, url := range rpcURL {
		client, err := ethclient.Dial(url)
		if err != nil {
			panic(err)
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
		panic(fmt.Errorf("invalid kafka topic: %w", err))
	}

	// kafka publisher for broadcasting trade logs
	kafkaPublisher, err := kafka.NewPublisher(libapp.KafkaConfigFromFlags(c))
	if err != nil {
		panic(fmt.Errorf("cannot create kafka publisher: %w", err))
	}

	// trade log handler
	tradeLogHandler := handler.NewTradeLogHandler(l, rpcNode, s, parsers, broadcastTopic, kafkaPublisher)

	// parse log worker
	w := worker.NewParseLog(tradeLogHandler)

	// redis config for evm listener's block keeper
	redisCfg := libapp.RedisConfigFromFlags(c)

	// subscribe evm listener with worker as a consumer
	return evmlistenerclient.SubscribeEvent(
		l,
		c.String(libapp.EVMListenerWsRPCUrlFlag.Name),
		c.String(libapp.EVMListenerHTTPRPCUrlFlag.Name),
		c.Int(libapp.EVMListenerMaxTrackingBlock.Name),
		c.Duration(libapp.EVMListenerBlockExpiration.Name),
		redisCfg,
		w,
	)
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
