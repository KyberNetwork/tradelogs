package zxrfqv3

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/KyberNetwork/tradelogs/pkg/rpcnode"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"github.com/KyberNetwork/tradelogs/pkg/tracecall"
	"github.com/KyberNetwork/tradelogs/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	rpcURL         = ""
	fallbackRPCURL = ""
)

func newParserTest(t *testing.T, contractABI ContractABI, needRpc bool) *Parser {
	var cache *tracecall.Cache
	if needRpc {
		ethClient, err := ethclient.Dial(rpcURL)
		if err != nil {
			panic(err)
		}
		fallbackClient, err := ethclient.Dial(fallbackRPCURL)
		if err != nil {
			panic(err)
		}
		cache = tracecall.NewCache(rpcnode.NewClient(ethClient), rpcnode.NewClient(fallbackClient))
	}
	return MustNewParser(cache, contractABI)
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
	parser := newParserTest(t, ContractABI{
		Address:      contractAddress,
		ContractType: DevContract,
	}, false)
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
	parser := newParserTest(t, ContractABI{
		Address:      contractAddress,
		ContractType: DevContract,
	}, true)
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
			//x, _ := json.MarshalIndent(tradeLog, "", "\t")
			//fmt.Println(string(x))
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
