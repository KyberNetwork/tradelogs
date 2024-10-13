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
	rpc             rpcnode.IClient
	parsers         []parser.Parser
}

func NewBackFiller(handler *handler.TradeLogHandler, backfillStorage backfill.IStorage, stateStorage state.Storage,
	l *zap.SugaredLogger, rpc rpcnode.IClient, parsers []parser.Parser) *BackFiller {
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
		first, last, exclusions, err := w.GetBlockRanges()
		if err != nil {
			return fmt.Errorf("cannot get block ranges: %w", err)
		}
		if first >= last {
			return nil
		}
		w.l.Infow("backfill blocks", "first", first, "last", last)
		err = w.processBlock(last-1, exclusions)
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

func (w *BackFiller) IsValidExchange(exchange string) bool {
	for _, p := range w.parsers {
		if p.Exchange() == exchange {
			return true
		}
	}
	return false
}

// GetBlockRanges return separated and non-overlapping ranges that cover all backfill ranges
func (w *BackFiller) GetBlockRanges() (uint64, uint64, sets.Set[string], error) {
	states, err := w.backfillStorage.Get()
	if err != nil {
		return 0, 0, nil, err
	}
	if len(states) == 0 {
		return 0, 0, sets.New[string](), nil
	}

	// the result sorted by deploy block ascending, so when all exchanges are completely backfilled,
	// the first will be the oldest deploy block among exchanges
	first, last := states[0].DeployBlock, states[0].BackFilledBlock
	exclusions := sets.New[string]()

	// get the oldest deploy block
	for _, s := range states {
		if s.DeployBlock >= s.BackFilledBlock && s.BackFilledBlock != 0 {
			exclusions.Insert(s.Exchange)
			continue
		}
		first = min(first, s.DeployBlock)
	}

	// get the newest filled block
	for _, s := range states {
		if exclusions.Has(s.Exchange) {
			continue
		}
		// fill new exchange
		if s.BackFilledBlock <= 0 {
			blockNumber, err := w.getRecentBlock()
			if err != nil {
				return 0, 0, nil, err
			}
			last = blockNumber
			break
		}
		last = max(last, s.BackFilledBlock)
	}
	return first, last, exclusions, nil
}

// getRecentBlock get the newest processed block to backfill for new deployed exchange
func (w *BackFiller) getRecentBlock() (uint64, error) {
	var blockNumber uint64
	block, err := w.stateStorage.GetState(state.ProcessedBlockKey)
	if err == nil {
		blockNumber, err = strconv.ParseUint(block, 10, 64)
		if err == nil {
			return blockNumber + 1, nil
		}
	}
	w.l.Errorw("cannot get from db", "err", err)
	blockNumber, err = w.rpc.GetBlockNumber(context.Background())
	if err != nil {
		return 0, fmt.Errorf("cannot get from node: %w", err)
	}
	return blockNumber, nil
}

func (w *BackFiller) processBlock(blockNumber uint64, exclusions sets.Set[string]) error {
	block, err := w.rpc.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		return fmt.Errorf("cannot get block %d: %w", blockNumber, err)
	}

	err = w.handler.ProcessBlockWithExclusion(block.Hash().String(), blockNumber, block.Time(), exclusions)
	if err != nil {
		return fmt.Errorf("cannot process block %d: %w", blockNumber, err)
	}

	w.l.Infow("successfully backfill block", "block", blockNumber)
	return nil
}

func (w *BackFiller) BackfillByExchange(task backfill.Task) {
	from, to := task.FromBlock, task.ToBlock
	if task.ProcessedBlock > 0 {
		to = task.ProcessedBlock - 1
	}
	if from > to {
		return
	}

	// update the status before starting to backfill
	err := w.backfillStorage.UpdateTask(task.ID, nil, backfill.StatusTypeProcessing)
	if err != nil {
		w.l.Errorw("cannot update task status", "task", task.ID, "err", err)
		return
	}

	blocks, exclusions, err := w.getBlockByExchange(from, to, task.Exchange)
	if err != nil {
		w.l.Errorw("cannot get block", "task", task, "err", err)
		err = w.backfillStorage.UpdateTask(task.ID, nil, backfill.StatusTypeFailed)
		if err != nil {
			w.l.Errorw("cannot update task status", "task", task, "err", err)
		}
		return
	}

	w.l.Infow("start to backfill blocks", "task_id", task.ID, "blocks", blocks)

	// backfill from the newest blocks, if error occurs we can continue backfill from error block
	for _, b := range blocks {
		// check the task status, stop if canceled
		task, err = w.backfillStorage.GetTaskByID(task.ID)
		if err != nil {
			w.l.Errorw("cannot get backfill task", "task_id", task.ID, "err", err)
			return
		}
		if task.Status == backfill.StatusTypeCanceled {
			w.l.Infow("cannot backfill because task is canceled", "task_id", task.ID)
			return
		}

		err = w.processBlock(b, exclusions)
		if err != nil {
			w.l.Errorw("cannot backfill block", "task_id", task.ID, "block", b, "err", err)
			err = w.backfillStorage.UpdateTask(task.ID, nil, backfill.StatusTypeFailed)
			if err != nil {
				w.l.Errorw("cannot update task status", "task_id", task.ID, "err", err)
			}
			return
		}

		// update the processed block
		err = w.backfillStorage.UpdateTask(task.ID, &b, "")
		if err != nil {
			w.l.Errorw("cannot update task in db", "id", task.ID, "err", err)
			return
		}
	}

	err = w.backfillStorage.UpdateTask(task.ID, &task.FromBlock, backfill.StatusTypeDone)
	if err != nil {
		w.l.Errorw("cannot update task status", "task_id", task.ID, "err", err)
	}
}

// getBlockByExchange get the blocks having logs of specific exchange, the block number sorted descending
func (w *BackFiller) getBlockByExchange(from, to uint64, exchange string) ([]uint64, sets.Set[string], error) {
	var (
		address    string
		topics     []string
		exclusions = sets.New[string]()
	)
	// get exchange address and topics to filter logs
	for _, p := range w.parsers {
		if strings.EqualFold(p.Exchange(), exchange) {
			address = p.Address()
			topics = p.Topics()
			continue
		}
		exclusions.Insert(p.Exchange())
	}

	// get logs
	logs, err := w.rpc.FetchLogs(context.Background(), from, to, address, topics)
	if err != nil {
		return nil, nil, err
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

	return blocks, exclusions, nil
}
