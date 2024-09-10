package worker

import (
	"context"
	"fmt"

	"github.com/KyberNetwork/evmlistener/pkg/types"
	"github.com/KyberNetwork/tradelogs/v2/pkg/handler"
)

type ParseLog struct {
	handler *handler.TradeLogHandler
}

func NewParseLog(handler *handler.TradeLogHandler) *ParseLog {
	return &ParseLog{
		handler: handler,
	}
}

func (w *ParseLog) Publish(ctx context.Context, topic string, data interface{}) error {
	err := w.processMessage(data.(types.Message))
	if err != nil {
		return fmt.Errorf("failed to process message: %w", err)
	}
	return nil
}

func (w *ParseLog) processMessage(msg types.Message) error {
	var deleteBlocks []uint64

	for _, block := range msg.RevertedBlocks {
		deleteBlocks = append(deleteBlocks, block.Number.Uint64())
	}
	if err := w.handler.RevertBlock(deleteBlocks); err != nil {
		return fmt.Errorf("failed to revert blocks: %w", err)
	}

	for _, block := range msg.NewBlocks {
		err := w.handler.ProcessBlock(block.Hash, block.Number.Uint64(), block.Timestamp)
		if err != nil {
			return fmt.Errorf("failed to process new block: %w", err)
		}
	}
	return nil
}
