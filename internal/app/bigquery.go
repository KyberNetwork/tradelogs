package app

import (
	"github.com/urfave/cli"
)

//nolint:gochecknoglobals
var (
	bigqueryProjectIDFlag = &cli.StringFlag{
		Name:     "bigquery-project-id",
		EnvVar:   "BIGQUERY_PROJECT_ID",
		Required: true,
		Usage:    "Bigquery project id for setup bigquery client",
	}
)

// NewEvmListenerFlags returns flags for evmlistener.
func BigqueryFlags() []cli.Flag {
	return []cli.Flag{
		bigqueryProjectIDFlag,
	}
}

func BigqueryProjectIDFFromCli(c *cli.Context) string {
	return c.String(bigqueryProjectIDFlag.Name)
}
