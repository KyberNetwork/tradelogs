package kyberswap

import (
	"strings"

	"github.com/KyberNetwork/tradelogs/v2/pkg/constant"
	"github.com/KyberNetwork/tradelogs/v2/pkg/parser"
	storageTypes "github.com/KyberNetwork/tradelogs/v2/pkg/storage/tradelogs/types"
	"github.com/KyberNetwork/tradelogs/v2/pkg/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
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

func (p *Parser) Parse(log ethereumTypes.Log, blockTime uint64) ([]storageTypes.TradeLog, error) {
	if len(log.Topics) > 0 && log.Topics[0].Hex() != p.eventHash {
		return nil, parser.ErrInvalidTopic
	}
	e, err := p.ps.ParseSwapped(log)
	if err != nil {
		return nil, err
	}
	res := storageTypes.TradeLog{
		Taker:            e.Sender.String(),
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
	return []storageTypes.TradeLog{res}, nil
}

func (p *Parser) Exchange() string {
	return constant.ExKs
}

func (p *Parser) UseTraceCall() bool {
	return false
}

func (p *Parser) ParseWithCallFrame(_ types.CallFrame, log ethereumTypes.Log, blockTime uint64) ([]storageTypes.TradeLog, error) {
	return p.Parse(log, blockTime)
}

func (p *Parser) LogFromExchange(log ethereumTypes.Log) bool {
	return strings.EqualFold(log.Address.String(), constant.AddrKyberswap) &&
		len(log.Topics) > 0 &&
		strings.EqualFold(log.Topics[0].String(), p.eventHash)
}

func (p *Parser) Address() string {
	return constant.AddrKyberswap
}
