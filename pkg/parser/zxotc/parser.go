package zxotc

import (
	"errors"

	"github.com/KyberNetwork/tradelogs/pkg/parser"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	OtcOrderFilledEvent = "OtcOrderFilled"
)

var ErrInvalidOTCTopic = errors.New("invalid OTCFilled topic")

type Parser struct {
	abi       *abi.ABI
	ps        *OtcOrderFilledFilterer
	eventHash string
}

func MustNewParser() *Parser {
	ps, err := NewOtcOrderFilledFilterer(common.Address{}, nil)
	if err != nil {
		panic(err)
	}
	ab, err := OtcOrderFilledMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	event, ok := ab.Events[OtcOrderFilledEvent]
	if !ok {
		panic("no such event: OtcOrderFilled")
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

func (p *Parser) Parse(log types.Log, blockTime uint64) (storage.TradeLog, error) {
	if len(log.Topics) > 0 && log.Topics[0].Hex() != p.eventHash {
		return storage.TradeLog{}, ErrInvalidOTCTopic
	}
	o, err := p.ps.ParseOtcOrderFilled(log)
	if err != nil {
		return storage.TradeLog{}, err
	}
	res := storage.TradeLog{
		OrderHash:        common.Hash(o.OrderHash).String(),
		Maker:            o.Maker.Hex(),
		Taker:            o.Taker.Hex(),
		MakerToken:       o.MakerToken.Hex(),
		TakerToken:       o.TakerToken.Hex(),
		MakerTokenAmount: o.MakerTokenFilledAmount.String(),
		TakerTokenAmount: o.TakerTokenFilledAmount.String(),
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
	return parser.ExZeroX
}

func (p *Parser) UseTraceCall() bool {
	return false
}
