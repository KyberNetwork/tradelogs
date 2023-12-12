package oneinch

import (
	"context"
	"encoding/json"
	"math/big"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

func TestFetchEvent(t *testing.T) {
	t.Skip()
	p := MustNewParser(os.Getenv("RPC_URL"))
	require.Equal(t, p.abi.Events[FilledEvent].ID, common.HexToHash("0xc3b639f02b125bfa160e50739b8c44eb2d1b6908e2b6d5925c6d770f2ca78127"))
	client, err := ethclient.Dial("https://ethereum.kyberengineering.io")
	require.NoError(t, err)
	logs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: big.NewInt(18217653),
		ToBlock:   big.NewInt(18217653),
		Addresses: nil,
		Topics: [][]common.Hash{
			{
				common.HexToHash("0xc3b639f02b125bfa160e50739b8c44eb2d1b6908e2b6d5925c6d770f2ca78127"),
			},
		},
	})
	require.NoError(t, err)
	d, err := json.Marshal(logs)
	require.NoError(t, err)
	t.Log(string(d))
}

func TestParseEvent(t *testing.T) {
	eventRaw := `{"address":"0x1111111254eeb25477b68fb85ed929f73a960582",
	"topics":["0xc3b639f02b125bfa160e50739b8c44eb2d1b6908e2b6d5925c6d770f2ca78127"],
	"data":"0x59c1fba2238e5e61f38c5860e11430b52c4dc7eb787b84dba76a8aff988ff7660000000000000000000000000000000000000000000000000000000015ebf670",
	"blockNumber":"0x115fab5",
	"transactionHash":"0x17fe56f39bb51880a61570f14c8a55e7f7658054327cd7a1936d50eb633aaf5d",
	"transactionIndex":"0x67",
	"blockHash":"0x8507d04effec357768853e9f81df41d288af53e5377a57ad53c80091a0218524",
	"logIndex":"0x7c",
	"removed":false}`
	event := types.Log{}
	err := json.Unmarshal([]byte(eventRaw), &event)
	require.NoError(t, err)
	p := MustNewParser(os.Getenv("RPC_URL"))
	log, err := p.Parse(event, uint64(time.Now().Unix()))
	require.NoError(t, err)
	require.Equal(t, log.EventHash, p.eventHash)
	t.Log(log)
}

func TestParseOneinchTradeLog(t *testing.T) {
	p := MustNewParser(os.Getenv("RPC_URL"))

	client, err := ethclient.Dial(os.Getenv("RPC_URL"))
	require.NoError(t, err)

	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash("0xda842bb20353947c3363384fd25f89e1f72c8039bd23efb2d796bd7363892333"))
	require.NoError(t, err)

	for _, eventLog := range receipt.Logs {
		if len(eventLog.Topics) == 0 {
			continue
		}
		firstTopic := eventLog.Topics[0]

		if !strings.EqualFold(firstTopic.String(), p.eventHash) {
			continue
		}

		order, err := p.Parse(*eventLog, uint64(time.Now().Unix()))
		require.NoError(t, err)

		require.Equal(t, true, strings.EqualFold("0x807cF9A772d5a3f9CeFBc1192e939D62f0D9bD38", order.Maker))
		require.Equal(t, strings.ToLower("0xdac3a1ba8fa517c1d98ffecf092b2ad167440131b19dee3d782d0d7eadce01a2"), order.OrderHash)
	}

}
