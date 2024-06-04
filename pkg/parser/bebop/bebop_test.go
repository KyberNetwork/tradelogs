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
		FromBlock: big.NewInt(20009519),
		ToBlock:   big.NewInt(20009519),
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
	eventRaw := `[{"address":"0xbbbbbbb520d69a9775e85b458c58c648259fad5f","topics":["0xadd7095becdaa725f0f33243630938c861b0bba83dfd217d4055701aa768ec2e","0x00000000000000000000000000000000bb6e24300f0812800000000000000000"],"data":"0x","blockNumber":"0x131522f","transactionHash":"0x1ced3e62f6e8083901b3d9365947f53edd28cf7901f79b25a9c8161f7b242b2c","transactionIndex":"0x3c","blockHash":"0xc41f37ec02be129d81f9f81527d52b37e00b2368a9ef418eda14f640662a600a","logIndex":"0xe3","removed":false}]`
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
