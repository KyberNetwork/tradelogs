package uniswapx // nolint: testpackage

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/KyberNetwork/tradelogs/pkg/rpcnode"
	"github.com/KyberNetwork/tradelogs/pkg/tracecall"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

const rpcURL = ""

func TestFetchEvent(t *testing.T) {
	t.Skip("Need to add the rpc url that enables the trace call JSON-RPC")
	ethClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		panic(err)
	}
	traceCalls := tracecall.NewCache(rpcnode.NewClient(zap.S(), ethClient))
	p := MustNewParser(traceCalls)
	require.Equal(t, p.abi.Events[FilledEvent].ID, common.HexToHash("0x78ad7ec0e9f89e74012afa58738b6b661c024cb0fd185ee2f616c0a28924bd66"))
	logs, err := ethClient.FilterLogs(context.Background(), ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: big.NewInt(20878510),
		ToBlock:   big.NewInt(20878510),
		Addresses: nil,
		Topics: [][]common.Hash{
			{
				p.abi.Events[FilledEvent].ID,
			},
		},
	})
	require.NoError(t, err)
	d, err := json.Marshal(logs)
	require.NoError(t, err)
	t.Log(string(d))
}

func TestParseEvent(t *testing.T) {
	t.Skip("Need to add the rpc url that enables the trace call JSON-RPC")
	eventRaw := `{"address":"0x00000011f84b9aa48e5f8aa8b9897600006289be","topics":["0x78ad7ec0e9f89e74012afa58738b6b661c024cb0fd185ee2f616c0a28924bd66","0xa09fd699de09d6492bd94a552d1411802fcce905984e90fe5720039931546cfc","0x000000000000000000000000807cf9a772d5a3f9cefbc1192e939d62f0d9bd38","0x000000000000000000000000cf2828bb821e935292f1d3ff7f996e19efe4910c"],"data":"0x046832975f1fff57796138f8a594ca97d887a6d4fef01f55f262cc210b31ba04","blockNumber":"0x13e94ae","transactionHash":"0xa0f1fb44e76f666e79bc4b290f226dac851e60c8eb57226012167b2cd313cc98","transactionIndex":"0x34","blockHash":"0xea65325c25577b35c89d5b7fee914f16ae9e3a4683fe7abf3ec5eeba0f17e128","logIndex":"0x11f","removed":false}`
	event := types.Log{}
	err := json.Unmarshal([]byte(eventRaw), &event)
	require.NoError(t, err)
	ethClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		panic(err)
	}
	block, err := ethClient.BlockByHash(context.Background(), event.BlockHash)
	require.NoError(t, err)
	traceCalls := tracecall.NewCache(rpcnode.NewClient(zap.S(), ethClient))
	p := MustNewParser(traceCalls)
	log, err := p.Parse(event, block.Time())
	require.NoError(t, err)
	fmt.Printf("%+v\n", log)
	require.Equal(t, log.EventHash, p.eventHash)
	require.Equal(t, log.Maker, "0x807cF9A772d5a3f9CeFBc1192e939D62f0D9bD38")
	require.Equal(t, log.Taker, "0xcF2828bb821E935292f1D3Ff7F996e19eFe4910c")
	require.Equal(t, log.MakerTokenAmount, "10195811258")
	require.Equal(t, log.TakerTokenAmount, "10200030000")
	require.Equal(t, log.Expiry, uint64(0x66fd5e45))
	t.Log(log)
}

func TestParseBatchEvent(t *testing.T) {
	t.Skip("Need to add the rpc url that enables the trace call JSON-RPC")
	eventRaw := `[{"address":"0x00000011f84b9aa48e5f8aa8b9897600006289be","topics":["0x78ad7ec0e9f89e74012afa58738b6b661c024cb0fd185ee2f616c0a28924bd66","0x1d62255d2fd90374be846f2114e2540a8b80242c4f849fd6dfb58a316479e9ca","0x00000000000000000000000043ee35d5542f5b826fa92832bd97caff153675a8","0x0000000000000000000000002d2e4848fa164911f935e713c58a3ce251f95ddf"],"data":"0x046832f5aa4b6a391a75b577e69c7409cf31d676a3983732f7cce4a8100f4001","blockNumber":"0x1414756","transactionHash":"0xe2530de46bf952c22a0c094586b6c2408a6b4900f07488f838a76874bde6cbf7","transactionIndex":"0x15","blockHash":"0x81279855a5a4d20e4369dc27928e08d8076ff342c4af02710862672c853659e4","logIndex":"0xa7","removed":false},{"address":"0x00000011f84b9aa48e5f8aa8b9897600006289be","topics":["0x78ad7ec0e9f89e74012afa58738b6b661c024cb0fd185ee2f616c0a28924bd66","0x7a70f8eb7d8a241a9016a9bee57b5166fa1aab9bdfe22bdf2a08d32b0ac4dd13","0x00000000000000000000000043ee35d5542f5b826fa92832bd97caff153675a8","0x00000000000000000000000085e3d047e92feb7879c302b92b689d7538b34691"],"data":"0x0468321cf5cd8349ef1d43cbdc795973c2c38a7392e7bc6e246e42051702840f","blockNumber":"0x1414756","transactionHash":"0xe2530de46bf952c22a0c094586b6c2408a6b4900f07488f838a76874bde6cbf7","transactionIndex":"0x15","blockHash":"0x81279855a5a4d20e4369dc27928e08d8076ff342c4af02710862672c853659e4","logIndex":"0xa8","removed":false}]`
	var events []types.Log
	err := json.Unmarshal([]byte(eventRaw), &events)
	require.NoError(t, err)
	ethClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		panic(err)
	}
	block, err := ethClient.BlockByHash(context.Background(), events[0].BlockHash)
	require.NoError(t, err)
	traceCalls := tracecall.NewCache(rpcnode.NewClient(zap.S(), ethClient))
	p := MustNewParser(traceCalls)
	for _, e := range events {
		log, err := p.Parse(e, block.Time())
		require.NoError(t, err)
		require.Equal(t, log.EventHash, p.eventHash)
		t.Log("maker", log.Maker)
		t.Log("maker_token", log.MakerToken)
		t.Log("taker", log.Taker)
		t.Log("taker_token", log.TakerToken)
		t.Log("maker_amount", log.MakerTokenAmount)
		t.Log("taker_amount", log.TakerTokenAmount)
		t.Log("------------------------")
	}
}
