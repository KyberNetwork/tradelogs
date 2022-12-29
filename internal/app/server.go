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

func HTTPServerFlags() []cli.Flag {
	return []cli.Flag{
		HTTPServerFlag,
	}
}
