package cowprotocol // nolint: testpackage

import (
	"context"
	"encoding/json"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

var rpcURL = os.Getenv("TEST_RPC_URL")

func TestFetchEvent(t *testing.T) {
	t.Skip("Need to add the rpc url that enables the trace call JSON-RPC")
	ethClient, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)
	p, err := MustNewParser()
	require.NoError(t, err)
	require.Equal(t, p.abi.Events[TradeEvent].ID, common.HexToHash("0xa07a543ab8a018198e99ca0184c93fe9050a79400a0a723441f84de1d972cc17"))
	logs, err := ethClient.FilterLogs(context.Background(), ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: big.NewInt(22185932),
		ToBlock:   big.NewInt(22185932),
		Addresses: nil,
		Topics: [][]common.Hash{
			{
				p.abi.Events[TradeEvent].ID,
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

	eventRaw := `{"address":"0x9008d19f58aabd9ed0d60971565aa8510560ab41","topics":["0xa07a543ab8a018198e99ca0184c93fe9050a79400a0a723441f84de1d972cc17","0x0000000000000000000000000cde8a2ec22ff9073dce5fcf35dafe0963afff8d"],"data":"0x000000000000000000000000d1d2eb1b1e90b638588728b4130137d262c87cae000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48000000000000000000000000000000000000000000000000000051dac207a000000000000000000000000000000000000000000000000000000000032cfab47b000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c00000000000000000000000000000000000000000000000000000000000000038d3ad598dc57f26efad0ae780014475aeded181245a35c33dce21ae7dbcb405d30cde8a2ec22ff9073dce5fcf35dafe0963afff8d67ee12000000000000000000","blockNumber":"0x15287cc","transactionHash":"0x71bea1574fcd04ad16ec5223a66fa00ecfcc759fa32d0ac5ee4b176b0f6ad909","transactionIndex":"0x19","blockHash":"0x247bd9fd838c3a7ca6acbd83ed4f118dfff714320b0021cb07e240e91442b5fe","logIndex":"0x28","removed":false}`
	event := types.Log{}
	err := json.Unmarshal([]byte(eventRaw), &event)
	require.NoError(t, err)

	ethClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		panic(err)
	}
	block, err := ethClient.BlockByHash(context.Background(), event.BlockHash)
	require.NoError(t, err)
	p, err := MustNewParser()
	require.NoError(t, err)

	log, err := p.Parse(event, block.Time())
	require.NoError(t, err)
	require.Equal(t, log.EventHash, p.tradeEventHash)
	t.Log("owner", log.Owner)
	t.Log("sell token", log.SellToken)
	t.Log("sell amount", log.SellAmount)
	t.Log("buy token", log.BuyToken)
	t.Log("buy amount", log.BuyAmount)
	t.Log("contract", log.ContractAddress)
	t.Log("------------------------")
}
