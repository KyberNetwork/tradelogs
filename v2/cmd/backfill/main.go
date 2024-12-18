package main

import (
	"fmt"
	"log"
	"os"

	"github.com/KyberNetwork/tradelogs/v2/internal/server"
	"github.com/KyberNetwork/tradelogs/v2/internal/service"
	"github.com/KyberNetwork/tradelogs/v2/internal/worker"
	libapp "github.com/KyberNetwork/tradelogs/v2/pkg/app"
	"github.com/KyberNetwork/tradelogs/v2/pkg/constant"
	"github.com/KyberNetwork/tradelogs/v2/pkg/handler"
	"github.com/KyberNetwork/tradelogs/v2/pkg/kafka"
	"github.com/KyberNetwork/tradelogs/v2/pkg/parser"
	"github.com/KyberNetwork/tradelogs/v2/pkg/parser/bebop"
	hashflowv3 "github.com/KyberNetwork/tradelogs/v2/pkg/parser/hashflow_v3"
	"github.com/KyberNetwork/tradelogs/v2/pkg/parser/kyberswap"
	kyberswaprfq "github.com/KyberNetwork/tradelogs/v2/pkg/parser/kyberswap_rfq"
	"github.com/KyberNetwork/tradelogs/v2/pkg/parser/oneinchv6"
	"github.com/KyberNetwork/tradelogs/v2/pkg/parser/pancakeswap"
	"github.com/KyberNetwork/tradelogs/v2/pkg/parser/paraswap"
	"github.com/KyberNetwork/tradelogs/v2/pkg/parser/uniswapx"
	"github.com/KyberNetwork/tradelogs/v2/pkg/parser/zxotc"
	"github.com/KyberNetwork/tradelogs/v2/pkg/parser/zxrfqv3"
	"github.com/KyberNetwork/tradelogs/v2/pkg/promotionparser"
	promotion1inchv2 "github.com/KyberNetwork/tradelogs/v2/pkg/promotionparser/oneinchv2"
	"github.com/KyberNetwork/tradelogs/v2/pkg/rpcnode"
	"github.com/KyberNetwork/tradelogs/v2/pkg/storage/backfill"
	promotee_storage "github.com/KyberNetwork/tradelogs/v2/pkg/storage/promotees"
	"github.com/KyberNetwork/tradelogs/v2/pkg/storage/state"
	"github.com/KyberNetwork/tradelogs/v2/pkg/storage/tradelogs"
	bebopStorage "github.com/KyberNetwork/tradelogs/v2/pkg/storage/tradelogs/bebop"
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
	"github.com/KyberNetwork/tradelogs/v2/pkg/storage/zerox_deployment"
	"github.com/KyberNetwork/tradinglib/pkg/dbutil"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jmoiron/sqlx"
	"github.com/urfave/cli"
	"go.uber.org/zap"
)

func main() {
	app := libapp.NewApp()
	app.Name = "trade log backfill service"
	app.Action = run

	app.Flags = append(app.Flags, libapp.RPCNodeFlags()...)
	app.Flags = append(app.Flags, libapp.PostgresSQLFlags("tradelogs_v2")...)
	app.Flags = append(app.Flags, libapp.KafkaFlag()...)
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
	l.Infow("Starting backfill service")

	db, err := initDB(c)
	l.Infow("init db successfully")
	if err != nil {
		return fmt.Errorf("cannot init DB: %w", err)
	}

	// trade log manager
	storages := []storageTypes.Storage{
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
	}
	manager := tradelogs.NewManager(l, storages)

	//promotee storage
	promoteeStorage := promotee_storage.New(l, db)

	// backfill storage
	backfillStorage := backfill.New(l, db)

	// state storage
	stateStorage := state.New(l, db)

	// zerox deployment storage
	zxv3DeployStorage := zerox_deployment.NewStorage(l, db)

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
		kyberswap.MustNewParser(),
		zxotc.MustNewParser(),
		paraswap.MustNewParser(),
		kyberswaprfq.MustNewParser(),
		hashflowv3.MustNewParser(),
		oneinchv6.MustNewParser(promoteeStorage),
		uniswapx.MustNewParser(),
		bebop.MustNewParser(),
		zxrfqv3.MustNewParserWithDeployer(l, zxv3DeployStorage, ethClients[0], common.HexToAddress(constant.Deployer0xV3)),
		pancakeswap.MustNewParser(),
	}

	promotionParsers := []promotionparser.Parser{promotion1inchv2.MustNewParser()}

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
	w := worker.NewBackFiller(tradeLogHandler, backfillStorage, stateStorage, l, rpcNode, parsers)

	w.Run()

	srv, err := service.NewBackfillService(backfillStorage, l, w)
	if err != nil {
		return fmt.Errorf("cannot create backfill service: %w", err)
	}

	s := server.NewBackfill(l, c.String(libapp.HTTPBackfillServerFlag.Name), srv)

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
