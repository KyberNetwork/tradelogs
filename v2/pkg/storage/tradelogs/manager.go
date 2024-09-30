package tradelogs

import (
	"fmt"

	"github.com/KyberNetwork/tradelogs/v2/pkg/storage/tradelogs/types"
	"go.uber.org/zap"
)

type Manager struct {
	l        *zap.SugaredLogger
	storages []types.Storage
}

func NewManager(l *zap.SugaredLogger, storages []types.Storage) *Manager {
	return &Manager{
		l:        l,
		storages: storages,
	}
}

func (m *Manager) Insert(orders []types.TradeLog) error {
	if len(orders) == 0 {
		return nil
	}
	tradeLogsByExchange := make(map[string][]types.TradeLog)
	for _, order := range orders {
		exchange := order.Exchange
		tradeLogsByExchange[exchange] = append(tradeLogsByExchange[exchange], order)
	}
	for _, storage := range m.storages {
		err := storage.Insert(tradeLogsByExchange[storage.Exchange()])
		if err != nil {
			m.l.Errorw("insert trade logs failed", "exchange", storage.Exchange())
			return fmt.Errorf("insert trade logs for exchange %v failed: %w", storage.Exchange(), err)
		}
		m.l.Infow("insert trade logs", "exchange", storage.Exchange(), "number", len(tradeLogsByExchange[storage.Exchange()]))
	}
	return nil
}

func (m *Manager) Get(query types.TradeLogsQuery) ([]types.TradeLog, error) {
	var result []types.TradeLog
	for _, storage := range m.storages {
		tradeLogs, err := storage.Get(query)
		if err != nil {
			return nil, fmt.Errorf("get trade logs failed: %w", err)
		}
		result = append(result, tradeLogs...)
	}
	return result, nil
}

func (m *Manager) Delete(blocks []uint64) error {
	if len(blocks) == 0 {
		return nil
	}
	for _, storage := range m.storages {
		err := storage.Delete(blocks)
		if err != nil {
			return fmt.Errorf("delete trade logs failed: %w", err)
		}
	}
	m.l.Infow("deleted trade logs", "blocks", blocks)
	return nil
}
