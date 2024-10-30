package app

import "github.com/urfave/cli"

var HTTPBackfillServerFlag = cli.StringFlag{
	Name:   "backfill-server-address",
	Usage:  "Run the rest for backfill server",
	EnvVar: "BACKFILL_SERVER_ADDRESS",
	Value:  "localhost:8081",
}

var HTTPPromoteeServerFlag = cli.StringFlag{
	Name:   "promotee-server-address",
	Usage:  "Run the rest for promotees server",
	EnvVar: "PROMOTEE_SERVER_ADDRESS",
	Value:  "localhost:8082",
}

func HTTPServerFlags() []cli.Flag {
	return []cli.Flag{
		HTTPBackfillServerFlag,
		HTTPPromoteeServerFlag,
	}
}
