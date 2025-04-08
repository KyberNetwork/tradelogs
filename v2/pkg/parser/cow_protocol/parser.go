package cowprotocol

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
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
	TradeEvent = "Trade"
)

type Parser struct {
	abi       *abi.ABI
	ps        *CowProtocolFilterer
	eventHash string
}

func MustNewParser() *Parser {
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
	return &Parser{
		ps:        ps,
		abi:       ab,
		eventHash: tradeEvent.ID.String(),
	}
}

func (p *Parser) Topics() []string {
	return []string{
		p.eventHash,
	}
}

func (p *Parser) Parse(log ethereumTypes.Log, blockTime uint64) ([]storageTypes.TradeLog, error) {
	res, err := p.parse(log, blockTime)
	return []storageTypes.TradeLog{res}, err
}

func (p *Parser) parse(log ethereumTypes.Log, blockTime uint64) (storageTypes.TradeLog, error) {
	if len(log.Topics) > 0 && log.Topics[0].Hex() != p.eventHash {
		return storageTypes.TradeLog{}, parser.ErrInvalidTopic
	}
	e, err := p.ps.ParseTrade(log)
	if err != nil {
		return storageTypes.TradeLog{}, err
	}
	var rawData string
	rawDataJson, err := json.Marshal(e.Raw)
	rawData = string(rawDataJson)
	if err != nil {
		rawData = fmt.Sprintf("%+v", e.Raw)
	}
	res := storageTypes.TradeLog{
		Exchange:        p.Exchange(),
		Owner:           e.Owner.String(),
		SellToken:       e.SellToken.String(),
		BuyToken:        e.BuyToken.String(),
		SellAmount:      e.SellAmount.String(),
		BuyAmount:       e.BuyAmount.String(),
		FeeAmount:       e.FeeAmount.String(),
		OrderUid:        hex.EncodeToString(e.OrderUid),
		RawTradeData:    rawData,
		ContractAddress: e.Raw.Address.String(),
		BlockNumber:     e.Raw.BlockNumber,
		TxHash:          e.Raw.TxHash.String(),
		LogIndex:        uint64(e.Raw.Index),
		Timestamp:       blockTime * 1000,
		EventHash:       p.eventHash,
	}
	return res, nil
}

func (p *Parser) Exchange() string {
	return constant.CowProtocol
}

func (p *Parser) UseTraceCall() bool {
	return false
}

func (p *Parser) ParseWithCallFrame(callFrame types.CallFrame, log ethereumTypes.Log, blockTime uint64) ([]storageTypes.TradeLog, error) {
	res, err := p.parse(log, blockTime)
	if err != nil {
		return nil, err
	}
	return []storageTypes.TradeLog{res}, nil
}

func (p *Parser) LogFromExchange(log ethereumTypes.Log) bool {
	return strings.EqualFold(log.Address.String(), constant.AddrCowProtocol) &&
		len(log.Topics) > 0 &&
		strings.EqualFold(log.Topics[0].String(), p.eventHash)
}

func (p *Parser) Address() string {
	return constant.AddrCowProtocol
}
