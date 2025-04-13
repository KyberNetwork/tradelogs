package promotee

import (
	"fmt"

	promotionparser "github.com/KyberNetwork/tradelogs/v2/pkg/parser/promotion"
	promoteeTypes "github.com/KyberNetwork/tradelogs/v2/pkg/storage/promotees"
	"github.com/KyberNetwork/tradelogs/v2/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
)

type PromoteeHandler struct {
	l       *zap.SugaredLogger
	storage *promoteeTypes.Storage
	parsers []promotionparser.Parser
}

type logMetadata struct {
	blockNumber uint64
	blockHash   string
	txHash      string
	txIndex     int
	timestamp   uint64
}

func NewPromoteeHandler(
	l *zap.SugaredLogger,
	promoteeStorage *promoteeTypes.Storage,
	promotionParsers []promotionparser.Parser,
) *PromoteeHandler {
	return &PromoteeHandler{
		l:       l,
		storage: promoteeStorage,
		parsers: promotionParsers,
	}
}

func (h *PromoteeHandler) ProcessBlock(blockHash string, blockNumber uint64, timestamp uint64, calls []types.TransactionCallFrame) error {
	err := h.processForPromotion(calls, blockHash, blockNumber, timestamp)
	if err != nil {
		return fmt.Errorf("error when process block for promotee: %d", blockNumber)
	}

	h.l.Infow("successfully process block for promotee", "blockNumber", blockNumber)
	return nil
}

func (h *PromoteeHandler) processForPromotion(calls []types.TransactionCallFrame, blockHash string, blockNumber uint64, timestamp uint64) error {
	for i, call := range calls {
		metadata := logMetadata{
			blockNumber: blockNumber,
			blockHash:   blockHash,
			txHash:      call.TxHash,
			txIndex:     i,
			timestamp:   timestamp,
		}

		promotees := h.processCallFrame(call.CallFrame, metadata)
		if len(promotees) == 0 {
			continue
		}

		err := h.storage.Insert(promotees)
		if err != nil {
			return fmt.Errorf("write to storage error: %w", err)
		}
		h.l.Infow("successfully insert promotees", "blockNumber", blockNumber)

	}

	return nil
}

func (h *PromoteeHandler) processCallFrame(call types.CallFrame, metadata logMetadata) []promoteeTypes.Promotee {
	result := make([]promoteeTypes.Promotee, 0)

	// process the sub trace calls
	for _, traceCall := range call.Calls {
		promotees := h.processCallFrame(traceCall, metadata)
		result = append(result, promotees...)
	}

	// process current trace call
	for _, log := range call.Logs {
		ethLog := ethereumTypes.Log{
			Address:     log.Address,
			Topics:      log.Topics,
			Data:        common.FromHex(log.Data),
			TxIndex:     uint(metadata.txIndex),
			TxHash:      common.HexToHash(metadata.txHash),
			BlockHash:   common.HexToHash(metadata.blockHash),
			BlockNumber: metadata.blockNumber,
			Index:       uint(log.Index),
		}
		// find the corresponding promotion parser
		p := h.findMatchingParser(ethLog)
		if p == nil {
			continue
		}

		// parse
		promotee, err := p.Parse(ethLog, metadata.timestamp)
		if err != nil {
			h.l.Errorw("error when parse log", "log", ethLog, "err", err, "parser", "promotion")
			continue
		}
		result = append(result, promotee)
	}
	return result
}

func (h *PromoteeHandler) findMatchingParser(log ethereumTypes.Log) promotionparser.Parser {
	for _, p := range h.parsers {
		if p.IsMatchLog(log) {
			return p
		}
	}
	return nil
}

func (h *PromoteeHandler) RevertBlock(blocks []uint64) error {
	if len(blocks) == 0 {
		return nil
	}

	err := h.storage.Delete(blocks)
	if err != nil {
		return fmt.Errorf("delete blocks for promotee storage error: %w", err)
	}

	return nil
}
