package app

import "github.com/urfave/cli"

const (
	RPCUrlFlagName = "rpc-url"
)

var (
	RPCUrlFlag = &cli.StringSliceFlag{
		Name:   RPCUrlFlagName,
		EnvVar: "RPC_URL",
		Usage:  "RPC node url",
		Value:  &cli.StringSlice{"https://ethereum.kyberengineering.io/trading-tokyo"},
	}
)

func RPCNodeFlags() []cli.Flag {
	return []cli.Flag{
		RPCUrlFlag,
	}
}
