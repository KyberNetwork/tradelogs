package zxotc

import (
	"github.com/KyberNetwork/tradelogs/internal/evmlistenerclient"
	"github.com/KyberNetwork/tradelogs/internal/storage"
	"go.uber.org/zap"
)

type Worker struct {
	l *zap.SugaredLogger
	s *storage.Storage
	p *Parser
}

func NewWorker(l *zap.SugaredLogger, s *storage.Storage) (*Worker, error) {
	p, err := NewParser()
	if err != nil {
		return nil, err
	}
	return &Worker{
		l: l,
		s: s,
		p: p,
	}, nil
}

func (w *Worker) Run(m []evmlistenerclient.Message) error {
	var (
		insertOrders []storage.OtcOrderFilled
		deleteOrders []string
	)
	for _, message := range m {
		for _, block := range message.NewBlocks {
			for _, log := range block.Logs {
				if len(log.Topics) > 0 && log.Topics[0].Hex() == OtcOrderFilledLog {
					order, err := w.p.Parse(log, block.Number.Uint64(), block.Timestamp)
					if err != nil {
						return err
					}
					insertOrders = append(insertOrders, order)
				}
			}
		}
		for _, block := range message.RevertedBlocks {
			for _, log := range block.Logs {
				if len(log.Topics) > 0 && log.Topics[0].Hex() == OtcOrderFilledLog {
					order, err := w.p.Parse(log, block.Number.Uint64(), block.Timestamp)
					if err != nil {
						return err
					}
					deleteOrders = append(deleteOrders, order.OrderHash)
				}
			}
		}
	}
	err := w.s.Delete(deleteOrders)
	if err != nil {
		return err
	}
	err = w.s.Insert(insertOrders)
	if err != nil {
		return err
	}
	return nil
}
