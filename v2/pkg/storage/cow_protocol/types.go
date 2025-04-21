package cowprotocol

type CowTransfer struct {
	TransferId   uint64   `db:"transfer_id" json:"transfer_id"`
	TxHash       string   `db:"tx_hash" json:"tx_hash"`
	LogIndex     uint64   `db:"log_index" json:"log_index"`
	Timestamp    uint64   `db:"timestamp" json:"timestamp"`
	BlockNumber  uint64   `db:"block_number" json:"block_number"`
	FromAddress  string   `db:"from_address" json:"from_address"`
	ToAddress    string   `db:"to_address" json:"to_address"`
	Token        string   `db:"token" json:"token"`
	Amount       string   `db:"amount" json:"amount"`
	TransferType string   `db:"transfer_type" json:"transfer_type"`
	TokenPrice   *float64 `db:"token_price" json:"token_price,omitempty"`
	AmountUsd    *float64 `db:"amount_usd" json:"amount_usd,omitempty"`
}
type CowTransferQuery struct {
	FromTime uint64 `form:"from_time" json:"from_time,omitempty"`
	ToTime   uint64 `form:"to_time" json:"to_time,omitempty"`
	TxHash   string `form:"tx_hash" json:"tx_hash,omitempty"`
}

type CowTradeQuery struct {
	FromTime uint64 `form:"from_time" json:"from_time,omitempty" binding:"required"`
	ToTime   uint64 `form:"to_time" json:"to_time,omitempty" binding:"required"`
	TxHash   string `form:"tx_hash" json:"tx_hash,omitempty"`
	Limit    uint64 `form:"limit" json:"limit,omitempty"`
}

type CowTrade struct {
	ContractAddress  string   `db:"contract_address" json:"contract_address"`
	BlockNumber      uint64   `db:"block_number" json:"block_number"`
	TxHash           string   `db:"tx_hash" json:"tx_hash"`
	LogIndex         uint64   `db:"log_index" json:"log_index"`
	Timestamp        uint64   `db:"timestamp" json:"timestamp"`
	EventHash        string   `db:"event_hash" json:"event_hash"`
	MessageSender    string   `db:"message_sender" json:"message_sender"`
	TxOrigin         string   `db:"tx_origin" json:"tx_origin"`
	InteractContract string   `db:"interact_contract" json:"interact_contract"`
	Owner            string   `db:"owner" json:"owner"`
	SellToken        string   `db:"sell_token" json:"sell_token"`
	BuyToken         string   `db:"buy_token" json:"buy_token"`
	SellAmount       string   `db:"sell_amount" json:"sell_amount"`
	BuyAmount        string   `db:"buy_amount" json:"buy_amount"`
	FeeAmount        string   `db:"fee_amount" json:"fee_amount"`
	BuyTokenPrice    *float64 `db:"buy_token_price" json:"buy_token_price,omitempty"`
	SellTokenPrice   *float64 `db:"sell_token_price" json:"sell_token_price,omitempty"`
	BuyUsdAmount     *float64 `db:"buy_usd_amount" json:"buy_usd_amount,omitempty"`
	SellUsdAmount    *float64 `db:"sell_usd_amount" json:"sell_usd_amount,omitempty"`
	OrderUid         string   `db:"order_uid" json:"order_uid"`
}

type CowTradeCallFrame struct {
	TxHash    string `db:"tx_hash" json:"tx_hash"`
	CallFrame string `db:"call_frame" json:"call_frame"`
}
