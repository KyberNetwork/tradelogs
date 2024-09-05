package app

import "github.com/urfave/cli"

const (
	RPCUrlFlagName         = "rpc-url"
	FallbackRPCUrlFlagName = "fallback-rpc-url"
)

var (
	RPCUrlFlag = &cli.StringFlag{
		Name:   RPCUrlFlagName,
		EnvVar: "RPC_URL",
		Usage:  "RPC node url",
		Value:  "https://ethereum.kyberengineering.io/trading-tokyo",
	}
	FallbackRPCUrlFlag = &cli.StringFlag{
		Name:   FallbackRPCUrlFlagName,
		EnvVar: "FALLBACK_RPC_URL",
		Usage:  "Fallback RPC node url",
		Value:  "https://ethereum.kyberengineering.io/trading-tokyo",
	}
)

func RPCNodeFlags() []cli.Flag {
	return []cli.Flag{
		RPCUrlFlag,
		FallbackRPCUrlFlag,
	}
}
