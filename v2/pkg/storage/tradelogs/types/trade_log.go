package types

import (
	"encoding/json"
	"strings"
)

type TradeLog struct {
	Exchange               string          `json:"exchange"`
	OrderHash              string          `db:"order_hash" json:"order_hash,omitempty"`
	Maker                  string          `db:"maker" json:"maker,omitempty"`
	Taker                  string          `db:"taker" json:"taker,omitempty"`
	MakerToken             string          `db:"maker_token" json:"maker_token,omitempty"`
	TakerToken             string          `db:"taker_token" json:"taker_token,omitempty"`
	MakerTokenAmount       string          `db:"maker_token_amount" json:"maker_token_amount,omitempty"`
	TakerTokenAmount       string          `db:"taker_token_amount" json:"taker_token_amount,omitempty"`
	MakerTokenOriginAmount string          `db:"maker_token_origin_amount" json:"maker_token_origin_amount,omitempty"`
	TakerTokenOriginAmount string          `db:"taker_token_origin_amount" json:"taker_token_origin_amount,omitempty"`
	ContractAddress        string          `db:"contract_address" json:"contract_address,omitempty"`
	BlockNumber            uint64          `db:"block_number" json:"block_number,omitempty"`
	TxHash                 string          `db:"tx_hash" json:"tx_hash,omitempty"`
	LogIndex               uint64          `db:"log_index" json:"log_index,omitempty"`
	TradeIndex             uint64          `db:"trade_index" json:"trade_index,omitempty"`
	Timestamp              uint64          `db:"timestamp" json:"timestamp,omitempty"`
	EventHash              string          `db:"event_hash" json:"event_hash,omitempty"`
	MessageSender          string          `db:"message_sender" json:"message_sender,omitempty"`
	TxOrigin               string          `db:"tx_origin" json:"tx_origin,omitempty"`
	InteractContract       string          `db:"interact_contract" json:"interact_contract,omitempty"`
	MakerTraits            json.RawMessage `db:"maker_traits" json:"maker_traits,omitempty"`
	Expiry                 uint64          `db:"expiration_date" json:"expiration_date"`
	MakerTokenPrice        *float64        `db:"maker_token_price" json:"maker_token_price"`
	TakerTokenPrice        *float64        `db:"taker_token_price" json:"taker_token_price"`
	MakerUsdAmount         *float64        `db:"maker_usd_amount" json:"maker_usd_amount"`
	TakerUsdAmount         *float64        `db:"taker_usd_amount" json:"taker_usd_amount"`
	Type                   OneInchType     `db:"type" json:"type"`
}

// CommonTradeLogSerialize used for exchanges only storing fields in common trade logs,
// if these exchanges need to store extra fields, they have to use themselves serialize function
func CommonTradeLogSerialize(o *TradeLog) []interface{} {
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
		o.TxOrigin,
		o.MessageSender,
		o.InteractContract,
		o.MakerTokenPrice,
		o.TakerTokenPrice,
		o.MakerUsdAmount,
		o.TakerUsdAmount,
	}
}

// CommonTradeLogColumns used for query in db and corresponding to CommonTradeLogSerialize function above when insert new row in db
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
		"tx_origin",
		"message_sender",
		"interact_contract",
		"maker_token_price",
		"taker_token_price",
		"maker_usd_amount",
		"taker_usd_amount",
	}
}

// RFQTradeLogSerialize used for exchanges supporting RFQ trades and partial fill
func RFQTradeLogSerialize(o *TradeLog) []interface{} {
	return []interface{}{
		o.OrderHash,
		strings.ToLower(o.Maker),
		strings.ToLower(o.Taker),
		strings.ToLower(o.MakerToken),
		strings.ToLower(o.TakerToken),
		o.MakerTokenAmount,
		o.TakerTokenAmount,
		o.MakerTokenOriginAmount,
		o.TakerTokenOriginAmount,
		strings.ToLower(o.ContractAddress),
		o.BlockNumber,
		o.TxHash,
		o.LogIndex,
		o.Timestamp,
		o.EventHash,
		o.TxOrigin,
		o.MessageSender,
		o.InteractContract,
		o.MakerTokenPrice,
		o.TakerTokenPrice,
		o.MakerUsdAmount,
		o.TakerUsdAmount,
	}
}

// RFQTradeLogColumns used for query in db and corresponding to RFQTradeLogSerialize function above when insert new row in db
func RFQTradeLogColumns() []string {
	return []string{
		"order_hash",
		"maker",
		"taker",
		"maker_token",
		"taker_token",
		"maker_token_amount",
		"taker_token_amount",
		"maker_token_origin_amount",
		"taker_token_origin_amount",
		"contract_address",
		"block_number",
		"tx_hash",
		"log_index",
		"timestamp",
		"event_hash",
		"tx_origin",
		"message_sender",
		"interact_contract",
		"maker_token_price",
		"taker_token_price",
		"maker_usd_amount",
		"taker_usd_amount",
	}
}
