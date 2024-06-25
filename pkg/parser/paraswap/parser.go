package paraswap

import (
	"errors"
	"github.com/KyberNetwork/tradelogs/pkg/types"

	"github.com/KyberNetwork/tradelogs/pkg/parser"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	OrderFilledEvent = "OrderFilled"
)

var ErrInvalidRFQTopic = errors.New("invalid RfqFilled topic")

type Parser struct {
	abi       *abi.ABI
	ps        *OrderFilledFilterer
	eventHash string
}

func MustNewParser() *Parser {
	ps, err := NewOrderFilledFilterer(common.Address{}, nil)
	if err != nil {
		panic(err)
	}
	ab, err := OrderFilledMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	event, ok := ab.Events[OrderFilledEvent]
	if !ok {
		panic("no such event: OrderFilled")
	}
	return &Parser{
		ps:        ps,
		abi:       ab,
		eventHash: event.ID.String(),
	}
}

func (p *Parser) Topics() []string {
	return []string{
		p.eventHash,
	}
}

func (p *Parser) Parse(log ethereumTypes.Log, blockTime uint64) (storage.TradeLog, error) {
	if len(log.Topics) > 0 && log.Topics[0].Hex() != p.eventHash {
		return storage.TradeLog{}, ErrInvalidRFQTopic
	}
	o, err := p.ps.ParseOrderFilled(log)
	if err != nil {
		return storage.TradeLog{}, err
	}
	res := storage.TradeLog{
		OrderHash:        common.Hash(o.OrderHash).String(),
		Maker:            o.Maker.Hex(),
		Taker:            o.Taker.Hex(),
		MakerToken:       o.MakerAsset.Hex(),
		TakerToken:       o.TakerAsset.Hex(),
		MakerTokenAmount: o.MakerAmount.String(),
		TakerTokenAmount: o.TakerAmount.String(),
		ContractAddress:  o.Raw.Address.String(),
		BlockNumber:      o.Raw.BlockNumber,
		TxHash:           o.Raw.TxHash.String(),
		LogIndex:         uint64(o.Raw.Index),
		Timestamp:        blockTime * 1000,
		EventHash:        p.eventHash,
	}
	return res, nil
}

func (p *Parser) Exchange() string {
	return parser.ExParaswap
}

func (p *Parser) UseTraceCall() bool {
	return false
}

func (p *Parser) ParseWithCallFrame(_ types.CallFrame, log ethereumTypes.Log) (storage.TradeLog, error) {
	return p.Parse(log, 0)
}
