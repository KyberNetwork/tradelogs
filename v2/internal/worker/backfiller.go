package worker

import (
	"context"
	"fmt"
	"time"

	"github.com/KyberNetwork/tradelogs/v2/pkg/handler"
	"github.com/KyberNetwork/tradelogs/v2/pkg/rpcnode"
	"github.com/KyberNetwork/tradelogs/v2/pkg/storage/backfill"
	"go.uber.org/zap"
)

type BackFiller struct {
	handler *handler.TradeLogHandler
	s       backfill.IStorage
	l       *zap.SugaredLogger
	rpc     *rpcnode.Client
}

func NewBackFiller(handler *handler.TradeLogHandler, s backfill.IStorage, l *zap.SugaredLogger, rpc *rpcnode.Client) *BackFiller {
	return &BackFiller{
		handler: handler,
		s:       s,
		l:       l,
		rpc:     rpc,
	}
}

func (w *BackFiller) Run() error {
	ticker := time.NewTicker(time.Hour)
	defer ticker.Stop()
	for ; ; <-ticker.C {
		for {
			first, last, err := w.getBlockRanges()
			if err != nil {
				return fmt.Errorf("cannot get block ranges: %w", err)
			}
			if first >= last {
				break
			}
			err = w.processBlock(last - 1)
			if err != nil {
				return fmt.Errorf("cannot process block: %w", err)
			}
		}
	}
}

// getBlockRanges return separated and non-overlapping ranges that cover all backfill ranges
func (w *BackFiller) getBlockRanges() (uint64, uint64, error) {
	states, err := w.s.Get()
	if err != nil {
		return 0, 0, err
	}
	if len(states) == 0 {
		return 0, 0, nil
	}

	first, last := states[0].DeployBlock, states[0].BackFilledBlock

	// get the oldest deploy block
	for _, state := range states {
		first = min(first, state.DeployBlock)
	}

	// get the newest filled block
	for _, state := range states {
		// fill new exchange
		if state.BackFilledBlock <= 0 {
			blockNumber, err := w.rpc.GetBlockNumber(context.Background())
			if err != nil {
				return 0, 0, fmt.Errorf("get backfill block number: %w", err)
			}
			last = blockNumber
			break
		}
		last = max(last, state.BackFilledBlock)
	}
	return first, last, nil
}

func (w *BackFiller) processBlock(blockNumber uint64) error {
	block, err := w.rpc.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		return fmt.Errorf("cannot get block %d: %w", blockNumber, err)
	}

	err = w.handler.ProcessBlock(block.Hash().String(), blockNumber, block.Time())
	if err != nil {
		return fmt.Errorf("cannot process block %d: %w", blockNumber, err)
	}

	err = w.s.Update(blockNumber)
	if err != nil {
		return fmt.Errorf("cannot update backfill table with %d: %w", blockNumber, err)
	}

	w.l.Infow("successfully backfill block", "block", blockNumber)
	return nil
}
