package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/KyberNetwork/tradelogs/v2/pkg/kafka"
	"github.com/KyberNetwork/tradelogs/v2/pkg/parser"
	"github.com/KyberNetwork/tradelogs/v2/pkg/rpcnode"
	"github.com/KyberNetwork/tradelogs/v2/pkg/storage/state"
	"github.com/KyberNetwork/tradelogs/v2/pkg/storage/tradelogs"
	storageTypes "github.com/KyberNetwork/tradelogs/v2/pkg/storage/tradelogs/types"
	"github.com/KyberNetwork/tradelogs/v2/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
)

type TradeLogHandler struct {
	l          *zap.SugaredLogger
	rpcClient  *rpcnode.Client
	storage    *tradelogs.Manager
	parsers    []parser.Parser
	kafkaTopic string
	publisher  kafka.Publisher
	state      state.Storage
}

type logMetadata struct {
	blockNumber uint64
	blockHash   string
	txHash      string
	txIndex     int
	timestamp   uint64
}

func NewTradeLogHandler(l *zap.SugaredLogger, rpc *rpcnode.Client, storage *tradelogs.Manager, parsers []parser.Parser,
	kafkaTopic string, publisher kafka.Publisher, state state.Storage) *TradeLogHandler {
	return &TradeLogHandler{
		l:          l,
		rpcClient:  rpc,
		storage:    storage,
		parsers:    parsers,
		kafkaTopic: kafkaTopic,
		publisher:  publisher,
		state:      state,
	}
}

func (h *TradeLogHandler) ProcessBlock(blockHash string, blockNumber uint64, timestamp uint64) error {
	// remove old trade log in db of processing block
	err := h.RevertBlock([]uint64{blockNumber})
	if err != nil {
		return fmt.Errorf("error when revert block number %d before processing: %w", blockNumber, err)
	}

	// fetch trace call
	calls, err := h.rpcClient.FetchTraceCalls(context.Background(), blockHash)
	if err != nil {
		return fmt.Errorf("fetch calls error: %w", err)
	}

	logIndexStart := 0
	for i, call := range calls {
		logIndexStart = assignLogIndexes(&call.CallFrame, logIndexStart)
		metadata := logMetadata{
			blockNumber: blockNumber,
			blockHash:   blockHash,
			txHash:      call.TxHash,
			txIndex:     i,
			timestamp:   timestamp,
		}

		tradeLogs := h.processCallFrame(call.CallFrame, metadata)
		if len(tradeLogs) == 0 {
			continue
		}

		for j := range tradeLogs {
			tradeLogs[j].MessageSender = call.CallFrame.From
			tradeLogs[j].InteractContract = call.CallFrame.To
		}

		err = h.storage.Insert(tradeLogs)
		if err != nil {
			return fmt.Errorf("write to storage error: %w", err)
		}
		h.l.Infow("successfully insert trade logs", "blockNumber", blockNumber)

		passCount, failCount := 0, 0
		for _, log := range tradeLogs {
			msgBytes, err := json.Marshal(kafka.Message{
				Type: kafka.MessageTypeTradeLog,
				Data: log,
			})
			if err != nil {
				h.l.Errorw(" error when marshal trade log to json", "blockNumber", blockNumber, "log", log, "err", err)
				failCount++
				continue
			}
			err = h.publisher.Publish(h.kafkaTopic, msgBytes)
			if err != nil {
				h.l.Errorw("error when publish trade log to kafka", "blockNumber", blockNumber, "log", log, "err", err)
				failCount++
				continue
			}
			passCount++
		}
		h.l.Infow("successfully publish trade logs", "blockNumber", blockNumber, "success", passCount, "fail", failCount)
	}

	// persist most recent processed block
	err = h.state.SetState(state.ProcessedBlockKey, strconv.FormatUint(blockNumber, 10))
	if err != nil {
		h.l.Errorw("cannot persist processed block", "blockNumber", blockNumber, "err", err)
		return fmt.Errorf("cannot persist processed block: %w", err)
	}
	h.l.Infow("successfully persist processed block", "blockNumber", blockNumber)

	return nil
}

func (h *TradeLogHandler) processCallFrame(call types.CallFrame, metadata logMetadata) []storageTypes.TradeLog {
	result := make([]storageTypes.TradeLog, 0)

	// process the sub trace calls
	for _, traceCall := range call.Calls {
		tradeLogs := h.processCallFrame(traceCall, metadata)
		result = append(result, tradeLogs...)
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

		// find the corresponding parser
		p := h.findMatchingParser(ethLog)
		if p == nil {
			continue
		}

		// parse trade log
		tradeLogs, err := p.ParseWithCallFrame(call, ethLog, metadata.timestamp)
		if err != nil {
			h.l.Errorw("error when parse log", "log", ethLog, "err", err, "parser", p.Exchange())
			continue
		}
		result = append(result, tradeLogs...)
	}
	return result
}

func (h *TradeLogHandler) findMatchingParser(log ethereumTypes.Log) parser.Parser {
	for _, p := range h.parsers {
		if p.LogFromExchange(log) {
			return p
		}
	}
	return nil
}

func (h *TradeLogHandler) RevertBlock(blocks []uint64) error {
	if len(blocks) == 0 {
		return nil
	}

	err := h.storage.Delete(blocks)
	if err != nil {
		return fmt.Errorf("delete blocks error: %w", err)
	}

	msgBytes, err := json.Marshal(kafka.Message{
		Type: kafka.MessageTypeRevert,
		Data: blocks,
	})
	if err != nil {
		h.l.Errorw(" error when marshal revert message to json", "err", err)
	}
	err = h.publisher.Publish(h.kafkaTopic, msgBytes)
	if err != nil {
		h.l.Errorw("error when publish revert message", "err", err)
	}
	h.l.Infow("published revert message", "message", msgBytes)

	return nil
}

func assignLogIndexes(cf *types.CallFrame, index int) int {
	subCallIndex := hexutil.Uint(0)
	for i := range cf.Logs {
		for subCallIndex < cf.Logs[i].Position {
			index = assignLogIndexes(&cf.Calls[subCallIndex], index)
			subCallIndex++
		}
		cf.Logs[i].Index = index
		index++
	}
	for subCallIndex < hexutil.Uint(len(cf.Calls)) {
		index = assignLogIndexes(&cf.Calls[subCallIndex], index)
		subCallIndex++
	}
	return index
}