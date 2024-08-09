package zxotc

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
	OtcOrderFilledEvent = "OtcOrderFilled"
	paramName           = "order"
)

type Parser struct {
	abi       *abi.ABI
	ps        *ZeroXOTCFilterer
	eventHash string
}

func MustNewParser() *Parser {
	ps, err := NewZeroXOTCFilterer(common.Address{}, nil)
	if err != nil {
		panic(err)
	}
	ab, err := ZeroXOTCMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	event, ok := ab.Events[OtcOrderFilledEvent]
	if !ok {
		panic("no such event: OtcOrderFilled")
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
		return storage.TradeLog{}, parser.ErrInvalidTopic
	}
	o, err := p.ps.ParseOtcOrderFilled(log)
	if err != nil {
		return storage.TradeLog{}, err
	}
	res := storage.TradeLog{
		OrderHash:        common.Hash(o.OrderHash).String(),
		Maker:            o.Maker.Hex(),
		Taker:            o.Taker.Hex(),
		MakerToken:       o.MakerToken.Hex(),
		TakerToken:       o.TakerToken.Hex(),
		MakerTokenAmount: o.MakerTokenFilledAmount.String(),
		TakerTokenAmount: o.TakerTokenFilledAmount.String(),
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
	return parser.ExZeroX
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
	tradeLog.Expiry = orderRfq.GetExpiry()
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
	return strings.EqualFold(log.Address.String(), parser.Addr0x) &&
		len(log.Topics) > 0 &&
		strings.EqualFold(log.Topics[0].String(), p.eventHash)
}
