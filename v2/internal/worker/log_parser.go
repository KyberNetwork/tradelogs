package worker

import (
	"context"
	"fmt"
	"strconv"

	"github.com/KyberNetwork/evmlistener/pkg/types"
	"github.com/KyberNetwork/tradelogs/v2/pkg/handler"
	"github.com/KyberNetwork/tradelogs/v2/pkg/rpcnode"
	"github.com/KyberNetwork/tradelogs/v2/pkg/storage/state"
	"go.uber.org/zap"
)

type LogParser struct {
	rpcClient        rpcnode.IClient
	tradelogsHandler *handler.TradeLogHandler
	promoteeHandler  *handler.PromoteeHandler
	cowtradesHandler *handler.CowTradesHandler
	state            state.Storage
	l                *zap.SugaredLogger
}

func NewParseLog(
	rpcClient rpcnode.IClient,
	tradelogsHandler *handler.TradeLogHandler,
	promoteeHandler *handler.PromoteeHandler,
	cowtradesHandler *handler.CowTradesHandler,
	s state.Storage,
	l *zap.SugaredLogger,
) *LogParser {
	return &LogParser{
		rpcClient:        rpcClient,
		tradelogsHandler: tradelogsHandler,
		promoteeHandler:  promoteeHandler,
		cowtradesHandler: cowtradesHandler,
		state:            s,
		l:                l,
	}
}

func (w *LogParser) Publish(ctx context.Context, topic string, data interface{}) error {
	_, ok := data.(types.Message)
	if !ok {
		return fmt.Errorf("invalid data: not a message")
	}
	err := w.processMessage(data.(types.Message))
	if err != nil {
		return fmt.Errorf("failed to process message: %w", err)
	}
	return nil
}

func (w *LogParser) processMessage(msg types.Message) error {
	var deleteBlocks []uint64

	w.l.Infow("process message", "msg", msg)

	for _, block := range msg.RevertedBlocks {
		deleteBlocks = append(deleteBlocks, block.Number.Uint64())
	}
	if err := w.tradelogsHandler.RevertBlock(deleteBlocks); err != nil {
		return fmt.Errorf("failed to revert blocks: %w", err)
	}

	for _, block := range msg.NewBlocks {
		blockNumber := block.Number.Uint64()
		// fetch trace call
		calls, err := w.rpcClient.FetchTraceCalls(context.Background(), block.Hash)
		if err != nil {
			return fmt.Errorf("fetch calls error: %w", err)
		}
		err = w.promoteeHandler.ProcessBlock(block.Hash, blockNumber, block.Timestamp, calls)
		if err != nil {
			return fmt.Errorf("failed to process new block: %w", err)
		}
		err = w.tradelogsHandler.ProcessBlock(block.Hash, blockNumber, block.Timestamp, calls)
		if err != nil {
			return fmt.Errorf("failed to process new block: %w", err)
		}
		err = w.cowtradesHandler.ProcessBlock(block.Hash, blockNumber, block.Timestamp, calls)
		if err != nil {
			return fmt.Errorf("failed to process new block: %w", err)
		}

		// persist most recent processed block
		err = w.state.SetState(state.ProcessedBlockKey, strconv.FormatUint(blockNumber, 10))
		if err != nil {
			w.l.Errorw("cannot persist processed block", "blockNumber", blockNumber, "err", err)
			return fmt.Errorf("cannot persist processed block: %w", err)
		}

		w.l.Infow("successfully persist processed block", "blockNumber", blockNumber)
	}
	return nil
}
