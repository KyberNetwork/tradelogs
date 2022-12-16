package storage

type OtcOrderFilled struct {
	OrderHash              string `db:"order_hash" json:"order_hash,omitempty"`
	Maker                  string `db:"maker" json:"maker,omitempty"`
	Taker                  string `db:"taker" json:"taker,omitempty"`
	MakerToken             string `db:"maker_token" json:"maker_token,omitempty"`
	TakerToken             string `db:"taker_token" json:"taker_token,omitempty"`
	MakerTokenFilledAmount string `db:"maker_token_filled_amount" json:"maker_token_filled_amount,omitempty"`
	TakerTokenFilledAmount string `db:"taker_token_filled_amount" json:"taker_token_filled_amount,omitempty"`
	Address                string `db:"address" json:"address,omitempty"`
	BlockNumber            uint64 `db:"block_number" json:"block_number,omitempty"`
	TxHash                 string `db:"tx_hash" json:"tx_hash,omitempty"`
	Timestamp              uint64 `db:"timestamp" json:"timestamp,omitempty"`
}

func (o *OtcOrderFilled) Serialize() []interface{} {
	return []interface{}{
		o.OrderHash,
		o.Maker,
		o.Taker,
		o.MakerToken,
		o.TakerToken,
		o.MakerTokenFilledAmount,
		o.TakerTokenFilledAmount,
		o.Address,
		o.BlockNumber,
		o.TxHash,
		o.Timestamp,
	}
}
