package optimex

import (
	"strings"
	"time"

	"github.com/KyberNetwork/tradelogs/pkg/parser"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	TradeEvent = "SelectPMM"
)

type Parser struct {
	abi       *abi.ABI
	filter    *OptimexFilterer
	caller    *OptimexCaller
	eventHash string
}

func MustNewParser(ethClient *ethclient.Client) *Parser {
	filter, err := NewOptimexFilterer(common.Address{}, nil)
	if err != nil {
		panic(err)
	}
	caller, err := NewOptimexCaller(common.HexToAddress(parser.AddrOptimex), ethClient)
	if err != nil {
		panic(err)
	}
	abi, err := OptimexMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	event, ok := abi.Events[TradeEvent]
	if !ok {
		panic("no such event: SelectPMM")
	}
	return &Parser{
		filter:    filter,
		caller:    caller,
		abi:       abi,
		eventHash: event.ID.String(),
	}
}

func (p *Parser) Topics() []string {
	return []string{
		p.eventHash,
	}
}

func (p *Parser) Parse(log ethereumTypes.Log, blockTime uint64) (storage.OptimexTradeLog, error) {
	if len(log.Topics) > 0 && log.Topics[0].Hex() != p.eventHash {
		return storage.OptimexTradeLog{}, parser.ErrInvalidTopic
	}
	pmm, err := p.filter.ParseSelectPMM(log)
	if err != nil {
		return storage.OptimexTradeLog{}, err
	}
	trade, err := p.caller.GetTradeData(nil, pmm.TradeId)
	if err != nil {
		return storage.OptimexTradeLog{}, err
	}

	selection, err := p.caller.GetPMMSelection(nil, pmm.TradeId)
	if err != nil {
		return storage.OptimexTradeLog{}, err
	}

	tradeID := hexutil.Encode(pmm.TradeId[:])
	fromChainID := string(trade.TradeInfo.FromChain[1])
	toChainID := string(trade.TradeInfo.ToChain[1])
	fromUser, fromTokenID := tradeInfoToData(fromChainID, trade.TradeInfo.FromChain)
	toUser, toTokenID := tradeInfoToData(toChainID, trade.TradeInfo.ToChain)

	res := storage.OptimexTradeLog{
		TradeID:          tradeID,
		Maker:            hexutil.Encode(selection.PmmInfo.SelectedPMMId[:]),
		FromTaker:        fromUser,
		ToTaker:          toUser,
		MakerToken:       fromTokenID,
		TakerToken:       toTokenID,
		MakerTokenAmount: selection.PmmInfo.AmountOut.String(),
		TakerTokenAmount: trade.TradeInfo.AmountIn.String(),
		ContractAddress:  pmm.Raw.Address.String(),
		BlockNumber:      pmm.Raw.BlockNumber,
		TxHash:           pmm.Raw.TxHash.String(),
		LogIndex:         uint64(pmm.Raw.Index),
		LogTime:          time.Unix(int64(blockTime), 0),
		EventHash:        p.eventHash,
		FromChain:        fromChainID,
		ToChain:          toChainID,
		TradeTimeout:     time.Unix(int64(selection.RfqInfo.TradeTimeout), 0),
	}
	return res, nil
}

func (p *Parser) Exchange() string {
	return parser.ExOptimex
}

func tradeInfoToData(chainID string, tradeInfo [3][]byte) (string, string) {
	var tokenID string
	if len(tradeInfo[2]) == common.AddressLength {
		tokenID = common.BytesToAddress(tradeInfo[2]).String()
	} else {
		tokenID = string(tradeInfo[2])
	}
	switch chainID {
	case "ethereum", "ethereum_sepolia":
		return common.Address(tradeInfo[0]).String(), tokenID
	}
	return string(tradeInfo[0]), tokenID
}

func (p *Parser) LogFromExchange(log ethereumTypes.Log) bool {
	return strings.EqualFold(parser.AddrOptimex, log.Address.String()) &&
		len(log.Topics) > 0 && strings.EqualFold(log.Topics[0].String(), p.eventHash)
}
