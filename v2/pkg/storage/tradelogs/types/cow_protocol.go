package types

type CowTransfer struct {
	TxHash      string   `db:"tx_hash" json:"tx_hash,omitempty"`
	Timestamp   uint64   `db:"timestamp" json:"timestamp,omitempty"`
	BlockNumber uint64   `db:"block_number" json:"block_number,omitempty"`
	FromAddress string   `db:"from_address" json:"from_address,omitempty"`
	ToAddress   string   `db:"to_address" json:"to_address,omitempty"`
	Token       string   `db:"token" json:"token,omitempty"`
	Amount      string   `db:"amount" json:"amount,omitempty"`
	TokenPrice  *float64 `db:"token_price" json:"token_price,omitempty"`
	AmountUsd   *float64 `db:"amount_usd" json:"amount_usd,omitempty"`
}
type CowTransferQuery struct {
	FromTime uint64 `form:"from_time" json:"from_time,omitempty"`
	ToTime   uint64 `form:"to_time" json:"to_time,omitempty"`
	TxHash   string `form:"tx_hash" json:"tx_hash,omitempty"`
}
