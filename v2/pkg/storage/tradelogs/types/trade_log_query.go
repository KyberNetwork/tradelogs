package types

type TradeLogsQuery struct {
	FromTime         uint64 `form:"from_time" json:"from_time,omitempty" binding:"required"`
	ToTime           uint64 `form:"to_time" json:"to_time,omitempty" binding:"required"`
	BlockNumber      string `form:"block_number" json:"block_number,omitempty"`
	ContractAddress  string `form:"contract_address" json:"contract_address,omitempty"`
	TxHash           string `form:"tx_hash" json:"tx_hash,omitempty"`
	Maker            string `form:"maker" json:"maker,omitempty"`
	Taker            string `form:"taker" json:"taker,omitempty"`
	MakerToken       string `form:"maker_token" json:"maker_token,omitempty"`
	TakerToken       string `form:"taker_token" json:"taker_token,omitempty"`
	OrderHash        string `form:"order_hash" json:"order_hash,omitempty"`
	EventHash        string `form:"event_hash" json:"event_hash,omitempty"`
	MessageSender    string `form:"message_sender" json:"message_sender,omitempty"`
	InteractContract string `form:"interact_contract" json:"interact_contract,omitempty"`
	Limit            uint64 `form:"limit" json:"limit,omitempty"`
}
