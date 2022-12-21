package worker

import (
	"context"

	"github.com/KyberNetwork/tradelogs/internal/evmlistenerclient"
	"github.com/KyberNetwork/tradelogs/internal/parser"
	"github.com/KyberNetwork/tradelogs/internal/storage"
	"go.uber.org/zap"
)

type Worker struct {
	listener *evmlistenerclient.Client
	l        *zap.SugaredLogger
	s        *storage.Storage
	p        map[string]parser.Parser
}

func New(l *zap.SugaredLogger, s *storage.Storage, listener *evmlistenerclient.Client, parsers ...parser.Parser) (*Worker, error) {
	p := make(map[string]parser.Parser)
	for _, parser := range parsers {
		for _, topic := range parser.Topics() {
			p[topic] = parser
		}
	}
	return &Worker{
		listener: listener,
		l:        l,
		s:        s,
		p:        p,
	}, nil
}

func (w *Worker) Run(ctx context.Context) error {
	for {
		m, err := w.listener.GConsume(ctx)
		if err != nil {
			w.l.Errorw("Error while consume in group")
			return err
		}
		w.l.Infow("Begin process msg")
		if err := w.run(m); err != nil {
			return err
		}
		if err := w.listener.Ack(ctx, m); err != nil {
			return err
		}
	}
}
func (w *Worker) run(m []evmlistenerclient.Message) error {
	var (
		insertOrders []storage.TradeLogs
		deleteBlocks []uint64
	)

	for _, message := range m {
		for _, block := range message.NewBlocks {
			for _, log := range block.Logs {
				if len(log.Topics) == 0 {
					continue
				}
				parser := w.p[log.Topics[0].Hex()]
				if parser == nil {
					continue
				}
				order, err := parser.Parse(log, block.Number.Uint64(), block.Timestamp)
				if err != nil {
					return err
				}
				insertOrders = append(insertOrders, order)
			}
		}
		for _, block := range message.RevertedBlocks {
			deleteBlocks = append(deleteBlocks, block.Number.Uint64())
		}
	}
	err := w.s.Delete(deleteBlocks)
	if err != nil {
		return err
	}
	err = w.s.Insert(insertOrders)
	if err != nil {
		return err
	}
	return nil
}
