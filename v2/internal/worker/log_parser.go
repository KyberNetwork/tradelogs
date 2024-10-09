package worker

import (
	"context"
	"fmt"
	"strconv"

	"github.com/KyberNetwork/evmlistener/pkg/types"
	"github.com/KyberNetwork/tradelogs/v2/pkg/handler"
	"github.com/KyberNetwork/tradelogs/v2/pkg/storage/state"
	"go.uber.org/zap"
)

type LogParser struct {
	handler *handler.TradeLogHandler
	state   state.Storage
	l       *zap.SugaredLogger
}

func NewParseLog(handler *handler.TradeLogHandler, s state.Storage, l *zap.SugaredLogger) *LogParser {
	return &LogParser{
		handler: handler,
		state:   s,
		l:       l,
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

	for _, block := range msg.RevertedBlocks {
		deleteBlocks = append(deleteBlocks, block.Number.Uint64())
	}
	if err := w.handler.RevertBlock(deleteBlocks); err != nil {
		return fmt.Errorf("failed to revert blocks: %w", err)
	}

	for _, block := range msg.NewBlocks {
		blockNumber := block.Number.Uint64()

		err := w.handler.ProcessBlock(block.Hash, blockNumber, block.Timestamp)
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
