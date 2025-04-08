package cowprotocol

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
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
	TradeEvent        = "Trade"
	TransferEventHash = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	CowAddressInLog   = "0x0000000000000000000000009008d19f58aabd9ed0d60971565aa8510560ab41"
)

type Parser struct {
	abi            *abi.ABI
	ps             *CowProtocolFilterer
	tradeEventHash string
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
		ps:             ps,
		abi:            ab,
		tradeEventHash: tradeEvent.ID.String(),
	}
}

func (p *Parser) Topics() []string {
	return []string{
		p.tradeEventHash,
	}
}

func (p *Parser) Parse(log ethereumTypes.Log, blockTime uint64) ([]storageTypes.TradeLog, error) {
	res, err := p.parse(log, blockTime)
	return []storageTypes.TradeLog{res}, err
}

func (p *Parser) parse(log ethereumTypes.Log, blockTime uint64) (storageTypes.TradeLog, error) {
	if len(log.Topics) > 0 && log.Topics[0].Hex() != p.tradeEventHash {
		return storageTypes.TradeLog{}, parser.ErrInvalidTopic
	}
	e, err := p.ps.ParseTrade(log)
	if err != nil {
		return storageTypes.TradeLog{}, err
	}
	var rawData string
	rawDataJson, err := json.Marshal(log)
	rawData = string(rawDataJson)
	if err != nil {
		rawData = fmt.Sprintf("%+v", log)
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
		ContractAddress: log.Address.String(),
		BlockNumber:     log.BlockNumber,
		TxHash:          log.TxHash.String(),
		LogIndex:        uint64(log.Index),
		Timestamp:       blockTime * 1000,
		EventHash:       p.tradeEventHash,
	}
	return res, nil
}

func (p *Parser) Exchange() string {
	return constant.CowProtocol
}

func (p *Parser) UseTraceCall() bool {
	return false
}

func (p *Parser) ParseWithCallFrame(_ types.CallFrame, log ethereumTypes.Log, blockTime uint64) ([]storageTypes.TradeLog, error) {
	res, err := p.parse(log, blockTime)
	if err != nil {
		return nil, err
	}
	return []storageTypes.TradeLog{res}, nil
}

func (p *Parser) ParseTransferEvent(
	txHash string,
	block_number, timestamp uint64,
	call types.CallFrame,
) []interface{} {
	var result []interface{}

	// process the sub trace calls
	for _, traceCall := range call.Calls {
		events := p.ParseTransferEvent(txHash, block_number, timestamp, traceCall)
		result = append(result, events...)

	}
	for _, log := range call.Logs {
		if len(log.Topics) <= 3 {
			continue
		}
		if log.Topics[0].String() != TransferEventHash {
			continue
		}
		amount := new(big.Int)
		amount.SetString(log.Data[2:], 16)
		cowTransferEvent := storageTypes.CowTransfer{
			TxHash:      txHash,
			BlockNumber: block_number,
			Timestamp:   timestamp,
			FromAddress: log.Topics[1].String(),
			ToAddress:   log.Topics[2].String(),
			Token:       log.Address.String(),
			Amount:      amount.String(),
		}
		if cowTransferEvent.FromAddress == CowAddressInLog || cowTransferEvent.ToAddress == CowAddressInLog {
			result = append(result, cowTransferEvent)
		}

	}
	return result
}

func (p *Parser) LogFromExchange(log ethereumTypes.Log) bool {
	return strings.EqualFold(log.Address.String(), constant.AddrCowProtocol) &&
		len(log.Topics) > 0 &&
		strings.EqualFold(log.Topics[0].String(), p.tradeEventHash)
}

func (p *Parser) Address() string {
	return constant.AddrCowProtocol
}
