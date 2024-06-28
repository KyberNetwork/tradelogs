package uniswapx

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/KyberNetwork/tradelogs/pkg/types"
	tradingTypes "github.com/KyberNetwork/tradinglib/pkg/types"
	"math/big"

	"github.com/KyberNetwork/tradelogs/pkg/decoder"
	"github.com/KyberNetwork/tradelogs/pkg/parser"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"github.com/KyberNetwork/tradelogs/pkg/tracecall"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	FilledEvent   = "Fill"
	TransferEvent = "Transfer"
)

type ResolvedOrder struct {
	Info struct {
		Reactor                      common.Address "json:\"reactor\""
		Swapper                      common.Address "json:\"swapper\""
		Nonce                        *big.Int       "json:\"nonce\""
		Deadline                     *big.Int       "json:\"deadline\""
		AdditionalValidationContract common.Address "json:\"additionalValidationContract\""
		AdditionalValidationData     []uint8        "json:\"additionalValidationData\""
	} "json:\"info\""
	Cosigner   common.Address "json:\"cosigner\""
	InputToken struct {
		InputToken       common.Address "json:\"inputToken\""
		InputStartAmount *big.Int       "json:\"inputStartAmount\""
		InputEndAmount   *big.Int       "json:\"inputEndAmount\""
	} "json:\"inputToken\""
	Outputs []struct {
		Token       common.Address "json:\"token\""
		StartAmount *big.Int       "json:\"startAmount\""
		EndAmount   *big.Int       "json:\"endAmount\""
		Recipient   common.Address "json:\"recipient\""
	} "json:\"outputs\""
	CosignerData struct {
		DecayStartTime         *big.Int       "json:\"decayStartTime\""
		DecayEndTime           *big.Int       "json:\"decayEndTime\""
		ExclusiveFiller        common.Address "json:\"exclusiveFiller\""
		ExclusivityOverrideBps *big.Int       "json:\"exclusivityOverrideBps\""
		InputOverride          *big.Int       "json:\"inputOverride\""
		OutputOverrides        []struct {
			Override *big.Int "json:\"override\""
		} "json:\"outputOverrides\""
	}
	Cosignature []uint8 "json:\"cosignature\""
}

var (
	ErrInvalidOneInchFilledTopic = errors.New("invalid uniswapx order filled topic")
	ErrNotFoundLog               = errors.New("not found log")
)

type Parser struct {
	abi            *abi.ABI
	ps             *UniswapxFilterer
	eventHash      string
	traceCalls     *tracecall.Cache
	orderArguments abi.Arguments
}

func MustNewParser(cache *tracecall.Cache) *Parser {
	ps, err := NewUniswapxFilterer(common.Address{}, nil)
	if err != nil {
		panic(err)
	}
	ab, err := UniswapxMetaData.GetAbi()
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
		{Name: "cosigner", Type: "address"},
		{Name: "inputToken", Type: "tuple", Components: []abi.ArgumentMarshaling{
			{Name: "inputToken", Type: "address"},
			{Name: "inputStartAmount", Type: "uint256"},
			{Name: "inputEndAmount", Type: "uint256"}},
		},
		{Name: "outputs", Type: "tuple[]", Components: []abi.ArgumentMarshaling{
			{Name: "token", Type: "address"},
			{Name: "startAmount", Type: "uint256"},
			{Name: "endAmount", Type: "uint256"},
			{Name: "recipient", Type: "address"},
		}},
		{Name: "cosignerData", Type: "tuple", Components: []abi.ArgumentMarshaling{
			{Name: "decayStartTime", Type: "uint256"},
			{Name: "decayEndTime", Type: "uint256"},
			{Name: "exclusiveFiller", Type: "address"},
			{Name: "exclusivityOverrideBps", Type: "uint256"},
			{Name: "inputOverride", Type: "uint256"},
			{Name: "outputOverrides", Type: "tuple[]", Components: []abi.ArgumentMarshaling{
				{Name: "override", Type: "uint256"}}}},
		},
		{Name: "cosignature", Type: "bytes"},
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
		return storage.TradeLog{}, ErrInvalidOneInchFilledTopic
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
	for _, l := range call.Logs {
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
			inputOrder = inputOrders[0]
		}

		parsedOrder, err := p.orderArguments.Unpack(inputOrder.Order)
		if err != nil {
			continue
		}
		finalOrder, err := p.updateOrder(order, parsedOrder)
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
	return order, fmt.Errorf("%w %s", ErrNotFoundLog, string(traceData))
}

func (p *Parser) updateOrder(internal storage.TradeLog, parsed []interface{}) (storage.TradeLog, error) {
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
	if order.CosignerData.InputOverride.Cmp(big.NewInt(0)) != 0 {
		order.InputToken.InputStartAmount = order.CosignerData.InputOverride
	}
	for i := range order.Outputs {
		if order.CosignerData.OutputOverrides[i].Override.Cmp(big.NewInt(0)) != 0 {
			order.Outputs[i].StartAmount = order.CosignerData.OutputOverrides[i].Override
		}
	}

	internal.TakerToken = order.InputToken.InputToken.String()
	internal.MakerToken = order.Outputs[0].Token.String()

	internal.TakerTokenAmount = decay(order.InputToken.InputStartAmount,
		order.InputToken.InputEndAmount,
		order.CosignerData.DecayStartTime,
		order.CosignerData.DecayEndTime,
		big.NewInt(int64(internal.Timestamp/1000))).String()

	makerAmount := big.NewInt(0)
	for _, o := range order.Outputs {
		makerAmount = makerAmount.Add(makerAmount, decay(o.StartAmount, o.EndAmount,
			order.CosignerData.DecayStartTime, order.CosignerData.DecayEndTime,
			big.NewInt(int64(internal.Timestamp/1000))))
	}
	internal.MakerTokenAmount = makerAmount.String()
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

// func handleOverride(amount, exclusivityOverrideBps, BPS *big.Int) *big.Int {
// 	return mulDiv(amount, new(big.Int).Add(exclusivityOverrideBps, BPS), BPS)
// }

func mulDiv(a, b, c *big.Int) *big.Int {
	return new(big.Int).Div(new(big.Int).Mul(a, b), c)
}

// func checkExclusivity(exclusive, sender common.Address, exclusivityEndTime, ts uint64) bool {
// 	return exclusive == common.Address{} || ts > exclusivityEndTime || exclusive == sender
// }

func (p *Parser) Exchange() string {
	return parser.ExUniswapX
}

func (p *Parser) UseTraceCall() bool {
	return true
}

func (p *Parser) ParseWithCallFrame(callFrame *tradingTypes.CallFrame, log ethereumTypes.Log, blockTime uint64) (storage.TradeLog, error) {
	if callFrame == nil {
		return storage.TradeLog{}, errors.New("missing call frame")
	}
	order, err := p.buildOrderByLog(log, blockTime)
	if err != nil {
		return storage.TradeLog{}, err
	}
	order, err = p.recursiveDetectRFQTrades(order, types.ConvertCallFrame(callFrame))
	if err != nil {
		return order, err
	}
	return order, nil
}
