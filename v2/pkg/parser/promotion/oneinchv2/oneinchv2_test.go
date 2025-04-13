package oneinchv2

import (
	"context"
	"encoding/json"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

var rpcURL = os.Getenv("TEST_RPC_URL")

func TestFetchEvent(t *testing.T) {
	t.Skip()
	p := MustNewParser()
	require.Equal(t, p.abi.Events[PromotionEvent].ID, common.HexToHash("0xb863cf86b291171e4b0332ea12b59af030f98a2c74a6d51effaf1109ae4c7f1e"))
	client, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)
	logs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: big.NewInt(20583304),
		ToBlock:   big.NewInt(20583304),
		Addresses: nil,
		Topics:    [][]common.Hash{{common.HexToHash("0xb863cf86b291171e4b0332ea12b59af030f98a2c74a6d51effaf1109ae4c7f1e")}},
	})
	require.NoError(t, err)
	d, _ := json.Marshal(logs)
	t.Log(string(d))
}

func TestParseEvent(t *testing.T) {
	t.Skip("Need to add the rpc url that enables the trace call JSON-RPC")
	eventRaw := `{"address":"0xf55684bc536487394b423e70567413fab8e45e26",
		"topics": [
			"0xb863cf86b291171e4b0332ea12b59af030f98a2c74a6d51effaf1109ae4c7f1e"
		],
		"data":"0x000000000000000000000000dcdf16a03360d4971ca4c1fd9967a47125f3c995000000000000000000000000000000000000000000000000000000000000a4b1000000000000000000000000a8c1c98aaf99a5dfc907d61b892b2ad624901185",
		"blockNumber":"0x13a1388",
		"transactionHash":"0x0cfa3b71536bff882a99cc5896d5653b1f7438e313d5b91e9ec0ba74fa9b1957",
		"transactionIndex":"0xac",
		"blockHash":"0x016715ee70579c2023b59a14eb3b8911ff7bbb40198e0c0273e54470164bbff4",
		"logIndex":"0x18d",
		"removed":false
	}`
	event := types.Log{}
	err := json.Unmarshal([]byte(eventRaw), &event)
	require.NoError(t, err)
	p := MustNewParser()
	log, err := p.Parse(event, uint64(time.Now().Unix()))
	require.NoError(t, err)
	require.Equal(t, "0xb863cf86b291171e4b0332ea12b59af030f98a2c74a6d51effaf1109ae4c7f1e", p.eventHash)
	t.Log("event hash", p.eventHash)
	t.Log("promoter", log.Promoter)
	t.Log("promotee", log.Promotee)
	t.Log("chain id", log.ChainId)
	// t.Log(log)
}
