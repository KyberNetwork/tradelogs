package optimex

import (
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
	SelectEvent = "SelectPMM"
	SettleEvent = "ConfirmSettlement"
)

type Parser struct {
	abi    *abi.ABI
	filter *OptimexFilterer
	caller *OptimexCaller
	events map[common.Hash]string
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
	tradeEvent, ok := abi.Events[SelectEvent]
	if !ok {
		panic("no such event: SelectPMM")
	}
	settleEvent, ok := abi.Events[SettleEvent]
	if !ok {
		panic("no such event: ConfirmSettlement")
	}
	return &Parser{
		filter: filter,
		caller: caller,
		abi:    abi,
		events: map[common.Hash]string{
			tradeEvent.ID:  SelectEvent,
			settleEvent.ID: SettleEvent,
		},
	}
}

func (p *Parser) ParseSelected(log ethereumTypes.Log, blockTime uint64) (storage.OptimexTradeLog, error) {
	if len(log.Topics) > 0 && p.events[log.Topics[0]] != SelectEvent {
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
	fromChainID, fromUser, fromTokenID := tradeInfoToData(trade.TradeInfo.FromChain)
	toChainID, toUser, toTokenID := tradeInfoToData(trade.TradeInfo.ToChain)

	res := storage.OptimexTradeLog{
		TradeID:          tradeID,
		Maker:            hexutil.Encode(selection.PmmInfo.SelectedPMMId[:]),
		FromTaker:        fromUser,
		ToTaker:          toUser,
		MakerToken:       toTokenID,
		TakerToken:       fromTokenID,
		MakerTokenAmount: selection.PmmInfo.AmountOut.String(),
		TakerTokenAmount: trade.TradeInfo.AmountIn.String(),
		ContractAddress:  log.Address.String(),
		BlockNumber:      log.BlockNumber,
		TxHash:           log.TxHash.String(),
		LogIndex:         uint64(log.Index),
		LogTime:          time.Unix(int64(blockTime), 0),
		FromChain:        fromChainID,
		ToChain:          toChainID,
		TradeTimeout:     time.Unix(int64(selection.RfqInfo.TradeTimeout), 0),
	}
	return res, nil
}

func (p *Parser) Exchange() string {
	return parser.ExOptimex
}

func tradeInfoToData(tradeInfo [3][]byte) (string, string, string) { // chain, user, token
	chainID := string(tradeInfo[1])
	var tokenID string
	if len(tradeInfo[2]) == common.AddressLength {
		tokenID = common.BytesToAddress(tradeInfo[2]).String()
	} else {
		tokenID = string(tradeInfo[2])
	}
	switch chainID {
	case "ethereum", "ethereum_sepolia":
		return chainID, common.Address(tradeInfo[0]).String(), tokenID
	}
	return chainID, string(tradeInfo[0]), tokenID
}

func (p *Parser) LogFromExchange(log ethereumTypes.Log) (string, bool) {
	if len(log.Topics) == 0 {
		return "", false
	}
	name, ok := p.events[log.Topics[0]]
	return name, ok
}

func (p *Parser) ParseSettle(log ethereumTypes.Log, blockTime uint64) (string, error) {
	if len(log.Topics) > 0 && p.events[log.Topics[0]] != SettleEvent {
		return "", parser.ErrInvalidTopic
	}
	settle, err := p.filter.ParseConfirmSettlement(log)
	if err != nil {
		return "", err
	}
	return hexutil.Encode(settle.TradeId[:]), nil
}
