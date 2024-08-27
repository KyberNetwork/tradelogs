package oneinchv6

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/KyberNetwork/tradelogs/pkg/abitypes"
	"github.com/KyberNetwork/tradelogs/pkg/decoder"
	"github.com/KyberNetwork/tradelogs/pkg/parser"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"github.com/KyberNetwork/tradelogs/pkg/tracecall"
	"github.com/KyberNetwork/tradelogs/pkg/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/lru"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	FilledEvent = "OrderFilled"
)

var (
	RFQOrderOutputArgument abi.Arguments
)

func init() {
	// uint256 filledMakingAmount, uint256 filledTakingAmount, bytes32 orderHash
	RFQOrderOutputArgument = abi.Arguments{
		{Name: "filledMakingAmount", Type: abitypes.Uint256},
		{Name: "filledTakingAmount", Type: abitypes.Uint256},
		{Name: "orderHash", Type: abitypes.Bytes32},
	}
}

type Parser struct {
	abi            *abi.ABI
	ps             *Oneinchv6Filterer
	eventHash      string
	traceCalls     *tracecall.Cache
	orderHashCount lru.BasicLRU[string, int]
}

func MustNewParser(cache *tracecall.Cache) *Parser {
	ps, err := NewOneinchv6Filterer(common.Address{}, nil)
	if err != nil {
		panic(err)
	}
	ab, err := Oneinchv6MetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	event, ok := ab.Events[FilledEvent]
	if !ok {
		panic(fmt.Sprintf("no such event: %s", FilledEvent))
	}

	return &Parser{
		ps:             ps,
		abi:            ab,
		eventHash:      event.ID.String(),
		traceCalls:     cache,
		orderHashCount: lru.NewBasicLRU[string, int](100),
	}
}

func (p *Parser) Topics() []string {
	return []string{
		p.eventHash,
	}
}

func (p *Parser) Parse(log ethereumTypes.Log, blockTime uint64) (storage.TradeLog, error) {
	order, err := p.buildOrderByLog(log)
	if err != nil {
		return storage.TradeLog{}, err
	}
	order.Timestamp = blockTime * 1000
	order, err = p.detectOneInchRfqTrade(order)
	if err != nil {
		return order, err
	}
	return order, nil
}

func (p *Parser) buildOrderByLog(log ethereumTypes.Log) (storage.TradeLog, error) {
	if len(log.Topics) > 0 && log.Topics[0].Hex() != p.eventHash {
		return storage.TradeLog{}, parser.ErrInvalidTopic
	}
	e, err := p.ps.ParseOrderFilled(log)
	if err != nil {
		return storage.TradeLog{}, fmt.Errorf("error when parse log %w", err)
	}
	order := storage.TradeLog{
		OrderHash:       common.Hash(e.OrderHash).String(),
		ContractAddress: e.Raw.Address.String(),
		BlockNumber:     e.Raw.BlockNumber,
		TxHash:          e.Raw.TxHash.String(),
		LogIndex:        uint64(e.Raw.Index),
		EventHash:       p.eventHash,
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
		return order, fmt.Errorf("error when decode output %w", err)
	}

	order.OrderHash = orderHash
	order.MakerTokenAmount = filledMakingAmount
	order.TakerTokenAmount = filledTakingAmount
	order.EventHash = p.eventHash

	contractCall, err := decoder.Decode(p.abi, internalCall.Input)
	if err != nil {
		return order, fmt.Errorf("error when decode input %w", err)
	}

	order, err = ToTradeLog(order, contractCall)
	if err != nil {
		return order, fmt.Errorf("error when parse contract call to order %w", err)
	}

	return order, nil
}

func (p *Parser) detectOneInchRfqTrade(order storage.TradeLog) (storage.TradeLog, error) {
	var (
		traceCall types.CallFrame
		err       error
	)

	traceCall, err = p.traceCalls.GetTraceCall(order.TxHash)
	if err != nil {
		return order, fmt.Errorf("error when get tracecall %w", err)
	}

	count := 0
	order, err = p.recursiveDetectOneInchRFQTrades(order, traceCall, &count)
	if err != nil {
		traceData, _ := json.Marshal(traceCall)
		return order, fmt.Errorf("error when parse tracecall %s %w", string(traceData), err)
	}

	return order, nil
}

func (p *Parser) recursiveDetectOneInchRFQTrades(tradeLog storage.TradeLog, traceCall types.CallFrame, count *int) (storage.TradeLog, error) {
	var (
		err error
	)
	if p.isOneInchRFQTrades(tradeLog.TxHash, tradeLog.OrderHash, traceCall, count) {
		return p.ParseFromInternalCall(tradeLog, traceCall)
	}

	for _, subCall := range traceCall.Calls {
		tradeLog, err = p.recursiveDetectOneInchRFQTrades(tradeLog, subCall, count)
		if err == nil {
			return tradeLog, nil
		}
	}

	return tradeLog, parser.ErrNotFoundTrade
}

func (p *Parser) isOneInchRFQTrades(txHash, orderHash string, traceCall types.CallFrame, count *int) bool {
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

		if orderHash != orderHashFromOutput {
			return false
		}
		last, ok := p.orderHashCount.Get(fmt.Sprintf("%s-%s", txHash, orderHash))
		if !ok {
			last = 0
		}
		if last != *count {
			*count += 1
			return false
		}
		p.orderHashCount.Add(fmt.Sprintf("%s-%s", txHash, orderHash), last+1)
		return true
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

func (p *Parser) Exchange() string {
	return parser.Ex1InchV6
}

func (p *Parser) UseTraceCall() bool {
	return true
}

func (p *Parser) ParseWithCallFrame(callFrame types.CallFrame, log ethereumTypes.Log, blockTime uint64) (storage.TradeLog, error) {
	order, err := p.buildOrderByLog(log)
	if err != nil {
		return storage.TradeLog{}, err
	}
	order.Timestamp = blockTime * 1000
	count := 0
	return p.recursiveDetectOneInchRFQTrades(order, callFrame, &count)
}

func (p *Parser) LogFromExchange(log ethereumTypes.Log) bool {
	return strings.EqualFold(log.Address.String(), parser.Addr1InchV6) &&
		len(log.Topics) > 0 &&
		strings.EqualFold(log.Topics[0].String(), p.eventHash)
}
