package pancakeswap

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/KyberNetwork/tradelogs/pkg/decoder"
	"github.com/KyberNetwork/tradelogs/pkg/parser"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"github.com/KyberNetwork/tradelogs/pkg/tracecall"
	"github.com/KyberNetwork/tradelogs/pkg/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	FilledEvent = "Fill"
)

var BPS = big.NewInt(10000)

type ResolvedOrder struct {
	Info struct {
		Reactor                      common.Address "json:\"reactor\""
		Swapper                      common.Address "json:\"swapper\""
		Nonce                        *big.Int       "json:\"nonce\""
		Deadline                     *big.Int       "json:\"deadline\""
		AdditionalValidationContract common.Address "json:\"additionalValidationContract\""
		AdditionalValidationData     []uint8        "json:\"additionalValidationData\""
	} "json:\"info\""
	DecayStartTime         *big.Int       "json:\"decayStartTime\""
	DecayEndTime           *big.Int       "json:\"decayEndTime\""
	ExclusiveFiller        common.Address "json:\"exclusiveFiller\""
	ExclusivityOverrideBps *big.Int       "json:\"exclusivityOverrideBps\""
	InputToken             common.Address "json:\"inputToken\""
	InputStartAmount       *big.Int       "json:\"inputStartAmount\""
	InputEndAmount         *big.Int       "json:\"inputEndAmount\""
	Outputs                []struct {
		Token       common.Address "json:\"token\""
		StartAmount *big.Int       "json:\"startAmount\""
		EndAmount   *big.Int       "json:\"endAmount\""
		Recipient   common.Address "json:\"recipient\""
	} "json:\"outputs\""
}

type Parser struct {
	abi            *abi.ABI
	ps             *PancakeswapFilterer
	eventHash      string
	traceCalls     *tracecall.Cache
	orderArguments abi.Arguments
}

func MustNewParser(cache *tracecall.Cache) *Parser {
	ps, err := NewPancakeswapFilterer(common.Address{}, nil)
	if err != nil {
		panic(err)
	}
	ab, err := PancakeswapMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	event, ok := ab.Events[FilledEvent]
	if !ok {
		panic(fmt.Sprintf("no such event: %s", FilledEvent))
	}

	orderTuple, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{Name: "info", Type: "tuple", Components: []abi.ArgumentMarshaling{
			{Name: "reactor", Type: "address"},
			{Name: "swapper", Type: "address"},
			{Name: "nonce", Type: "uint256"},
			{Name: "deadline", Type: "uint256"},
			{Name: "additionalValidationContract", Type: "address"},
			{Name: "additionalValidationData", Type: "bytes"}},
		},
		{Name: "decayStartTime", Type: "uint256"},
		{Name: "decayEndTime", Type: "uint256"},
		{Name: "exclusiveFiller", Type: "address"},
		{Name: "exclusivityOverrideBps", Type: "uint256"},
		{Name: "inputToken", Type: "address"},
		{Name: "inputStartAmount", Type: "uint256"},
		{Name: "inputEndAmount", Type: "uint256"},
		{Name: "outputs", Type: "tuple[]", Components: []abi.ArgumentMarshaling{
			{Name: "token", Type: "address"},
			{Name: "startAmount", Type: "uint256"},
			{Name: "endAmount", Type: "uint256"},
			{Name: "recipient", Type: "address"},
		}},
	})
	if err != nil {
		panic("cant create order abi type")
	}

	return &Parser{
		ps:             ps,
		abi:            ab,
		eventHash:      event.ID.String(),
		traceCalls:     cache,
		orderArguments: abi.Arguments{{Type: orderTuple}},
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

	order, err = p.detectRfqTrade(order)
	if err != nil {
		return order, err
	}
	return order, nil
}

func (p *Parser) buildOrderByLog(log ethereumTypes.Log, blockTime uint64) (storage.TradeLog, error) {
	if len(log.Topics) > 0 && log.Topics[0].Hex() != p.eventHash {
		return storage.TradeLog{}, parser.ErrInvalidTopic
	}
	e, err := p.ps.ParseFill(log)
	if err != nil {
		return storage.TradeLog{}, err
	}
	order := storage.TradeLog{
		OrderHash:       common.Hash(e.OrderHash).String(),
		ContractAddress: e.Raw.Address.String(),
		BlockNumber:     e.Raw.BlockNumber,
		TxHash:          e.Raw.TxHash.String(),
		LogIndex:        uint64(e.Raw.Index),
		Timestamp:       blockTime * 1000,
		EventHash:       p.eventHash,
		Maker:           e.Filler.String(),
		Taker:           e.Swapper.String(),
	}
	return order, nil
}

func (p *Parser) detectRfqTrade(order storage.TradeLog) (storage.TradeLog, error) {
	traceCall, err := p.traceCalls.GetTraceCall(order.TxHash)
	if err != nil {
		return order, err
	}

	order, err = p.recursiveDetectRFQTrades(order, traceCall)
	if err != nil {
		return order, err
	}
	return order, nil
}

func (p *Parser) recursiveDetectRFQTrades(order storage.TradeLog, call types.CallFrame) (storage.TradeLog, error) {
	for i, l := range call.Logs {
		if len(l.Topics) < 2 {
			continue
		}
		if l.Topics[0].String() != p.eventHash || l.Topics[1].String() != order.OrderHash {
			continue
		}
		input, err := decoder.Decode(p.abi, call.Input)
		if err != nil {
			continue
		}
		if len(input.Params) == 0 {
			continue
		}
		inputOrder, ok := input.Params[0].Value.(struct {
			Order []uint8 `json:"order"`
			Sig   []uint8 `json:"sig"`
		})
		if !ok {
			inputOrders, ok := input.Params[0].Value.([]struct {
				Order []uint8 `json:"order"`
				Sig   []uint8 `json:"sig"`
			})
			if !ok || len(inputOrders) == 0 {
				continue
			}
			inputOrder = inputOrders[i]
		}
		parsedOrder, err := p.orderArguments.Unpack(inputOrder.Order)
		if err != nil {
			continue
		}
		finalOrder, err := p.updateOrder(order, parsedOrder, common.HexToAddress(call.From))
		if err != nil {
			continue
		}
		return finalOrder, nil
	}

	for i := range call.Calls {
		order, err := p.recursiveDetectRFQTrades(order, call.Calls[i])
		if err == nil {
			return order, nil
		}
	}
	traceData, _ := json.Marshal(call)
	return order, fmt.Errorf("%w %s", parser.ErrNotFoundTrade, string(traceData))
}

func (p *Parser) updateOrder(internal storage.TradeLog, parsed []interface{}, sender common.Address) (storage.TradeLog, error) {
	data, err := json.Marshal(parsed)
	if err != nil {
		return storage.TradeLog{}, err
	}
	var resolvedOrder []ResolvedOrder
	if err = json.Unmarshal(data, &resolvedOrder); err != nil {
		return storage.TradeLog{}, err
	}
	if len(resolvedOrder) == 0 || len(resolvedOrder[0].Outputs) == 0 {
		return storage.TradeLog{}, err
	}
	order := resolvedOrder[0]
	internal.TakerToken = order.InputToken.String()
	blockTime := big.NewInt(int64(internal.Timestamp / 1000))
	internal.TakerTokenAmount = decay(order.InputStartAmount,
		order.InputEndAmount,
		order.DecayStartTime,
		order.DecayEndTime,
		blockTime).String()
	internal.MakerToken = order.Outputs[0].Token.String()

	makerAmount := big.NewInt(0)
	for _, o := range order.Outputs {
		output := decay(o.StartAmount, o.EndAmount, order.DecayStartTime, order.DecayEndTime, blockTime)
		if !hasFillingRights(order.ExclusiveFiller, sender, order.DecayStartTime, blockTime) {
			output = mulDivUp(output, new(big.Int).Add(BPS, order.ExclusivityOverrideBps), BPS)
		}
		makerAmount = makerAmount.Add(makerAmount, output)
	}
	internal.MakerTokenAmount = makerAmount.String()
	if order.Info.Deadline != nil {
		internal.Expiry = order.Info.Deadline.Uint64()
	}
	return internal, nil
}

func decay(startAmount, endAmount, decayStartTime, decayEndTime, blockTime *big.Int) (decayedAmount *big.Int) {
	decayedAmount = new(big.Int)

	if decayEndTime.Cmp(blockTime) <= 0 {
		decayedAmount.Set(endAmount)
	} else if decayStartTime.Cmp(blockTime) >= 0 {
		decayedAmount.Set(startAmount)
	} else {
		elapsed := new(big.Int).Sub(blockTime, decayStartTime)
		duration := new(big.Int).Sub(decayEndTime, decayStartTime)

		if endAmount.Cmp(startAmount) < 0 {
			diff := new(big.Int).Sub(startAmount, endAmount)
			decayedAmount = new(big.Int).Sub(startAmount, mulDiv(diff, elapsed, duration))
		} else {
			diff := new(big.Int).Sub(endAmount, startAmount)
			decayedAmount = new(big.Int).Add(startAmount, mulDiv(diff, elapsed, duration))
		}
	}
	return
}

func mulDiv(a, b, c *big.Int) *big.Int {
	return new(big.Int).Div(new(big.Int).Mul(a, b), c)
}

func mulDivUp(x, y, denominator *big.Int) *big.Int {
	product := new(big.Int).Mul(x, y)
	remainder := new(big.Int).Mod(product, denominator)
	result := new(big.Int).Div(product, denominator)
	if remainder.Cmp(big.NewInt(0)) > 0 {
		result.Add(result, big.NewInt(1))
	}
	return result
}

func hasFillingRights(exclusive, sender common.Address, exclusivityEndTime, blockTime *big.Int) bool {
	return exclusive == common.Address{} || blockTime.Cmp(exclusivityEndTime) > 0 || exclusive == sender
}

func (p *Parser) Exchange() string {
	return parser.ExPancackeSwap
}

func (p *Parser) UseTraceCall() bool {
	return true
}

func (p *Parser) ParseWithCallFrame(callFrame types.CallFrame, log ethereumTypes.Log, blockTime uint64) (storage.TradeLog, error) {
	order, err := p.buildOrderByLog(log, blockTime)
	if err != nil {
		return storage.TradeLog{}, err
	}
	order, err = p.recursiveDetectRFQTrades(order, callFrame)
	if err != nil {
		return order, err
	}
	return order, nil
}

func (p *Parser) LogFromExchange(log ethereumTypes.Log) bool {
	return strings.EqualFold(log.Address.String(), parser.AddrPancakewap) &&
		len(log.Topics) > 0 &&
		strings.EqualFold(log.Topics[0].String(), p.eventHash)
}

func (p *Parser) Address() string {
	return parser.AddrPancakewap
}
