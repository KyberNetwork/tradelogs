package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/KyberNetwork/tradelogs/v2/pkg/kafka"
	"github.com/KyberNetwork/tradelogs/v2/pkg/parser"
	"github.com/KyberNetwork/tradelogs/v2/pkg/promotionparser"
	"github.com/KyberNetwork/tradelogs/v2/pkg/rpcnode"
	promoteeTypes "github.com/KyberNetwork/tradelogs/v2/pkg/storage/promotees"
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
	l                *zap.SugaredLogger
	rpcClient        rpcnode.IClient
	storage          *tradelogs.Manager
	promoteeStorage  *promoteeTypes.Storage
	parsers          []parser.Parser
	promotionParsers []promotionparser.Parser
	kafkaTopic       string
	publisher        kafka.Publisher
}

type logMetadata struct {
	blockNumber uint64
	blockHash   string
	txHash      string
	txIndex     int
	timestamp   uint64
}

func NewTradeLogHandler(l *zap.SugaredLogger, rpc rpcnode.IClient,
	storage *tradelogs.Manager, promoteeStorage *promoteeTypes.Storage,
	parsers []parser.Parser, promotionParsers []promotionparser.Parser,
	kafkaTopic string, publisher kafka.Publisher) *TradeLogHandler {
	return &TradeLogHandler{
		l:                l,
		rpcClient:        rpc,
		storage:          storage,
		promoteeStorage:  promoteeStorage,
		parsers:          parsers,
		promotionParsers: promotionParsers,
		kafkaTopic:       kafkaTopic,
		publisher:        publisher,
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

	err = h.processForPromotion(calls, blockHash, blockNumber, timestamp)
	if err != nil {
		return fmt.Errorf("error when process block: %d", blockNumber)
	}

	err = h.processForTradelog(calls, blockHash, blockNumber, timestamp, exclusions)
	if err != nil {
		return fmt.Errorf("error when process block: %d", blockNumber)
	}

	h.l.Infow("successfully process block", "blockNumber", blockNumber)
	return nil
}

func (h *TradeLogHandler) processForTradelog(calls []types.TransactionCallFrame, blockHash string, blockNumber uint64, timestamp uint64, exclusions sets.Set[string]) error {
	logIndexStart := 0
	var result []storageTypes.TradeLog

	for i, call := range calls {
		logIndexStart = assignLogIndexes(&call.CallFrame, logIndexStart)
		metadata := logMetadata{
			blockNumber: blockNumber,
			blockHash:   blockHash,
			txHash:      call.TxHash,
			txIndex:     i,
			timestamp:   timestamp,
		}

		tradeLogs := h.processCallFrameForTradelog(call.CallFrame, metadata, exclusions)
		if len(tradeLogs) == 0 {
			continue
		}

		for j := range tradeLogs {
			tradeLogs[j].TxOrigin = call.CallFrame.From // txOrigin, messageSender == internalCall.From
			tradeLogs[j].InteractContract = call.CallFrame.To
		}

		result = append(result, tradeLogs...)
	}

	if len(result) == 0 {
		return nil
	}

	err := h.storage.Insert(result)
	if err != nil {
		return fmt.Errorf("write to storage error: %w", err)
	}
	h.l.Infow("successfully insert trade logs", "blockNumber", blockNumber, "number", len(result))

	passCount, failCount := 0, 0
	for _, log := range result {
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

	return nil
}

func (h *TradeLogHandler) processForPromotion(calls []types.TransactionCallFrame, blockHash string, blockNumber uint64, timestamp uint64) error {
	for i, call := range calls {
		metadata := logMetadata{
			blockNumber: blockNumber,
			blockHash:   blockHash,
			txHash:      call.TxHash,
			txIndex:     i,
			timestamp:   timestamp,
		}

		promotees := h.processCallFrameForPromotion(call.CallFrame, metadata)
		if len(promotees) == 0 {
			continue
		}

		err := h.promoteeStorage.Insert(promotees)
		if err != nil {
			return fmt.Errorf("write to storage error: %w", err)
		}
		h.l.Infow("successfully insert promotees", "blockNumber", blockNumber)

	}

	return nil
}

func (h *TradeLogHandler) processCallFrameForTradelog(call types.CallFrame, metadata logMetadata, exclusions sets.Set[string]) []storageTypes.TradeLog {
	result := make([]storageTypes.TradeLog, 0)

	// process the sub trace calls
	for _, traceCall := range call.Calls {
		tradeLogs := h.processCallFrameForTradelog(traceCall, metadata, exclusions)
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

		for i := range tradeLogs {
			tradeLogs[i].MessageSender = call.From
		}

		result = append(result, tradeLogs...)
	}
	return result
}

func (h *TradeLogHandler) processCallFrameForPromotion(call types.CallFrame, metadata logMetadata) []promoteeTypes.Promotee {
	result := make([]promoteeTypes.Promotee, 0)

	// process the sub trace calls
	for _, traceCall := range call.Calls {
		promotees := h.processCallFrameForPromotion(traceCall, metadata)
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
		p := h.findMatchingPromotionParser(ethLog)
		if p == nil {
			continue
		}

		// parse
		promotee, err := p.Parse(ethLog, metadata.timestamp)
		if err != nil {
			h.l.Errorw("error when parse log", "log", ethLog, "err", err, "parser", p.Contract())
			continue
		}
		result = append(result, promotee)
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

func (h *TradeLogHandler) findMatchingPromotionParser(log ethereumTypes.Log) promotionparser.Parser {
	for _, p := range h.promotionParsers {
		if p.LogFromContract(log) {
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

	err = h.promoteeStorage.Delete(blocks)
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
