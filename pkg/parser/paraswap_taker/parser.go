package pstaker

import (
	"strings"

	"github.com/KyberNetwork/tradelogs/pkg/parser"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	tradingTypes "github.com/KyberNetwork/tradelogs/pkg/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	SwapEvent  = "SwappedV3"
	SwapDirect = "SwappedDirect"
)

type Parser struct {
	abi             *abi.ABI
	ps              *PsTakerFilterer
	eventHash       string
	directEventHash string
}

func MustNewParser() *Parser {
	ps, err := NewPsTakerFilterer(common.Address{}, nil)
	if err != nil {
		panic(err)
	}
	ab, err := PsTakerMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	event, ok := ab.Events[SwapEvent]
	if !ok {
		panic("no such event: SwappedV3")
	}
	directEvent, ok := ab.Events[SwapDirect]
	if !ok {
		panic("no such event: SwappedDirect")
	}
	return &Parser{
		ps:              ps,
		abi:             ab,
		eventHash:       event.ID.String(),
		directEventHash: directEvent.ID.String(),
	}
}

func (p *Parser) Topics() []string {
	return []string{
		p.eventHash,
		p.directEventHash,
	}
}

func (p *Parser) Parse(log types.Log, blockTime uint64) (storage.TradeLog, error) {
	if len(log.Topics) == 0 {
		return storage.TradeLog{}, parser.ErrInvalidTopic
	}

	if log.Topics[0].Hex() == p.eventHash {
		e, err := p.ps.ParseSwappedV3(log)
		if err != nil {
			return storage.TradeLog{}, err
		}
		res := storage.TradeLog{
			Taker:            e.Beneficiary.String(),
			MakerToken:       e.DestToken.String(),
			TakerToken:       e.SrcToken.String(),
			MakerTokenAmount: e.ReceivedAmount.String(),
			TakerTokenAmount: e.SrcAmount.String(),
			ContractAddress:  e.Raw.Address.String(),
			BlockNumber:      e.Raw.BlockNumber,
			TxHash:           e.Raw.TxHash.String(),
			LogIndex:         uint64(e.Raw.Index),
			Timestamp:        blockTime * 1000,
			EventHash:        p.eventHash,
		}
		return res, nil
	}
	if log.Topics[0].Hex() == p.directEventHash {
		e, err := p.ps.ParseSwappedDirect(log)
		if err != nil {
			return storage.TradeLog{}, err
		}
		res := storage.TradeLog{
			Taker:            e.Initiator.String(),
			MakerToken:       e.DestToken.String(),
			TakerToken:       e.SrcToken.String(),
			MakerTokenAmount: e.ReceivedAmount.String(),
			TakerTokenAmount: e.SrcAmount.String(),
			ContractAddress:  e.Raw.Address.String(),
			BlockNumber:      e.Raw.BlockNumber,
			TxHash:           e.Raw.TxHash.String(),
			LogIndex:         uint64(e.Raw.Index),
			Timestamp:        blockTime * 1000,
			EventHash:        p.directEventHash,
		}
		return res, nil
	}
	return storage.TradeLog{}, parser.ErrInvalidTopic
}

func (p *Parser) Exchange() string {
	return parser.ExParaswapTaker
}

func (p *Parser) UseTraceCall() bool {
	return false
}

func (p *Parser) ParseWithCallFrame(
	callFrame tradingTypes.CallFrame,
	log types.Log,
	blockTime uint64,
) (storage.TradeLog, error) {
	return p.Parse(log, blockTime)
}

func (p *Parser) LogFromExchange(log types.Log) bool {
	return strings.EqualFold(log.Address.String(), parser.AddrParaswapTaker) &&
		len(log.Topics) > 0 && (strings.EqualFold(log.Topics[0].String(), p.eventHash) || strings.EqualFold(log.Topics[0].String(), p.directEventHash))
}

func (p *Parser) Address() string {
	return parser.AddrParaswapTaker
}
