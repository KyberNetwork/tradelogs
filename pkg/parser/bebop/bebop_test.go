package bebop

import (
	"context"
	"encoding/json"
	"math/big"
	"net/http"
	"testing"
	"time"

	"github.com/KyberNetwork/tradelogs/pkg/rpcnode"
	"github.com/KyberNetwork/tradelogs/pkg/tracecall"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/test-go/testify/require"
)

const rpcURL = ""

func TestFetchEvent(t *testing.T) {
	t.Skip("Need to add the rpc url that enables the trace call JSON-RPC")
	rpcClient, err := rpcnode.NewClient(http.DefaultClient, rpcURL)
	if err != nil {
		panic(err)
	}
	traceCalls := tracecall.NewCache(rpcClient)
	p := MustNewParser(traceCalls)
	require.Equal(t, p.abi.Events[TradeEvent].ID, common.HexToHash("0xadd7095becdaa725f0f33243630938c861b0bba83dfd217d4055701aa768ec2e"))
	client, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)
	logs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: big.NewInt(20025442),
		ToBlock:   big.NewInt(20025442),
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

func TestParseEvent(t *testing.T) {
	t.Skip("Need to add the rpc url that enables the trace call JSON-RPC")
	eventRaw := `[{"address":"0xbbbbbbb520d69a9775e85b458c58c648259fad5f","topics":["0xadd7095becdaa725f0f33243630938c861b0bba83dfd217d4055701aa768ec2e","0x00000000000000000000000000000000b05c492e829fd9c00000000000000000"],"data":"0x","blockNumber":"0x1319062","transactionHash":"0x0f7567ffd5fc92aa552c5e6ebf8173514693aadd5a888c816b4735a931e8c8c3","transactionIndex":"0xe","blockHash":"0x3015d9d1960cab845f333c49f8cb0d81fa89b2c837de8ac22cd645810bb18c83","logIndex":"0x5a","removed":false}]`
	events := []types.Log{}
	err := json.Unmarshal([]byte(eventRaw), &events)
	require.NoError(t, err)
	rpcClient, err := rpcnode.NewClient(http.DefaultClient, rpcURL)
	if err != nil {
		panic(err)
	}
	traceCalls := tracecall.NewCache(rpcClient)
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
