package bebop

import (
	"context"
	"encoding/json"
	"math/big"
	"testing"
	"time"

	"github.com/KyberNetwork/tradelogs/pkg/rpcnode"
	"github.com/KyberNetwork/tradelogs/pkg/tracecall"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/test-go/testify/require"
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
	require.Equal(t, p.abi.Events[TradeEvent].ID, common.HexToHash("0xadd7095becdaa725f0f33243630938c861b0bba83dfd217d4055701aa768ec2e"))
	client, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)
	logs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: big.NewInt(19932179),
		ToBlock:   big.NewInt(19932179),
		Addresses: nil,
		Topics: [][]common.Hash{
			{
				common.HexToHash("0xadd7095becdaa725f0f33243630938c861b0bba83dfd217d4055701aa768ec2e"),
			},
		},
	})
	require.NoError(t, err)
	d, err := json.Marshal(logs)
	require.NoError(t, err)
	t.Log(string(d))
}

func TestParseAggregateOrderEvent(t *testing.T) {
	t.Skip("Need to add the rpc url that enables the trace call JSON-RPC")
	eventRaw := `{"address":"0xbbbbbbb520d69a9775e85b458c58c648259fad5f","topics":["0xadd7095becdaa725f0f33243630938c861b0bba83dfd217d4055701aa768ec2e","0x000000000000000000000000000000006f0760aefdfe33400000000000000000"],"data":"0x","blockNumber":"0x138de04","transactionHash":"0xb9d13c7057d6d0779a14023d8ce469a8266ea89eb2ca4e6f4c085ef0929c6f08","transactionIndex":"0x8","blockHash":"0x54a780bfa027b8cfb16a40bb86be066795e84d6a124582f9a85b8c662f27fda9","logIndex":"0x46","removed":false}`
	events := types.Log{}
	err := json.Unmarshal([]byte(eventRaw), &events)
	require.NoError(t, err)
	ethClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		panic(err)
	}
	traceCalls := tracecall.NewCache(rpcnode.NewClient(zap.S(), ethClient))
	p := MustNewParser(traceCalls)
	log, err := p.Parse(events, uint64(time.Now().Unix()))
	require.NoError(t, err)
	require.Equal(t, log.EventHash, p.eventHash)
	t.Log(log.Maker)
	t.Log(log.MakerToken)
	t.Log(log.MakerTokenAmount)
	t.Log(log.Taker)
	t.Log(log.TakerToken)
	t.Log(log.TakerTokenAmount)
}

func TestParseMultiOrderEvent(t *testing.T) {
	t.Skip("Need to add the rpc url that enables the trace call JSON-RPC")
	eventRaw := `{"address":"0xbbbbbbb520d69a9775e85b458c58c648259fad5f","topics":["0xadd7095becdaa725f0f33243630938c861b0bba83dfd217d4055701aa768ec2e","0x0000000000000000000000000000000027611cb25303fa700000000000000001"],"data":"0x","blockNumber":"0x133842c","transactionHash":"0xff46ac555ec7da7aa484864dc0df90217b7f46dfa51627c20ef6e25451c64b15","transactionIndex":"0x10","blockHash":"0xb03cd01d48456c60cdf445e35b0a4d7672b80d801245e516336596f0d5b77951","logIndex":"0x85","removed":false}`
	events := types.Log{}
	err := json.Unmarshal([]byte(eventRaw), &events)
	require.NoError(t, err)
	ethClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		panic(err)
	}
	traceCalls := tracecall.NewCache(rpcnode.NewClient(zap.S(), ethClient))
	p := MustNewParser(traceCalls)
	log, err := p.Parse(events, uint64(time.Now().Unix()))
	require.NoError(t, err)
	require.Equal(t, log.EventHash, p.eventHash)
	t.Log(log.Maker)
	t.Log(log.MakerToken)
	t.Log(log.MakerTokenAmount)
	t.Log(log.Taker)
	t.Log(log.TakerToken)
	t.Log(log.TakerTokenAmount)
}

func TestParseSingleOrderEvent(t *testing.T) {
	t.Skip("Need to add the rpc url that enables the trace call JSON-RPC")
	eventRaw := `{"address":"0xbbbbbbb520d69a9775e85b458c58c648259fad5f","topics":["0xadd7095becdaa725f0f33243630938c861b0bba83dfd217d4055701aa768ec2e","0x0000000000000000000000000000000084e2ab041ba07e400000000000000000"],"data":"0x","blockNumber":"0x1302413","transactionHash":"0xa227c4aed6f3ae86cd8ba02349d0ae5be78337266416ac65cf8a8d095a2f572e","transactionIndex":"0x47","blockHash":"0xbb4a5ae13e2a66f83f444f24078ac3fb6d0b09f610c48d920d3b89d4f20ea86c","logIndex":"0xd8","removed":false}`
	events := types.Log{}
	err := json.Unmarshal([]byte(eventRaw), &events)
	require.NoError(t, err)
	ethClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		panic(err)
	}
	traceCalls := tracecall.NewCache(rpcnode.NewClient(zap.S(), ethClient))
	p := MustNewParser(traceCalls)
	log, err := p.Parse(events, uint64(time.Now().Unix()))
	require.NoError(t, err)
	require.Equal(t, log.EventHash, p.eventHash)
	t.Log(log.Maker)
	t.Log(log.MakerToken)
	t.Log(log.MakerTokenAmount)
	t.Log(log.Taker)
	t.Log(log.TakerToken)
	t.Log(log.TakerTokenAmount)
}
