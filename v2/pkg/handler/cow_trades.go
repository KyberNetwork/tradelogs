package handler

import (
	"fmt"

	cowTradeParser "github.com/KyberNetwork/tradelogs/v2/pkg/parser/cow_protocol/cowtrade_parser"
	cowTransferParser "github.com/KyberNetwork/tradelogs/v2/pkg/parser/cow_protocol/cowtransfer_parser"
	cowStorage "github.com/KyberNetwork/tradelogs/v2/pkg/storage/cow_protocol"
	"github.com/KyberNetwork/tradelogs/v2/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
)

type CowTradesHandler struct {
	l              *zap.SugaredLogger
	storage        *cowStorage.CowTradeStorage
	tradeParser    *cowTradeParser.CowTradeParser
	transferParser *cowTransferParser.CowTransferParser
}

func NewCowTradeHandler(
	l *zap.SugaredLogger,
	cowTradeStorage *cowStorage.CowTradeStorage,
	tradeParser *cowTradeParser.CowTradeParser,
	transferParser *cowTransferParser.CowTransferParser,
) *CowTradesHandler {
	return &CowTradesHandler{
		l:              l,
		storage:        cowTradeStorage,
		tradeParser:    tradeParser,
		transferParser: transferParser,
	}
}

func (h *CowTradesHandler) ProcessBlock(blockHash string, blockNumber uint64, timestamp uint64, calls []types.TransactionCallFrame) error {
	err := h.processForCowTrade(calls, blockHash, blockNumber, timestamp)
	if err != nil {
		return fmt.Errorf("error when process block for cow trade: %d", blockNumber)
	}
	h.l.Infow("successfully process block for cow trade", "blockNumber", blockNumber)
	return nil
}

func (h *CowTradesHandler) processForCowTrade(calls []types.TransactionCallFrame, blockHash string, blockNumber uint64, timestamp uint64) error {
	logIndexStart := 0
	var (
		tradesResult    []cowStorage.CowTrade
		transfersResult []cowStorage.CowTransfer
	)

	for i, call := range calls {
		logIndexStart = assignLogIndexes(&call.CallFrame, logIndexStart)
		metadata := logMetadata{
			blockNumber: blockNumber,
			blockHash:   blockHash,
			txHash:      call.TxHash,
			txIndex:     i,
			timestamp:   timestamp,
		}

		cowTrades := h.processCallFrameForCowTrades(call.CallFrame, metadata)
		if len(cowTrades) == 0 {
			continue
		}
		cowTransfers := h.processCallFrameForCowTransfers(call.CallFrame, metadata)
		for j := range cowTrades {
			cowTrades[j].TxOrigin = call.CallFrame.From // txOrigin, messageSender == internalCall.From
			cowTrades[j].InteractContract = call.CallFrame.To
		}

		tradesResult = append(tradesResult, cowTrades...)
		transfersResult = append(transfersResult, cowTransfers...)
	}

	if len(tradesResult) == 0 {
		return nil
	}

	err := h.storage.InsertCowTrades(tradesResult)
	if err != nil {
		return fmt.Errorf("insert cow trades error: %w", err)
	}
	h.l.Infow("successfully insert cow trades", "blockNumber", blockNumber, "number", len(tradesResult))
	err = h.storage.InsertCowTransfers(transfersResult)
	if err != nil {
		return fmt.Errorf("insert cow transfers error: %w", err)
	}
	h.l.Infow("successfully insert cow transfers", "blockNumber", blockNumber, "number", len(transfersResult))

	return nil
}

func (h *CowTradesHandler) processCallFrameForCowTrades(call types.CallFrame, metadata logMetadata) []cowStorage.CowTrade {
	result := make([]cowStorage.CowTrade, 0)

	// process the sub trace calls
	for _, traceCall := range call.Calls {
		cowTrades := h.processCallFrameForCowTrades(traceCall, metadata)
		result = append(result, cowTrades...)
	}
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

		if !h.tradeParser.IsMatchLog(ethLog) {
			continue
		}
		// parse trade log
		cowTrade, err := h.tradeParser.Parse(ethLog, metadata.timestamp)
		if err != nil {
			h.l.Errorw("error when parse log", "log", ethLog, "err", err, "parser", "cowTradeParser")
			continue
		}
		cowTrade.MessageSender = call.From
		result = append(result, cowTrade)
	}
	return result
}

func (h *CowTradesHandler) processCallFrameForCowTransfers(call types.CallFrame, metadata logMetadata) []cowStorage.CowTransfer {
	result := make([]cowStorage.CowTransfer, 0)

	// process the sub trace calls
	for _, traceCall := range call.Calls {
		cowTransfers := h.processCallFrameForCowTransfers(traceCall, metadata)
		result = append(result, cowTransfers...)
	}
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

		if !h.transferParser.IsMatchLog(ethLog) {
			continue
		}
		// parse trade log
		cowTransfer, err := h.transferParser.Parse(ethLog, metadata.timestamp)
		if err != nil {
			h.l.Errorw("error when parse log", "log", ethLog, "err", err, "parser", "cowTransferParesr")
			continue
		}
		result = append(result, cowTransfer)
	}
	return result
}

func (h *CowTradesHandler) RevertBlock(blocks []uint64) error {
	if len(blocks) == 0 {
		return nil
	}

	err := h.storage.DeleteCowTrades(blocks)
	if err != nil {
		return fmt.Errorf("delete blocks error: %w", err)
	}

	err = h.storage.DeleteCowTransfers(blocks)
	if err != nil {
		return fmt.Errorf("delete blocks error: %w", err)
	}

	return nil
}
