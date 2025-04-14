package worker

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/KyberNetwork/tradelogs/v2/pkg/constant"
	"github.com/KyberNetwork/tradelogs/v2/pkg/handler"
	"github.com/KyberNetwork/tradelogs/v2/pkg/parser"
	cowParser "github.com/KyberNetwork/tradelogs/v2/pkg/parser/cow_protocol/cowtrade_parser"
	"github.com/KyberNetwork/tradelogs/v2/pkg/rpcnode"
	"github.com/KyberNetwork/tradelogs/v2/pkg/storage/backfill"
	"github.com/KyberNetwork/tradelogs/v2/pkg/storage/state"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/util/sets"
)

const maxQueryBlockRange = 10000

type BackFiller struct {
	mu               sync.Mutex
	tradelogsHandler *handler.TradeLogHandler
	promoteeHandler  *handler.PromoteeHandler
	cowtradesHandler *handler.CowTradesHandler
	backfillStorage  backfill.IStorage
	stateStorage     state.Storage
	l                *zap.SugaredLogger
	rpc              rpcnode.IClient
	parsers          []parser.Parser
	cowtradeParsers  []cowParser.Parser
}

func NewBackFiller(
	tradelogsHandler *handler.TradeLogHandler,
	promoteeHandler *handler.PromoteeHandler,
	cowtradesHandler *handler.CowTradesHandler,
	backfillStorage backfill.IStorage, stateStorage state.Storage,
	l *zap.SugaredLogger, rpc rpcnode.IClient, parsers []parser.Parser,
	cowtradeParsers []cowParser.Parser,
) *BackFiller {
	return &BackFiller{
		tradelogsHandler: tradelogsHandler,
		promoteeHandler:  promoteeHandler,
		cowtradesHandler: cowtradesHandler,
		backfillStorage:  backfillStorage,
		stateStorage:     stateStorage,
		l:                l,
		rpc:              rpc,
		parsers:          parsers,
		cowtradeParsers:  cowtradeParsers,
	}
}

func (w *BackFiller) Run() {
	go func() {
		if err := w.run(); err != nil {
			w.l.Panicw("run system backfill failed", "err", err)
		}
	}()
}

func (w *BackFiller) run() error {
	// run only one process at a time
	if !w.mu.TryLock() {
		return nil
	}
	defer w.mu.Unlock()

	for {
		enable, err := w.isEnableBackfill()
		if err != nil {
			return err
		}
		if !enable {
			return nil
		}
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

		time.Sleep(time.Second)
	}
}

// isEnableBackfill check the state of enabling system backfill task
func (w *BackFiller) isEnableBackfill() (bool, error) {
	v, err := w.stateStorage.GetState(state.EnableSystemBackfillKey)
	if err != nil {
		return false, fmt.Errorf("cannot get system backfill status: %w", err)
	}
	enable, err := strconv.ParseBool(v)
	if err != nil {
		return false, fmt.Errorf("cannot parse system backfill status: %w", err)
	}
	return enable, nil
}

// CancelSystemBackfill stop the running system backfill task
func (w *BackFiller) CancelSystemBackfill() error {
	return w.stateStorage.SetState(state.EnableSystemBackfillKey, strconv.FormatBool(false))
}

// RunSystemBackfill run the system backfill task, if there is another running system backfill task, this task need to wait
func (w *BackFiller) RunSystemBackfill() error {
	enable, err := w.isEnableBackfill()
	if err != nil {
		return err
	}
	if enable {
		return nil
	}
	err = w.stateStorage.SetState(state.EnableSystemBackfillKey, strconv.FormatBool(true))
	if err != nil {
		return err
	}
	w.Run()
	return nil
}

func (w *BackFiller) IsValidExchange(exchange string) bool {
	for _, p := range w.parsers {
		if p.Exchange() == exchange {
			return true
		}
	}
	return false
}

// GetBlockRanges return the first and last block that cover all backfill ranges and the exclusive exchange that finishing backfill
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
	calls, err := w.rpc.FetchTraceCalls(context.Background(), block.Hash().String())
	if err != nil {
		return fmt.Errorf("fetch calls error: %w", err)
	}

	err = w.promoteeHandler.ProcessBlock(block.Hash().String(), blockNumber, block.Time(), calls)
	if err != nil {
		return fmt.Errorf("cannot process block %d: %w", blockNumber, err)
	}

	err = w.tradelogsHandler.ProcessBlockWithExclusion(block.Hash().String(), blockNumber, block.Time(), exclusions, calls)
	if err != nil {
		return fmt.Errorf("cannot process block %d: %w", blockNumber, err)
	}

	w.l.Infow("successfully backfill block", "block", blockNumber)
	return nil
}

func (w *BackFiller) processBlockForCowTrade(blockNumber uint64) error {
	block, err := w.rpc.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		return fmt.Errorf("cannot get block %d: %w", blockNumber, err)
	}
	calls, err := w.rpc.FetchTraceCalls(context.Background(), block.Hash().String())
	if err != nil {
		return fmt.Errorf("fetch calls error: %w", err)
	}

	err = w.cowtradesHandler.ProcessBlock(block.Hash().String(), blockNumber, block.Time(), calls)
	if err != nil {
		return fmt.Errorf("cannot process block %d: %w", blockNumber, err)
	}

	w.l.Infow("successfully backfill block", "block", blockNumber)
	return nil
}

func (w *BackFiller) BackfillByTask(task backfill.Task) {
	switch task.Exchange {
	case constant.CowProtocol:
		w.BackfillForCowProtocol(task)
	default:
		w.BackfillByExchange(task)
	}
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

	// we need to handle by batch to avoid too large block range
	for currentTo := to; currentTo >= from; {
		currentFrom := max(currentTo-maxQueryBlockRange, from)

		// error can occur while filtering blocks, we handle partial result and
		// mark the status as `failed` after processing received blocks.
		// else if there are no error, mark the status as `done` after processing all blocks
		blocks, exclusions, getBlockErr := w.getBlockByExchange(currentFrom, currentTo, task.Exchange)
		if getBlockErr != nil {
			w.l.Errorw("error when get block by exchange", "task", task, "err", getBlockErr)
		}

		w.l.Infow("start to backfill blocks", "task_id", task.ID,
			"current_from", currentFrom, "current_to", currentTo, "num_block", len(blocks))

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

		// mark the status = `failed`
		if getBlockErr != nil {
			err = w.backfillStorage.UpdateTask(task.ID, nil, backfill.StatusTypeFailed)
			if err != nil {
				w.l.Errorw("cannot update task status", "task", task, "err", err)
			}
			return
		}

		currentTo = currentFrom - 1
	}

	// else mark the status = `done`
	err = w.backfillStorage.UpdateTask(task.ID, &task.FromBlock, backfill.StatusTypeDone)
	if err != nil {
		w.l.Errorw("cannot update task status", "task_id", task.ID, "err", err)
	}
}

// getBlockByExchange get the blocks having logs of specific exchange, the block number sorted descending
func (w *BackFiller) getBlockByExchange(from, to uint64, exchange string) ([]uint64, sets.Set[string], error) {
	var (
		address      string
		topics       []string
		exclusions   = sets.New[string]()
		blocksNumber = sets.New[uint64]()
		logs         []ethereumTypes.Log
		err          error
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
	for currentTo := to; currentTo >= from; {
		currentFrom := max(currentTo-maxQueryBlockRange, from)

		// get logs
		logs, err = w.rpc.FetchLogs(context.Background(), currentFrom, currentTo, address, topics)
		if err != nil {
			break
		}

		// get blocks need to backfill
		for _, l := range logs {
			blocksNumber.Insert(l.BlockNumber)
		}

		// update new block range
		currentTo = currentFrom - 1
	}

	// sort the blocks descending
	blocks := blocksNumber.UnsortedList()
	sort.Slice(blocks, func(i, j int) bool {
		return blocks[i] > blocks[j]
	})

	return blocks, exclusions, err
}

func (w *BackFiller) BackfillForCowProtocol(task backfill.Task) {
	from, to := task.FromBlock, task.ToBlock
	if task.ProcessedBlock > 0 {
		to = task.ProcessedBlock - 1
	}

	// update the status before starting to backfill
	err := w.backfillStorage.UpdateTask(task.ID, nil, backfill.StatusTypeProcessing)
	if err != nil {
		w.l.Errorw("cannot update task status", "task", task.ID, "err", err)
		return
	}

	// we need to handle by batch to avoid too large block range
	for currentTo := to; currentTo >= from; {
		currentFrom := max(currentTo-maxQueryBlockRange, from)

		// error can occur while filtering blocks, we handle partial result and
		// mark the status as `failed` after processing received blocks.
		// else if there are no error, mark the status as `done` after processing all blocks
		blocks, getBlockErr := w.getBlockByCowProtocol(currentFrom, currentTo)
		if getBlockErr != nil {
			w.l.Errorw("error when get block by exchange", "task", task, "err", getBlockErr)
		}

		w.l.Infow("start to backfill blocks", "task_id", task.ID,
			"current_from", currentFrom, "current_to", currentTo, "num_block", len(blocks))

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

			err = w.processBlockForCowTrade(b)
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

		// mark the status = `failed`
		if getBlockErr != nil {
			err = w.backfillStorage.UpdateTask(task.ID, nil, backfill.StatusTypeFailed)
			if err != nil {
				w.l.Errorw("cannot update task status", "task", task, "err", err)
			}
			return
		}

		currentTo = currentFrom - 1
	}

	// else mark the status = `done`
	err = w.backfillStorage.UpdateTask(task.ID, &task.FromBlock, backfill.StatusTypeDone)
	if err != nil {
		w.l.Errorw("cannot update task status", "task_id", task.ID, "err", err)
	}
}

func (w *BackFiller) getBlockByCowProtocol(from, to uint64) ([]uint64, error) {
	var (
		address      string
		topics       []string
		blocksNumber = sets.New[uint64]()
		logs         []ethereumTypes.Log
		err          error
	)
	// get exchange address and topics to filter logs
	for _, p := range w.cowtradeParsers {
		address = constant.AddrCowProtocol
		topics = p.Topics()
		break
	}

	for currentTo := to; currentTo >= from; {
		currentFrom := max(currentTo-maxQueryBlockRange, from)

		// get logs
		logs, err = w.rpc.FetchLogs(context.Background(), currentFrom, currentTo, address, topics)
		if err != nil {
			break
		}

		// get blocks need to backfill
		for _, l := range logs {
			blocksNumber.Insert(l.BlockNumber)
		}

		// update new block range
		currentTo = currentFrom - 1
	}

	// sort the blocks descending
	blocks := blocksNumber.UnsortedList()
	sort.Slice(blocks, func(i, j int) bool {
		return blocks[i] > blocks[j]
	})

	return blocks, err
}
