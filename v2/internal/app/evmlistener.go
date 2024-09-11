package app

import (
	"github.com/urfave/cli"
)

const (
	EVMHTTPRPCUrlFlagName = "http-rpc-url"
	EVMwsRPCUrlFlagName   = "ws-rpc-url"
	EVMMaxTrackingBlock   = "max-tracking-block"
	EVMBlockExpiration    = "block-expiration"
)

//nolint:gochecknoglobals
var (
	EVMListenerHTTPRPCUrlFlag = &cli.StringFlag{
		Name:   EVMHTTPRPCUrlFlagName,
		EnvVar: "EVM_HTTP_RPC_URL",
		Usage:  "HTTP RPC node url for EVM Listener",
		Value:  "https://ethereum.kyberengineering.io/trading-tokyo",
	}
	EVMListenerWsRPCUrlFlag = &cli.StringFlag{
		Name:   EVMwsRPCUrlFlagName,
		EnvVar: "EVM_WS_RPC_URL",
		Usage:  "Web socket RPC node url for EVM Listener",
		Value:  "wss://ethereum.kyberengineering.io/trading-tokyo",
	}
	EVMListenerMaxTrackingBlock = &cli.IntFlag{
		Name:   EVMMaxTrackingBlock,
		EnvVar: "EVM_MAX_TRACKING_BLOCK",
		Usage:  "Max tracking block number for block keeper",
		Value:  32,
	}
)

// NewEvmListenerFlags returns flags for evmlistener.
func EvmListenerFlags() []cli.Flag {
	return []cli.Flag{
		EVMListenerHTTPRPCUrlFlag, EVMListenerWsRPCUrlFlag,
		EVMListenerMaxTrackingBlock,
	}
}
