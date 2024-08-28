package kyberswap

import (
	"strings"

	ethereumTypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/KyberNetwork/tradelogs/pkg/parser"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"github.com/KyberNetwork/tradelogs/pkg/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

const (
	SwappedEvent = "Swapped"
)

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

func (p *Parser) Parse(log ethereumTypes.Log, blockTime uint64) (storage.TradeLog, error) {
	if len(log.Topics) > 0 && log.Topics[0].Hex() != p.eventHash {
		return storage.TradeLog{}, parser.ErrInvalidTopic
	}
	e, err := p.ps.ParseSwapped(log)
	if err != nil {
		return storage.TradeLog{}, err
	}
	res := storage.TradeLog{
		Taker:            e.DstReceiver.String(),
		MakerToken:       e.DstToken.String(),
		TakerToken:       e.SrcToken.String(),
		MakerTokenAmount: e.ReturnAmount.String(),
		TakerTokenAmount: e.SpentAmount.String(),
		ContractAddress:  e.Raw.Address.String(),
		BlockNumber:      e.Raw.BlockNumber,
		TxHash:           e.Raw.TxHash.String(),
		LogIndex:         uint64(e.Raw.Index),
		Timestamp:        blockTime * 1000,
		EventHash:        p.eventHash,
	}
	return res, nil
}

func (p *Parser) Exchange() string {
	return parser.ExKs
}

func (p *Parser) UseTraceCall() bool {
	return false
}

func (p *Parser) ParseWithCallFrame(_ types.CallFrame, log ethereumTypes.Log, blockTime uint64) (storage.TradeLog, error) {
	return p.Parse(log, blockTime)
}

func (p *Parser) LogFromExchange(log ethereumTypes.Log) bool {
	return strings.EqualFold(log.Address.String(), parser.AddrKyberswap) &&
		len(log.Topics) > 0 &&
		strings.EqualFold(log.Topics[0].String(), p.eventHash)
}
