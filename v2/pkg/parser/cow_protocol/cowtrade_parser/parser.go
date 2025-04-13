package cowprotocol

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/KyberNetwork/tradelogs/v2/pkg/constant"
	cowStorage "github.com/KyberNetwork/tradelogs/v2/pkg/storage/cow_protocol"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	TradeEvent = "Trade"
)

type CowTradeParser struct {
	abi            *abi.ABI
	ps             *CowProtocolFilterer
	tradeEventHash string
}

func MustNewParser() *CowTradeParser {
	ps, err := NewCowProtocolFilterer(common.Address{}, nil)
	if err != nil {
		panic(err)
	}
	ab, err := CowProtocolMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	tradeEvent, ok := ab.Events[TradeEvent]
	if !ok {
		panic("no such event: Trade")
	}
	return &CowTradeParser{
		ps:             ps,
		abi:            ab,
		tradeEventHash: tradeEvent.ID.String(),
	}
}

func (p *CowTradeParser) Topics() []string {
	return []string{
		p.tradeEventHash,
	}
}

func (p *CowTradeParser) Parse(log ethereumTypes.Log, blockTime uint64) (cowStorage.CowTrade, error) {
	if len(log.Topics) > 0 && log.Topics[0].Hex() != p.tradeEventHash {
		return cowStorage.CowTrade{}, errors.New("invalid order topic")
	}
	e, err := p.ps.ParseTrade(log)
	if err != nil {
		return cowStorage.CowTrade{}, err
	}
	var rawData string
	rawDataJson, err := json.Marshal(log)
	rawData = string(rawDataJson)
	if err != nil {
		rawData = fmt.Sprintf("%+v", log)
	}
	res := cowStorage.CowTrade{
		Owner:           e.Owner.String(),
		SellToken:       e.SellToken.String(),
		BuyToken:        e.BuyToken.String(),
		SellAmount:      e.SellAmount.String(),
		BuyAmount:       e.BuyAmount.String(),
		FeeAmount:       e.FeeAmount.String(),
		OrderUid:        hex.EncodeToString(e.OrderUid),
		RawTradeData:    rawData,
		ContractAddress: log.Address.String(),
		BlockNumber:     log.BlockNumber,
		TxHash:          log.TxHash.String(),
		LogIndex:        uint64(log.Index),
		Timestamp:       blockTime * 1000,
		EventHash:       p.tradeEventHash,
	}
	return res, nil
}

func (p *CowTradeParser) IsMatchLog(log ethereumTypes.Log) bool {
	return strings.EqualFold(log.Address.String(), constant.AddrCowProtocol) &&
		len(log.Topics) > 0 &&
		strings.EqualFold(log.Topics[0].String(), p.tradeEventHash)
}
