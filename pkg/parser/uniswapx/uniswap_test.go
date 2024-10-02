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
		FromBlock: big.NewInt(20655762),
		ToBlock:   big.NewInt(20655762),
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
	eventRaw := `{"address":"0x00000011f84b9aa48e5f8aa8b9897600006289be","topics":["0x78ad7ec0e9f89e74012afa58738b6b661c024cb0fd185ee2f616c0a28924bd66","0x9e57c601cc099a21f2adebee775f5c87197f64449d8dc14f4daa363a0acccd82","0x000000000000000000000000807cf9a772d5a3f9cefbc1192e939d62f0d9bd38","0x00000000000000000000000097e1e6c70fe9f2209792640336534bfc9f93aa7a"],"data":"0x046832e0bb8b67660fa0dab20dd6c717f3c188907d1b33cf6bb432d0434b4631","blockNumber":"0x13b2e92","transactionHash":"0x82f36abfb0f492ce6f6363cd9f233b2a86626234bfac5b1efe54e1c25a6fc410","transactionIndex":"0x1e","blockHash":"0xe6877169821ccccaf15256d0f802b4de56f11d6ca86783eb29a2530cc19fcdbe","logIndex":"0xc3","removed":false}`
	event := types.Log{}
	err := json.Unmarshal([]byte(eventRaw), &event)
	require.NoError(t, err)
	ethClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		panic(err)
	}
	traceCalls := tracecall.NewCache(rpcnode.NewClient(zap.S(), ethClient))
	p := MustNewParser(traceCalls)
	log, err := p.Parse(event, 1713889895)
	require.NoError(t, err)
	fmt.Printf("%+v\n", log)
	require.Equal(t, log.EventHash, p.eventHash)
	require.Equal(t, log.Maker, "0x807cF9A772d5a3f9CeFBc1192e939D62f0D9bD38")
	require.Equal(t, log.Taker, "0x97e1e6C70fe9F2209792640336534bFc9F93AA7A")
	require.Equal(t, log.MakerTokenAmount, "197069169177783279401")
	require.Equal(t, log.TakerTokenAmount, "10000000000000000000")
	require.Equal(t, log.Expiry, uint64(0x66d46102))
	t.Log(log)
}
