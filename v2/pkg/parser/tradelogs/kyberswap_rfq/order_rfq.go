package kyberswaprfq

import (
	"math/big"
)

type OrderRFQ struct {
	Info          *big.Int `json:"info"`
	MakerAsset    string   `json:"makerAsset"`
	TakerAsset    string   `json:"takerAsset"`
	Maker         string   `json:"maker"`
	AllowedSender string   `json:"allowedSender"`
	MakingAmount  *big.Int `json:"makingAmount"`
	TakingAmount  *big.Int `json:"takingAmount"`
}

func (o *OrderRFQ) GetExpiry() uint64 {
	expiry := o.Info.Rsh(o.Info, 64)

	return expiry.Uint64()

}
