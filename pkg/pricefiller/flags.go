package pricefiller

import (
	"github.com/urfave/cli"
)

var BinanceAPIKeyFlag = cli.StringFlag{
	Name:   "binance-api-key",
	EnvVar: "BINANCE_API_KEY",
}

var BinanceSecretKeyFlag = cli.StringFlag{
	Name:   "binance-secret-key",
	EnvVar: "BINANCE_SECRET_KEY",
}

func PriceFillerFlags() []cli.Flag {
	return []cli.Flag{
		BinanceAPIKeyFlag,
		BinanceSecretKeyFlag,
	}
}
