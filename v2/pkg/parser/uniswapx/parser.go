package uniswapx

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
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	FilledEvent = "Fill"
)

var BPS = big.NewInt(10000)

type (
	ResolvedOrder struct {
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
	InputOrder struct {
		Order []uint8 `json:"order"`
		Sig   []uint8 `json:"sig"`
	}
)

type Parser struct {
	abi            *abi.ABI
	ps             *UniswapxFilterer
	eventHash      string
	orderArguments abi.Arguments
}

func MustNewParser() *Parser {
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
		orderArguments: abi.Arguments{{Type: orderTuple}},
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

func (p *Parser) buildOrderByLog(log ethereumTypes.Log, blockTime uint64) (storageTypes.TradeLog, error) {
	if len(log.Topics) > 0 && log.Topics[0].Hex() != p.eventHash {
		return storageTypes.TradeLog{}, parser.ErrInvalidTopic
	}
	e, err := p.ps.ParseFill(log)
	if err != nil {
		return storageTypes.TradeLog{}, err
	}
	order := storageTypes.TradeLog{
		Exchange:        p.Exchange(),
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

func (p *Parser) updateOrder(tradeLog storageTypes.TradeLog, call types.CallFrame) (storageTypes.TradeLog, error) {
	input, err := decoder.Decode(p.abi, call.Input)
	if err != nil {
		return storageTypes.TradeLog{}, err
	}
	if len(input.Params) == 0 {
		return storageTypes.TradeLog{}, err
	}

	// determine index of current log's resolved order
	var internalIndex int
	for i, l := range call.Logs {
		if l.Topics[0].String() == p.eventHash && l.Topics[1].String() == tradeLog.OrderHash {
			internalIndex = i
		}
	}

	var inputOrder InputOrder
	switch t := input.Params[0].Value.(type) {
	case struct {
		Order []uint8 `json:"order"`
		Sig   []uint8 `json:"sig"`
	}:
		inputOrder = t
	case []struct {
		Order []uint8 `json:"order"`
		Sig   []uint8 `json:"sig"`
	}:
		if len(t) <= internalIndex {
			return storageTypes.TradeLog{}, errors.New("empty input order")
		}
		inputOrder = t[internalIndex]
	default:
		return storageTypes.TradeLog{}, errors.New("invalid input order")
	}

	parsedOrder, err := p.orderArguments.Unpack(inputOrder.Order)
	if err != nil {
		return storageTypes.TradeLog{}, err
	}

	data, err := json.Marshal(parsedOrder)
	if err != nil {
		return storageTypes.TradeLog{}, err
	}

	var resolvedOrder []ResolvedOrder
	if err = json.Unmarshal(data, &resolvedOrder); err != nil {
		return storageTypes.TradeLog{}, err
	}

	if len(resolvedOrder) == 0 || len(resolvedOrder[0].Outputs) == 0 {
		return storageTypes.TradeLog{}, errors.New("empty order")
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

	tradeLog.TakerToken = order.InputToken.InputToken.String()
	tradeLog.MakerToken = order.Outputs[0].Token.String()
	blockTime := big.NewInt(int64(tradeLog.Timestamp / 1000))
	tradeLog.TakerTokenAmount = decay(order.InputToken.InputStartAmount,
		order.InputToken.InputEndAmount,
		order.CosignerData.DecayStartTime,
		order.CosignerData.DecayEndTime,
		blockTime).String()

	makerAmount := big.NewInt(0)
	for _, o := range order.Outputs {
		output := decay(o.StartAmount, o.EndAmount,
			order.CosignerData.DecayStartTime, order.CosignerData.DecayEndTime,
			blockTime)
		if !hasFillingRights(order.CosignerData.ExclusiveFiller, common.HexToAddress(call.From), order.CosignerData.DecayStartTime, blockTime) {
			output = mulDivUp(output, new(big.Int).Add(BPS, order.CosignerData.ExclusivityOverrideBps), BPS)
		}
		makerAmount = makerAmount.Add(makerAmount, output)
	}
	tradeLog.MakerTokenAmount = makerAmount.String()
	if order.Info.Deadline != nil {
		tradeLog.Expiry = order.Info.Deadline.Uint64()
	}
	if order.CosignerData.DecayStartTime != nil {
		var decayStartTime uint64 = order.CosignerData.DecayStartTime.Uint64()
		tradeLog.DecayStartTime = &decayStartTime
	}
	if order.CosignerData.DecayEndTime != nil {
		var decayEndTime uint64 = order.CosignerData.DecayEndTime.Uint64()
		tradeLog.DecayEndTime = &decayEndTime
	}
	tradeLog.ExclusiveFiller = order.CosignerData.ExclusiveFiller.Hex()
	return tradeLog, nil
}

func decay(startAmount, endAmount, decayStartTime, decayEndTime, blockTime *big.Int) (decayedAmount *big.Int) {
	decayedAmount = new(big.Int)

	if startAmount.Cmp(endAmount) == 0 {
		return startAmount
	} else if decayEndTime.Cmp(blockTime) <= 0 {
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
			decayedAmount = new(big.Int).Add(startAmount, mulDivUp(diff, elapsed, duration))
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
	return constant.ExUniswapX
}

func (p *Parser) UseTraceCall() bool {
	return true
}

func (p *Parser) ParseWithCallFrame(callFrame types.CallFrame, log ethereumTypes.Log, blockTime uint64) ([]storageTypes.TradeLog, error) {
	order, err := p.buildOrderByLog(log, blockTime)
	if err != nil {
		return nil, fmt.Errorf("failed to build order: %w", err)
	}
	order, err = p.updateOrder(order, callFrame)
	if err != nil {
		return []storageTypes.TradeLog{order}, fmt.Errorf("failed to update order: %w", err)
	}
	return []storageTypes.TradeLog{order}, nil
}

func (p *Parser) LogFromExchange(log ethereumTypes.Log) bool {
	return strings.EqualFold(log.Address.String(), constant.AddrUniswapX) &&
		len(log.Topics) > 0 &&
		strings.EqualFold(log.Topics[0].String(), p.eventHash)
}

func (p *Parser) Address() string {
	return constant.AddrUniswapX
}
