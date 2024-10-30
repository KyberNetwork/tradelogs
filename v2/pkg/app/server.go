package app

import "github.com/urfave/cli"

var (
	HTTPBackfillServerFlag = cli.StringFlag{
		Name:   "backfill-server-address",
		Usage:  "Run the rest for backfill server",
		EnvVar: "BACKFILL_SERVER_ADDRESS",
		Value:  "localhost:8081",
	}
	HTTPBroadcastServerFlag = cli.StringFlag{
		Name:   "broadcast-server-address",
		Usage:  "Run the rest for broadcast server",
		EnvVar: "BROADCAST_SERVER_ADDRESS",
		Value:  "localhost:8082",
	}
	HTTPTradeLogsServerFlag = cli.StringFlag{
		Name:   "tradelogs-server-address",
		Usage:  "Run the rest for tradelogs server",
		EnvVar: "TRADELOGS_SERVER_ADDRESS",
		Value:  "localhost:8080",
	}
	HTTPPromoteeServerFlag = cli.StringFlag{
		Name:   "promotee-server-address",
		Usage:  "Run the rest for promotees server",
		EnvVar: "PROMOTEE_SERVER_ADDRESS",
		Value:  "localhost:8082",
	}
)

func HTTPServerFlags() []cli.Flag {
	return []cli.Flag{
		HTTPBackfillServerFlag,
		HTTPBroadcastServerFlag,
		HTTPTradeLogsServerFlag,
		HTTPPromoteeServerFlag,
	}
}
