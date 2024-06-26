package storage

import (
	"encoding/json"
	"math/big"
	"strings"
)

type TradeLog struct {
	OrderHash        string `db:"order_hash" json:"order_hash,omitempty"`
	Maker            string `db:"maker" json:"maker,omitempty"`
	Taker            string `db:"taker" json:"taker,omitempty"`
	MakerToken       string `db:"maker_token" json:"maker_token,omitempty"`
	TakerToken       string `db:"taker_token" json:"taker_token,omitempty"`
	MakerTokenAmount string `db:"maker_token_amount" json:"maker_token_amount,omitempty"`
	TakerTokenAmount string `db:"taker_token_amount" json:"taker_token_amount,omitempty"`
	ContractAddress  string `db:"contract_address" json:"contract_address,omitempty"`
	BlockNumber      uint64 `db:"block_number" json:"block_number,omitempty"`
	TxHash           string `db:"tx_hash" json:"tx_hash,omitempty"`
	LogIndex         uint64 `db:"log_index" json:"log_index,omitempty"`
	Timestamp        uint64 `db:"timestamp" json:"timestamp,omitempty"`
	EventHash        string `db:"event_hash" json:"event_hash,omitempty"`
	MakerTraits      string `db:"maker_traits" json:"maker_traits,omitempty"`
	Expiry           uint64 `json:"expiration_date" db:"expiration_date"`
}

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
}

func (o *TradeLog) Serialize() []interface{} {
	return []interface{}{
		o.OrderHash,
		strings.ToLower(o.Maker),
		strings.ToLower(o.Taker),
		strings.ToLower(o.MakerToken),
		strings.ToLower(o.TakerToken),
		o.MakerTokenAmount,
		o.TakerTokenAmount,
		strings.ToLower(o.ContractAddress),
		o.BlockNumber,
		o.TxHash,
		o.LogIndex,
		o.Timestamp,
		o.EventHash,
		o.MakerTraits,
		o.Expiry,
	}
}

type BackfillQuery struct {
	QueryID   int64  `json:"query_id" binding:"required"`
	FromBlock uint64 `json:"from_block" binding:"required"`
	ToBlock   uint64 `json:"to_block" binding:"required"`
	EventHash string `json:"event_hash"`
	Exchange  string `json:"exchange"`
}

type BigInt struct {
	*big.Int
}

func (b *BigInt) UnmarshalJSON(data []byte) error {
	var num json.Number
	if err := json.Unmarshal(data, &num); err != nil {
		return err
	}

	b.Int = new(big.Int)
	b.Int, _ = b.Int.SetString(num.String(), 10)
	return nil
}

func (b *BigInt) MarshalJSON() ([]byte, error) {
	if b.Int == nil {
		return []byte("null"), nil
	}
	return []byte(b.String()), nil
}

func (b *BigInt) Hex() string {
	return "0x" + b.Text(16)
}
