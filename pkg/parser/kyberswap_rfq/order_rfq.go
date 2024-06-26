package kyberswaprfq

import (
	"math/big"
)

type OrderRFQ struct {
	Info          float64 `json:"info"`
	MakerAsset    string  `json:"makerAsset"`
	TakerAsset    string  `json:"takerAsset"`
	Maker         string  `json:"maker"`
	AllowedSender string  `json:"allowedSender"`
	MakingAmount  float64 `json:"makingAmount"`
	TakingAmount  float64 `json:"takingAmount"`
}

func (o *OrderRFQ) GetExpiry() uint64 {
	infoFloat := big.NewFloat(o.Info)
	info := new(big.Int)
	info, _ = infoFloat.Int(info)
	expiry := info.Rsh(info, 64)

	return expiry.Uint64()

}
