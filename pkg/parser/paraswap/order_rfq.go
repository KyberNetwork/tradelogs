package paraswap

type OrderRFQ struct {
	NonceAndMeta float64 `json:"nonceAndMeta"`
	Expiry       float64 `json:"expiry"`
	MakerAsset   string  `json:"makerAsset"`
	TakerAsset   string  `json:"takerAsset"`
	Maker        string  `json:"maker"`
	Taker        string  `json:"taker"`
	MakerAmount  float64 `json:"makerAmount"`
	TakerAmount  float64 `json:"takerAmount"`
}
