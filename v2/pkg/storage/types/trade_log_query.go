package types

type TradeLogsQuery struct {
	FromTime        uint64 `form:"from_time" json:"from_time,omitempty" binding:"required"`
	ToTime          uint64 `form:"to_time" json:"to_time,omitempty" binding:"required"`
	ContractAddress string `form:"contract_address" json:"contract_address,omitempty"`
	Maker           string `form:"maker" json:"maker,omitempty"`
	Taker           string `form:"taker" json:"taker,omitempty"`
	MakerToken      string `form:"maker_token" json:"maker_token,omitempty"`
	TakerToken      string `form:"taker_token" json:"taker_token,omitempty"`
	OrderHash       string `form:"order_hash" json:"order_hash,omitempty"`
	EventHash       string `form:"event_hash" json:"event_hash,omitempty"`
	State           string `form:"state" json:"state,omitempty"`
	Limit           uint64 `form:"limit" json:"limit,omitempty"`
}
