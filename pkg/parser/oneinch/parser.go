package oneinch

import (
	"context"
	"errors"
	"fmt"
	"github.com/KyberNetwork/tradelogs/internal/types"
	"github.com/KyberNetwork/tradelogs/pkg/abitypes"
	"github.com/KyberNetwork/tradelogs/pkg/decoder"
	"github.com/KyberNetwork/tradelogs/pkg/rpcnode"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/lru"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strings"
)

const (
	FilledEvent = "OrderFilledRFQ"
)

var (
	RFQOrderOutputArgument abi.Arguments
)

var ErrInvalidOneInchFilledTopic = errors.New("invalid oneinch order filled topic")

func init() {
	// uint256 filledMakingAmount, uint256 filledTakingAmount, bytes32 orderHash
	RFQOrderOutputArgument = abi.Arguments{
		{Name: "filledMakingAmount", Type: abitypes.Uint256},
		{Name: "filledTakingAmount", Type: abitypes.Uint256},
		{Name: "orderHash", Type: abitypes.Bytes32},
	}
}

type Parser struct {
	abi       *abi.ABI
	ps        *OneinchFilterer
	eventHash string
	rpcNode   *rpcnode.Client
	// stores the last trace_calls by the tx_hash to reduce the requests when the transaction has more than one oneinch RFQ order
	latestTraceCall lru.BasicLRU[string, types.CallFrame]
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
		ps:              ps,
		abi:             ab,
		eventHash:       event.ID.String(),
		rpcNode:         rpcNode,
		latestTraceCall: lru.NewBasicLRU[string, types.CallFrame](10), // assumes one block hasn't more than 10 oneinch RFQ orders
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

	order, err = p.detectOneInchRfqTrade(order)
	if err != nil {
		return storage.TradeLog{}, err
	}
	return order, nil
}

func (p *Parser) ParseFromInternalCall(order storage.TradeLog, internalCall types.CallFrame) (storage.TradeLog, error) {
	output := internalCall.Output

	if output == "" {
		return order, errors.New("the output is blank")
	}
	filledMakingAmount, filledTakingAmount, orderHash, err := p.decodeOutput(output)
	if err != nil {
		return order, err
	}

	order.OrderHash = orderHash
	order.MakerTokenAmount = filledMakingAmount
	order.TakerTokenAmount = filledTakingAmount
	order.EventHash = p.eventHash

	contractCall, err := decoder.Decode(p.abi, internalCall.Input)
	if err != nil {
		return order, err
	}

	order, err = ToTradeLog(order, contractCall)
	if err != nil {
		return order, err
	}

	return order, nil
}

func (p *Parser) detectOneInchRfqTrade(order storage.TradeLog) (storage.TradeLog, error) {
	var (
		traceCall types.CallFrame
		err       error
	)
	if value, found := p.latestTraceCall.Get(order.TxHash); found {
		traceCall = value
	} else {
		traceCall, err = p.rpcNode.FetchTraceCall(context.Background(), order.TxHash)
		if err != nil {
			return order, err
		}
		p.latestTraceCall.Add(order.TxHash, traceCall)
	}

	order, err = p.recursiveDetectOneInchRFQTrades(order, traceCall)
	if err != nil {
		return order, err
	}

	return order, nil
}

func (p *Parser) recursiveDetectOneInchRFQTrades(tradeLog storage.TradeLog, traceCall types.CallFrame) (storage.TradeLog, error) {
	var (
		err error
	)
	isOneInchRFQTrade := p.isOneInchRFQTrades(tradeLog.OrderHash, traceCall)

	if isOneInchRFQTrade {
		return p.ParseFromInternalCall(tradeLog, traceCall)
	}

	for _, subCall := range traceCall.Calls {
		tradeLog, err = p.recursiveDetectOneInchRFQTrades(tradeLog, subCall)
		if err != nil {
			return tradeLog, err
		}
	}
	return tradeLog, nil
}

func (p *Parser) isOneInchRFQTrades(orderHash string, traceCall types.CallFrame) bool {
	for _, eventLog := range traceCall.Logs {
		if len(eventLog.Topics) == 0 {
			continue
		}

		if !strings.EqualFold(eventLog.Topics[0].String(), p.eventHash) {
			continue
		}

		_, _, orderHashFromOutput, err := p.decodeOutput(traceCall.Output)
		if err != nil {
			return false
		}
		return orderHash == orderHashFromOutput
	}
	return false
}

func (p *Parser) decodeOutput(output string) (string, string, string, error) {
	bytes, err := hexutil.Decode(output)
	if err != nil {
		return "", "", "", err
	}
	outputParams, err := RFQOrderOutputArgument.Unpack(bytes)
	if err != nil {
		return "", "", "", err
	}

	if len(outputParams) < 3 {
		return "", "", "", err
	}

	filledMakingAmountFromOutput, ok := outputParams[0].(*big.Int)
	if !ok {
		return "", "", "", errors.New("cant convert the filledMakingAmount from the output")
	}

	filledTakingAmountFromOutput, ok := outputParams[1].(*big.Int)
	if !ok {
		return "", "", "", errors.New("cant convert the filledTakingAmount from the output")
	}
	orderHashParams, ok := outputParams[2].([32]byte)
	if !ok {
		return "", "", "", errors.New("cant convert the order hash from the output")
	}

	orderHash := hexutil.Encode(orderHashParams[:])
	return filledMakingAmountFromOutput.String(), filledTakingAmountFromOutput.String(), orderHash, nil
}
