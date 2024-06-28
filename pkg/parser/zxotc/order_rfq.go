package zxotc

import (
	"math/big"
)

type OrderRFQ struct {
	MakerToken     string  `json:"makerToken"`
	TakerToken     string  `json:"takerToken"`
	MakerAmount    float64 `json:"makerAmount"`
	TakerAmount    float64 `json:"takerAmount"`
	Maker          string  `json:"maker"`
	Taker          string  `json:"taker"`
	TxOrigin       string  `json:"txOrigin"`
	ExpiryAndNonce float64 `json:"expiryAndNonce"`
}

func (o *OrderRFQ) GetExpiry() uint64 {
	expiryAndNonceFloat := big.NewFloat(o.ExpiryAndNonce)
	expiryAndNonce := new(big.Int)
	expiryAndNonce, _ = expiryAndNonceFloat.Int(expiryAndNonce)
	expiry := expiryAndNonce.Rsh(expiryAndNonce, 192)
	return expiry.Uint64()
}
