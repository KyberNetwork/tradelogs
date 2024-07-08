package worker

import (
	"context"
	"strings"
	"time"

	"github.com/KyberNetwork/evmlistener/pkg/types"
	"github.com/KyberNetwork/tradelogs/pkg/convert"
	"github.com/KyberNetwork/tradelogs/pkg/evmlistenerclient"
	"github.com/KyberNetwork/tradelogs/pkg/parser"
	"github.com/KyberNetwork/tradelogs/pkg/pricefiller"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"go.uber.org/zap"
)

type Worker struct {
	listener     *evmlistenerclient.Client
	l            *zap.SugaredLogger
	s            *storage.Storage
	p            map[string]parser.Parser
	priceFiller  *pricefiller.PriceFiller
	tradeLogChan chan storage.TradeLog
}

func New(l *zap.SugaredLogger, s *storage.Storage, listener *evmlistenerclient.Client,
	priceFiller *pricefiller.PriceFiller, tradeLogChan chan storage.TradeLog,
	parsers ...parser.Parser) (*Worker, error) {
	p := make(map[string]parser.Parser)
	for _, ps := range parsers {
		for _, topic := range ps.Topics() {
			p[topic] = ps
		}
	}
	return &Worker{
		listener:     listener,
		l:            l,
		s:            s,
		p:            p,
		priceFiller:  priceFiller,
		tradeLogChan: tradeLogChan,
	}, nil
}

func (w *Worker) Run(ctx context.Context) error {
	retryTimer := time.NewTicker(evmlistenerclient.RetryTime)
	defer retryTimer.Stop()
	for {
		select {
		case <-retryTimer.C:
			if err := w.retryParseLog(); err != nil {
				w.l.Errorw("error when retry parse log", "err", err)
				return err
			}
		default:
		}
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
	}
}
func (w *Worker) processMessages(m []evmlistenerclient.Message) error {
	for _, message := range m {
		var (
			insertOrders []storage.TradeLog
			deleteBlocks []uint64
		)
		for _, block := range message.NewBlocks {
			for _, block := range message.RevertedBlocks {
				deleteBlocks = append(deleteBlocks, block.Number.Uint64())
				if err := w.s.DeleteErrorLogsWithBlock(block.Hash); err != nil {
					return err
				}
			}
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
					if err := w.s.InsertErrorLog(storage.EVMLog{
						Address:     log.Address,
						Topics:      strings.Join(log.Topics, ","),
						Data:        log.Data,
						BlockNumber: log.BlockNumber,
						TxHash:      log.TxHash,
						TxIndex:     log.TxIndex,
						BlockHash:   log.BlockHash,
						Index:       log.Index,
						Time:        block.Timestamp,
					}); err != nil {
						return err
					}
					continue
				}
				insertOrders = append(insertOrders, order)
			}
		}

		if err := w.s.Delete(deleteBlocks); err != nil {
			return err
		}
		if err := w.s.Insert(w.priceFiller.FullFillTradeLogs(insertOrders)); err != nil {
			return err
		}
		for _, log := range insertOrders {
			w.tradeLogChan <- log
		}
	}

	return nil
}

func (w *Worker) retryParseLog() error {
	insertOrders := []storage.TradeLog{}
	logs, err := w.s.GetErrorLogsSince(time.Now().Add(-time.Minute * 5).Unix())
	if err != nil {
		return err
	}
	w.l.Infow("start retry logs", "len", len(logs))
	for _, l := range logs {
		topics := strings.Split(l.Topics, ",")
		if len(topics) == 0 {
			continue
		}
		ps := w.p[topics[0]]
		if ps == nil {
			continue
		}
		order, err := ps.Parse(convert.ToETHLog(types.Log{
			Address:     l.Address,
			Topics:      topics,
			Data:        l.Data,
			BlockNumber: l.BlockNumber,
			TxHash:      l.TxHash,
			TxIndex:     l.TxIndex,
			BlockHash:   l.BlockHash,
			Index:       l.Index,
		}), l.Time)
		if err != nil {
			w.l.Errorw("error when retry log", "log", l, "err", err)
			continue
		}

		w.l.Infow("retry log successfully", "tx", order.TxHash, "index", order.LogIndex, "parser", ps.Exchange())
		if err := w.s.DeleteErrorLogsWithLogIndex(l.BlockNumber, l.Index); err != nil {
			return err
		}
		insertOrders = append(insertOrders, order)
	}

	if err := w.s.Insert(w.priceFiller.FullFillTradeLogs(insertOrders)); err != nil {
		return err
	}
	for _, log := range insertOrders {
		w.tradeLogChan <- log
	}
	return nil
}
