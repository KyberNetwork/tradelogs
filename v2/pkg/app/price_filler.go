package app

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

var MarkToMarketURLFlag = cli.StringFlag{
	Name:   "mark-to-market-url",
	EnvVar: "MARK_TO_MARKET_URL",
}

var FillPriceTimeIntervalFlag = cli.StringFlag{
	Name:   "fill-price-time-interval",
	EnvVar: "FILL_PRICE_TIME_INTERVAL",
}

func PriceFillerFlags() []cli.Flag {
	return []cli.Flag{
		BinanceAPIKeyFlag,
		BinanceSecretKeyFlag,
		MarkToMarketURLFlag,
		FillPriceTimeIntervalFlag,
	}
}
