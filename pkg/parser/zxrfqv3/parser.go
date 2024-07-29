package zxrfqv3

import (
	"fmt"
	"github.com/KyberNetwork/tradelogs/pkg/decoder"
	"github.com/KyberNetwork/tradelogs/pkg/parser"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"github.com/KyberNetwork/tradelogs/pkg/tracecall"
	"github.com/KyberNetwork/tradelogs/pkg/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
	"os"
)

type Parser struct {
	abi             *abi.ABI
	customAbi       *abi.ABI
	traceCalls      *tracecall.Cache
	contractAddress common.Address
	selectorAction  map[decoder.Bytes4]FunctionName
}

func loadABI(filePath string) *abi.ABI {
	abiByteValue, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open path file %v, error: %s", err, filePath)
	}
	ab, err := abi.JSON(abiByteValue)
	if err != nil {
		log.Fatalf("failed to get abi: %s", err)
	}
	return &ab
}

func mustNewParser(cache *tracecall.Cache, contractAddress common.Address, filePath string) *Parser {
	// read abi from file path and get
	ab := loadABI(filePath)

	customAbi := loadABI("./abi/custom_abi.json")

	if isZeroAddress(contractAddress) {
		log.Fatalf("contract address is zero address")
	}
	return &Parser{
		abi:             ab,
		customAbi:       customAbi,
		traceCalls:      cache,
		contractAddress: contractAddress,
		selectorAction:  getSettlerAction(),
	}
}

func MustNewDevParser(cache *tracecall.Cache, contractAddress common.Address) *Parser {
	return mustNewParser(cache, contractAddress, "./abi/dev_abi.json")
}

func MustNewSwapParser(cache *tracecall.Cache, contractAddress common.Address) *Parser {
	return mustNewParser(cache, contractAddress, "./abi/swap_abi.json")
}

func MustNewGaslessParser(cache *tracecall.Cache, contractAddress common.Address) *Parser {
	return mustNewParser(cache, contractAddress, "./abi/gasless_abi.json")
}

func isZeroAddress(address common.Address) bool {
	return address == common.Address{}
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
	tokenMakerAmount, err := getMakerAmount(log.Data)
	if err != nil {
		return storage.TradeLog{}, fmt.Errorf("get maker amount error %w", err)
	}
	tradeLog.MakerTokenAmount = tokenMakerAmount.String()

	callFrame, err := p.traceCalls.GetTraceCall(log.TxHash.Hex())
	if err != nil {
		return tradeLog, err
	}
	return p.recursiveDetectRFQTrades(tradeLog, callFrame, log)
}

func getOrderHash(data []byte) (string, error) {
	if len(data) < common.HashLength {
		return "", fmt.Errorf("data length is less than 32")
	}
	return common.BytesToHash(data[:32]).String(), nil
}

func getMakerAmount(data []byte) (*big.Int, error) {
	if len(data) < common.HashLength {
		return nil, fmt.Errorf("data length is less than 32")
	}

	return new(big.Int).SetBytes(data[32:]), nil
}

func (p *Parser) getRfqTradeIndex(callFrame types.CallFrame, log ethereumTypes.Log) (int, bool) {
	// action index is the index of []byte data of param actions which one has function call _logRfqOrder -> log0(0x00, 0x30)
	actionIndex := 0
	// if the log of tx have no topics, it will be zero0xv3 rfq trade
	for _, l := range callFrame.Logs {
		if len(l.Topics) == 0 && l.Address == p.contractAddress {
			actionIndex++
			if l.Data == hexutil.Encode(log.Data) {
				return actionIndex, true
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
		return p.ParseFromInternalCall(tradeLog, callFrame, actionIndex)
	}

	for _, subCall := range callFrame.Calls {
		tradeLog, err = p.recursiveDetectRFQTrades(tradeLog, subCall, log)
		if err == nil {
			return tradeLog, nil
		}
	}
	return tradeLog, fmt.Errorf("%w, tx_hash: %s", ErrNotFoundLogWithEmptyTopic, tradeLog.TxHash)
}

func (p *Parser) ParseFromInternalCall(tradeLog storage.TradeLog, callFrame types.CallFrame, actionIndex int) (storage.TradeLog, error) {
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
			currentActionIndex++
			if currentActionIndex != actionIndex {
				continue
			}
			// with 0x7f6ceE965959295cC64d0E6c00d99d6532d8e86b production
			// for now we only support parse rfq with this actionName: METATXN_RFQ_VIP, RFQ, RFQ_VIP

			// with 0x7966aF62034313D87Ede39380bf60f1A84c62BE7 abi
			// for now we only support parse rfq with this actionName: SETTLER_OTC_SELF_FUNDED, METATXN_SETTLER_OTC_PERMIT2
			// because when action in these enum SC will call function  fillOtcOrderSelfFunded, fillOtcOrder and call log _logRfqOrder
			switch functionName {
			case settlerOtcSelfFundedName:
				return getTradeLogFromSettlerOtcSelfFundedName(callFrame, p.customAbi, tradeLog, data)
			case metatxnRFQVipName, rfqVIPName:
				return getTradeLogFromMetatxnRFQVipName(callFrame, p.customAbi, tradeLog, data)
			}

		}

	}

	// if we can not find any action, we will return error
	return tradeLog, ErrDetectRfqButCanNotParse
}

func getTradeLogFromSettlerOtcSelfFundedName(callFrame types.CallFrame, abi *abi.ABI, tradeLog storage.TradeLog, data []byte) (storage.TradeLog, error) {
	input, err := DecodeInputParamsOfFillRfqOrderSelfFunded(abi, methodIdDecodeParamOfFillOrderSelfFunded, data)
	if err != nil {
		return tradeLog, fmt.Errorf("get input param of fill rfq order self funded failed: %w", err)
	}
	tradeLog.Maker = input.Maker.String()
	tradeLog.Taker = callFrame.From
	tradeLog.MakerToken = input.Permit.Permitted.Token.String()
	tradeLog.TakerToken = input.TakerToken.String()
	tradeLog.Expiry = input.Permit.Deadline.Uint64()
	makerTokenAmount, ok := new(big.Int).SetString(tradeLog.MakerTokenAmount, 10)
	if !ok {
		return tradeLog, fmt.Errorf("failed to convert maker token amount to big.Int")
	}
	takerTokenAmount := calculateTakerTokenAmount(makerTokenAmount, input.MaxTakerAmount, input.Permit.Permitted.Amount)
	tradeLog.TakerTokenAmount = takerTokenAmount.String()
	return tradeLog, nil
}

func getTradeLogFromMetatxnRFQVipName(callFrame types.CallFrame, abi *abi.ABI, tradeLog storage.TradeLog, data []byte) (storage.TradeLog, error) {
	input, err := DecodeInputParamsOfFillRfqOrderVIP(abi, methodIdDecodeParamOfFillOrderVIP, data)
	if err != nil {
		return tradeLog, fmt.Errorf("get input param of fill rfq order vip failed: %w", err)
	}
	tradeLog.Maker = input.Maker.String()
	tradeLog.Taker = callFrame.From
	tradeLog.MakerToken = input.MakerPermit.Permitted.Token.String()
	tradeLog.TakerToken = input.TakerPermit.Permitted.Token.String()
	tradeLog.Expiry = input.MakerPermit.Deadline.Uint64()
	takerTokenAmount := input.TakerPermit.Permitted.Amount
	tradeLog.TakerTokenAmount = takerTokenAmount.String()
	return tradeLog, nil
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

func calculateTakerTokenAmount(makerTokenAmount *big.Int, maxTakerTokenAmount *big.Int, permittedTokenAmount *big.Int) *big.Int {
	if permittedTokenAmount.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0)
	}
	makerTokenAmountCal := new(big.Int).Set(makerTokenAmount)
	maxTakerTokenAmountCal := new(big.Int).Set(maxTakerTokenAmount)
	permittedTokenAmountCal := new(big.Int).Set(permittedTokenAmount)

	//makerAmount = permittedTokenAmount.unsafeMulDiv(takerAmount, maxTakerAmount);
	// -> takerAmount = (makerAmount * maxTakerAmount) / permittedTokenAmount
	tmp := makerTokenAmountCal.Mul(makerTokenAmountCal, maxTakerTokenAmountCal)
	return tmp.Div(tmp, permittedTokenAmountCal)
}
