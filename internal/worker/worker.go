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
	"github.com/KyberNetwork/tradelogs/pkg/rpcnode"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"github.com/KyberNetwork/tradinglib/pkg/metrics"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
)

const (
	RetryInterval          = 4 * time.Second
	RemoveInterval         = time.Hour
	ParsingErrorMetricName = "trading_tradelogs_parse_error"
	BlockCountMetricName   = "trading_tradelogs_block_count"
)

type Worker struct {
	listener      *evmlistenerclient.Client
	l             *zap.SugaredLogger
	s             *storage.Storage
	p             []parser.Parser
	priceFiller   *pricefiller.PriceFiller
	tradeLogChan  chan storage.TradeLog
	rpcNodeClient *rpcnode.Client
}

func New(l *zap.SugaredLogger, s *storage.Storage, listener *evmlistenerclient.Client,
	priceFiller *pricefiller.PriceFiller, tradeLogChan chan storage.TradeLog,
	parsers []parser.Parser, rpcNodeClient *rpcnode.Client) (*Worker, error) {
	return &Worker{
		listener:      listener,
		l:             l,
		s:             s,
		p:             parsers,
		priceFiller:   priceFiller,
		tradeLogChan:  tradeLogChan,
		rpcNodeClient: rpcNodeClient,
	}, nil
}

func (w *Worker) Run(ctx context.Context) error {
	retryTimer := time.NewTicker(RetryInterval)
	defer retryTimer.Stop()
	removeTimer := time.NewTicker(RemoveInterval)
	defer removeTimer.Stop()
	for {
		select {
		case <-retryTimer.C:
			if err := w.retryParseLog(); err != nil {
				w.l.Errorw("error when retry parse log", "err", err)
			}
		case <-removeTimer.C:
			if err := w.removeOldErrorLog(); err != nil {
				w.l.Errorw("error when remove old error log", "err", err)
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
		w.l.Infow("about to process new message", "msg", message)
		var (
			insertOrders []storage.TradeLog
			deleteBlocks []uint64
		)
		// delete revert blocks
		for _, block := range message.RevertedBlocks {
			deleteBlocks = append(deleteBlocks, block.Number.Uint64())
			if err := w.s.DeleteErrorLogsWithBlock(block.Hash); err != nil {
				w.l.Infow("error when revert block", "blocks", deleteBlocks, "err", err)
				return err
			}
		}

		// record block count metric
		if err := metrics.RecordCounter(context.Background(), BlockCountMetricName, int64(len(message.NewBlocks))); err != nil {
			w.l.Errorw("error when record block count metric", "err", err)
		}

		// handle new block
		for _, block := range message.NewBlocks {
			for _, log := range block.Logs {
				ethLog := convert.ToETHLog(log)
				ps := w.findMatchingParser(ethLog)
				if ps == nil {
					continue
				}
				order, err := ps.Parse(ethLog, block.Timestamp)
				if err != nil {
					w.l.Errorw("error when parse log", "log", log, "order", order, "err", err)
					// record log parsing error metric
					if err = metrics.RecordCounter(context.Background(), ParsingErrorMetricName, 1); err != nil {
						w.l.Errorw("error when record parsing error metric", "err", err)
					}
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
				txOrigin, err := w.rpcNodeClient.GetTxOriginByTxHash(context.Background(), order.TxHash)
				if err != nil {
					w.l.Errorw("error when get tx origin", "txHash", order.TxHash, "err", err)
				}
				if txOrigin != (common.Address{}) {
					order.TxOrigin = strings.ToLower(txOrigin.Hex())
				}
				insertOrders = append(insertOrders, order)
			}
		}

		if err := w.s.Delete(deleteBlocks); err != nil {
			return err
		}
		w.priceFiller.FullFillTradeLogs(insertOrders)
		if err := w.s.Insert(insertOrders); err != nil {
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
		ethLog := convert.ToETHLog(types.Log{
			Address:     l.Address,
			Topics:      topics,
			Data:        l.Data,
			BlockNumber: l.BlockNumber,
			TxHash:      l.TxHash,
			TxIndex:     l.TxIndex,
			BlockHash:   l.BlockHash,
			Index:       l.Index,
		})
		ps := w.findMatchingParser(ethLog)
		if ps == nil {
			continue
		}
		order, err := ps.Parse(ethLog, l.Time)
		if err != nil {
			w.l.Errorw("error when retry log", "log", l, "err", err)
			if err := metrics.RecordCounter(context.Background(), ParsingErrorMetricName, 1); err != nil {
				w.l.Errorw("error when record parsing error", "err", err)
			}
			continue
		}

		w.l.Infow("retry log successfully", "tx", order.TxHash, "index", order.LogIndex, "parser", ps.Exchange())
		if err := w.s.DeleteErrorLogsWithLogIndex(l.BlockNumber, l.Index); err != nil {
			return err
		}
		insertOrders = append(insertOrders, order)
	}

	w.priceFiller.FullFillTradeLogs(insertOrders)
	if err := w.s.Insert(insertOrders); err != nil {
		return err
	}
	for _, log := range insertOrders {
		w.tradeLogChan <- log
	}
	return nil
}

func (w *Worker) findMatchingParser(log ethTypes.Log) parser.Parser {
	var ps parser.Parser
	for _, p := range w.p {
		if !p.LogFromExchange(log) {
			continue
		}
		ps = p
		break
	}
	return ps
}

func (w *Worker) removeOldErrorLog() error {
	w.l.Info("start to remove old error log")
	return w.s.RemoveLogUtil(time.Now().Add(-time.Hour * 24).Unix())
}
