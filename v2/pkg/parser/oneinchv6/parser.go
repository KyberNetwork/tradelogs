package oneinchv6

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/KyberNetwork/tradelogs/v2/pkg/abitypes"
	"github.com/KyberNetwork/tradelogs/v2/pkg/constant"
	"github.com/KyberNetwork/tradelogs/v2/pkg/decoder"
	"github.com/KyberNetwork/tradelogs/v2/pkg/parser"
	storageTypes "github.com/KyberNetwork/tradelogs/v2/pkg/storage/tradelogs/types"
	"github.com/KyberNetwork/tradelogs/v2/pkg/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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
	abi        *abi.ABI
	ps         *Oneinchv6Filterer
	eventHash  string
	markFusion *MarkFusion
}

func MustNewParser(markFusion *MarkFusion) *Parser {
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
		ps:         ps,
		abi:        ab,
		eventHash:  event.ID.String(),
		markFusion: markFusion,
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

func (p *Parser) buildOrderByLog(log ethereumTypes.Log) (storageTypes.TradeLog, error) {
	if len(log.Topics) > 0 && log.Topics[0].Hex() != p.eventHash {
		return storageTypes.TradeLog{}, parser.ErrInvalidTopic
	}
	e, err := p.ps.ParseOrderFilled(log)
	if err != nil {
		return storageTypes.TradeLog{}, fmt.Errorf("error when parse log %w", err)
	}
	order := storageTypes.TradeLog{
		Exchange:        p.Exchange(),
		OrderHash:       common.Hash(e.OrderHash).String(),
		ContractAddress: e.Raw.Address.String(),
		BlockNumber:     e.Raw.BlockNumber,
		TxHash:          e.Raw.TxHash.String(),
		LogIndex:        uint64(e.Raw.Index),
		EventHash:       p.eventHash,
	}
	return order, nil
}

func (p *Parser) ParseFromInternalCall(order storageTypes.TradeLog, internalCall types.CallFrame) (storageTypes.TradeLog, error) {
	output := internalCall.Output

	if output == "" {
		return order, errors.New("the output is blank")
	}
	filledMakingAmount, filledTakingAmount, orderHash, err := p.decodeOutput(output)
	if err != nil {
		return order, fmt.Errorf("error when decode output %w", err)
	}

	order.OrderHash = orderHash
	order.MakerTokenAmount = filledMakingAmount.String()
	order.TakerTokenAmount = filledTakingAmount.String()
	order.EventHash = p.eventHash

	contractCall, err := decoder.Decode(p.abi, internalCall.Input)
	if err != nil {
		return order, fmt.Errorf("error when decode input %w", err)
	}
	order, err = ToTradeLog(order, contractCall)
	if err != nil {
		return order, fmt.Errorf("error when parse contract call to order %w", err)
	}
	order.Taker = internalCall.From

	if order.Type == storageTypes.RFQType {
		return order, nil
	}

	fusion := p.markFusion.CheckPromoteeExist(strings.ToLower(order.Taker))
	if fusion {
		order.Type = storageTypes.FusionType
	}
	return order, nil
}

func (p *Parser) decodeOutput(output string) (*big.Int, *big.Int, string, error) {
	bytes, err := hexutil.Decode(output)
	if err != nil {
		return nil, nil, "", err
	}
	outputParams, err := RFQOrderOutputArgument.Unpack(bytes)
	if err != nil {
		return nil, nil, "", err
	}

	if len(outputParams) < 3 {
		return nil, nil, "", err
	}

	filledMakingAmountFromOutput, ok := outputParams[0].(*big.Int)
	if !ok {
		return nil, nil, "", errors.New("cant convert the filledMakingAmount from the output")
	}

	filledTakingAmountFromOutput, ok := outputParams[1].(*big.Int)
	if !ok {
		return nil, nil, "", errors.New("cant convert the filledTakingAmount from the output")
	}
	orderHashParams, ok := outputParams[2].([32]byte)
	if !ok {
		return nil, nil, "", errors.New("cant convert the order hash from the output")
	}

	orderHash := hexutil.Encode(orderHashParams[:])
	return filledMakingAmountFromOutput, filledTakingAmountFromOutput, orderHash, nil
}

func (p *Parser) Exchange() string {
	return constant.Ex1InchV6
}

func (p *Parser) UseTraceCall() bool {
	return true
}

func (p *Parser) ParseWithCallFrame(callFrame types.CallFrame, log ethereumTypes.Log, blockTime uint64) ([]storageTypes.TradeLog, error) {
	order, err := p.buildOrderByLog(log)
	if err != nil {
		return []storageTypes.TradeLog{}, err
	}
	order.Timestamp = blockTime * 1000
	tradeLog, err := p.ParseFromInternalCall(order, callFrame)
	if err != nil {
		return nil, err
	}
	return []storageTypes.TradeLog{tradeLog}, nil
}

func (p *Parser) LogFromExchange(log ethereumTypes.Log) bool {
	return strings.EqualFold(log.Address.String(), constant.Addr1InchV6) &&
		len(log.Topics) > 0 &&
		strings.EqualFold(log.Topics[0].String(), p.eventHash)
}

func (p *Parser) Address() string {
	return constant.Addr1InchV6
}
