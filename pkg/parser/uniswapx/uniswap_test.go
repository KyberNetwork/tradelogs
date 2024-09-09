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
		FromBlock: big.NewInt(19719300),
		ToBlock:   big.NewInt(19719300),
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
	eventRaw := `{"address":"0x3867393cc6ea7b0414c2c3e1d9fe7cea987fd066","topics":["0x78ad7ec0e9f89e74012afa58738b6b661c024cb0fd185ee2f616c0a28924bd66","0x510fd6dd82e657a1bc4e007ee563925923c5896fecc6996e491adcaff6c8a528","0x000000000000000000000000ff8ba4d1fc3762f6154cc942ccf30049a2a0cec6","0x000000000000000000000000250a94c03b9b57c93cc5549760d59d6eacfb136d"],"data":"0x04683239f5ccb91d719180e8d856523bf095571534b3cd850e678dd61919c153","blockNumber":"0x12ce484","transactionHash":"0xcbd70c12e81d3521ca0e96dc887ae6653063f75dbdcd943a5ee406fa30446619","transactionIndex":"0x81","blockHash":"0x32fc731de2f9509f6c829995e31b843c28594b866b39f9009cac65dbb8bee173","logIndex":"0x10c","removed":false}`
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
	require.Equal(t, log.Maker, "0xff8Ba4D1fC3762f6154cc942CCF30049A2A0cEC6")
	require.Equal(t, log.Taker, "0x250A94C03b9b57C93CC5549760D59d6eAcfB136d")
	require.Equal(t, log.MakerTokenAmount, "42282994361466557")
	require.Equal(t, log.TakerTokenAmount, "146889265")
	require.Equal(t, log.Expiry, uint64(1713890185))
	t.Log(log)
}
