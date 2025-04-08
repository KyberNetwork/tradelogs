package bebop

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/KyberNetwork/tradelogs/v2/pkg/constant"
	"github.com/KyberNetwork/tradelogs/v2/pkg/decoder"
	"github.com/KyberNetwork/tradelogs/v2/pkg/parser"
	storageTypes "github.com/KyberNetwork/tradelogs/v2/pkg/storage/tradelogs/types"
	"github.com/KyberNetwork/tradelogs/v2/pkg/types"
	tradingTypes "github.com/KyberNetwork/tradinglib/pkg/types"
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
	singleOrderFunc    sets.Set[string]
	multiOrderFunc     sets.Set[string]
	aggregateOrderFunc sets.Set[string]
}

func MustNewParser() *Parser {
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

func (p *Parser) Parse(log ethereumTypes.Log, blockTime uint64) ([]storageTypes.TradeLog, error) {
	return nil, fmt.Errorf("you need to use the ParseWithCallFrame function")
}

func (p *Parser) ParseFromInternalCall(order storageTypes.TradeLog, internalCall types.CallFrame) ([]storageTypes.TradeLog, error) {
	contractCall, err := decoder.Decode(p.abi, internalCall.Input)
	if err != nil {
		return []storageTypes.TradeLog{order}, err
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
		return []storageTypes.TradeLog{order}, ErrParamNotFound
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

	return []storageTypes.TradeLog{order}, nil
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
	return constant.ExBebop
}

func (p *Parser) UseTraceCall() bool {
	return true
}

func (p *Parser) buildOrderByLog(log ethereumTypes.Log, blockTime uint64) (storageTypes.TradeLog, error) {
	if len(log.Topics) > 0 && log.Topics[0].Hex() != p.eventHash {
		return storageTypes.TradeLog{}, parser.ErrInvalidTopic
	}
	o, err := p.ps.ParseBebopOrder(log)
	if err != nil {
		return storageTypes.TradeLog{}, err
	}
	order := storageTypes.TradeLog{
		Exchange:        p.Exchange(),
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

func (p *Parser) ParseWithCallFrame(callFrame types.CallFrame, log ethereumTypes.Log, blockTime uint64) ([]storageTypes.TradeLog, error) {
	order, err := p.buildOrderByLog(log, blockTime)
	if err != nil {
		return nil, fmt.Errorf("fail to build order by log: %w", err)
	}
	return p.ParseFromInternalCall(order, callFrame)
}

func (p *Parser) LogFromExchange(log ethereumTypes.Log) bool {
	return strings.EqualFold(log.Address.String(), constant.AddrBebop) &&
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

func (p *Parser) parseSingleSwap(order storageTypes.TradeLog,
	contractCall *tradingTypes.ContractCall,
	orderParam tradingTypes.ContractCallParam,
	fillTakerAmount *big.Int,
	internalCall types.CallFrame) ([]storageTypes.TradeLog, error) {
	var rfqOrder SingleOrder
	if err := unpackOrder(orderParam.Value, &rfqOrder); err != nil {
		return []storageTypes.TradeLog{order}, err
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
			return []storageTypes.TradeLog{order}, ErrParamNotFound
		}
	}

	order.MakerTokenOriginAmount = rfqOrder.MakerAmount.String()
	order.TakerTokenOriginAmount = rfqOrder.TakerAmount.String()

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

	return []storageTypes.TradeLog{order}, nil
}

func (p *Parser) parseMultiSwap(order storageTypes.TradeLog,
	contractCall *tradingTypes.ContractCall,
	orderParam tradingTypes.ContractCallParam,
	fillTakerAmount *big.Int) ([]storageTypes.TradeLog, error) {
	var rfqOrder MultiOrder
	if err := unpackOrder(orderParam.Value, &rfqOrder); err != nil {
		return []storageTypes.TradeLog{order}, err
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

	order.MakerTokenOriginAmount = joinAmounts(rfqOrder.MakerAmounts)
	order.TakerTokenOriginAmount = joinAmounts(rfqOrder.TakerAmounts)

	if len(rfqOrder.TakerTokens) == 1 { // many to one don't support partial, just handle one - many
		if fillTakerAmount.Cmp(big.NewInt(0)) > 0 && fillTakerAmount.Cmp(rfqOrder.TakerAmounts[0]) < 0 {
			for j := range rfqOrder.MakerAmounts {
				tmp := big.NewInt(0).Mul(rfqOrder.MakerAmounts[j], fillTakerAmount)
				rfqOrder.MakerAmounts[j] = tmp.Div(tmp, rfqOrder.TakerAmounts[0])
				rfqOrder.TakerAmounts[0] = fillTakerAmount
			}
		}
	}

	order.MakerToken = strings.Join(rfqOrder.MakerTokens, ",")
	order.TakerToken = strings.Join(rfqOrder.TakerTokens, ",")
	order.Maker = rfqOrder.MakerAddress
	order.Taker = rfqOrder.TakerAddress

	order.MakerTokenAmount = joinAmounts(rfqOrder.MakerAmounts)
	order.TakerTokenAmount = joinAmounts(rfqOrder.TakerAmounts)

	if rfqOrder.Expiry != nil {
		order.Expiry = rfqOrder.Expiry.Uint64()
	}

	return []storageTypes.TradeLog{order}, nil
}

func (p *Parser) parseAggregateSwap(order storageTypes.TradeLog,
	contractCall *tradingTypes.ContractCall,
	orderParam tradingTypes.ContractCallParam,
	filledTakerAmount *big.Int) ([]storageTypes.TradeLog, error) {

	var rfqOrder AggregateOrder
	if err := unpackOrder(orderParam.Value, &rfqOrder); err != nil {
		return []storageTypes.TradeLog{order}, err
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

	var orders []storageTypes.TradeLog

	// parse each trade
	for i := range rfqOrder.MakerAddresses {
		newOrder := storageTypes.TradeLog{
			Exchange:               p.Exchange(),
			OrderHash:              order.OrderHash,
			Maker:                  rfqOrder.MakerAddresses[i],
			Taker:                  rfqOrder.TakerAddress,
			MakerToken:             strings.Join(rfqOrder.MakerTokens[i], ","),
			TakerToken:             strings.Join(rfqOrder.TakerTokens[i], ","),
			MakerTokenAmount:       joinAmounts(rfqOrder.MakerAmounts[i]),
			TakerTokenAmount:       joinAmounts(rfqOrder.TakerAmounts[i]),
			MakerTokenOriginAmount: joinAmounts(rfqOrder.MakerAmounts[i]),
			TakerTokenOriginAmount: joinAmounts(rfqOrder.TakerAmounts[i]),
			ContractAddress:        order.ContractAddress,
			BlockNumber:            order.BlockNumber,
			TxHash:                 order.TxHash,
			LogIndex:               order.LogIndex,
			TradeIndex:             uint64(i),
			Timestamp:              order.Timestamp,
			EventHash:              order.EventHash,
			Expiry:                 rfqOrder.Expiry.Uint64(),
		}

		if len(rfqOrder.TakerTokens[i]) > 1 { // many-to-one not support partial fill
			orders = append(orders, newOrder)
			continue
		}

		// partial fill
		if filledTakerAmount.Cmp(big.NewInt(0)) > 0 && filledTakerAmount.Cmp(quoteTakerAmount) < 0 {
			for j := range rfqOrder.MakerAmounts[i] {
				tmp := big.NewInt(0).Mul(rfqOrder.MakerAmounts[i][j], filledTakerAmount)
				rfqOrder.MakerAmounts[i][j] = tmp.Div(tmp, quoteTakerAmount)
			}
			for j := range rfqOrder.TakerAmounts[i] {
				tmp := big.NewInt(0).Mul(rfqOrder.TakerAmounts[i][j], filledTakerAmount)
				rfqOrder.TakerAmounts[i][j] = tmp.Div(tmp, quoteTakerAmount)
			}
		}

		newOrder.MakerTokenAmount = joinAmounts(rfqOrder.MakerAmounts[i])
		newOrder.TakerTokenAmount = joinAmounts(rfqOrder.TakerAmounts[i])

		orders = append(orders, newOrder)
	}

	return orders, nil
}

func joinAmounts(input []*big.Int) string {
	amounts := make([]string, len(input))
	for i, amount := range input {
		amounts[i] = amount.String()
	}
	return strings.Join(amounts, ",")
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
	}
	return quoteTakerAmount
}

func (p *Parser) Address() string {
	return constant.AddrBebop
}

func (p *Parser) ParseTransferEvent(
	txHash string,
	block_number, timestamp uint64,
	call types.CallFrame,
) []interface{} {
	return nil
}
