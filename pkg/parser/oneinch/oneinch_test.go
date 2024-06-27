package oneinch

import (
	"context"
	"encoding/json"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	tradingTypes "github.com/KyberNetwork/tradinglib/pkg/types"
	"math/big"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/KyberNetwork/tradelogs/pkg/rpcnode"
	"github.com/KyberNetwork/tradelogs/pkg/tracecall"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

const rpcURL = "https://eth-mainnet.nodereal.io/v1/1659dfb40aa24bbb8153a677b98064d7"

func TestFetchEvent(t *testing.T) {
	t.Skip("Need to add the rpc url that enables the trace call JSON-RPC")
	rpcClient, err := rpcnode.NewClient(http.DefaultClient, rpcURL)
	if err != nil {
		panic(err)
	}
	traceCalls := tracecall.NewCache(rpcClient)
	p := MustNewParser(traceCalls)
	require.Equal(t, p.abi.Events[FilledEvent].ID, common.HexToHash("0xc3b639f02b125bfa160e50739b8c44eb2d1b6908e2b6d5925c6d770f2ca78127"))
	client, err := ethclient.Dial(rpcURL)
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
	//t.Skip("Need to add the rpc url that enables the trace call JSON-RPC")
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
	rpcClient, err := rpcnode.NewClient(http.DefaultClient, rpcURL)
	if err != nil {
		panic(err)
	}
	traceCalls := tracecall.NewCache(rpcClient)
	p := MustNewParser(traceCalls)
	log, err := p.Parse(event, uint64(time.Now().Unix()))
	require.NoError(t, err)
	require.Equal(t, log.EventHash, p.eventHash)
	t.Log(log)
}

func TestParseOneinchTradeLog(t *testing.T) {
	t.Skip("Need to add the rpc url that enables the trace call JSON-RPC")
	rpcClient, err := rpcnode.NewClient(http.DefaultClient, rpcURL)
	if err != nil {
		panic(err)
	}
	traceCalls := tracecall.NewCache(rpcClient)
	p := MustNewParser(traceCalls)
	client, err := ethclient.Dial(rpcURL)
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
		require.Equal(t, "25583325494215510852", order.MakerTokenAmount)
		require.Equal(t, "40450976402181226773533785758", order.TakerTokenAmount)
	}
}

func TestParseWithCallFrame(t *testing.T) {
	t.Skip("Need to add the rpc url that enables the trace call JSON-RPC")
	expectedTradeLog := storage.TradeLog{
		OrderHash:        "0x01928a3ab4a7b94517fb0c20eff41ebbb589f17b184a99c4a5e84b8375f0b597",
		Maker:            "0x807cF9A772d5a3f9CeFBc1192e939D62f0D9bD38",
		Taker:            "",
		MakerToken:       "0xdAC17F958D2ee523a2206206994597C13D831ec7",
		TakerToken:       "0x6E2a43be0B1d33b726f0CA3b8de60b3482b8b050",
		MakerTokenAmount: "1123140395",
		TakerTokenAmount: "580000000000000000000",
		ContractAddress:  "0x1111111254EEB25477B68fb85Ed929f73A960582",
		BlockNumber:      20182172,
		TxHash:           "0x18bf6f4cc05989a2f401035aa54e188cebeb8deb64a810fa727b672cd7842101",
		LogIndex:         56,
		Timestamp:        0,
		EventHash:        "0xc3b639f02b125bfa160e50739b8c44eb2d1b6908e2b6d5925c6d770f2ca78127",
		MakerTraits:      "",
		Expiry:           1719481031,
	}
	rpcClient, err := rpcnode.NewClient(http.DefaultClient, rpcURL)
	if err != nil {
		panic(err)
	}
	traceCalls := tracecall.NewCache(rpcClient)
	p := MustNewParser(traceCalls)
	client, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)

	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash("0x18bf6f4cc05989a2f401035aa54e188cebeb8deb64a810fa727b672cd7842101"))
	require.NoError(t, err)

	rawData := "{\"from\":\"0x74de5d4fcbf63e00296fd95d33236b9794016631\",\"gas\":161263,\"gasUsed\":103681,\"to\":\"0x1111111254eeb25477b68fb85ed929f73a960582\",\"input\":\"lXDu7gAAAAAAAAAAAAAAAAAAAAAAAAAAZn0yx6/aB9HlMsUJAAAAAAAAAAAAAAAA2sF/lY0u5SOiIGIGmUWXwT2DHscAAAAAAAAAAAAAAABuKkO+Cx0ztybwyjuN5gs0griwUAAAAAAAAAAAAAAAAIB8+ady1aP5zvvBGS6TnWLw2b04AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQvHDKwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAH3Ed7wc+kAAAHFEqpt24w4zmBNb9eoZMOV2ZM2TnjtstZm3645WsXnBCmM3w5npJAXRcYhdnyL0WUV+f2AUoEPnqxOLwrKbkEgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAH3Ed7wc+kAAAfcvqfA==\",\"output\":\"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAELxwysAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAB9xHe8HPpAAAAGSijq0p7lFF/sMIO/0Hru1ifF7GEqZxKXoS4N18LWX\",\"calls\":[{\"from\":\"0x1111111254eeb25477b68fb85ed929f73a960582\",\"gas\":157059,\"gasUsed\":3000,\"to\":\"0x0000000000000000000000000000000000000001\",\"input\":\"AZKKOrSnuUUX+wwg7/Qeu7WJ8XsYSpnEpehLg3XwtZcAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGxxRKqbduMOM5gTW/XqGTDldmTNk547bLWZt+uOVrF5wQpjN8OZ6SQF0XGIXZ8i9FlFfn9gFKBD56sTi8Kym5BI=\",\"output\":\"AAAAAAAAAAAAAAAAcMjuYoxvpOw8CupiJt9bm90Yawg=\"},{\"from\":\"0x1111111254eeb25477b68fb85ed929f73a960582\",\"gas\":151320,\"gasUsed\":14324,\"to\":\"0x807cf9a772d5a3f9cefbc1192e939d62f0d9bd38\",\"input\":\"Fia6fgGSijq0p7lFF/sMIO/0Hru1ifF7GEqZxKXoS4N18LWXAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQBxRKqbduMOM5gTW/XqGTDldmTNk547bLWZt+uOVrF5wQpjN8OZ6SQF0XGIXZ8i9FlFfn9gFKBD56sTi8Kym5BI=\",\"output\":\"Fia6fgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=\",\"calls\":[{\"from\":\"0x807cf9a772d5a3f9cefbc1192e939d62f0d9bd38\",\"gas\":141861,\"gasUsed\":7075,\"to\":\"0xec3c808e408026a346f2752c990854aca61ad838\",\"input\":\"Fia6fgGSijq0p7lFF/sMIO/0Hru1ifF7GEqZxKXoS4N18LWXAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQBxRKqbduMOM5gTW/XqGTDldmTNk547bLWZt+uOVrF5wQpjN8OZ6SQF0XGIXZ8i9FlFfn9gFKBD56sTi8Kym5BI=\",\"output\":\"Fia6fgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=\",\"calls\":[{\"from\":\"0x807cf9a772d5a3f9cefbc1192e939d62f0d9bd38\",\"gas\":136125,\"gasUsed\":3000,\"to\":\"0x0000000000000000000000000000000000000001\",\"input\":\"a6AH07XIBawwOlbER1VDAEmQtypso+eiEj7kRKrQfiUAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGxxRKqbduMOM5gTW/XqGTDldmTNk547bLWZt+uOVrF5wQpjN8OZ6SQF0XGIXZ8i9FlFfn9gFKBD56sTi8Kym5BI=\",\"output\":\"AAAAAAAAAAAAAAAA6fGRwJZ5p0Wh16fyBlhoGxjZoxo=\"}],\"value\":0}]},{\"from\":\"0x1111111254eeb25477b68fb85ed929f73a960582\",\"gas\":111580,\"gasUsed\":43630,\"to\":\"0xdac17f958d2ee523a2206206994597c13d831ec7\",\"input\":\"I7hy3QAAAAAAAAAAAAAAAIB8+ady1aP5zvvBGS6TnWLw2b04AAAAAAAAAAAAAAAAdN5dT8v2PgApb9ldMyNrl5QBZjEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQvHDKw==\",\"logs\":[{\"address\":\"0xdac17f958d2ee523a2206206994597c13d831ec7\",\"topics\":[\"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef\",\"0x000000000000000000000000807cf9a772d5a3f9cefbc1192e939d62f0d9bd38\",\"0x00000000000000000000000074de5d4fcbf63e00296fd95d33236b9794016631\"],\"data\":\"0x0000000000000000000000000000000000000000000000000000000042f1c32b\",\"blockNumber\":\"0x0\",\"transactionHash\":\"0x18bf6f4cc05989a2f401035aa54e188cebeb8deb64a810fa727b672cd7842101\",\"transactionIndex\":\"0x0\",\"blockHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"logIndex\":\"0x0\",\"removed\":false}],\"value\":0,\"contract_call\":{\"name\":\"transferFrom\",\"params\":[{\"name\":\"src\",\"value\":\"0x807cf9a772d5a3f9cefbc1192e939d62f0d9bd38\",\"type\":\"address\"},{\"name\":\"dst\",\"value\":\"0x74de5d4fcbf63e00296fd95d33236b9794016631\",\"type\":\"address\"},{\"name\":\"wad\",\"value\":1123140395,\"type\":\"uint256\"}]}},{\"from\":\"0x1111111254eeb25477b68fb85ed929f73a960582\",\"gas\":68071,\"gasUsed\":9958,\"to\":\"0x6e2a43be0b1d33b726f0ca3b8de60b3482b8b050\",\"input\":\"I7hy3QAAAAAAAAAAAAAAAHTeXU/L9j4AKW/ZXTMja5eUAWYxAAAAAAAAAAAAAAAAgHz5p3LVo/nO+8EZLpOdYvDZvTgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAB9xHe8HPpAAAA==\",\"output\":\"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=\",\"calls\":[{\"from\":\"0x6e2a43be0b1d33b726f0ca3b8de60b3482b8b050\",\"gas\":66648,\"gasUsed\":9551,\"to\":\"0x2291323cf23d1553c6f79dc30b4a8865c03a90cf\",\"input\":\"I7hy3QAAAAAAAAAAAAAAAHTeXU/L9j4AKW/ZXTMja5eUAWYxAAAAAAAAAAAAAAAAgHz5p3LVo/nO+8EZLpOdYvDZvTgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAB9xHe8HPpAAAA==\",\"output\":\"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=\",\"logs\":[{\"address\":\"0x6e2a43be0b1d33b726f0ca3b8de60b3482b8b050\",\"topics\":[\"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef\",\"0x00000000000000000000000074de5d4fcbf63e00296fd95d33236b9794016631\",\"0x000000000000000000000000807cf9a772d5a3f9cefbc1192e939d62f0d9bd38\"],\"data\":\"0x00000000000000000000000000000000000000000000001f711def073e900000\",\"blockNumber\":\"0x0\",\"transactionHash\":\"0x18bf6f4cc05989a2f401035aa54e188cebeb8deb64a810fa727b672cd7842101\",\"transactionIndex\":\"0x0\",\"blockHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"logIndex\":\"0x0\",\"removed\":false}],\"value\":0,\"contract_call\":{\"name\":\"transferFrom\",\"params\":[{\"name\":\"src\",\"value\":\"0x74de5d4fcbf63e00296fd95d33236b9794016631\",\"type\":\"address\"},{\"name\":\"dst\",\"value\":\"0x807cf9a772d5a3f9cefbc1192e939d62f0d9bd38\",\"type\":\"address\"},{\"name\":\"wad\",\"value\":580000000000000000000,\"type\":\"uint256\"}]}}],\"value\":0,\"contract_call\":{\"name\":\"transferFrom\",\"params\":[{\"name\":\"src\",\"value\":\"0x74de5d4fcbf63e00296fd95d33236b9794016631\",\"type\":\"address\"},{\"name\":\"dst\",\"value\":\"0x807cf9a772d5a3f9cefbc1192e939d62f0d9bd38\",\"type\":\"address\"},{\"name\":\"wad\",\"value\":580000000000000000000,\"type\":\"uint256\"}]}}],\"logs\":[{\"address\":\"0x1111111254eeb25477b68fb85ed929f73a960582\",\"topics\":[\"0xc3b639f02b125bfa160e50739b8c44eb2d1b6908e2b6d5925c6d770f2ca78127\"],\"data\":\"0x01928a3ab4a7b94517fb0c20eff41ebbb589f17b184a99c4a5e84b8375f0b5970000000000000000000000000000000000000000000000000000000042f1c32b\",\"blockNumber\":\"0x0\",\"transactionHash\":\"0x18bf6f4cc05989a2f401035aa54e188cebeb8deb64a810fa727b672cd7842101\",\"transactionIndex\":\"0x0\",\"blockHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"logIndex\":\"0x0\",\"removed\":false}],\"value\":0}"
	var callFrame tradingTypes.CallFrame
	err = json.Unmarshal([]byte(rawData), &callFrame)
	require.NoError(t, err)

	for _, eventLog := range receipt.Logs {
		if len(eventLog.Topics) == 0 {
			continue
		}
		firstTopic := eventLog.Topics[0]
		if !strings.EqualFold(firstTopic.String(), p.eventHash) {
			continue
		}

		parse, err := p.ParseWithCallFrame(&callFrame, *eventLog, 0)
		require.NoError(t, err)
		t.Log(parse)
		require.Equal(t, expectedTradeLog, parse)
	}
}
