package handler

import (
	"fmt"

	"github.com/KyberNetwork/tradelogs/v2/pkg/constant"
	erc20Parser "github.com/KyberNetwork/tradelogs/v2/pkg/parser/ERC20_transfer"
	cowParser "github.com/KyberNetwork/tradelogs/v2/pkg/parser/cow_protocol"
	cowStorage "github.com/KyberNetwork/tradelogs/v2/pkg/storage/cow_protocol"
	"github.com/KyberNetwork/tradelogs/v2/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
)

type CowTradesHandler struct {
	l              *zap.SugaredLogger
	storage        *cowStorage.CowTradeStorage
	tradeParser    cowParser.TradeParser
	transferParser *erc20Parser.ERC20TransferParser
}

func NewCowTradeHandler(
	l *zap.SugaredLogger,
	cowTradeStorage *cowStorage.CowTradeStorage,
	tradeParser cowParser.TradeParser,
	transferParser *erc20Parser.ERC20TransferParser,
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
	var (
		tradesResult    []cowStorage.CowTrade
		transfersResult []cowStorage.CowTransfer
	)

	for i, call := range calls {
		metadata := logMetadata{
			blockNumber: blockNumber,
			blockHash:   blockHash,
			txHash:      call.TxHash,
			txIndex:     i,
			timestamp:   timestamp,
		}

		cowTrades, cowTransfers := h.processCallFrameForCowTrades(call.CallFrame, metadata)
		if len(cowTrades) == 0 {
			continue
		}
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

	err := h.storage.UpsertCowTrades(tradesResult)
	if err != nil {
		return fmt.Errorf("insert cow trades error: %w", err)
	}
	h.l.Infow("successfully insert cow trades", "blockNumber", blockNumber, "number", len(tradesResult))
	err = h.storage.UpsertCowTransfers(transfersResult, false)
	if err != nil {
		return fmt.Errorf("insert cow transfers error: %w", err)
	}
	h.l.Infow("successfully insert cow transfers", "blockNumber", blockNumber, "number", len(transfersResult))

	return nil
}

func (h *CowTradesHandler) processCallFrameForCowTrades(call types.CallFrame, metadata logMetadata) ([]cowStorage.CowTrade, []cowStorage.CowTransfer) {
	tradesResult := make([]cowStorage.CowTrade, 0)
	transfersResult := make([]cowStorage.CowTransfer, 0)

	// process the sub trace calls
	for _, traceCall := range call.Calls {
		cowTrades, cowTransfers := h.processCallFrameForCowTrades(traceCall, metadata)
		tradesResult = append(tradesResult, cowTrades...)
		transfersResult = append(transfersResult, cowTransfers...)
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

		if h.tradeParser.IsContractLog(ethLog) {
			cowTrade, err := h.tradeParser.Parse(ethLog, metadata.timestamp)
			if err != nil {
				h.l.Errorw("error when parse log", "log", ethLog, "err", err, "parser", "cowTransferParesr")
			} else {
				cowTrade.MessageSender = call.From
				tradesResult = append(tradesResult, cowTrade)
			}
		}

		if h.transferParser.IsContractLog(ethLog) {
			cowTransfer, err := h.transferParser.Parse(ethLog, metadata.timestamp)
			if err != nil {
				h.l.Errorw("error when parse log", "log", ethLog, "err", err, "parser", "cowTransferParesr")
			} else {
				if cowTransfer.FromAddress == constant.AddrCowProtocol || cowTransfer.ToAddress == constant.AddrCowProtocol {
					transfersResult = append(transfersResult, cowTransfer)
				}
			}
		}

	}
	return tradesResult, transfersResult
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
