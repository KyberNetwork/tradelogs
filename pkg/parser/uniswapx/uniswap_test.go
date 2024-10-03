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
