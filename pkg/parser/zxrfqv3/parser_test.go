package zxrfqv3

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/KyberNetwork/tradelogs/pkg/parser"
	"github.com/KyberNetwork/tradelogs/pkg/rpcnode"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"github.com/KyberNetwork/tradelogs/pkg/tracecall"
	"github.com/KyberNetwork/tradelogs/pkg/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	etypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

const rpcURL = ""

func newParserTest(t *testing.T, needRpc bool, contractABI ...ContractABI) *Parser {
	var cache *tracecall.Cache
	if needRpc {
		ethClient, err := ethclient.Dial(rpcURL)
		if err != nil {
			t.Fatalf("failed to dial to rpc url: %s, err: %s", rpcURL, err)
		}
		cache = tracecall.NewCache(rpcnode.NewClient(zap.S(), ethClient))
	}
	return MustNewParser(cache, contractABI...)
}

func getTestCaseData(t *testing.T, path string, result interface{}) {
	byteValue, err := os.Open(path)
	assert.NoError(t, err)
	assert.NoError(t, json.NewDecoder(byteValue).Decode(&result), "failed to parse call frame has rfq")
}

func TestGetActionDataFromCallFame(t *testing.T) {
	selectorAction := getSettlerAction()

	var callFrame types.CallFrame
	getTestCaseData(t, "./test/call_frame_rfq.json", &callFrame)
	var expectedInput InputParamOfFillRfqOrderSelfFunded
	getTestCaseData(t, "./test/expected_input_rfq.json", &expectedInput)

	contractAddress := common.HexToAddress("0x7966aF62034313D87Ede39380bf60f1A84c62BE7")
	parser := newParserTest(t, false, ContractABI{
		Address:      contractAddress,
		ContractType: DevContract,
	})
	data, err := parser.getExecuteActionData(contractAddress, callFrame)
	assert.NoError(t, err, "failed to get execute action data")
	actionType, rawData, err := decodeCall(data[1])
	assert.NoError(t, err, "failed to decode call")
	_, ok := selectorAction[actionType]
	assert.True(t, ok, "action type not found")
	assert.NoError(t, err, "failed to get method id")
	input, err := parser.rfqArguments.UnpackInputParamsOfFillRfqOrderSelfFunded(rawData)
	assert.NoError(t, err, "failed to decode data")

	assert.Equal(t, expectedInput.Recipient, input.Recipient, "recipient not match")
	assert.Equal(t, expectedInput.PermitTransferFrom.Permitted.Token, input.PermitTransferFrom.Permitted.Token, "permitted token not match")
	assert.Equal(t, expectedInput.PermitTransferFrom.Permitted.Amount, input.PermitTransferFrom.Permitted.Amount, "permitted amount not match")
	assert.Equal(t, expectedInput.PermitTransferFrom.Deadline, input.PermitTransferFrom.Deadline, "deadline not match")
	assert.Equal(t, expectedInput.PermitTransferFrom.Nonce, input.PermitTransferFrom.Nonce, "nonce not match")
	assert.Equal(t, expectedInput.Maker, input.Maker, "maker not match")
	assert.Equal(t, expectedInput.MakerSig, input.MakerSig, "maker sig not match")
	assert.Equal(t, expectedInput.TakerToken, input.TakerToken, "taker token not match")
	assert.Equal(t, expectedInput.MaxTakerAmount, input.MaxTakerAmount, "max taker amount not match")
}

func TestGetTransactionAndParseZeroxV3Rfq(t *testing.T) {
	t.Skip("Need to add the rpc url that enables the trace call JSON-RPC")
	var expectedTradelog storage.TradeLog
	getTestCaseData(t, "./test/expected_rfq.json", &expectedTradelog)

	txHash := common.HexToHash("0xb244877d0cf0badcf2ac82dbb3cbf338b420ab5d1c6e6630fce4a4874121e427")
	contractAddress := common.HexToAddress("0x7966aF62034313D87Ede39380bf60f1A84c62BE7")
	parser := newParserTest(t, true, ContractABI{
		Address:      contractAddress,
		ContractType: DevContract,
	})
	ethClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		t.Fatalf("failed to dial to rpc url: %s, err: %s", rpcURL, err)
	}

	txReceipt, err := ethClient.TransactionReceipt(context.Background(), txHash)
	assert.NoError(t, err, "failed to get tx receipt")
	rfqCount := 0
	for _, log := range txReceipt.Logs {
		if parser.contractABIs.containAddress(log.Address) && len(log.Topics) == 0 {
			rfqCount++
			tradeLog, err := parser.Parse(*log, 1)
			assert.NoError(t, err, "failed to parse trade log")
			assert.Equal(t, expectedTradelog, tradeLog, "trade log not match")
		}
	}
	assert.Greater(t, rfqCount, 0, "no rfq trade found")
}

func TestMustNewParserWithDeployer(t *testing.T) {
	t.Skip("Need to add the rpc url that enables the trace call JSON-RPC")
	contractAddress := common.HexToAddress("0x00000000000004533Fe15556B1E086BB1A72cEae")
	ethClient, err := ethclient.Dial(rpcURL)
	require.NoError(t, err, "failed to dial to rpc url: %s, err: %s", rpcURL, err)
	parser := MustNewParserWithDeployer(nil, ethClient, contractAddress)
	assert.NotNil(t, parser, "failed to create parser")
}

func TestParseTradeLogsZeroxv3(t *testing.T) {
	t.Skip("Need to add the rpc url that enables the trace call JSON-RPC")
	type testCase struct {
		TxHash           common.Hash
		ExpectedRfqTrade int
	}

	testCases := []testCase{
		{
			TxHash:           common.HexToHash("0x8775544670f4c40271b3c9ab1b48e0648a1862bedd600c46f017e999b132e6a8"),
			ExpectedRfqTrade: 1,
		},
		{
			TxHash:           common.HexToHash("0x99aba817ba6171488417da6990947413064eaff5200bc8d0be9f0c87580f19cc"),
			ExpectedRfqTrade: 1,
		},
		{
			TxHash:           common.HexToHash("0xd90590fbf2799cbe63dac8a8a3dd05798b4ee0bd9077654fc1fc85cc09ad5bd1"),
			ExpectedRfqTrade: 2,
		},
		{
			TxHash:           common.HexToHash("0x3fff39d6025924755ecfe4787e09cec3f87926791ce31ac79e58ba615463b564"),
			ExpectedRfqTrade: 4,
		},
		{
			TxHash:           common.HexToHash("0x56ecb98917f20cfe1864a45e18d2afadbdf43b16b397464fda32bc08bb3018b6"),
			ExpectedRfqTrade: 2,
		},
		{
			TxHash:           common.HexToHash("0xb81b2c0467c3c7b8f85a46c25b4b5fdb10256b36d0b2e1426c11f8ed1d710b61"),
			ExpectedRfqTrade: 1,
		},
	}

	swapContractAddress := common.HexToAddress("0x70bf6634eE8Cb27D04478f184b9b8BB13E5f4710")
	gasLessContractAddress := common.HexToAddress("0x12D737470fB3ec6C3DeEC9b518100Bec9D520144")
	parser := newParserTest(t, true, ContractABI{
		Address:      swapContractAddress,
		ContractType: SwapContract,
	}, ContractABI{
		Address:      gasLessContractAddress,
		ContractType: SwapContract,
	})
	ethClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		t.Fatalf("failed to dial to rpc url: %s, err: %s", rpcURL, err)
	}

	for _, tc := range testCases {
		txReceipt, err := ethClient.TransactionReceipt(context.Background(), tc.TxHash)
		require.NoError(t, err, "failed to get tx receipt")
		rfqCount := 0
		for _, log := range txReceipt.Logs {
			if parser.contractABIs.containAddress(log.Address) && len(log.Topics) == 0 {
				rfqCount++
				tradeLog, err := parser.Parse(*log, 1)
				assert.NoError(t, err, "failed to parse trade log")
				t.Log(tradeLog)
			}
		}
		assert.Equal(t, tc.ExpectedRfqTrade, rfqCount, fmt.Sprintf("Tx hash %s number rfq is not expected", tc.TxHash.Hex()))
	}
}

func TestParserActionsWithCuttingBytes(t *testing.T) {
	parser := MustNewParser(nil)
	type inputTest struct {
		InputData string   `json:"input_data"`
		Actions   []string `json:"actions"`
	}
	var inputTestCases []inputTest
	getTestCaseData(t, "./test/input_execute_data.json", &inputTestCases)
	for _, input := range inputTestCases {
		inputData := input.InputData
		actions, ok, err := parser.DecodeExecuteInput(inputData)
		assert.NoError(t, err, "failed to parse action data")
		assert.True(t, ok, "failed to parse action data")
		require.Equal(t, len(actions), len(input.Actions), "action length not match")
		for i, action := range actions {
			assert.Equal(t, input.Actions[i], hex.EncodeToString(action), "action not match")
		}
	}
}

func TestExtractLogData(t *testing.T) {
	p := MustNewParser(nil)
	data, err := hexutil.Decode("0x0fe145508fd6aa299e55f8b299284d6fef19bfaffe235a22120caee307bc582e0000000000000000349837334b321000")
	require.NoError(t, err)
	orderHash, amount, err := p.ExtractLogData(etypes.Log{
		Data: data,
	})
	require.NoError(t, err)
	t.Log(orderHash, amount)
	require.Equal(t, "0x0fe145508fd6aa299e55f8b299284d6fef19bfaffe235a22120caee307bc582e", orderHash)
}

func TestNewGaslessParser(t *testing.T) {
	t.Skip()
	ethclient, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)
	logs, err := ethclient.FilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: big.NewInt(22190258),
		ToBlock:   big.NewInt(22190258),
		Addresses: []common.Address{
			common.HexToAddress("0x2d2646b619697e3696D7413694a6a2A5f63d5a27"),
		},
	})
	require.NoError(t, err)
	s, _ := json.Marshal(logs)
	t.Log(string(s))
	cache := tracecall.NewCache(rpcnode.NewClient(zap.S(), ethclient))
	p := MustNewParserWithDeployer(cache, ethclient, common.HexToAddress(parser.Deployer0xV3))
	tradeLog, err := p.Parse(logs[0], 1000)
	require.NoError(t, err)
	s, _ = json.Marshal(tradeLog)
	t.Log(string(s))
}
