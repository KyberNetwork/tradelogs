package kyberswap

import (
	"errors"

	"github.com/KyberNetwork/tradelogs/internal/storage"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	SwappedEvent = "Swapped"
)

var ErrInvalidKSSwappedTopic = errors.New("invalid KS Swapped topic")

type Parser struct {
	abi       *abi.ABI
	ps        *SwappedFilterer
	eventHash string
}

func MustNewParser() *Parser {
	ps, err := NewSwappedFilterer(common.Address{}, nil)
	if err != nil {
		panic(err)
	}
	ab, err := SwappedMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	swapEvent, ok := ab.Events[SwappedEvent]
	if !ok {
		panic("no such event: Swapped")
	}
	return &Parser{
		ps:        ps,
		abi:       ab,
		eventHash: swapEvent.ID.String(),
	}
}

func (p *Parser) Topics() []string {
	return []string{
		p.eventHash,
	}
}

func (p *Parser) Parse(log types.Log, blockTime uint64) (storage.TradeLog, error) {
	if len(log.Topics) > 0 && log.Topics[0].Hex() != p.eventHash {
		return storage.TradeLog{}, ErrInvalidKSSwappedTopic
	}
	e, err := p.ps.ParseSwapped(log)
	if err != nil {
		return storage.TradeLog{}, err
	}
	res := storage.TradeLog{
		Taker:            e.DstReceiver.String(),
		MakerToken:       e.SrcToken.String(),
		TakerToken:       e.DstToken.String(),
		MakerTokenAmount: e.SpentAmount.String(),
		TakerTokenAmount: e.ReturnAmount.String(),
		ContractAddress:  e.Raw.Address.String(),
		BlockNumber:      e.Raw.BlockNumber,
		TxHash:           e.Raw.TxHash.String(),
		LogIndex:         uint64(e.Raw.Index),
		Timestamp:        blockTime * 1000,
	}
	return res, nil
}
