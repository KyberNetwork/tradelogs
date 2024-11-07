package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/KyberNetwork/tradelogs/v2/internal/worker"
	libapp "github.com/KyberNetwork/tradelogs/v2/pkg/app"
	"github.com/KyberNetwork/tradelogs/v2/pkg/evmlistenerclient"
	"github.com/KyberNetwork/tradelogs/v2/pkg/handler"
	"github.com/KyberNetwork/tradelogs/v2/pkg/kafka"
	"github.com/KyberNetwork/tradelogs/v2/pkg/parser"
	"github.com/KyberNetwork/tradelogs/v2/pkg/parser/zxotc"
	"github.com/KyberNetwork/tradelogs/v2/pkg/promotionparser"
	pro1inchv2 "github.com/KyberNetwork/tradelogs/v2/pkg/promotionparser/oneinchv2"
	"github.com/KyberNetwork/tradelogs/v2/pkg/rpcnode"
	promoteeTypes "github.com/KyberNetwork/tradelogs/v2/pkg/storage/promotees"
	"github.com/KyberNetwork/tradelogs/v2/pkg/storage/state"
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

	app.Flags = append(app.Flags, libapp.PostgresSQLFlags("tradelogs_v2")...)
	app.Flags = append(app.Flags, libapp.EvmListenerFlags()...)
	app.Flags = append(app.Flags, libapp.RPCNodeFlags()...)
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
	l.Infow("Starting log parser service")

	db, err := initDB(c)
	l.Infow("init db successfully")
	if err != nil {
		return fmt.Errorf("cannot init DB: %w", err)
	}

	// trade log manager
	manager := tradelogs.NewManager(l, []storageTypes.Storage{
		zxotcStorage.New(l, db),
	})

	//promotee storage
	promoteeStorage := promoteeTypes.New(l, db)

	// state storage
	s := state.New(l, db)

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

	promotionParsers := []promotionparser.Parser{pro1inchv2.MustNewParser()}

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
	tradeLogHandler := handler.NewTradeLogHandler(l, rpcNode, manager, promoteeStorage, parsers, promotionParsers, broadcastTopic, kafkaPublisher)

	// parse log worker
	w := worker.NewParseLog(tradeLogHandler, s, l)

	mostRecentBlock, err := getMostRecentBlock(l, s, rpcNode)
	if err != nil {
		return fmt.Errorf("cannot get most recent block: %w", err)
	}

	// subscribe evm listener with worker as a consumer
	return evmlistenerclient.SubscribeEvent(
		l,
		c.String(libapp.EVMListenerWsRPCUrlFlag.Name),
		c.String(libapp.EVMListenerHTTPRPCUrlFlag.Name),
		c.Int(libapp.EVMListenerMaxTrackingBlock.Name),
		w,
		mostRecentBlock,
		c.Int(libapp.EVMListenerClientTimeoutSecondFlag.Name),
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

func getMostRecentBlock(l *zap.SugaredLogger, s state.Storage, rpcClient rpcnode.IClient) (uint64, error) {
	var blockNumber uint64
	block, err := s.GetState(state.ProcessedBlockKey)
	if err == nil {
		blockNumber, err = strconv.ParseUint(block, 10, 64)
		if err == nil {
			return blockNumber, nil
		}
	}
	l.Errorw("cannot get from db", "err", err)
	blockNumber, err = rpcClient.GetBlockNumber(context.Background())
	if err != nil {
		return 0, fmt.Errorf("cannot get from node: %w", err)
	}
	return blockNumber, nil
}
