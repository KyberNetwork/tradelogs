package bebop

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/KyberNetwork/tradelogs/pkg/types"
	tradingTypes "github.com/KyberNetwork/tradinglib/pkg/types"

	"github.com/KyberNetwork/tradelogs/pkg/decoder"
	"github.com/KyberNetwork/tradelogs/pkg/parser"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"github.com/KyberNetwork/tradelogs/pkg/tracecall"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
	"k8s.io/apimachinery/pkg/util/sets"
)

const (
	TradeEvent = "BebopOrder"
	OrderParam = "order"

	balanceOfMethodID                  = "0x70a08231"
	swapSingleFromContractFunctionName = "swapSingleFromContract"
)

var (
	ErrParamNotFound = errors.New("param not found")
)

type SingleOrder struct {
	Expiry         *big.Int `json:"expiry"`
	TakerAddress   string   `json:"taker_address"`
	MakerAddress   string   `json:"maker_address"`
	MakerNonce     *big.Int `json:"maker_nonce"`
	TakerToken     string   `json:"taker_token"`
	MakerToken     string   `json:"maker_token"`
	TakerAmount    *big.Int `json:"taker_amount"`
	MakerAmount    *big.Int `json:"maker_amount"`
	Receiver       string   `json:"receiver"`
	PackedCommands *big.Int `json:"packed_commands"`
	Flags          *big.Int `json:"flags"`
}

type MultiOrder struct {
	Expiry       *big.Int   `json:"expiry"`
	TakerAddress string     `json:"taker_address"`
	MakerAddress string     `json:"maker_address"`
	MakerNonce   *big.Int   `json:"maker_nonce"`
	TakerTokens  []string   `json:"taker_tokens"`
	MakerTokens  []string   `json:"maker_tokens"`
	TakerAmounts []*big.Int `json:"taker_amounts"`
	MakerAmounts []*big.Int `json:"maker_amounts"`
	Receiver     string     `json:"receiver"`
	Commands     []byte     `json:"commands"`
	Flags        *big.Int   `json:"flags"`
}

type AggregateOrder struct {
	Expiry         *big.Int     `json:"expiry"`
	TakerAddress   string       `json:"taker_address"`
	MakerAddresses []string     `json:"maker_addresses"`
	MakerNonces    []*big.Int   `json:"maker_nonces"`
	TakerTokens    [][]string   `json:"taker_tokens"`
	MakerTokens    [][]string   `json:"maker_tokens"`
	TakerAmounts   [][]*big.Int `json:"taker_amounts"`
	MakerAmounts   [][]*big.Int `json:"maker_amounts"`
	Receiver       string       `json:"receiver"`
	Commands       []byte       `json:"commands"`
	Flags          *big.Int     `json:"flags"`
}

type OldSingleQuote struct {
	UseOldAmount bool     `json:"useOldAmount"`
	MakerAmount  *big.Int `json:"makerAmount"`
	MakerNonce   *big.Int `json:"makerNonce"`
}

type OldMultiQuote struct {
	UseOldAmount bool       `json:"useOldAmount"`
	MakerAmounts []*big.Int `json:"makerAmounts"`
	MakerNonce   *big.Int   `json:"makerNonce"`
}

type OldAggregateQuote struct {
	UseOldAmount bool         `json:"useOldAmount"`
	MakerAmounts [][]*big.Int `json:"makerAmounts"`
	MakerNonces  []*big.Int   `json:"makerNonces"`
}

type Parser struct {
	abi                *abi.ABI
	ps                 *BebopFilterer
	eventHash          string
	traceCalls         *tracecall.Cache
	singleOrderFunc    sets.Set[string]
	multiOrderFunc     sets.Set[string]
	aggregateOrderFunc sets.Set[string]
}

func MustNewParser(tracecall *tracecall.Cache) *Parser {
	ps, err := NewBebopFilterer(common.Address{}, nil)
	if err != nil {
		panic(err)
	}
	ab, err := abi.JSON(strings.NewReader(BebopMetaData.ABI))
	if err != nil {
		panic(err)
	}
	event, ok := ab.Events[TradeEvent]
	if !ok {
		panic("no such event: " + TradeEvent)
	}
	return &Parser{
		ps:                 ps,
		abi:                &ab,
		eventHash:          event.ID.String(),
		traceCalls:         tracecall,
		singleOrderFunc:    sets.New("swapSingle", "swapSingleFromContract", "settleSingle", "settleSingleAndSignPermit", "settleSingleAndSignPermit2"),
		multiOrderFunc:     sets.New("swapMulti", "settleMulti", "settleMultiAndSignPermit", "settleMultiAndSignPermit2"),
		aggregateOrderFunc: sets.New("swapAggregate", "settleAggregate", "settleAggregateAndSignPermit", "settleAggregateAndSignPermit2"),
	}
}

func (p *Parser) Topics() []string {
	return []string{
		p.eventHash,
	}
}

func (p *Parser) Parse(log ethereumTypes.Log, blockTime uint64) (storage.TradeLog, error) {
	order, err := p.buildOrderByLog(log, blockTime)
	if err != nil {
		return storage.TradeLog{}, err
	}
	return p.parseTraceCall(order)
}

func (p *Parser) parseTraceCall(order storage.TradeLog) (storage.TradeLog, error) {
	traceCall, err := p.traceCalls.GetTraceCall(order.TxHash)
	if err != nil {
		return order, err
	}

	order, err = p.searchTradeLog(order, traceCall)
	if err != nil {
		return order, err
	}

	return order, nil
}

func (p *Parser) searchTradeLog(order storage.TradeLog, traceCall types.CallFrame) (storage.TradeLog, error) {
	if p.checkBebopTrade(traceCall, order.OrderHash) {
		return p.ParseFromInternalCall(order, traceCall)
	}

	for _, subCall := range traceCall.Calls {
		tradeLog, err := p.searchTradeLog(order, subCall)
		if err == nil {
			return tradeLog, nil
		}
	}
	traceData, _ := json.Marshal(traceCall)
	return order, fmt.Errorf("%w %s", parser.ErrNotFoundTrade, string(traceData))
}

func (p *Parser) checkBebopTrade(traceCall types.CallFrame, orderHash string) bool {
	for _, eventLog := range traceCall.Logs {
		if len(eventLog.Topics) < 2 {
			continue
		}
		if !strings.EqualFold(eventLog.Topics[0].String(), p.eventHash) {
			continue
		}
		x, ok := big.NewInt(0).SetString(orderHash, 10)
		if !ok {
			continue
		}
		y, ok := big.NewInt(0).SetString(eventLog.Topics[1].String()[2:], 16)
		if !ok {
			continue
		}
		if x.Cmp(y) != 0 {
			continue
		}
		return true
	}
	return false
}

func (p *Parser) ParseFromInternalCall(order storage.TradeLog, internalCall types.CallFrame) (storage.TradeLog, error) {
	contractCall, err := decoder.Decode(p.abi, internalCall.Input)
	if err != nil {
		return order, err
	}
	var filledTakerAmount *big.Int
	for _, param := range contractCall.Params {
		if param.Name != "filledTakerAmount" {
			continue
		}
		v, ok := param.Value.(*big.Int)
		if ok {
			filledTakerAmount = v
		}
		break
	}
	if filledTakerAmount == nil && contractCall.Name != swapSingleFromContractFunctionName {
		return order, ErrParamNotFound
	}
	for _, param := range contractCall.Params {
		if param.Name != OrderParam {
			continue
		}
		switch {
		case p.singleOrderFunc.Has(contractCall.Name):
			return p.parseSingleSwap(order, contractCall, param, filledTakerAmount, internalCall)
		case p.multiOrderFunc.Has(contractCall.Name):
			return p.parseMultiSwap(order, contractCall, param, filledTakerAmount)
		case p.aggregateOrderFunc.Has(contractCall.Name):
			return p.parseAggregateSwap(order, contractCall, param, filledTakerAmount)
		}
		break
	}

	return order, nil
}

func unpackOrder(in, out interface{}) error {
	bytes, err := json.Marshal(in)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bytes, out); err != nil {
		return err
	}
	return nil
}

func (p *Parser) Exchange() string {
	return parser.ExBebop
}

func (p *Parser) UseTraceCall() bool {
	return true
}

func (p *Parser) buildOrderByLog(log ethereumTypes.Log, blockTime uint64) (storage.TradeLog, error) {
	if len(log.Topics) > 0 && log.Topics[0].Hex() != p.eventHash {
		return storage.TradeLog{}, parser.ErrInvalidTopic
	}
	o, err := p.ps.ParseBebopOrder(log)
	if err != nil {
		return storage.TradeLog{}, err
	}
	order := storage.TradeLog{
		OrderHash:       o.EventId.String(),
		Maker:           log.Address.Hex(),
		ContractAddress: o.Raw.Address.String(),
		BlockNumber:     o.Raw.BlockNumber,
		TxHash:          o.Raw.TxHash.String(),
		LogIndex:        uint64(o.Raw.Index),
		EventHash:       p.eventHash,
		Timestamp:       blockTime * 1000,
	}
	return order, nil
}

func (p *Parser) ParseWithCallFrame(callFrame types.CallFrame, log ethereumTypes.Log, blockTime uint64) (storage.TradeLog, error) {
	order, err := p.buildOrderByLog(log, blockTime)
	if err != nil {
		return storage.TradeLog{}, err
	}
	return p.searchTradeLog(order, callFrame)
}

func (p *Parser) LogFromExchange(log ethereumTypes.Log) bool {
	return strings.EqualFold(log.Address.String(), parser.AddrBebop) &&
		len(log.Topics) > 0 &&
		strings.EqualFold(log.Topics[0].String(), p.eventHash)
}

func (p *Parser) getFilledTakerAmount(order SingleOrder, traceCall types.CallFrame) *big.Int {
	for _, call := range traceCall.Calls {
		if call.From != strings.ToLower(p.Address()) {
			continue
		}
		if call.To != order.TakerToken {
			continue
		}
		if !strings.HasPrefix(call.Input, balanceOfMethodID) {
			continue
		}
		if len(call.Output) <= 2 {
			continue
		}
		balance, ok := new(big.Int).SetString(call.Output[2:], 16)
		if !ok {
			continue
		}
		return balance
	}
	return nil
}

func (p *Parser) parseSingleSwap(order storage.TradeLog,
	contractCall *tradingTypes.ContractCall,
	orderParam tradingTypes.ContractCallParam,
	fillTakerAmount *big.Int,
	internalCall types.CallFrame) (storage.TradeLog, error) {
	var rfqOrder SingleOrder
	if err := unpackOrder(orderParam.Value, &rfqOrder); err != nil {
		return order, err
	}

	for _, param := range contractCall.Params {
		if param.Name != "takerQuoteInfo" {
			continue
		}
		var oldOrder OldSingleQuote
		if err := unpackOrder(param.Value, &oldOrder); err == nil && oldOrder.UseOldAmount {
			rfqOrder.MakerAmount = oldOrder.MakerAmount
		}
		break
	}

	if contractCall.Name == swapSingleFromContractFunctionName {
		fillTakerAmount = p.getFilledTakerAmount(rfqOrder, internalCall)
		if fillTakerAmount == nil {
			return order, ErrParamNotFound
		}
	}

	if fillTakerAmount.Cmp(big.NewInt(0)) > 0 && fillTakerAmount.Cmp(rfqOrder.TakerAmount) < 0 {
		tmp := big.NewInt(0).Mul(rfqOrder.MakerAmount, fillTakerAmount)
		rfqOrder.MakerAmount = tmp.Div(tmp, rfqOrder.TakerAmount)
		rfqOrder.TakerAmount = fillTakerAmount
	}

	order.MakerToken = rfqOrder.MakerToken
	order.TakerToken = rfqOrder.TakerToken
	order.Maker = rfqOrder.MakerAddress
	order.Taker = rfqOrder.TakerAddress
	order.MakerTokenAmount = rfqOrder.MakerAmount.String()
	order.TakerTokenAmount = rfqOrder.TakerAmount.String()
	if rfqOrder.Expiry != nil {
		order.Expiry = rfqOrder.Expiry.Uint64()
	}

	return order, nil
}

func (p *Parser) parseMultiSwap(order storage.TradeLog,
	contractCall *tradingTypes.ContractCall,
	orderParam tradingTypes.ContractCallParam,
	fillTakerAmount *big.Int) (storage.TradeLog, error) {
	var rfqOrder MultiOrder
	if err := unpackOrder(orderParam.Value, &rfqOrder); err != nil {
		return order, err
	}

	for _, param := range contractCall.Params {
		if param.Name != "takerQuoteInfo" {
			continue
		}
		var oldOrder OldMultiQuote
		if err := unpackOrder(param.Value, &oldOrder); err == nil && oldOrder.UseOldAmount {
			rfqOrder.MakerAmounts = oldOrder.MakerAmounts
		}
		break
	}

	if len(rfqOrder.TakerTokens) == 1 { // many to one don't support partial, just handle one - many
		if fillTakerAmount.Cmp(big.NewInt(0)) > 0 && fillTakerAmount.Cmp(rfqOrder.TakerAmounts[0]) < 0 {
			for j := range rfqOrder.MakerAmounts {
				tmp := big.NewInt(0).Mul(rfqOrder.MakerAmounts[j], fillTakerAmount)
				rfqOrder.MakerAmounts[j] = tmp.Div(tmp, rfqOrder.TakerAmounts[0])
				rfqOrder.TakerAmounts[0] = fillTakerAmount
			}
		}
	}

	makerTokens, _ := json.Marshal(rfqOrder.MakerTokens)
	order.MakerToken = string(makerTokens)
	takerTokens, _ := json.Marshal(rfqOrder.TakerTokens)
	order.TakerToken = string(takerTokens)
	order.Maker = rfqOrder.MakerAddress
	order.Taker = rfqOrder.TakerAddress
	makerAmounts, _ := json.Marshal(rfqOrder.MakerAmounts)
	order.MakerTokenAmount = string(makerAmounts)
	takerAmounts, _ := json.Marshal(rfqOrder.TakerAmounts)
	order.TakerTokenAmount = string(takerAmounts)
	if rfqOrder.Expiry != nil {
		order.Expiry = rfqOrder.Expiry.Uint64()
	}

	return order, nil
}

func (p *Parser) parseAggregateSwap(order storage.TradeLog,
	contractCall *tradingTypes.ContractCall,
	orderParam tradingTypes.ContractCallParam,
	filledTakerAmount *big.Int) (storage.TradeLog, error) {

	var rfqOrder AggregateOrder
	if err := unpackOrder(orderParam.Value, &rfqOrder); err != nil {
		return order, err
	}

	for _, param := range contractCall.Params {
		if param.Name != "takerQuoteInfo" {
			continue
		}
		var oldOrder OldAggregateQuote
		if err := unpackOrder(param.Value, &oldOrder); err == nil && oldOrder.UseOldAmount {
			rfqOrder.MakerAmounts = oldOrder.MakerAmounts
		}
		break
	}
	quoteTakerAmount := getAggregateOrderInfo(rfqOrder)
	if filledTakerAmount.Cmp(big.NewInt(0)) > 0 && filledTakerAmount.Cmp(quoteTakerAmount) < 0 {
		for i := range rfqOrder.MakerAddresses {
			for j := range rfqOrder.MakerAmounts[i] {
				tmp := big.NewInt(0).Mul(rfqOrder.MakerAmounts[i][j], filledTakerAmount)
				rfqOrder.MakerAmounts[i][j] = tmp.Div(tmp, quoteTakerAmount)
			}
			for j := range rfqOrder.TakerAmounts[i] {
				tmp := big.NewInt(0).Mul(rfqOrder.TakerAmounts[i][j], filledTakerAmount)
				rfqOrder.TakerAmounts[i][j] = tmp.Div(tmp, quoteTakerAmount)
			}
		}
	}

	makerTokens, _ := json.Marshal(rfqOrder.MakerTokens)
	order.MakerToken = string(makerTokens)
	takerTokens, _ := json.Marshal(rfqOrder.TakerTokens)
	order.TakerToken = string(takerTokens)
	makerAddress, _ := json.Marshal(rfqOrder.MakerAddresses)
	order.Maker = string(makerAddress)
	order.Taker = rfqOrder.TakerAddress
	makerAmounts, _ := json.Marshal(rfqOrder.MakerAmounts)
	order.MakerTokenAmount = string(makerAmounts)
	takerAmounts, _ := json.Marshal(rfqOrder.TakerAmounts)
	order.TakerTokenAmount = string(takerAmounts)
	if rfqOrder.Expiry != nil {
		order.Expiry = rfqOrder.Expiry.Uint64()
	}

	return order, nil
}

func getAggregateOrderInfo(order AggregateOrder) *big.Int {
	commandsInd := 0
	quoteTakerAmount := big.NewInt(0)
	for i := range order.TakerTokens {
		commandsInd += len(order.MakerTokens[i])
		for j := range order.TakerTokens[i] {
			curCommand := order.Commands[commandsInd+j]
			if curCommand != 0x08 {
				quoteTakerAmount = quoteTakerAmount.Add(quoteTakerAmount, order.TakerAmounts[i][j])
			}
		}
		commandsInd += len(order.TakerTokens[i])
	}
	return quoteTakerAmount
}

func (p *Parser) Address() string {
	return parser.AddrBebop
}
