package worker

import (
	"context"
	"fmt"
	"time"

	etype "github.com/KyberNetwork/evmlistener/pkg/types"
	"github.com/KyberNetwork/tradelogs/pkg/convert"
	"github.com/KyberNetwork/tradelogs/pkg/evmlistenerclient"
	"github.com/KyberNetwork/tradelogs/pkg/parser"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"github.com/ethereum/go-ethereum/common/lru"
	"go.uber.org/zap"
)

type EVMLog struct {
	log etype.Log
	ts  uint64
}

type Worker struct {
	listener *evmlistenerclient.Client
	l        *zap.SugaredLogger
	s        *storage.Storage
	p        map[string]parser.Parser
	errLogs  lru.BasicLRU[string, EVMLog]
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
		errLogs:  lru.NewBasicLRU[string, EVMLog](1000),
	}, nil
}

func (w *Worker) Run(ctx context.Context) error {
	retryTimer := time.NewTicker(evmlistenerclient.BlockTime)
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
			w.l.Errorw("Error when proccess msg", "error", err)
			return err
		}
		if err := w.listener.Ack(ctx, m); err != nil {
			w.l.Errorw("Error when ack msg", "error", err)
			return err
		}
		select {
		case <-retryTimer.C:
			if err := w.retryParseLog(); err != nil {
				w.l.Errorw("error when retry parse log", "err", err)
				return err
			}
		default:
		}
	}
}
func (w *Worker) processMessages(m []evmlistenerclient.Message) error {
	var (
		insertOrders []storage.TradeLog
		deleteBlocks []uint64
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
				order, err := ps.Parse(convert.ToETHLog(log), block.Timestamp)
				if err != nil {
					w.l.Errorw("error when parse log", "log", log, "order", order, "err", err)
					w.errLogs.Add(fmt.Sprintf("%d-%d", log.BlockNumber, log.Index), EVMLog{
						log: log,
						ts:  block.Timestamp,
					})
					continue
				}
				insertOrders = append(insertOrders, order)
			}
		}
		for _, block := range message.RevertedBlocks {
			deleteBlocks = append(deleteBlocks, block.Number.Uint64())
			for _, k := range w.errLogs.Keys() {
				l, ok := w.errLogs.Peek(k)
				if !ok {
					continue
				}
				if l.log.BlockNumber == block.Number.Uint64() {
					w.errLogs.Remove(k)
				}
			}
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

func (w *Worker) retryParseLog() error {
	insertOrders := []storage.TradeLog{}
	for _, k := range w.errLogs.Keys() {
		l, ok := w.errLogs.Peek(k)
		if !ok {
			continue
		}
		ps := w.p[l.log.Topics[0]]
		if ps == nil {
			continue
		}
		order, err := ps.Parse(convert.ToETHLog(l.log), l.ts)
		if err != nil {
			continue
		}
		w.l.Infow("retry log successfully", "key", k, "parser", ps.Exchange())
		w.errLogs.Remove(k)
		insertOrders = append(insertOrders, order)
	}

	if err := w.s.Insert(insertOrders); err != nil {
		return err
	}
	return nil
}
