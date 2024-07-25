package zxrfqv3

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"github.com/KyberNetwork/tradelogs/pkg/decoder"
	"github.com/KyberNetwork/tradelogs/pkg/parser/zxrfqv3/helper"
	"github.com/KyberNetwork/tradelogs/pkg/rpcnode"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"github.com/KyberNetwork/tradelogs/pkg/tracecall"
	"github.com/KyberNetwork/tradelogs/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

const rpcURL = ""

type parserType int

const (
	dev parserType = iota
	swap
	gasless
)

func newParserTest(t *testing.T, contractAddress common.Address, needRpc bool, parserType parserType) *Parser {
	var cache *tracecall.Cache
	if needRpc {
		rpcClient, err := rpcnode.NewClient(http.DefaultClient, rpcURL)
		if err != nil {
			t.Fatal(err)
		}
		cache = tracecall.NewCache(rpcClient)
	}
	switch parserType {
	case dev:
		return MustNewDevParser(cache, contractAddress)
	case swap:
		return MustNewSwapParser(cache, contractAddress)
	case gasless:
		return MustNewGaslessParser(cache, contractAddress)
	default:
		return MustNewDevParser(cache, contractAddress)
	}
}

func TestInitNewParser(t *testing.T) {
	// no fatal
	address := common.HexToAddress("0x7966aF62034313D87Ede39380bf60f1A84c62BE7")
	_ = MustNewDevParser(nil, address)
	_ = MustNewGaslessParser(nil, address)
	_ = MustNewSwapParser(nil, address)
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
	var expectedInput helper.InputParamOfFillRfqOrderSelfFunded
	getTestCaseData(t, "./test/expected_input_rfq.json", &expectedInput)

	contractAddress := common.HexToAddress("0x7966aF62034313D87Ede39380bf60f1A84c62BE7")
	parser := newParserTest(t, contractAddress, false, dev)
	data, err := parser.getExecuteActionData(callFrame)
	assert.NoError(t, err, "failed to get execute action data")
	actionType, rawData, err := decodeCall(data[1])
	assert.NoError(t, err, "failed to decode call")
	_, ok := selectorAction[actionType]
	assert.True(t, ok, "action type not found")
	byteMethodId, err := hex.DecodeString(MethodIdDecodeParamOfFillOrderSelfFundedHex)
	assert.NoError(t, err, "failed to decode method id")
	methodId, err := decoder.GetBytes4(byteMethodId)
	assert.NoError(t, err, "failed to get method id")
	input, err := helper.GetInputParamsOfFillRfqOrderSelfFunded(parser.customAbi, methodId, rawData)
	assert.NoError(t, err, "failed to decode data")

	assert.Equal(t, expectedInput.Recipient, input.Recipient, "recipient not match")
	assert.Equal(t, expectedInput.Permit.Permitted.Token, input.Permit.Permitted.Token, "permitted token not match")
	assert.Equal(t, expectedInput.Permit.Permitted.Amount, input.Permit.Permitted.Amount, "permitted amount not match")
	assert.Equal(t, expectedInput.Permit.Deadline, input.Permit.Deadline, "deadline not match")
	assert.Equal(t, expectedInput.Permit.Nonce, input.Permit.Nonce, "nonce not match")
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
	parser := newParserTest(t, contractAddress, true, dev)
	ethClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		t.Fatalf("failed to dial to rpc url: %s, err: %s", rpcURL, err)
	}

	txReceipt, err := ethClient.TransactionReceipt(context.Background(), txHash)
	assert.NoError(t, err, "failed to get tx receipt")
	rfqCount := 0
	for _, log := range txReceipt.Logs {
		if log.Address == parser.contractAddress && len(log.Topics) == 0 {
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
