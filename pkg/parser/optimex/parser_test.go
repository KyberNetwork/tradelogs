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
	require.Equal(t, p.abi.Events[SelectEvent].ID, common.HexToHash("0x997438f726850dc0f689a2037c76da85881ee939121c2dce2ef362d02850f758"))
	logs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: big.NewInt(2896920),
		ToBlock:   big.NewInt(2896920),
		Topics:    [][]common.Hash{{p.abi.Events[SelectEvent].ID}},
	})
	require.NoError(t, err)
	d, _ := json.Marshal(logs)
	t.Log(string(d))
	require.Equal(t, p.abi.Events[SettleEvent].ID, common.HexToHash("0x6292bb735dd6c41d01a708128a6cb40de00a8864b9208926d9fc1540497c90e6"))
	logs, err = client.FilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: big.NewInt(2896920),
		ToBlock:   big.NewInt(2897314),
		Topics:    [][]common.Hash{{p.abi.Events[SettleEvent].ID}},
	})
	require.NoError(t, err)
	d, _ = json.Marshal(logs)
	t.Log(string(d))
}

func TestParseSelectEventLog(t *testing.T) {
	t.Skip()
	client, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)
	eventRaw := `{"address":"0x272599ce3602a49b580a5c4a4d3c1067e30248d2","topics":["0x997438f726850dc0f689a2037c76da85881ee939121c2dce2ef362d02850f758","0x000000000000000000000000903f19a2796691f1abf21d27ae6c17e125b602d2","0xade2c870ece6d5f872ffe8d3d64ada5384fd159a1ad3b5fb82d4cba56e3c3be4"],"data":"0x","blockNumber":"0x295e96","transactionHash":"0xe7c2ac715b016730fbf5d5e2f1eb7df6222d730e99199ceeabaef5a535ac5847","transactionIndex":"0x1","blockHash":"0x39c03d9d0229b4de1f37cc9e6a97b490ea318758569346394e250c6f8b79d58f","logIndex":"0x0","removed":false}`
	event := types.Log{}
	err = json.Unmarshal([]byte(eventRaw), &event)
	require.NoError(t, err)
	p := MustNewParser(client)
	name, ok := p.LogFromExchange(event)
	require.Equal(t, true, ok)
	require.Equal(t, SelectEvent, name)
	log, err := p.ParseSelected(event, uint64(time.Now().Unix()))
	require.NoError(t, err)
	fmt.Printf("%+v", log)
}

func TestParseSettleEventLog(t *testing.T) {
	t.Skip()
	client, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)
	eventRaw := `{"address":"0x272599ce3602a49b580a5c4a4d3c1067e30248d2","topics":["0x6292bb735dd6c41d01a708128a6cb40de00a8864b9208926d9fc1540497c90e6","0x0000000000000000000000009d946bb9ef2a8f384d7f29159472670070926c99","0x2fa874118633aa8f373b7efc6159eb6d0ffad2ac6da937b6536e56eed78fdcfd"],"data":"0x","blockNumber":"0x2c35a2","transactionHash":"0x95a022d54cea0fdf2348b33e3233c276580001ff56688c999397e615ffb989f5","transactionIndex":"0x1","blockHash":"0x7ae81a19951e959a6c43f3804105b175eb6ec15ddf282de3fdc071c26bea65d1","logIndex":"0x0","removed":false}`
	event := types.Log{}
	err = json.Unmarshal([]byte(eventRaw), &event)
	require.NoError(t, err)
	p := MustNewParser(client)
	name, ok := p.LogFromExchange(event)
	require.Equal(t, true, ok)
	require.Equal(t, SettleEvent, name)
	trade, err := p.ParseSettle(event, uint64(time.Now().Unix()))
	require.NoError(t, err)
	fmt.Printf("%+v", trade)
}
