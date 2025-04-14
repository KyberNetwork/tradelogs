package cowtransferparser // nolint: testpackage

import (
	"context"
	"encoding/json"
	"math/big"
	"os"
	"testing"

	ethereum "github.com/ethereum/go-ethereum"
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
	require.Equal(t, p.abi.Events[TransferEvent].ID, common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"))
	logs, err := ethClient.FilterLogs(context.Background(), ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: big.NewInt(22185932),
		ToBlock:   big.NewInt(22185932),
		Addresses: nil,
		Topics: [][]common.Hash{
			{
				p.abi.Events[TransferEvent].ID,
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
	eventRaw := `{"address":"0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2","topics":["0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef","0x000000000000000000000000dfea03ed1ec53116cd6ea45373cbb6e8823ef59d","0x00000000000000000000000081463b0f960f247f704377661ec81c1fd65b5128"],"data":"0x0000000000000000000000000000000000000000000000000003efcf27000000","blockNumber":"0x15287cc","transactionHash":"0x7818931fe80b31b4324313636498bb3d5ee1c56a91d7b8b3c58199da0c645083","transactionIndex":"0x122","blockHash":"0x247bd9fd838c3a7ca6acbd83ed4f118dfff714320b0021cb07e240e91442b5fe","logIndex":"0x3f7","removed":false}`
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
	t.Log("block number", log.BlockNumber)
	t.Log("time", log.Timestamp)
	t.Log("token", log.Token)
	t.Log("from", log.FromAddress)
	t.Log("to", log.ToAddress)
	t.Log("amount", log.Amount)
	t.Log("------------------------")
}
