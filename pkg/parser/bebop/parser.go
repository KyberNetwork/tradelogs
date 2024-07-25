package bebop

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/KyberNetwork/tradelogs/pkg/types"
	tradingTypes "github.com/KyberNetwork/tradinglib/pkg/types"
	"math/big"
	"strings"

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
)

var (
	ErrTradeTopic  = errors.New("invalid trade topic")
	ErrNotFoundLog = errors.New("not found log")
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
	Expiry         *big.Int        `json:"expiry"`
	TakerAddress   string          `json:"taker_address"`
	MakerAddress   string          `json:"maker_address"`
	MakerNonce     *big.Int        `json:"maker_nonce"`
	TakerTokens    json.RawMessage `json:"taker_tokens"`
	MakerTokens    json.RawMessage `json:"maker_tokens"`
	TakerAmounts   json.RawMessage `json:"taker_amounts"`
	MakerAmounts   json.RawMessage `json:"maker_amounts"`
	Receiver       string          `json:"receiver"`
	PackedCommands *big.Int        `json:"packed_commands"`
	Flags          *big.Int        `json:"flags"`
}

type AggregateOrder struct {
	Expiry         *big.Int        `json:"expiry"`
	TakerAddress   string          `json:"taker_address"`
	MakerAddresses json.RawMessage `json:"maker_addresses"`
	MakerNonces    json.RawMessage `json:"maker_nonces"`
	TakerTokens    json.RawMessage `json:"taker_tokens"`
	MakerTokens    json.RawMessage `json:"maker_tokens"`
	TakerAmounts   json.RawMessage `json:"taker_amounts"`
	MakerAmounts   json.RawMessage `json:"maker_amounts"`
	Receiver       string          `json:"receiver"`
	PackedCommands *big.Int        `json:"packed_commands"`
	Flags          *big.Int        `json:"flags"`
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
		singleOrderFunc:    sets.New("swapSingle", "swapSingleFromContract", "settleSingle", "settleSingleAndSign`Permit", "settleSingleAndSignPermit2"),
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
	return order, fmt.Errorf("%w %s", ErrNotFoundLog, string(traceData))
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

	for _, param := range contractCall.Params {
		if param.Name != OrderParam {
			continue
		}
		switch {
		case p.singleOrderFunc.Has(contractCall.Name):
			var rfqOrder SingleOrder
			if err := unpackOrder(param.Value, &rfqOrder); err != nil {
				return order, err
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
		case p.multiOrderFunc.Has(contractCall.Name):
			var rfqOrder MultiOrder
			if err := unpackOrder(param.Value, &rfqOrder); err != nil {
				return order, err
			}
			order.MakerToken = string(rfqOrder.MakerTokens)
			order.TakerToken = string(rfqOrder.TakerTokens)
			order.Maker = rfqOrder.MakerAddress
			order.Taker = rfqOrder.TakerAddress
			order.MakerTokenAmount = string(rfqOrder.MakerAmounts)
			order.TakerTokenAmount = string(rfqOrder.TakerAmounts)
			if rfqOrder.Expiry != nil {
				order.Expiry = rfqOrder.Expiry.Uint64()
			}
		case p.aggregateOrderFunc.Has(contractCall.Name):
			var rfqOrder AggregateOrder
			if err := unpackOrder(param.Value, &rfqOrder); err != nil {
				return order, err
			}
			order.MakerToken = string(rfqOrder.MakerTokens)
			order.TakerToken = string(rfqOrder.TakerTokens)
			order.Maker = string(rfqOrder.MakerAddresses)
			order.Taker = rfqOrder.TakerAddress
			order.MakerTokenAmount = string(rfqOrder.MakerAmounts)
			order.TakerTokenAmount = string(rfqOrder.TakerAmounts)
			if rfqOrder.Expiry != nil {
				order.Expiry = rfqOrder.Expiry.Uint64()
			}
		}
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
		return storage.TradeLog{}, ErrTradeTopic
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

func (p *Parser) ParseWithCallFrame(callFrame *tradingTypes.CallFrame, log ethereumTypes.Log, blockTime uint64) (storage.TradeLog, error) {
	if callFrame == nil {
		return storage.TradeLog{}, errors.New("missing call frame")
	}
	order, err := p.buildOrderByLog(log, blockTime)
	if err != nil {
		return storage.TradeLog{}, err
	}
	return p.searchTradeLog(order, types.ConvertCallFrame(callFrame))
}
