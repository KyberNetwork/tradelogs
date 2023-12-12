package oneinch

import (
	"context"
	"errors"
	"fmt"
	"github.com/KyberNetwork/tradelogs/internal/types"
	"github.com/KyberNetwork/tradelogs/pkg/decoder"
	"github.com/KyberNetwork/tradelogs/pkg/rpcnode"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
	"strings"
)

const (
	FilledEvent = "OrderFilledRFQ"
)

var ErrInvalidOneInchFilledTopic = errors.New("invalid oneinch order filled topic")

type Parser struct {
	abi       *abi.ABI
	ps        *OneinchFilterer
	eventHash string
	rpcNode   *rpcnode.Client
}

func MustNewParser(rpcUrl string) *Parser {
	ps, err := NewOneinchFilterer(common.Address{}, nil)
	if err != nil {
		panic(err)
	}
	ab, err := OneinchMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	event, ok := ab.Events[FilledEvent]
	if !ok {
		panic(fmt.Sprintf("no such event: %s", FilledEvent))
	}

	rpcNode, err := rpcnode.NewClient(rpcUrl)
	if err != nil {
		panic(fmt.Errorf("failed to init the rpc node client, err: %w", err))
	}
	return &Parser{
		ps:        ps,
		abi:       ab,
		eventHash: event.ID.String(),
		rpcNode:   rpcNode,
	}
}

func (p *Parser) Topics() []string {
	return []string{
		p.eventHash,
	}
}

func (p *Parser) Parse(log ethereumTypes.Log, blockTime uint64) (storage.TradeLog, error) {
	if len(log.Topics) > 0 && log.Topics[0].Hex() != p.eventHash {
		return storage.TradeLog{}, ErrInvalidOneInchFilledTopic
	}
	e, err := p.ps.ParseOrderFilledRFQ(log)
	if err != nil {
		return storage.TradeLog{}, err
	}
	order := storage.TradeLog{
		OrderHash:        common.Hash(e.OrderHash).String(),
		MakerTokenAmount: e.MakingAmount.String(),
		ContractAddress:  e.Raw.Address.String(),
		BlockNumber:      e.Raw.BlockNumber,
		TxHash:           e.Raw.TxHash.String(),
		LogIndex:         uint64(e.Raw.Index),
		Timestamp:        blockTime * 1000,
		EventHash:        p.eventHash,
	}

	order, err = p.detectOneInchRfqTrade(order, log, blockTime)
	if err != nil {
		return storage.TradeLog{}, err
	}
	return order, nil
}

func (p *Parser) detectOneInchRfqTrade(order storage.TradeLog, log ethereumTypes.Log, blockTimestamp uint64) (storage.TradeLog, error) {
	traceCall, err := p.rpcNode.FetchTraceCall(context.Background(), order.TxHash)
	if err != nil {
		return order, err
	}

	order, err = p.recursiveDetectOneInchRFQTrades(order, log, traceCall, blockTimestamp)
	if err != nil {
		return order, err
	}

	return order, nil
}

func (p *Parser) recursiveDetectOneInchRFQTrades(tradeLog storage.TradeLog, log ethereumTypes.Log, traceCall types.CallFrame, blockTimestamp uint64) (storage.TradeLog, error) {
	var (
		err error
	)
	isOneInchRFQTrade := p.isOneInchRFQTrades(log, traceCall.Logs)

	if isOneInchRFQTrade {
		contractCall, err := decoder.Decode(OneinchMetaData.ABI, traceCall.Input)
		if err != nil {
			return tradeLog, err
		}

		return ToTradeLog(tradeLog, contractCall)
	}

	for _, subCall := range traceCall.Calls {
		tradeLog, err = p.recursiveDetectOneInchRFQTrades(tradeLog, log, subCall, blockTimestamp)
		if err != nil {
			return tradeLog, err
		}
	}
	return tradeLog, nil
}

func (p *Parser) isOneInchRFQTrades(originalLog ethereumTypes.Log, eventLogs []types.CallLog) bool {
	for _, eventLog := range eventLogs {
		if len(eventLog.Topics) == 0 {
			continue
		}

		if !strings.EqualFold(eventLog.Topics[0].String(), p.eventHash) {
			continue
		}

		if !p.isSameEventLog(originalLog, eventLog) {
			continue
		}

		return true
	}

	return false
}

// compares two event logs to determine the extract rfq order
// replaces by the log index in the future(if the debug_traceTransaction endpoint supports the log index)
func (p *Parser) isSameEventLog(originalLog ethereumTypes.Log, eventLog types.CallLog) bool {
	// need to compare with the original log to get extract the order
	if eventLog.Data != hexutil.Encode(originalLog.Data) || !strings.EqualFold(eventLog.Address.String(), eventLog.Address.String()) {
		return false
	}

	isSameTopics := true
	for idx, topic := range originalLog.Topics {
		if len(eventLog.Topics) < idx {
			isSameTopics = false
			break
		}
		comparedTopic := eventLog.Topics[idx]
		if topic.String() != comparedTopic.String() {
			isSameTopics = false
			break
		}
	}
	return isSameTopics
}
