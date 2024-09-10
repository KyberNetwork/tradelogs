package types

import "strings"

type TradeLog interface {
	Serialize() []interface{}
	Type() string
}

type CommonTradeLog struct {
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
	//MakerTraits      string        `db:"maker_traits" json:"maker_traits,omitempty"`
	Expiry          uint64        `db:"expiration_date" json:"expiration_date"`
	MakerTokenPrice float64       `db:"maker_token_price" json:"maker_token_price"`
	TakerTokenPrice float64       `db:"taker_token_price" json:"taker_token_price"`
	MakerUsdAmount  float64       `db:"maker_usd_amount" json:"maker_usd_amount"`
	TakerUsdAmount  float64       `db:"taker_usd_amount" json:"taker_usd_amount"`
	State           TradeLogState `db:"state" json:"state"`
}

type TradeLogState string

const (
	TradeLogStateNew       TradeLogState = "new"
	TradeLogStateProcessed TradeLogState = "processed"
)

func (o *CommonTradeLog) Serialize() []interface{} {
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
		o.MakerTokenPrice,
		o.TakerTokenPrice,
		o.MakerUsdAmount,
		o.TakerUsdAmount,
		o.State,
	}
}

func CommonTradeLogColumns() []string {
	return []string{
		"order_hash",
		"maker",
		"taker",
		"maker_token",
		"taker_token",
		"maker_token_amount",
		"taker_token_amount",
		"contract_address",
		"block_number",
		"tx_hash",
		"log_index",
		"timestamp",
		"event_hash",
		"maker_token_price",
		"taker_token_price",
		"maker_usd_amount",
		"taker_usd_amount",
		"state",
	}
}
