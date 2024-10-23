package main

import (
	"fmt"
	"log"
	"os"

	"github.com/KyberNetwork/tradelogs/v2/internal/server"
	"github.com/KyberNetwork/tradelogs/v2/internal/worker"
	libapp "github.com/KyberNetwork/tradelogs/v2/pkg/app"
	"github.com/KyberNetwork/tradelogs/v2/pkg/kafka"
	"github.com/urfave/cli"
	"go.uber.org/zap"
)

func main() {
	app := libapp.NewApp()
	app.Name = "trade log broadcast service"
	app.Action = run

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
	l.Infow("Starting broadcast service")

	// kafka broadcast topic
	broadcastTopic := c.String(libapp.KafkaBroadcastTopic.Name)
	err = kafka.ValidateTopicName(broadcastTopic)
	if err != nil {
		return fmt.Errorf("invalid kafka topic: %w", err)
	}

	cfg := libapp.KafkaConfigFromFlags(c)
	broadcaster := worker.NewBroadcaster(l, cfg, broadcastTopic)

	s := server.New(l, c.String(libapp.HTTPBroadcastServerFlag.Name), broadcaster)
	return s.Run()
}
