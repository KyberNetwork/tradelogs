package zxrfqv3

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/KyberNetwork/tradelogs/pkg/decoder"
	"github.com/KyberNetwork/tradelogs/pkg/parser"
	"github.com/KyberNetwork/tradelogs/pkg/parser/zxrfqv3/deployer"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"github.com/KyberNetwork/tradelogs/pkg/tracecall"
	"github.com/KyberNetwork/tradelogs/pkg/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

type Parser struct {
	contractABIs    *ContractABIs
	traceCalls      *tracecall.Cache
	deployerAddress common.Address
	selectorAction  map[decoder.Bytes4]FunctionName
	rfqArguments    rfqArguments
	l               *zap.SugaredLogger
}

const (
	altEventHash    = "0x0000000000000000000000000000000000000000000000000000000000000000"
	paddingByteSize = 32
	orderHashLen    = 32
	fillAmountLen   = 16
)

func MustNewParserWithDeployer(cache *tracecall.Cache, ethClient *ethclient.Client, deployerAddress common.Address, contractAbiSupported ...ContractABI) *Parser {
	if isZeroAddress(deployerAddress) {
		log.Fatalf("deployer Address is zero Address")
	}

	p := MustNewParser(cache, contractAbiSupported...)
	p.deployerAddress = deployerAddress

	deployer, err := deployer.NewDeployer(deployerAddress, ethClient)
	if err != nil {
		log.Fatalf("failed to create deployer %v", err)
	}

	p.getAndUpdateSupportedContractAddress(deployer)
	go p.monitorContractAddresses(deployer)
	return p
}

func MustNewParser(cache *tracecall.Cache, contractAbiSupported ...ContractABI) *Parser {
	contractABIs := &ContractABIs{
		mAddressABIs: make(map[common.Address]*abi.ABI),
	}

	for _, contract := range contractAbiSupported {
		newABI, err := NewABI(contract)
		if err != nil {
			log.Fatalf("get contract new ABI error %v", err)
		}
		contractABIs.addContractABI(contract.Address, newABI)
	}

	arguments, err := newRfqArguments()
	if err != nil {
		log.Fatalf("failed to create rfq arguments %v", err)
	}

	p := &Parser{
		traceCalls:     cache,
		selectorAction: getSettlerAction(),
		contractABIs:   contractABIs,
		rfqArguments:   arguments,
		l:              zap.S(),
	}
	return p
}

func (p *Parser) monitorContractAddresses(deployer *deployer.Deployer) {
	cronTime := time.NewTicker(1 * time.Minute)

	for range cronTime.C {
		p.getAndUpdateSupportedContractAddress(deployer)
	}
}

func (p *Parser) getAndUpdateSupportedContractAddress(deployer *deployer.Deployer) {
	contractTypeSupported := []ContractType{SwapContract, GaslessContract, NewGaslessContract}
	for _, contractType := range contractTypeSupported {
		err := p.getAndUpdateContractAddress(deployer, contractType)
		if err != nil {
			p.l.Errorw("error to get contract address", "error", err)
		}
	}
}

func (p *Parser) getAndUpdateContractAddress(deployer *deployer.Deployer, contractType ContractType) error {
	callOpts := &bind.CallOpts{Context: context.Background()}
	contractAddress, err := deployer.OwnerOf(callOpts, big.NewInt(int64(contractType)))
	if err != nil {
		return fmt.Errorf("contract type %v, failed to get contract address %v", contractType, err)
	}

	abi, err := NewABI(ContractABI{Address: contractAddress, ContractType: contractType})
	if err != nil {
		return err
	}
	if !p.contractABIs.containAddress(contractAddress) && !isZeroAddress(contractAddress) {
		p.l.Infow("add contract abi", "contractAddress", contractAddress)
		p.contractABIs.addContractABI(contractAddress, abi)
	}
	return nil
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
	tradeLog, err := p.buildOrderByLog(log, blockTime)
	if err != nil {
		return storage.TradeLog{}, err
	}

	callFrame, err := p.traceCalls.GetTraceCall(log.TxHash.Hex())
	if err != nil {
		return tradeLog, err
	}
	return p.recursiveDetectRFQTrades(tradeLog, callFrame, log)
}

func (p *Parser) buildOrderByLog(log ethereumTypes.Log, blockTime uint64) (storage.TradeLog, error) {
	if len(log.Topics) > 0 || !p.contractABIs.containAddress(log.Address) {
		return storage.TradeLog{}, ErrInvalidRfqTrade
	}
	var tradeLog storage.TradeLog
	tradeLog.BlockNumber = log.BlockNumber
	tradeLog.LogIndex = uint64(log.Index)
	tradeLog.TxHash = log.TxHash.Hex()
	tradeLog.Timestamp = blockTime * 1000
	tradeLog.ContractAddress = log.Address.Hex()
	tradeLog.EventHash = altEventHash
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

	return tradeLog, nil
}

func (p *Parser) ParseWithCallFrame(callFrame types.CallFrame, log ethereumTypes.Log, blockTime uint64) (storage.TradeLog, error) {
	tradeLog, err := p.buildOrderByLog(log, blockTime)
	if err != nil {
		return storage.TradeLog{}, err
	}
	return p.recursiveDetectRFQTrades(tradeLog, callFrame, log)
}

func getOrderHash(data []byte) (string, error) {
	if len(data) < common.HashLength {
		return "", fmt.Errorf("data length is less than 32")
	}
	return common.BytesToHash(data[:common.HashLength]).String(), nil
}

func getMakerAmount(data []byte) (*big.Int, error) {
	if len(data) < common.HashLength {
		return nil, fmt.Errorf("data length is less than 32")
	}

	return new(big.Int).SetBytes(data[common.HashLength:]), nil
}

func (p *Parser) getRfqTradeIndex(callFrame types.CallFrame, log ethereumTypes.Log) (int, bool) {
	// action index is the index of []byte data of param actions which one has function call _logRfqOrder -> log0(0x00, 0x30)
	actionIndex := 0
	// if the log of tx have no topics, it will be zero0xv3 rfq trade
	for _, l := range callFrame.Logs {
		if len(l.Topics) == 0 && p.contractABIs.containAddress(l.Address) {
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
	return tradeLog, fmt.Errorf("%w, tx_hash: %s %+v", ErrNotFoundLogWithEmptyTopic, tradeLog.TxHash, callFrame)
}

func (p *Parser) ParseFromInternalCall(tradeLog storage.TradeLog, callFrame types.CallFrame, actionIndex int) (storage.TradeLog, error) {
	actions, err := p.getExecuteActionData(common.HexToAddress(tradeLog.ContractAddress), callFrame)
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
			case settlerOtcSelfFundedName, rfqName:
				return p.getTradeLogFromFillRfqOrderSelfFunded(callFrame, tradeLog, data)
			case metatxnRFQVipName, rfqVIPName:
				return p.getTradeLogFromFillRfqOrderVIP(callFrame, tradeLog, data)
			}

		}

	}

	// if we can not find any action, we will return error
	return tradeLog, ErrDetectRfqButCanNotParse
}

func (p *Parser) getTradeLogFromFillRfqOrderSelfFunded(callFrame types.CallFrame, tradeLog storage.TradeLog, data []byte) (storage.TradeLog, error) {
	input, err := p.rfqArguments.UnpackInputParamsOfFillRfqOrderSelfFunded(data)
	if err != nil {
		return tradeLog, fmt.Errorf("get input param of fill rfq order self funded failed: %w", err)
	}
	tradeLog.Maker = input.Maker.String()
	tradeLog.Taker = callFrame.From
	tradeLog.MakerToken = input.PermitTransferFrom.Permitted.Token.String()
	tradeLog.TakerToken = input.TakerToken.String()
	tradeLog.Expiry = input.PermitTransferFrom.Deadline.Uint64()
	makerTokenAmount, ok := new(big.Int).SetString(tradeLog.MakerTokenAmount, 10)
	if !ok {
		return tradeLog, fmt.Errorf("failed to convert maker token amount to big.Int")
	}
	takerTokenAmount := calculateTakerTokenAmount(makerTokenAmount, input.MaxTakerAmount, input.PermitTransferFrom.Permitted.Amount)
	tradeLog.TakerTokenAmount = takerTokenAmount.String()
	return tradeLog, nil
}

func (p *Parser) getTradeLogFromFillRfqOrderVIP(callFrame types.CallFrame, tradeLog storage.TradeLog, data []byte) (storage.TradeLog, error) {
	input, err := p.rfqArguments.UnpackInputParamsOfFillRfqOrderVIP(data)
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

func (p *Parser) getExecuteActionData(contractAddress common.Address, callFrame types.CallFrame) ([][]byte, error) {
	abi, ok := p.contractABIs.getABI(contractAddress)
	if !ok {
		return nil, fmt.Errorf("no abi supported for contract %v", contractAddress)
	}
	contractCall, err := decoder.Decode(abi, callFrame.Input)
	if err != nil {
		actions, ok, errAction := p.DecodeExecuteInput(callFrame.Input)
		if errAction != nil {
			return nil, errAction
		}
		if ok {
			return actions, nil
		}
		return nil, err
	}
	if contractCall.Name == executeFunctionName || contractCall.Name == executeMetaTxnFunctionName {
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

func (p *Parser) LogFromExchange(log ethereumTypes.Log) bool {
	return p.contractABIs.containAddress(log.Address) &&
		len(log.Topics) == 0
}

func (p *Parser) Address() string {
	return ""
}

func (p *Parser) ExecutorAddresses() []string {
	return p.contractABIs.getAddresses()
}

func (p *Parser) DecodeExecuteInput(input string) ([][]byte, bool, error) {
	if len(input) < 10 {
		return nil, false, ErrorInputDataIsNotEnough
	}
	if !strings.HasPrefix(input, executeFunctionSelector) && !strings.HasPrefix(input, executeMetaTxnFunctionSelector) {
		return nil, false, nil
	}

	data, err := hex.DecodeString(input[10:])
	if err != nil {
		return nil, false, err
	}
	// function executeMetaTxn((address,address,uint256) slippage, bytes[] actions, bytes32 zid, address msgSender, bytes sig)
	// function execute((address,address,uint256) slippage, bytes[] actions, bytes32 zid)
	// skip: slippage (address,address,uint256): 32 + 32 + 32 = 96 bytes
	offset := paddingByteSize * 3
	if len(data) < offset {
		return nil, false, ErrorInputDataIsNotEnough
	}

	offsetActions := new(big.Int).SetBytes(data[offset : offset+paddingByteSize]).Uint64()
	// skip actions and zid
	offset += paddingByteSize * 2

	// with executeMetaTxn we need to skip msgSender and sig too
	if strings.HasPrefix(input, executeMetaTxnFunctionSelector) {
		offset += paddingByteSize * 2
	}
	if len(data) < offset {
		return nil, false, ErrorInputDataIsNotEnough
	}
	actionsSize := new(big.Int).SetBytes(data[offsetActions : offsetActions+paddingByteSize]).Uint64()
	baseOffset := int(offsetActions + paddingByteSize)
	// skip actions size
	offset += paddingByteSize
	eachActionOffset := make([]int, actionsSize)

	for i := 0; i < int(actionsSize); i++ {
		actionOffset := new(big.Int).SetBytes(data[offset : offset+paddingByteSize]).Uint64()
		eachActionOffset[i] = int(actionOffset)
		offset += paddingByteSize
	}

	actions := make([][]byte, actionsSize)
	for i := 0; i < int(actionsSize); i++ {
		startOffset := baseOffset + eachActionOffset[i] + paddingByteSize
		actionSize := int(new(big.Int).SetBytes(data[eachActionOffset[i]:startOffset]).Int64())
		endIndex := startOffset + actionSize
		if actionSize > len(data) {
			endIndex = len(data)
		}
		actions[i] = data[startOffset:endIndex]
	}
	return actions, true, nil
}

func (p *Parser) ExtractLogData(log ethereumTypes.Log) (string, *big.Int, error) {
	if len(log.Data) < orderHashLen+fillAmountLen {
		return "", nil, errors.New("invalid data")
	}
	return hexutil.Encode(log.Data[:orderHashLen]), new(big.Int).SetBytes(log.Data[orderHashLen : orderHashLen+fillAmountLen]), nil
}
