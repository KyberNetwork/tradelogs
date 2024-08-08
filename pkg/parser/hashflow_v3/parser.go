package hashflowv3

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/KyberNetwork/tradelogs/pkg/decoder"
	tradingTypes "github.com/KyberNetwork/tradinglib/pkg/types"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/KyberNetwork/tradelogs/pkg/parser"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	TradeEvent = "Trade"
	paramName  = "quote"
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

func (p *Parser) Parse(log ethereumTypes.Log, blockTime uint64) (storage.TradeLog, error) {
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

func (p *Parser) UseTraceCall() bool {
	return false
}

func (p *Parser) ParseWithCallFrame(callFrame *tradingTypes.CallFrame, log ethereumTypes.Log, blockTime uint64) (storage.TradeLog, error) {
	if callFrame == nil {
		return storage.TradeLog{}, errors.New("missing call frame")
	}
	tradeLog, err := p.Parse(log, blockTime)
	if err != nil {
		return storage.TradeLog{}, err
	}
	orderRfq, err := p.getRFQOrderParams(callFrame)
	if err != nil {
		return storage.TradeLog{}, err
	}
	tradeLog.Expiry = uint64(orderRfq.QuoteExpiry)
	return tradeLog, nil
}

func (p *Parser) getRFQOrderParams(callFrame *tradingTypes.CallFrame) (*OrderRFQ, error) {
	var (
		err error
	)
	contractCall := callFrame.ContractCall
	if contractCall == nil {
		contractCall, err = decoder.Decode(p.abi, hexutil.Encode(callFrame.Input))
		if err != nil {
			return nil, err
		}
		if contractCall == nil {
			return nil, errors.New("missing contract_call")
		}
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
