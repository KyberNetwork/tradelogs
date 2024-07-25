package zxrfq_v3

import (
	"fmt"
	"github.com/KyberNetwork/tradelogs/pkg/abitypes"
	"github.com/KyberNetwork/tradelogs/pkg/decoder"
	"github.com/KyberNetwork/tradelogs/pkg/parser"
	"github.com/KyberNetwork/tradelogs/pkg/parser/zxrfq_v3/zxrfq_v3_helper"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"github.com/KyberNetwork/tradelogs/pkg/tracecall"
	"github.com/KyberNetwork/tradelogs/pkg/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
	"strings"
)

var (
	balanceArgument abi.Arguments
)

func init() {
	// uint256 balance
	balanceArgument = abi.Arguments{
		{Name: "balance", Type: abitypes.Uint256},
	}
}

type Parser struct {
	abi             *abi.ABI
	customAbi       *abi.ABI
	traceCalls      *tracecall.Cache
	contractAddress common.Address
	selectorAction  map[decoder.Bytes4]string
}

func MustNewParser(cache *tracecall.Cache, contractAddress common.Address) *Parser {
	ab, err := ZxrfqV3MetaData.GetAbi()
	if err != nil {
		log.Fatalf("failed to get abi: %s", err)
	}
	customAb, err := zxrfq_v3_helper.ZxrfqV3HelperMetaData.GetAbi()
	if err != nil {
		log.Fatalf("failed to get custom abi: %s", err)
	}
	zeroAddress := common.Address{}
	if contractAddress == zeroAddress {
		log.Fatalf("contract address is zero address")
	}
	return &Parser{
		abi:             ab,
		customAbi:       customAb,
		traceCalls:      cache,
		contractAddress: contractAddress,
		selectorAction:  getSettlerAction(),
	}
}

// Zeroxv3 has no topics
func (p *Parser) Topics() []string {
	return []string{}
}

func (p *Parser) Exchange() string {
	return parser.EXZeroXV3
}

func (p *Parser) UseTraceCall() bool {
	return true
}

func (p *Parser) Parse(log ethereumTypes.Log, blockTime uint64) (storage.TradeLog, error) {
	if len(log.Topics) > 0 || log.Address != p.contractAddress {
		return storage.TradeLog{}, ErrInvalidRfqTrade
	}
	var tradeLog storage.TradeLog
	tradeLog.BlockNumber = log.BlockNumber
	tradeLog.LogIndex = uint64(log.Index)
	tradeLog.TxHash = log.TxHash.Hex()
	tradeLog.Timestamp = blockTime * 1000
	tradeLog.ContractAddress = p.contractAddress.Hex()
	orderHash, err := getOrderHash(log.Data)
	if err != nil {
		return storage.TradeLog{}, fmt.Errorf("get order hash error %w", err)
	}
	tradeLog.OrderHash = orderHash

	callFrame, err := p.traceCalls.GetTraceCall(log.TxHash.Hex())
	if err != nil {
		return storage.TradeLog{}, err
	}
	return p.recursiveDetectRFQTrades(tradeLog, callFrame, log)
}

func getOrderHash(data []byte) (string, error) {
	if len(data) < 32 {
		return "", fmt.Errorf("data length is less than 32")
	}
	return common.BytesToHash(data[:32]).String(), nil
}

func (p *Parser) getRfqTradeIndex(callFrame types.CallFrame, log ethereumTypes.Log) (int, bool) {
	// action index is the index of []byte data of param actions which one has function call _logRfqOrder -> log0(0x00, 0x30)
	actionIndex := 0
	// if the log of tx have no topics, it will be zero0xv3 rfq trade
	for _, l := range callFrame.Logs {
		if len(l.Topics) == 0 && l.Address == p.contractAddress {
			if l.Data == hexutil.Encode(log.Data) {
				return actionIndex, true
			} else {
				actionIndex++
			}
		}
	}
	return actionIndex, false
}

func (p *Parser) recursiveDetectRFQTrades(tradeLog storage.TradeLog, callFrame types.CallFrame, log ethereumTypes.Log) (storage.TradeLog, error) {
	var (
		err error
	)

	actionIndex, isRfqTrade := p.getRfqTradeIndex(callFrame, log)
	if isRfqTrade {
		return p.ParseFromInternalCall(tradeLog, callFrame, log, actionIndex)
	}

	for _, subCall := range callFrame.Calls {
		tradeLog, err = p.recursiveDetectRFQTrades(tradeLog, subCall, log)
		if err == nil {
			return tradeLog, nil
		}
	}
	return tradeLog, fmt.Errorf("%w, tx_hash: %s", ErrNotFoundLogWithEmptyTopic, tradeLog.TxHash)
}

func (p *Parser) ParseFromInternalCall(tradeLog storage.TradeLog, callFrame types.CallFrame, log ethereumTypes.Log, actionIndex int) (storage.TradeLog, error) {
	actions, err := p.getExecuteActionData(callFrame)
	if err != nil {
		return tradeLog, err
	}
	// currentActionIndex is the action data when SC call function _logRfqOrder
	currentActionIndex := 0
	for _, action := range actions {
		// example actions [][]bytes have 4 slice action []bytes
		// [0] unsupported actionName
		// [1] supported actionName
		// [2] unsupported actionName
		// [3] supported actionName
		// so the maximum of currentActionIndex is 2 (1 and 3)
		actionName, data, err := decodeCall(action)
		if err != nil {
			return tradeLog, err
		}

		if functionName, ok := p.selectorAction[actionName]; ok {
			// with 0x7f6ceE965959295cC64d0E6c00d99d6532d8e86b production
			// for now we only support parse rfq with this actionName: METATXN_RFQ_VIP, RFQ, RFQ_VIP

			// with 0x7966aF62034313D87Ede39380bf60f1A84c62BE7 dev
			// for now we only support parse rfq with this actionName: SETTLER_OTC_SELF_FUNDED, METATXN_SETTLER_OTC_PERMIT2
			// because when action in these enum SC will call function  fillOtcOrderSelfFunded, fillOtcOrder and call log _logRfqOrder
			switch functionName {
			// dev SC
			case settler_otc_self_funded_name:
				if currentActionIndex == actionIndex {
					input, err := zxrfq_v3_helper.GetInputParamsOfFillRfqOrderSelfFunded(p.customAbi, methodIdDecodeParamOfFillOrderSelfFunded, data)
					if err != nil {
						return tradeLog, fmt.Errorf("get input param of fill rfq order self funded failed: %w", err)
					}
					tradeLog.Maker = input.Maker.String()
					tradeLog.Taker = callFrame.From
					tradeLog.MakerToken = input.Permit.Permitted.Token.String()
					tradeLog.TakerToken = input.TakerToken.String()
					tradeLog.Expiry = input.Permit.Deadline.Uint64()
					var takerToken *big.Int
					for _, subFrame := range callFrame.Calls {
						x, _ := decoder.Decode(p.abi, subFrame.Input)
						// find first method get balance of current contract for taker token
						if x != nil && x.Name == balanceOf &&
							strings.EqualFold(subFrame.To, tradeLog.TakerToken) &&
							strings.EqualFold(subFrame.From, p.contractAddress.Hex()) {
							//max current balance of zeroxv3 SC
							takerToken, err = decodeOutputBalanceOf(subFrame.Output)
							if err != nil {
								return tradeLog, err
							}
							break
						}
					}
					if takerToken == nil {
						return tradeLog, fmt.Errorf("can not find current balance of contract")
					}
					if takerToken.Cmp(input.MaxTakerAmount) > 0 {
						takerToken = new(big.Int).Set(input.MaxTakerAmount)
					}
					newMakerAmount := calculateTokenAmount(input.Permit.Permitted.Amount, takerToken, input.MaxTakerAmount)
					tradeLog.TakerTokenAmount = takerToken.String()
					tradeLog.MakerTokenAmount = newMakerAmount.String()
					return tradeLog, nil
				}
				break
			case metatxn_settler_otc_permit2_function:
				//TODO
				break
			}
		}

	}

	// if we can not find any action, we will return error
	return tradeLog, ErrDetectRfqButCanNotParse
}

func (p *Parser) getExecuteActionData(callFrame types.CallFrame) ([][]byte, error) {
	contractCall, err := decoder.Decode(p.abi, callFrame.Input)
	if err != nil {
		return nil, err
	}
	if contractCall.Name == executeFunctionName {
		for _, param := range contractCall.Params {
			if param.Name == actionParamName {
				data, ok := param.Value.([][]byte)
				if !ok {
					return nil, fmt.Errorf("failed to convert action param to bytes")
				}
				return data, nil
			}
		}
	}
	return [][]byte{}, nil
}

func decodeCall(data []byte) (decoder.Bytes4, []byte, error) {
	var action decoder.Bytes4
	// First 4 bytes are the function selector
	action, err := decoder.GetBytes4(data)
	if err != nil {
		return action, nil, err
	}

	// Remaining bytes are the arguments
	args := data[4:]
	return action, args, nil
}

func decodeOutputBalanceOf(data string) (*big.Int, error) {
	bytes, err := hexutil.Decode(data)
	if err != nil {
		return nil, err
	}
	decoded, err := balanceArgument.UnpackValues(bytes)
	if err != nil {
		return nil, err
	}
	balance, ok := decoded[0].(*big.Int)
	if !ok {
		return nil, fmt.Errorf("failed to convert balance to big.Int")
	}
	return balance, nil
}

func calculateTokenAmount(makerAmount *big.Int, takerAmount *big.Int, maxTakerAmount *big.Int) *big.Int {
	takerAmountCal := new(big.Int).Set(takerAmount)
	makerAmountCal := new(big.Int).Set(makerAmount)
	maxTakerAmountCal := new(big.Int).Set(maxTakerAmount)

	//makerAmount.unsafeMulDiv(takerAmount, maxTakerAmount);
	return makerAmountCal.Mul(makerAmountCal, takerAmountCal.Div(takerAmountCal, maxTakerAmountCal))
}
