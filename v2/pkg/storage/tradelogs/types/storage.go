package types

import "time"

type Storage interface {
	Insert(orders []TradeLog) error
	Get(query TradeLogsQuery) ([]TradeLog, error)
	GetEmptyPrice(limit uint64) ([]TradeLog, error)
	Delete(blocks []uint64) error
	Exchange() string
	ResetTokenPriceToRefetch(token string, from, to time.Time) (int64, error)
}
