package oneinchv2

import (
	"strings"

	"github.com/KyberNetwork/tradelogs/v2/pkg/constant"
	"github.com/KyberNetwork/tradelogs/v2/pkg/promotionparser"
	storageTypes "github.com/KyberNetwork/tradelogs/v2/pkg/storage/promotees"
	"github.com/KyberNetwork/tradelogs/v2/pkg/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
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

func (p *Parser) Parse(log ethereumTypes.Log, blockTime uint64) (storageTypes.Promotee, error) {
	if len(log.Topics) > 0 && log.Topics[0].Hex() != p.eventHash {
		return storageTypes.Promotee{}, promotionparser.ErrInvalidTopic
	}
	e, err := p.ps.ParsePromotion(log)
	if err != nil {
		return storageTypes.Promotee{}, err
	}
	res := storageTypes.Promotee{
		Promoter:  strings.ToLower(e.Promoter.String()),
		ChainId:   e.ChainId.String(),
		Promotee:  strings.ToLower(e.Promotee.String()),
		Timestamp: blockTime * 1000,
		TxHash:    e.Raw.TxHash.String(),
	}
	return res, nil
}

func (p *Parser) Contract() string {
	return constant.Promotion1InchV2
}

func (p *Parser) UseTraceCall() bool {
	return false
}

func (p *Parser) ParseWithCallFrame(_ types.CallFrame, log ethereumTypes.Log, blockTime uint64) (storageTypes.Promotee, error) {
	return p.Parse(log, blockTime)
}

func (p *Parser) LogFromContract(log ethereumTypes.Log) bool {
	return strings.EqualFold(log.Address.String(), constant.AddrPr1InchV2) &&
		len(log.Topics) > 0 &&
		strings.EqualFold(log.Topics[0].String(), p.eventHash)
}

func (p *Parser) Address() string {
	return constant.AddrPr1InchV2
}
