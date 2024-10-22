package promotionworker

import (
	"context"

	"github.com/KyberNetwork/tradelogs/pkg/convert"
	"github.com/KyberNetwork/tradelogs/pkg/evmlistenerclient"
	"github.com/KyberNetwork/tradelogs/pkg/promotionparser"
	"github.com/KyberNetwork/tradelogs/pkg/promotionstorage"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
)

type Proworker struct {
	listener *evmlistenerclient.Client
	l        *zap.SugaredLogger
	s        *promotionstorage.Storage
	p        []promotionparser.Parser
}

func New(l *zap.SugaredLogger, s *promotionstorage.Storage, listener *evmlistenerclient.Client,
	parsers []promotionparser.Parser) (*Proworker, error) {
	return &Proworker{
		listener: listener,
		l:        l,
		s:        s,
		p:        parsers,
	}, nil
}

func (w *Proworker) Run(ctx context.Context) error {
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
	}
}

func (w *Proworker) processMessages(m []evmlistenerclient.Message) error {
	for _, message := range m {
		var (
			insertPromotees []promotionstorage.Promotee
			deleteBlocks    []uint64
		)
		for _, block := range message.NewBlocks {
			for _, block := range message.RevertedBlocks {
				deleteBlocks = append(deleteBlocks, block.Number.Uint64())
			}
			for _, log := range block.Logs {
				ethLog := convert.ToETHLog(log)
				ps := w.findMatchingParser(ethLog)
				if ps == nil {
					continue
				}
				promotee, err := ps.Parse(ethLog, block.Timestamp)
				if err != nil {
					continue
				}
				insertPromotees = append(insertPromotees, promotee)
			}
		}

		if err := w.s.Delete(deleteBlocks); err != nil {
			return err
		}
		if err := w.s.Insert(insertPromotees); err != nil {
			return err
		}
	}
	return nil
}

func (w *Proworker) findMatchingParser(log ethTypes.Log) promotionparser.Parser {
	var ps promotionparser.Parser
	for _, p := range w.p {
		if !p.LogFromContract(log) {
			continue
		}
		ps = p
		break
	}
	return ps
}
