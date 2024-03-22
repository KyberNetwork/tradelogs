package app

import (
	"github.com/urfave/cli"
)

var HTTPServerFlag = cli.StringFlag{
	Name:   "server-address",
	Usage:  "Run the rest server",
	EnvVar: "SERVER_ADDRESS",
	Value:  "localhost:8080",
}

var HTTPBackfillServerFlag = cli.StringFlag{
	Name:   "backfill-server-address",
	Usage:  "Run the rest for backfill server",
	EnvVar: "BACKFILL_SERVER_ADDRESS",
	Value:  "localhost:8081",
}

var DuneURLFlag = cli.StringFlag{
	Name:   "dune-url",
	EnvVar: "DUNE_URL",
	Value:  "https://api.dune.com/api",
}

var DuneKeyFlag = cli.StringFlag{
	Name:   "dune-key",
	EnvVar: "DUNE_KEY",
}

func HTTPServerFlags() []cli.Flag {
	return []cli.Flag{
		HTTPServerFlag,
		HTTPBackfillServerFlag,
		DuneURLFlag,
		DuneKeyFlag,
	}
}
