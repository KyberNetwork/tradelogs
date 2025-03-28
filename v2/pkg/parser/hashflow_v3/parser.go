package hashflowv3

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/KyberNetwork/tradelogs/v2/pkg/constant"
	"github.com/KyberNetwork/tradelogs/v2/pkg/decoder"
	"github.com/KyberNetwork/tradelogs/v2/pkg/parser"
	storageTypes "github.com/KyberNetwork/tradelogs/v2/pkg/storage/tradelogs/types"
	"github.com/KyberNetwork/tradelogs/v2/pkg/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	TradeEvent = "Trade"
	paramName  = "quote"
)

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

func (p *Parser) parseLog(log ethereumTypes.Log, blockTime uint64) (storageTypes.TradeLog, error) {
	if len(log.Topics) > 0 && log.Topics[0].Hex() != p.eventHash {
		return storageTypes.TradeLog{}, parser.ErrInvalidTopic
	}
	o, err := p.ps.ParseTrade(log)
	if err != nil {
		return storageTypes.TradeLog{}, err
	}
	res := storageTypes.TradeLog{
		Exchange:         p.Exchange(),
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

func (p *Parser) Parse(log ethereumTypes.Log, blockTime uint64) ([]storageTypes.TradeLog, error) {
	tradeLog, err := p.parseLog(log, blockTime)
	if err != nil {
		return nil, err
	}
	return []storageTypes.TradeLog{tradeLog}, nil
}

func (p *Parser) Exchange() string {
	return constant.ExHashflowV3
}

func (p *Parser) UseTraceCall() bool {
	return false
}

func (p *Parser) ParseWithCallFrame(callFrame types.CallFrame, log ethereumTypes.Log, blockTime uint64) ([]storageTypes.TradeLog, error) {
	tradeLog, err := p.parseLog(log, blockTime)
	if err != nil {
		return nil, err
	}
	orderRfq, err := p.getRFQOrderParams(callFrame)
	if err != nil {
		return nil, err
	}
	tradeLog.Expiry = uint64(orderRfq.QuoteExpiry)
	tradeLog.MakerTokenOriginAmount = orderRfq.QuoteTokenAmount.String()
	tradeLog.TakerTokenOriginAmount = orderRfq.BaseTokenAmount.String()
	return []storageTypes.TradeLog{tradeLog}, nil
}

func (p *Parser) getRFQOrderParams(callFrame types.CallFrame) (*OrderRFQ, error) {
	contractCall, err := decoder.Decode(p.abi, callFrame.Input)
	if err != nil {
		return nil, err
	}
	if contractCall == nil {
		return nil, errors.New("missing contract_call")
	}

	for _, param := range contractCall.Params {
		if param.Name != paramName {
			continue
		}

		var rfqOrder OrderRFQ
		bytes, err := json.Marshal(param.Value)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(bytes, &rfqOrder); err != nil {
			return nil, err
		}

		return &rfqOrder, nil
	}
	return nil, nil
}

func (p *Parser) LogFromExchange(log ethereumTypes.Log) bool {
	return len(log.Topics) > 0 &&
		strings.EqualFold(log.Topics[0].String(), p.eventHash)
}

func (p *Parser) Address() string {
	return ""
}
