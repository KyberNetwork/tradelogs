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
	res := storage.TradeLog{
		OrderHash:        common.Hash(e.OrderHash).String(),
		MakerTokenAmount: e.MakingAmount.String(),
		ContractAddress:  e.Raw.Address.String(),
		BlockNumber:      e.Raw.BlockNumber,
		TxHash:           e.Raw.TxHash.String(),
		LogIndex:         uint64(e.Raw.Index),
		Timestamp:        blockTime * 1000,
		EventHash:        p.eventHash,
	}

	res, err = p.detectOneInchRfqTrade(res, log.TxHash.String(), blockTime)
	if err != nil {
		return storage.TradeLog{}, err
	}
	return res, nil
}

func (p *Parser) detectOneInchRfqTrade(order storage.TradeLog, txHash string, blockTimestamp uint64) (storage.TradeLog, error) {
	traceCall, err := p.rpcNode.FetchTraceCall(context.Background(), txHash)
	if err != nil {
		return order, err
	}

	order, err = p.recursiveDetectOneInchRFQTrades(order, traceCall, blockTimestamp)
	if err != nil {
		return order, err
	}

	return order, nil
}

func (p *Parser) recursiveDetectOneInchRFQTrades(tradeLog storage.TradeLog, traceCall types.CallFrame, blockTimestamp uint64) (storage.TradeLog, error) {
	var (
		err error
	)
	isOneInchRFQTrade := p.isOneInchRFQTrades(traceCall.Logs)

	if isOneInchRFQTrade {
		contractCall, err := decoder.Decode(OneinchMetaData.ABI, traceCall.Input)
		if err != nil {
			return tradeLog, err
		}

		return ToTradeLog(tradeLog, contractCall)
	}

	for _, subCall := range traceCall.Calls {
		tradeLog, err = p.recursiveDetectOneInchRFQTrades(tradeLog, subCall, blockTimestamp)
		if err != nil {
			return tradeLog, err
		}
	}
	return tradeLog, nil
}

func (p *Parser) isOneInchRFQTrades(eventLogs []types.CallLog) bool {
	for _, eventLog := range eventLogs {
		if len(eventLog.Topics) == 0 {
			continue
		}

		if !strings.EqualFold(eventLog.Topics[0].String(), p.eventHash) {
			continue
		}
		return true
	}

	return false
}
