package uniswapxv1 // nolint: testpackage

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
)

const rpcURL = ""

func TestFetchEvent(t *testing.T) {
	t.Skip("Need to add the rpc url that enables the trace call JSON-RPC")
	ethClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		panic(err)
	}
	traceCalls := tracecall.NewCache(rpcnode.NewClient(ethClient))
	p := MustNewParser(traceCalls)
	require.Equal(t, p.abi.Events[FilledEvent].ID, common.HexToHash("0x78ad7ec0e9f89e74012afa58738b6b661c024cb0fd185ee2f616c0a28924bd66"))
	logs, err := ethClient.FilterLogs(context.Background(), ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: big.NewInt(20488655),
		ToBlock:   big.NewInt(20488655),
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
	event := types.Log{}
	eventRaw := `{"address":"0x6000da47483062a0d734ba3dc7576ce6a0b645c4","topics":["0x78ad7ec0e9f89e74012afa58738b6b661c024cb0fd185ee2f616c0a28924bd66","0xc356060ff3e60da206578d00e9cc9145cb36889fd412a7f34ddb1435e85e689e","0x000000000000000000000000def4a2438df14050d98c0e9d4d93000000890bb2","0x000000000000000000000000fe87dfcd8164adf6437edb236ab42fce19238ea6"],"data":"0x046832e66de00af695779aa4b76a837f8e31026be2ee4f4ffbdb06860ca6fd08","blockNumber":"0x138a1cf","transactionHash":"0x7a45f257e24145dd06d25dc8f77384a5e08c5c5c6220bb3f379764cfa7d30735","transactionIndex":"0x1","blockHash":"0x6bfbbcf9ec00f478c20ad6c706b0c1f9b35111f5dade8bf7ccac6671cb0d5764","logIndex":"0x8","removed":false}`
	err := json.Unmarshal([]byte(eventRaw), &event)
	require.NoError(t, err)
	ethClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		panic(err)
	}
	traceCalls := tracecall.NewCache(rpcnode.NewClient(ethClient))
	p := MustNewParser(traceCalls)
	log, err := p.Parse(event, 1713889895)
	require.NoError(t, err)
	fmt.Printf("%+v\n", log)
	require.Equal(t, log.EventHash, p.eventHash)
	require.Equal(t, log.Maker, "0xdEF4a2438DF14050D98C0e9d4D93000000890BB2")
	require.Equal(t, log.Taker, "0xfE87Dfcd8164aDF6437edb236aB42fcE19238Ea6")
	require.Equal(t, log.MakerTokenAmount, "100419300000000000")
	require.Equal(t, log.MakerToken, "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
	require.Equal(t, log.TakerTokenAmount, "30000000000000000")
	require.Equal(t, log.TakerToken, "0xEE16bd5e21cd5D27D0EAfdabbCCA0A438e97E46C")
	t.Log(log)
}
