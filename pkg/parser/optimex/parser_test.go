package optimex

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"testing"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/test-go/testify/require"
)

var rpcURL = ""

func TestFetchEventLog(t *testing.T) {
	t.Skip()
	client, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)
	p := MustNewParser(client)
	require.Equal(t, p.abi.Events[TradeEvent].ID, common.HexToHash("0x997438f726850dc0f689a2037c76da85881ee939121c2dce2ef362d02850f758"))
	logs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: big.NewInt(2896920),
		ToBlock:   big.NewInt(2896920),
		Addresses: nil,
		Topics:    [][]common.Hash{{p.abi.Events[TradeEvent].ID}},
	})
	require.NoError(t, err)
	d, _ := json.Marshal(logs)
	t.Log(string(d))
}

func TestParseEventLog(t *testing.T) {
	t.Skip()
	client, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)
	eventRaw := `{"address":"0x272599ce3602a49b580a5c4a4d3c1067e30248d2","topics":["0x997438f726850dc0f689a2037c76da85881ee939121c2dce2ef362d02850f758","0x000000000000000000000000903f19a2796691f1abf21d27ae6c17e125b602d2","0xade2c870ece6d5f872ffe8d3d64ada5384fd159a1ad3b5fb82d4cba56e3c3be4"],"data":"0x","blockNumber":"0x295e96","transactionHash":"0xe7c2ac715b016730fbf5d5e2f1eb7df6222d730e99199ceeabaef5a535ac5847","transactionIndex":"0x1","blockHash":"0x39c03d9d0229b4de1f37cc9e6a97b490ea318758569346394e250c6f8b79d58f","logIndex":"0x0","removed":false}`
	event := types.Log{}
	err = json.Unmarshal([]byte(eventRaw), &event)
	require.NoError(t, err)
	p := MustNewParser(client)
	log, err := p.Parse(event, uint64(time.Now().Unix()))
	require.NoError(t, err)
	fmt.Printf("%+v", log)
}
