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
)

const (
	rpcURL         = ""
	fallbackRPCURL = ""
)

func TestFetchEvent(t *testing.T) {
	t.Skip("Need to add the rpc url that enables the trace call JSON-RPC")
	ethClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		panic(err)
	}
	fallbackClient, err := ethclient.Dial(fallbackRPCURL)
	if err != nil {
		panic(err)
	}
	traceCalls := tracecall.NewCache(rpcnode.NewClient(ethClient), rpcnode.NewClient(fallbackClient))
	p := MustNewParser(traceCalls)
	require.Equal(t, p.abi.Events[FilledEvent].ID, common.HexToHash("0xfec331350fce78ba658e082a71da20ac9f8d798a99b3c79681c8440cbfe77e07"))
	client, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)
	logs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: big.NewInt(19517464),
		ToBlock:   big.NewInt(19517464),
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
	eventRaw := `[{"address":"0x111111125421ca6dc452d289314280a0f8842a65","topics":["0xfec331350fce78ba658e082a71da20ac9f8d798a99b3c79681c8440cbfe77e07"],"data":"0x2c680fc902966ed935ddf9dc75550a32f1b2d1fba7dcaa7e4ab3b541734acddb00000000000000000000000000000000000000000000021246e3d52a827b39b7","blockNumber":"0x129d018","transactionHash":"0x4fa77df92aa4726ddbf54a41f72ca78a57035793c50e86643a25a058a27e35f4","transactionIndex":"0x6","blockHash":"0xa8eaa16d7ca8d934343bfe5fbb14bfeab272143f65fbd51f9caecf67754cfd29","logIndex":"0x53","removed":false},{"address":"0x111111125421ca6dc452d289314280a0f8842a65","topics":["0xfec331350fce78ba658e082a71da20ac9f8d798a99b3c79681c8440cbfe77e07"],"data":"0x2c680fc902966ed935ddf9dc75550a32f1b2d1fba7dcaa7e4ab3b541734acddb00000000000000000000000000000000000000000000020c219b06bffa614286","blockNumber":"0x129d018","transactionHash":"0x4fa77df92aa4726ddbf54a41f72ca78a57035793c50e86643a25a058a27e35f4","transactionIndex":"0x6","blockHash":"0xa8eaa16d7ca8d934343bfe5fbb14bfeab272143f65fbd51f9caecf67754cfd29","logIndex":"0x60","removed":false},{"address":"0x111111125421ca6dc452d289314280a0f8842a65","topics":["0xfec331350fce78ba658e082a71da20ac9f8d798a99b3c79681c8440cbfe77e07"],"data":"0xe0f3e2f9d0dbc1d72d355d2f80328b4e432d316c91a011d2fa26b56b7a39c371000000000000000000000000000000000000000000000000e1394e6d5fa7bd71","blockNumber":"0x129d018","transactionHash":"0x4fa77df92aa4726ddbf54a41f72ca78a57035793c50e86643a25a058a27e35f4","transactionIndex":"0x6","blockHash":"0xa8eaa16d7ca8d934343bfe5fbb14bfeab272143f65fbd51f9caecf67754cfd29","logIndex":"0x6a","removed":false}]`
	events := []types.Log{}
	err := json.Unmarshal([]byte(eventRaw), &events)
	require.NoError(t, err)
	ethClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		panic(err)
	}
	fallbackClient, err := ethclient.Dial(fallbackRPCURL)
	if err != nil {
		panic(err)
	}
	traceCalls := tracecall.NewCache(rpcnode.NewClient(ethClient), rpcnode.NewClient(fallbackClient))
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
