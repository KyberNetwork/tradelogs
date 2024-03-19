package server

import (
	"github.com/KyberNetwork/tradelogs/pkg/parser/oneinch"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
)

var (
	oneinchAbi, _       = oneinch.OneinchMetaData.GetAbi()
	oneinchRfqEvent, _  = oneinchAbi.Events[oneinch.FilledEvent]
	oneinchRfqEventHash = oneinchRfqEvent.String()
)

type rfqOrder struct {
	OrderHash        string `json:"order_hash" binding:"required"`
	BlockNumber      uint64 `json:"block_number" binding:"required"`
	LogIndex         uint64 `json:"log_index" binding:"required"`
	EventHash        string `json:"event_hash" binding:"required"`
	TxHash           string `json:"tx_hash" binding:"required"`
	Maker            string `json:"maker"`
	Taker            string `json:"taker"`
	MakerToken       string `json:"maker_token"`
	TakerToken       string `json:"taker_token"`
	MakerTokenAmount string `json:"maker_token_amount"`
	TakerTokenAmount string `json:"taker_token_amount"`
	ContractAddress  string `json:"contract_address"`
	Timestamp        uint64 `json:"timestamp"`
}
type BackFillOneInchRequest []rfqOrder

func (rfqOrders BackFillOneInchRequest) ToTradeLogs() []storage.TradeLog {
	tradeLogs := make([]storage.TradeLog, 0, len(rfqOrders))
	for _, order := range rfqOrders {
		if order.EventHash != oneinchRfqEventHash {
			continue
		}
		if order.TxHash == "" {
			continue
		}
		tradeLogs = append(tradeLogs, storage.TradeLog{
			OrderHash:        order.OrderHash,
			BlockNumber:      order.BlockNumber,
			LogIndex:         order.LogIndex,
			EventHash:        order.EventHash,
			TxHash:           order.TxHash,
			Maker:            order.Maker,
			Taker:            order.Taker,
			MakerToken:       order.MakerToken,
			TakerToken:       order.TakerToken,
			MakerTokenAmount: order.MakerTokenAmount,
			TakerTokenAmount: order.TakerTokenAmount,
			ContractAddress:  order.ContractAddress,
			Timestamp:        order.Timestamp,
		})
	}
	return tradeLogs
}
