package app

import (
	"github.com/KyberNetwork/tradelogs/pkg/evmlistenerclient"
	"github.com/urfave/cli"
)

//nolint:gochecknoglobals
var (
	publisherTopicFlag = &cli.StringFlag{
		Name:     "publisher-topic",
		EnvVar:   "PUBLISHER_TOPIC",
		Required: true,
		Usage:    "Topic name of publisher to get message to (Required)",
	}

	groupNameFlag = &cli.StringFlag{
		Name:   "group-name",
		EnvVar: "XGROUP_NAME",
		Value:  "trading-tradelogs",
		Usage:  "Name of group handler",
	}

	groupConsumerFlag = &cli.StringFlag{
		Name:   "group-consumer",
		EnvVar: "XGROUP_CONSUMER",
		Value:  "trading-consumer",
		Usage:  "Name of group consumer",
	}
)

// NewEvmListenerFlags returns flags for evmlistener.
func EvmListenerFlags() []cli.Flag {
	return []cli.Flag{
		publisherTopicFlag, groupNameFlag, groupConsumerFlag,
	}
}

func EvmlistenerConfigFromCli(c *cli.Context) evmlistenerclient.Config {
	cfg := evmlistenerclient.Config{
		Topic:         c.String(publisherTopicFlag.Name),
		GroupName:     c.String(groupNameFlag.Name),
		GroupConsumer: c.String(groupConsumerFlag.Name),
	}
	return cfg
}
