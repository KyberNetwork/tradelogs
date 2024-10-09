package worker

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/KyberNetwork/tradelogs/v2/pkg/handler"
	"github.com/KyberNetwork/tradelogs/v2/pkg/parser"
	"github.com/KyberNetwork/tradelogs/v2/pkg/rpcnode"
	"github.com/KyberNetwork/tradelogs/v2/pkg/storage/backfill"
	"github.com/KyberNetwork/tradelogs/v2/pkg/storage/state"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/util/sets"
)

type BackFiller struct {
	handler         *handler.TradeLogHandler
	backfillStorage backfill.IStorage
	stateStorage    state.Storage
	l               *zap.SugaredLogger
	rpc             *rpcnode.Client
	parsers         []parser.Parser
}

func NewBackFiller(handler *handler.TradeLogHandler, backfillStorage backfill.IStorage, stateStorage state.Storage,
	l *zap.SugaredLogger, rpc *rpcnode.Client, parsers []parser.Parser) *BackFiller {
	return &BackFiller{
		handler:         handler,
		backfillStorage: backfillStorage,
		stateStorage:    stateStorage,
		l:               l,
		rpc:             rpc,
		parsers:         parsers,
	}
}

func (w *BackFiller) Run() error {
	for {
		first, last, err := w.getBlockRanges()
		if err != nil {
			return fmt.Errorf("cannot get block ranges: %w", err)
		}
		if first >= last {
			return nil
		}
		w.l.Infow("backfill blocks", "first", first, "last", last)
		err = w.processBlock(last - 1)
		if err != nil {
			w.l.Errorw("error when backfill block", "block", last-1, "err", err)
			return fmt.Errorf("cannot process block: %w", err)
		}

		err = w.backfillStorage.Update(last - 1)
		if err != nil {
			w.l.Errorw("error when update backfill table", "block", last-1, "err", err)
			return fmt.Errorf("cannot update backfill table with %d: %w", last-1, err)
		}
	}
}

// getBlockRanges return separated and non-overlapping ranges that cover all backfill ranges
func (w *BackFiller) getBlockRanges() (uint64, uint64, error) {
	states, err := w.backfillStorage.Get()
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
			blockNumber, err := w.getRecentBlock()
			if err != nil {
				return 0, 0, err
			}
			last = blockNumber
			break
		}
		last = max(last, state.BackFilledBlock)
	}
	return first, last, nil
}

// getRecentBlock get the newest processed block to backfill for new deployed exchange
func (w *BackFiller) getRecentBlock() (uint64, error) {
	var blockNumber uint64
	block, err := w.stateStorage.GetState(state.ProcessedBlockKey)
	if err == nil {
		blockNumber, err = strconv.ParseUint(block, 10, 64)
		if err == nil {
			return blockNumber, nil
		}
	}
	w.l.Errorw("cannot get from db", "err", err)
	blockNumber, err = w.rpc.GetBlockNumber(context.Background())
	if err != nil {
		return 0, fmt.Errorf("cannot get from node: %w", err)
	}
	return blockNumber, nil
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

	w.l.Infow("successfully backfill block", "block", blockNumber)
	return nil
}

func (w *BackFiller) BackfillByExchange(from, to uint64, exchange string) error {
	blocks, err := w.getBlockByExchange(from, to, exchange)
	if err != nil {
		return fmt.Errorf("cannot get block %d: %w", from, err)
	}

	w.l.Infow("start to backfill blocks", "blocks", blocks)

	// backfill from the newest blocks, if error occurs we can continue backfill from error block
	for _, b := range blocks {
		err = w.processBlock(b)
		if err != nil {
			w.l.Errorw("cannot backfill block", "block", b, "err", err)
			return fmt.Errorf("error when backfill to block %d: %w", b, err)
		}
	}

	return nil
}

// getBlockByExchange get the blocks having logs of specific exchange, the block number sorted descending
func (w *BackFiller) getBlockByExchange(from, to uint64, exchange string) ([]uint64, error) {
	var (
		address string
		topics  []string
	)
	// get exchange address and topics to filter logs
	for _, p := range w.parsers {
		if strings.EqualFold(p.Exchange(), exchange) {
			address = p.Address()
			topics = p.Topics()
			break
		}
	}

	// get logs
	logs, err := w.rpc.FetchLogs(context.Background(), from, to, address, topics)
	if err != nil {
		return nil, err
	}

	// get blocks need to backfill
	blocksNumber := sets.New[uint64]()
	for _, l := range logs {
		blocksNumber.Insert(l.BlockNumber)
	}

	// sort the blocks descending
	blocks := blocksNumber.UnsortedList()
	sort.Slice(blocks, func(i, j int) bool {
		return blocks[i] > blocks[j]
	})

	return blocks, nil
}
