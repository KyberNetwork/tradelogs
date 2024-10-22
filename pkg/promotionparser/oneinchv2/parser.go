package oneinchv2

import (
	"strings"

	ethereumTypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/KyberNetwork/tradelogs/pkg/promotionparser"
	"github.com/KyberNetwork/tradelogs/pkg/promotionstorage"
	"github.com/KyberNetwork/tradelogs/pkg/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

const (
	PromotionEvent = "Promotion"
)

type Parser struct {
	abi       *abi.ABI
	ps        *PromotionFilterer
	eventHash string
}

func MustNewParser() *Parser {
	ps, err := NewPromotionFilterer(common.Address{}, nil)
	if err != nil {
		panic(err)
	}
	ab, err := PromotionMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	promotionEvent, ok := ab.Events[PromotionEvent]
	if !ok {
		panic("no such event: Promotion")
	}
	return &Parser{
		ps:        ps,
		abi:       ab,
		eventHash: promotionEvent.ID.String(),
	}
}

func (p *Parser) Topics() []string {
	return []string{
		p.eventHash,
	}
}

func (p *Parser) Parse(log ethereumTypes.Log, blockTime uint64) (promotionstorage.Promotee, error) {
	if len(log.Topics) > 0 && log.Topics[0].Hex() != p.eventHash {
		return promotionstorage.Promotee{}, promotionparser.ErrInvalidTopic
	}
	e, err := p.ps.ParsePromotion(log)
	if err != nil {
		return promotionstorage.Promotee{}, err
	}
	res := promotionstorage.Promotee{
		Promoter:  strings.ToLower(e.Promoter.String()),
		ChainId:   e.ChainId.String(),
		Promotee:  strings.ToLower(e.Promotee.String()),
		Timestamp: blockTime * 1000,
		EventHash: p.eventHash,
	}
	return res, nil
}

func (p *Parser) Contract() string {
	return promotionparser.Promotion1InchV2
}

func (p *Parser) UseTraceCall() bool {
	return false
}

func (p *Parser) ParseWithCallFrame(_ types.CallFrame, log ethereumTypes.Log, blockTime uint64) (promotionstorage.Promotee, error) {
	return p.Parse(log, blockTime)
}

func (p *Parser) LogFromContract(log ethereumTypes.Log) bool {
	return strings.EqualFold(log.Address.String(), promotionparser.AddrPr1InchV2) &&
		len(log.Topics) > 0 &&
		strings.EqualFold(log.Topics[0].String(), p.eventHash)
}

func (p *Parser) Address() string {
	return promotionparser.AddrPr1InchV2
}
