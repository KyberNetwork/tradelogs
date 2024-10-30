package types

type Storage interface {
	Insert(orders []TradeLog) error
	Get(query TradeLogsQuery) ([]TradeLog, error)
	GetEmptyPrice(limit uint64) ([]TradeLog, error)
	Delete(blocks []uint64) error
	Exchange() string
}
