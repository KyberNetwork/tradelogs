package paraswap

import "math/big"

type OrderRFQ struct {
	NonceAndMeta float64  `json:"nonceAndMeta"`
	Expiry       float64  `json:"expiry"`
	MakerAsset   string   `json:"makerAsset"`
	TakerAsset   string   `json:"takerAsset"`
	Maker        string   `json:"maker"`
	Taker        string   `json:"taker"`
	MakerAmount  *big.Int `json:"makerAmount"`
	TakerAmount  *big.Int `json:"takerAmount"`
}
