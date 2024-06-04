package bebop

import (
	"encoding/json"
	"errors"
	"math/big"
	"strings"

	"github.com/KyberNetwork/tradelogs/internal/types"
	"github.com/KyberNetwork/tradelogs/pkg/decoder"
	"github.com/KyberNetwork/tradelogs/pkg/parser"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"github.com/KyberNetwork/tradelogs/pkg/tracecall"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	TradeEvent = "BebopOrder"
	OrderParam = "order"
)

var (
	ErrTradeTopic        = errors.New("invalid trade topic")
	ErrNotFoundTraceCall = errors.New("not found trace call")
)

type SingleOrder struct {
	Expiry         *big.Int `json:"expiry"`
	TakerAddress   string   `json:"taker_address"`
	MakerAddress   string   `json:"maker_address"`
	MakerNonce     *big.Int `json:"maker_nonce"`
	TakerToken     string   `json:"taker_token"`
	MakerToken     string   `json:"maker_token"`
	TakerAmount    *big.Int `json:"taker_amount"`
	MakerAmount    *big.Int `json:"maker_amount"`
	Receiver       string   `json:"receiver"`
	PackedCommands *big.Int `json:"packed_commands"`
	Flags          *big.Int `json:"flags"`
}

type Parser struct {
	abi        *abi.ABI
	ps         *BebopFilterer
	eventHash  string
	traceCalls *tracecall.Cache
}

func MustNewParser(tracecall *tracecall.Cache) *Parser {
	ps, err := NewBebopFilterer(common.Address{}, nil)
	if err != nil {
		panic(err)
	}
	ab, err := abi.JSON(strings.NewReader(BebopMetaData.ABI))
	if err != nil {
		panic(err)
	}
	event, ok := ab.Events[TradeEvent]
	if !ok {
		panic("no such event: " + TradeEvent)
	}
	return &Parser{
		ps:         ps,
		abi:        &ab,
		eventHash:  event.ID.String(),
		traceCalls: tracecall,
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
	o, err := p.ps.ParseBebopOrder(log)
	if err != nil {
		return storage.TradeLog{}, err
	}
	order := storage.TradeLog{
		OrderHash:       o.EventId.String(),
		Maker:           log.Address.Hex(),
		ContractAddress: o.Raw.Address.String(),
		BlockNumber:     o.Raw.BlockNumber,
		TxHash:          o.Raw.TxHash.String(),
		LogIndex:        uint64(o.Raw.Index),
		Timestamp:       blockTime * 1000,
		EventHash:       p.eventHash,
	}
	return p.parseTraceCall(order)
}

func (p *Parser) parseTraceCall(order storage.TradeLog) (storage.TradeLog, error) {
	traceCall, err := p.traceCalls.GetTraceCall(order.TxHash)
	if err != nil {
		return order, err
	}

	order, err = p.searchTradeLog(order, traceCall)
	if err != nil {
		return order, err
	}

	return order, nil
}

func (p *Parser) searchTradeLog(order storage.TradeLog, traceCall types.CallFrame) (storage.TradeLog, error) {
	if p.checkBebopTrade(traceCall) {
		return p.ParseFromInternalCall(order, traceCall)
	}

	for _, subCall := range traceCall.Calls {
		tradeLog, err := p.searchTradeLog(order, subCall)
		if err == nil {
			return tradeLog, nil
		}
	}
	return order, ErrNotFoundTraceCall
}

func (p *Parser) checkBebopTrade(traceCall types.CallFrame) bool {
	for _, eventLog := range traceCall.Logs {
		if len(eventLog.Topics) == 0 {
			continue
		}
		if !strings.EqualFold(eventLog.Topics[0].String(), p.eventHash) {
			continue
		}
		return true
	}
	return false
}

func (p *Parser) ParseFromInternalCall(order storage.TradeLog, internalCall types.CallFrame) (storage.TradeLog, error) {
	contractCall, err := decoder.Decode(p.abi, internalCall.Input)
	if err != nil {
		return order, err
	}

	for _, param := range contractCall.Params {
		if param.Name != OrderParam {
			continue
		}
		var rfqOrder SingleOrder
		bytes, err := json.Marshal(param.Value)
		if err != nil {
			return order, err
		}

		if err := json.Unmarshal(bytes, &rfqOrder); err != nil {
			return order, err
		}
		order.MakerToken = rfqOrder.MakerToken
		order.TakerToken = rfqOrder.TakerToken
		order.Maker = rfqOrder.MakerAddress
		order.Taker = rfqOrder.TakerAddress
		order.MakerTokenAmount = rfqOrder.MakerAmount.String()
		order.TakerTokenAmount = rfqOrder.TakerAmount.String()
	}

	return order, nil
}

func (p *Parser) Exchange() string {
	return parser.ExBebop
}

func (p *Parser) UseTraceCall() bool {
	return true
}
