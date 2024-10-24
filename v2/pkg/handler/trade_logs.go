package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/KyberNetwork/tradelogs/v2/pkg/kafka"
	"github.com/KyberNetwork/tradelogs/v2/pkg/parser"
	"github.com/KyberNetwork/tradelogs/v2/pkg/rpcnode"
	"github.com/KyberNetwork/tradelogs/v2/pkg/storage/tradelogs"
	storageTypes "github.com/KyberNetwork/tradelogs/v2/pkg/storage/tradelogs/types"
	"github.com/KyberNetwork/tradelogs/v2/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/util/sets"
)

type TradeLogHandler struct {
	l          *zap.SugaredLogger
	rpcClient  rpcnode.IClient
	storage    *tradelogs.Manager
	parsers    []parser.Parser
	kafkaTopic string
	publisher  kafka.Publisher
}

type logMetadata struct {
	blockNumber uint64
	blockHash   string
	txHash      string
	txIndex     int
	timestamp   uint64
}

func NewTradeLogHandler(l *zap.SugaredLogger, rpc rpcnode.IClient, storage *tradelogs.Manager, parsers []parser.Parser,
	kafkaTopic string, publisher kafka.Publisher) *TradeLogHandler {
	return &TradeLogHandler{
		l:          l,
		rpcClient:  rpc,
		storage:    storage,
		parsers:    parsers,
		kafkaTopic: kafkaTopic,
		publisher:  publisher,
	}
}

func (h *TradeLogHandler) ProcessBlock(blockHash string, blockNumber uint64, timestamp uint64) error {
	return h.ProcessBlockWithExclusion(blockHash, blockNumber, timestamp, sets.New[string]())
}

func (h *TradeLogHandler) ProcessBlockWithExclusion(blockHash string, blockNumber uint64, timestamp uint64, exclusions sets.Set[string]) error {
	// fetch trace call
	calls, err := h.rpcClient.FetchTraceCalls(context.Background(), blockHash)
	if err != nil {
		return fmt.Errorf("fetch calls error: %w", err)
	}

	// remove old trade log in db of processing block
	err = h.storage.DeleteWithExclusions([]uint64{blockNumber}, exclusions)
	if err != nil {
		return fmt.Errorf("delete blocks error: %w", err)
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

		tradeLogs := h.processCallFrame(call.CallFrame, metadata, exclusions)
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

	h.l.Infow("successfully process block", "blockNumber", blockNumber)

	return nil
}

func (h *TradeLogHandler) processCallFrame(call types.CallFrame, metadata logMetadata, exclusions sets.Set[string]) []storageTypes.TradeLog {
	result := make([]storageTypes.TradeLog, 0)

	// process the sub trace calls
	for _, traceCall := range call.Calls {
		tradeLogs := h.processCallFrame(traceCall, metadata, exclusions)
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
		if p == nil || exclusions.Has(p.Exchange()) {
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

	h.l.Infow("published revert message", "message", string(msgBytes))

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
