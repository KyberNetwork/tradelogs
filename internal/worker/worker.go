package worker

import (
	"context"

	"github.com/KyberNetwork/tradelogs/pkg/convert"
	"github.com/KyberNetwork/tradelogs/pkg/evmlistenerclient"
	"github.com/KyberNetwork/tradelogs/pkg/parser"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"github.com/gammazero/workerpool"
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
	for _, ps := range parsers {
		for _, topic := range ps.Topics() {
			p[topic] = ps
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
		w.l.Infow("process msg", "count", len(m))
		if len(m) == 0 {
			continue
		}
		if err := w.processMessages(m); err != nil {
			return err
		}
		if err := w.listener.Ack(ctx, m); err != nil {
			return err
		}
	}
}
func (w *Worker) processMessages(m []evmlistenerclient.Message) error {
	var (
		insertOrders         []storage.TradeLog
		deleteBlocks         []uint64
		wp                   = workerpool.New(100)
		insertOrdersResultCh = make(chan parseEventLogResult)
	)

	for _, message := range m {
		for _, block := range message.NewBlocks {
			for _, log := range block.Logs {
				if len(log.Topics) == 0 {
					continue
				}
				ps := w.p[log.Topics[0]]
				if ps == nil {
					continue
				}
				// Create new instance of req for the goroutine. Can remove these variables after go1.22
				block := block
				log := log
				wp.Submit(func() {
					order, err := ps.Parse(convert.ToETHLog(log), block.Timestamp)
					if err != nil {
						insertOrdersResultCh <- parseEventLogResult{
							err: err,
						}
						return
					}
					insertOrdersResultCh <- parseEventLogResult{
						tradeLogOrder: order,
						err:           nil,
					}
				})
			}
		}
		for _, block := range message.RevertedBlocks {
			deleteBlocks = append(deleteBlocks, block.Number.Uint64())
		}
	}

	go func() {
		wp.StopWait()
		close(insertOrdersResultCh)
	}()

	for result := range insertOrdersResultCh {
		if result.err != nil {
			return result.err
		}
		insertOrders = append(insertOrders, result.tradeLogOrder)
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
