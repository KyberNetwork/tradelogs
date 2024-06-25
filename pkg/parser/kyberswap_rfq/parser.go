package kyberswaprfq

import (
	"errors"
	"fmt"

	"github.com/KyberNetwork/tradelogs/pkg/parser"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"github.com/KyberNetwork/tradelogs/pkg/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	FilledEvent = "OrderFilledRFQ"
)

var ErrInvalidKSFilledTopic = errors.New("invalid KS order filled topic")

type Parser struct {
	abi       *abi.ABI
	ps        *KyberswaprfqFilterer
	eventHash string
}

func MustNewParser() *Parser {
	ps, err := NewKyberswaprfqFilterer(common.Address{}, nil)
	if err != nil {
		panic(err)
	}
	ab, err := KyberswaprfqMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	event, ok := ab.Events[FilledEvent]
	if !ok {
		panic(fmt.Sprintf("no such event: %s", FilledEvent))
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
		return storage.TradeLog{}, ErrInvalidKSFilledTopic
	}
	e, err := p.ps.ParseOrderFilledRFQ(log)
	if err != nil {
		return storage.TradeLog{}, err
	}
	res := storage.TradeLog{
		OrderHash:        common.Hash(e.OrderHash).String(),
		Maker:            e.Maker.String(),
		Taker:            e.Taker.String(),
		MakerTokenAmount: e.MakingAmount.String(),
		TakerTokenAmount: e.TakingAmount.String(),
		ContractAddress:  e.Raw.Address.String(),
		BlockNumber:      e.Raw.BlockNumber,
		TxHash:           e.Raw.TxHash.String(),
		LogIndex:         uint64(e.Raw.Index),
		Timestamp:        blockTime * 1000,
		EventHash:        p.eventHash,
		MakerToken:       e.MakerAsset.String(),
		TakerToken:       e.TakerAsset.String(),
	}
	return res, nil
}

func (p *Parser) Exchange() string {
	return parser.ExKsRFQ
}

func (p *Parser) UseTraceCall() bool {
	return false
}

func (p *Parser) ParseWithCallFrame(_ types.CallFrame, log ethereumTypes.Log) (storage.TradeLog, error) {
	return p.Parse(log, 0)
}
