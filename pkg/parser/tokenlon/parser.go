package tokenlon

import (
	"errors"

	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	FillOrderEvent = "FillOrder"
)

var ErrInvalidRFQTopic = errors.New("invalid RfqFilled topic")

type Parser struct {
	abi       *abi.ABI
	ps        *FillOrderFilterer
	eventHash string
}

func MustNewParser() *Parser {
	ps, err := NewFillOrderFilterer(common.Address{}, nil)
	if err != nil {
		panic(err)
	}
	ab, err := FillOrderMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	event, ok := ab.Events[FillOrderEvent]
	if !ok {
		panic("no such event: FillOrder")
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
		return storage.TradeLog{}, ErrInvalidRFQTopic
	}
	o, err := p.ps.ParseFillOrder(log)
	if err != nil {
		return storage.TradeLog{}, err
	}
	res := storage.TradeLog{
		OrderHash:        common.Hash(o.OrderHash).String(),
		Maker:            o.MakerAddr.Hex(),
		Taker:            o.ReceiverAddr.Hex(),
		MakerToken:       o.MakerAssetAddr.Hex(),
		TakerToken:       o.TakerAssetAddr.Hex(),
		MakerTokenAmount: o.MakerAssetAmount.String(),
		TakerTokenAmount: o.TakerAssetAmount.String(),
		ContractAddress:  o.Raw.Address.String(),
		BlockNumber:      o.Raw.BlockNumber,
		TxHash:           o.Raw.TxHash.String(),
		LogIndex:         uint64(o.Raw.Index),
		Timestamp:        blockTime * 1000,
		EventHash:        p.eventHash,
	}
	return res, nil
}
