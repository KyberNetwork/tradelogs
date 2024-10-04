package oneinchv6

import (
	"context"
	"encoding/json"
	"math/big"
	"testing"
	"time"

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
	require.Equal(t, p.abi.Events[FilledEvent].ID, common.HexToHash("0xfec331350fce78ba658e082a71da20ac9f8d798a99b3c79681c8440cbfe77e07"))
	client, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)
	logs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: big.NewInt(20890481),
		ToBlock:   big.NewInt(20890481),
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
	eventRaw := `[{"address":"0x111111125421ca6dc452d289314280a0f8842a65","topics":["0xfec331350fce78ba658e082a71da20ac9f8d798a99b3c79681c8440cbfe77e07"],"data":"0xb81725973077ecc6db7dd73510be46152917f31ab7d02210aace851edc0f3fbb000000000000000000000000000000000000000000000000a5953ec9be11bd73","blockNumber":"0x13ec371","transactionHash":"0x09b63a10295261192885b38af317ba138ad7fe5828ce25524e7f0f053e250458","transactionIndex":"0x7","blockHash":"0x291e59d96ccc8b661842cfd9c4dcce1d11ed30cd4cf2b8034bff2a8d7a595785","logIndex":"0x3e","removed":false}]`
	events := []types.Log{}
	err := json.Unmarshal([]byte(eventRaw), &events)
	require.NoError(t, err)
	ethClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		panic(err)
	}
	traceCalls := tracecall.NewCache(rpcnode.NewClient(zap.S(), ethClient))
	p := MustNewParser(traceCalls)
	for _, event := range events {
		log, err := p.Parse(event, uint64(time.Now().Unix()))
		require.NoError(t, err)
		require.Equal(t, log.EventHash, p.eventHash)
		t.Log(log.Maker)
		t.Log(log.MakerToken)
		t.Log(log.MakerTokenAmount)
		t.Log(log.Taker)
		t.Log(log.TakerToken)
		t.Log(log.TakerTokenAmount)
	}
}
