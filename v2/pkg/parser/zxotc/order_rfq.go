package zxotc

import (
	"math/big"
)

type OrderRFQ struct {
	MakerToken     string   `json:"makerToken"`
	TakerToken     string   `json:"takerToken"`
	MakerAmount    *big.Int `json:"makerAmount"`
	TakerAmount    *big.Int `json:"takerAmount"`
	Maker          string   `json:"maker"`
	Taker          string   `json:"taker"`
	TxOrigin       string   `json:"txOrigin"`
	ExpiryAndNonce *big.Int `json:"expiryAndNonce"`
}

func (o *OrderRFQ) GetExpiry() uint64 {
	expiry := o.ExpiryAndNonce.Rsh(o.ExpiryAndNonce, 192)
	return expiry.Uint64()
}
