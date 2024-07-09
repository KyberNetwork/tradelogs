package storage

import (
	"encoding/json"
	"math/big"
	"strings"
)

type TradeLog struct {
	OrderHash        string        `db:"order_hash" json:"order_hash,omitempty"`
	Maker            string        `db:"maker" json:"maker,omitempty"`
	Taker            string        `db:"taker" json:"taker,omitempty"`
	MakerToken       string        `db:"maker_token" json:"maker_token,omitempty"`
	TakerToken       string        `db:"taker_token" json:"taker_token,omitempty"`
	MakerTokenAmount string        `db:"maker_token_amount" json:"maker_token_amount,omitempty"`
	TakerTokenAmount string        `db:"taker_token_amount" json:"taker_token_amount,omitempty"`
	ContractAddress  string        `db:"contract_address" json:"contract_address,omitempty"`
	BlockNumber      uint64        `db:"block_number" json:"block_number,omitempty"`
	TxHash           string        `db:"tx_hash" json:"tx_hash,omitempty"`
	LogIndex         uint64        `db:"log_index" json:"log_index,omitempty"`
	Timestamp        uint64        `db:"timestamp" json:"timestamp,omitempty"`
	EventHash        string        `db:"event_hash" json:"event_hash,omitempty"`
	MakerTraits      string        `db:"maker_traits" json:"maker_traits,omitempty"`
	Expiry           uint64        `db:"expiration_date" json:"expiration_date"`
	MakerTokenPrice  float64       `db:"maker_token_price" json:"maker_token_price"`
	TakerTokenPrice  float64       `db:"taker_token_price" json:"taker_token_price"`
	MakerUsdAmount   float64       `db:"maker_usd_amount" json:"maker_usd_amount"`
	TakerUsdAmount   float64       `db:"taker_usd_amount" json:"taker_usd_amount"`
	State            TradeLogState `db:"state" json:"state"`
}

type TradeLogState string

const (
	TradeLogStateNew       TradeLogState = "new"
	TradeLogStateProcessed TradeLogState = "processed"
)

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

func (o *TradeLog) Serialize() []interface{} {
	// set default state is new
	if o.State == "" {
		o.State = TradeLogStateNew
	}
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
		o.MakerTokenPrice,
		o.TakerTokenPrice,
		o.MakerUsdAmount,
		o.TakerUsdAmount,
		o.State,
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
