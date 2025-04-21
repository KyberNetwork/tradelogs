package handler

import (
	"encoding/json"
	"fmt"
	"math/big"

	erc20Parser "github.com/KyberNetwork/tradelogs/v2/pkg/parser/ERC20_transfer"
	cowParser "github.com/KyberNetwork/tradelogs/v2/pkg/parser/cow_protocol"
	cowStorage "github.com/KyberNetwork/tradelogs/v2/pkg/storage/cow_protocol"
	"github.com/KyberNetwork/tradelogs/v2/pkg/types"
	"github.com/KyberNetwork/tradelogs/v2/pkg/util"
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
		return fmt.Errorf("error when process block %d for cow trade: %w", blockNumber, err)
	}
	h.l.Infow("successfully process block for cow trade", "blockNumber", blockNumber)
	return nil
}

func (h *CowTradesHandler) processForCowTrade(calls []types.TransactionCallFrame, blockHash string, blockNumber uint64, timestamp uint64) error {
	var (
		tradesResult    []cowStorage.CowTrade
		transfersResult []cowStorage.CowTransfer
		callFrameResult []cowStorage.CowTradeCallFrame
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
		callframeJson, err := json.Marshal(call)
		if err != nil {
			h.l.Errorw("cannot marshal callframe", "error", err)
			continue
		}
		callFrameResult = append(callFrameResult, cowStorage.CowTradeCallFrame{
			TxHash:      call.TxHash,
			CallFrame:   callframeJson,
			BlockNumber: blockNumber,
		})
	}

	if len(tradesResult) == 0 {
		return nil
	}

	err := h.storage.UpsertCowTrades(tradesResult)
	if err != nil {
		return fmt.Errorf("insert cow trades error: %w", err)
	}
	h.l.Infow("successfully insert cow trades", "blockNumber", blockNumber, "number", len(tradesResult))
	err = h.storage.InsertCowTransfers(transfersResult)
	if err != nil {
		return fmt.Errorf("insert cow transfers error: %w", err)
	}
	h.l.Infow("successfully insert cow transfers", "blockNumber", blockNumber, "number", len(transfersResult))
	err = h.storage.InsertCowCallFrame(callFrameResult)
	if err != nil {
		return fmt.Errorf("insert cow callframe error: %w", err)
	}
	h.l.Infow("successfully insert cow callframe", "blockNumber", blockNumber, "number", len(callFrameResult))

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

		switch {
		case h.tradeParser.IsContractLog(ethLog):
			cowTrade, err := h.tradeParser.Parse(ethLog, metadata.timestamp)
			if err != nil {
				h.l.Errorw("error when parse log", "log", ethLog, "err", err, "parser", "cowTransferParesr")
				continue
			}
			cowTrade.MessageSender = call.From
			tradesResult = append(tradesResult, cowTrade)
		case h.transferParser.IsContractLog(ethLog):
			cowTransfer, err := h.transferParser.Parse(ethLog, metadata.timestamp)
			if err != nil {
				h.l.Errorw("error when parse log", "log", ethLog, "err", err, "parser", "cowTransferParesr")
				continue
			}
			transfersResult = append(transfersResult, cowTransfer)
		}
	}
	if !isNativeTransfer(call.Value) {
		return tradesResult, transfersResult
	}
	amountStr := h.convertHexToDecimal(call.Value)
	nativeTransfer := cowStorage.CowTransfer{
		TxHash:       common.HexToHash(metadata.txHash).String(),
		BlockNumber:  metadata.blockNumber,
		Timestamp:    metadata.timestamp * 1000,
		FromAddress:  call.From,
		ToAddress:    call.To,
		Amount:       amountStr,
		Token:        util.NativeTokenAddress,
		TransferType: string(util.TransferTypeNative),
	}
	transfersResult = append(transfersResult, nativeTransfer)
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

	err = h.storage.DeleteCowCallFrame(blocks)
	if err != nil {
		return fmt.Errorf("delete blocks error: %w", err)
	}

	return nil
}

func (h *CowTradesHandler) convertHexToDecimal(value string) string {
	if len(value) <= 2 {
		return value
	}
	amount := new(big.Int)
	amount, success := amount.SetString(value[2:], 16)
	amountStr := amount.String()
	if !success {
		h.l.Error("cannot convert hex to decimal")
		amountStr = value
	}
	return amountStr
}

func isNativeTransfer(value string) bool {
	if value != "" && value != "0x0" {
		return true
	}
	return false
}
