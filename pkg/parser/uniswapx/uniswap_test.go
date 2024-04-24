package uniswapx // nolint: testpackage

import (
	"context"
	"encoding/json"
	"math/big"
	"net/http"
	"testing"

	"github.com/KyberNetwork/tradelogs/pkg/rpcnode"
	"github.com/KyberNetwork/tradelogs/pkg/tracecall"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
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
	require.Equal(t, p.abi.Events[FilledEvent].ID, common.HexToHash("0x78ad7ec0e9f89e74012afa58738b6b661c024cb0fd185ee2f616c0a28924bd66"))
	client, err := ethclient.Dial("https://ethereum.kyberengineering.io")
	require.NoError(t, err)
	logs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: big.NewInt(19332438),
		ToBlock:   big.NewInt(19332438),
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
	eventRaw := `{"address":"0x6000da47483062a0d734ba3dc7576ce6a0b645c4","topics":["0x78ad7ec0e9f89e74012afa58738b6b661c024cb0fd185ee2f616c0a28924bd66","0x901ef9625044477a348a1adb59bc4e63b033795a84ebf57054af16e5592ecaf3","0x000000000000000000000000ff8ba4d1fc3762f6154cc942ccf30049a2a0cec6","0x000000000000000000000000b63d9f1780fc78d73448716b77edd4dbb626ba4c"],"data":"0x0468326e1fdbb4d4140293cdb87fedf42e6a56fcf42b6390f60d37ada9c59203","blockNumber":"0x126fd56","transactionHash":"0xd9589209564eb0a28bad891472758472ceaf343918b62a572fcb030f04e0246b","transactionIndex":"0x3c","blockHash":"0xcc969ea040fc0108c1d4e8e5d694ca388973eee4f9ba3ec500d367c8cf6c2d21","logIndex":"0x7a","removed":false}`
	event := types.Log{}
	err := json.Unmarshal([]byte(eventRaw), &event)
	require.NoError(t, err)
	rpcClient, err := rpcnode.NewClient(http.DefaultClient, rpcURL)
	if err != nil {
		panic(err)
	}
	traceCalls := tracecall.NewCache(rpcClient)
	p := MustNewParser(traceCalls)
	log, err := p.Parse(event, 1709200859)
	require.NoError(t, err)
	require.Equal(t, log.EventHash, p.eventHash)
	require.Equal(t, log.Maker, "0xff8Ba4D1fC3762f6154cc942CCF30049A2A0cEC6")
	require.Equal(t, log.Taker, "0xB63D9f1780fc78d73448716B77Edd4DbB626ba4C")
	require.Equal(t, log.MakerTokenAmount, "4405963618")
	require.Equal(t, log.TakerTokenAmount, "100000000000000000000")
	t.Log(log)
}
