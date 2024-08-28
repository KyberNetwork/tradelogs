package native

import (
	"errors"
	"github.com/KyberNetwork/tradelogs/pkg/parser"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"github.com/KyberNetwork/tradelogs/pkg/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	SwapEvent = "Swap"
)

var ErrTradeTopic = errors.New("invalid Trade topic")

type Parser struct {
	abi       *abi.ABI
	ps        *NativeFilterer
	eventHash string
}

func MustNewParser() *Parser {
	ps, err := NewNativeFilterer(common.Address{}, nil)
	if err != nil {
		panic(err)
	}
	ab, err := NativeMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	event, ok := ab.Events[SwapEvent]
	if !ok {
		panic("no such event: Swap")
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
		return storage.TradeLog{}, ErrTradeTopic
	}
	o, err := p.ps.ParseSwap(log)
	if err != nil {
		return storage.TradeLog{}, err
	}
	res := storage.TradeLog{
		OrderHash:        common.BytesToHash(o.QuoteId[:]).String(),
		Maker:            o.Raw.Address.String(),
		Taker:            o.Recipient.String(),
		MakerToken:       o.TokenOut.String(),
		TakerToken:       o.TokenIn.String(),
		MakerTokenAmount: o.AmountOut.Abs(o.AmountOut).String(),
		TakerTokenAmount: o.AmountIn.String(),
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
	return parser.ExNative
}

func (p *Parser) UseTraceCall() bool {
	return false
}

func (p *Parser) ParseWithCallFrame(_ types.CallFrame, log ethereumTypes.Log, blockTime uint64) (storage.TradeLog, error) {
	return p.Parse(log, blockTime)
}
