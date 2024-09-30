package app

import (
	"github.com/urfave/cli"
)

const (
	EVMHTTPRPCUrlFlagName  = "http-rpc-url"
	EVMwsRPCUrlFlagName    = "ws-rpc-url"
	EVMMaxTrackingBlock    = "max-tracking-block"
	EVMClientTimeoutSecond = "timeout-second"
)

//nolint:gochecknoglobals
var (
	EVMListenerHTTPRPCUrlFlag = &cli.StringFlag{
		Name:   EVMHTTPRPCUrlFlagName,
		EnvVar: "EVM_HTTP_RPC_URL",
		Usage:  "HTTP RPC node url for EVM Listener",
	}
	EVMListenerWsRPCUrlFlag = &cli.StringFlag{
		Name:   EVMwsRPCUrlFlagName,
		EnvVar: "EVM_WS_RPC_URL",
		Usage:  "Web socket RPC node url for EVM Listener",
	}
	EVMListenerMaxTrackingBlock = &cli.IntFlag{
		Name:   EVMMaxTrackingBlock,
		EnvVar: "EVM_MAX_TRACKING_BLOCK",
		Usage:  "Max tracking block number for block keeper",
		Value:  32,
	}
	EVMListenerClientTimeoutSecondFlag = &cli.IntFlag{
		Name:   EVMClientTimeoutSecond,
		EnvVar: "EVM_CLIENT_TIMEOUT_SECOND",
		Usage:  "Client timeout second",
		Value:  15,
	}
)

// NewEvmListenerFlags returns flags for evmlistener.
func EvmListenerFlags() []cli.Flag {
	return []cli.Flag{
		EVMListenerHTTPRPCUrlFlag, EVMListenerWsRPCUrlFlag,
		EVMListenerMaxTrackingBlock, EVMListenerClientTimeoutSecondFlag,
	}
}
