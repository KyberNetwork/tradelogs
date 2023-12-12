package app

import "github.com/urfave/cli"

const (
	RPCUrlFlagName = "rpc-url"
)

var (
	RPCUrlFlag = &cli.StringFlag{
		Name:   RPCUrlFlagName,
		EnvVar: "RPC_URL",
		Usage:  "RPC node url",
	}
)

func RPCNodeFlags() []cli.Flag {
	return []cli.Flag{
		RPCUrlFlag,
	}
}
