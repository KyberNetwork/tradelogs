package pricefiller

import (
	"math/big"

	"github.com/KyberNetwork/tradelogs/pkg/convert"
)

// calculateAmountUsd returns raw / (10**decimals) * price
func calculateAmountUsd(raw string, decimals int64, price float64) float64 {
	rawAmt, ok := new(big.Int).SetString(raw, 10)
	if !ok {
		return 0
	}

	return convert.WeiToFloat(rawAmt, decimals) * price
}
