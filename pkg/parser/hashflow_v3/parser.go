package hashflowv3

import (
	"errors"

	"github.com/KyberNetwork/tradelogs/pkg/parser"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	TradeEvent = "Trade"
)

var ErrTradeTopic = errors.New("invalid Trade topic")

type Parser struct {
	abi       *abi.ABI
	ps        *Hashflowv3Filterer
	eventHash string
}

func MustNewParser() *Parser {
	ps, err := NewHashflowv3Filterer(common.Address{}, nil)
	if err != nil {
		panic(err)
	}
	ab, err := Hashflowv3MetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	event, ok := ab.Events[TradeEvent]
	if !ok {
		panic("no such event: Trade")
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
		return storage.TradeLog{}, ErrTradeTopic
	}
	o, err := p.ps.ParseTrade(log)
	if err != nil {
		return storage.TradeLog{}, err
	}
	res := storage.TradeLog{
		OrderHash:        common.Hash(o.Txid).String(),
		Maker:            log.Address.Hex(),
		Taker:            o.Trader.Hex(),
		MakerToken:       o.QuoteToken.Hex(),
		TakerToken:       o.BaseToken.Hex(),
		MakerTokenAmount: o.QuoteTokenAmount.String(),
		TakerTokenAmount: o.BaseTokenAmount.String(),
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
	return parser.ExHashflowV3
}
