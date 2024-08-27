package oneinch

import (
	"context"
	"encoding/json"
	tradelogs_type "github.com/KyberNetwork/tradelogs/pkg/types"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/KyberNetwork/tradelogs/pkg/storage"
	tradingTypes "github.com/KyberNetwork/tradinglib/pkg/types"

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
	t.Skip("Need to add the rpc url that enables the trace call JSON-RPC")
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
	ethClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		panic(err)
	}
	traceCalls := tracecall.NewCache(rpcnode.NewClient(ethClient))
	p := MustNewParser(traceCalls)
	log, err := p.Parse(event, uint64(time.Now().Unix()))
	require.NoError(t, err)
	require.Equal(t, log.EventHash, p.eventHash)
	t.Log(log)
}

func TestParseOneinchTradeLog(t *testing.T) {
	t.Skip("Need to add the rpc url that enables the trace call JSON-RPC")
	ethClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		panic(err)
	}
	traceCalls := tracecall.NewCache(rpcnode.NewClient(ethClient))
	p := MustNewParser(traceCalls)
	client, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)

	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash("0x702beb5ccd0c96ba12de36fda799d95a08ec4d174b6f09bb8bb7bcdcef7f5f0f"))
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
		t.Log(order)
		//require.Equal(t, true, strings.EqualFold("0x807cF9A772d5a3f9CeFBc1192e939D62f0D9bD38", order.Maker))
		//require.Equal(t, strings.ToLower("0xdac3a1ba8fa517c1d98ffecf092b2ad167440131b19dee3d782d0d7eadce01a2"), order.OrderHash)
		//require.Equal(t, "25583325494215510852", order.MakerTokenAmount)
		//require.Equal(t, "40450976402181226773533785758", order.TakerTokenAmount)
	}
}

func TestParseWithCallFrame(t *testing.T) {
	t.Skip("Need to add the rpc url that enables the trace call JSON-RPC")
	expectedTradeLog := storage.TradeLog{
		OrderHash:        "0xeb4ab4d3477697ae60514548f6d49f943611addc82f6b95b4aded5c2641cdd9e",
		Maker:            "0x807cF9A772d5a3f9CeFBc1192e939D62f0D9bD38",
		Taker:            "",
		MakerToken:       "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
		TakerToken:       "0x744d70FDBE2Ba4CF95131626614a1763DF805B9E",
		MakerTokenAmount: "83168360428091648",
		TakerTokenAmount: "10852746058000000000000",
		ContractAddress:  "0x1111111254EEB25477B68fb85Ed929f73A960582",
		BlockNumber:      20184416,
		TxHash:           "0x702beb5ccd0c96ba12de36fda799d95a08ec4d174b6f09bb8bb7bcdcef7f5f0f",
		LogIndex:         364,
		Timestamp:        0,
		EventHash:        "0xc3b639f02b125bfa160e50739b8c44eb2d1b6908e2b6d5925c6d770f2ca78127",
		MakerTraits:      "",
		Expiry:           1719508091,
	}
	ethClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		panic(err)
	}
	traceCalls := tracecall.NewCache(rpcnode.NewClient(ethClient))
	p := MustNewParser(traceCalls)
	client, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)

	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash("0x702beb5ccd0c96ba12de36fda799d95a08ec4d174b6f09bb8bb7bcdcef7f5f0f"))
	require.NoError(t, err)

	rawData := "{\"from\":\"0x5f515f6c524b18ca30f7783fb58dd4be2e9904ec\",\"gas\":182294,\"gasUsed\":109176,\"to\":\"0x1111111254eeb25477b68fb85ed929f73a960582\",\"input\":\"PsqcCgAAAAAAAAAAAAAAAAAAAAAAAAAAZn2ce3lofoB8830BAAAAAAAAAAAAAAAAwCqqObIj/o0KDlxPJ+rZCDx1bMIAAAAAAAAAAAAAAAB0TXD9viukz5UTFiZhShdj34BbngAAAAAAAAAAAAAAAIB8+ady1aP5zvvBGS6TnWLw2b04AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABJ3ku4945AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACTFQbX3Ai5qAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAASAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAkxUG19wIuagAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABBxM+OYoIwqhi626Mbleczi9CFeIFmaQ076PDB5rpYnfJqDk/Eg831HsnkqE8BUXTq66rDcvuIuo0snvPvnmiA5RsAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA\",\"output\":\"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAASd5LuPeOQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAkxUG19wIuagAOtKtNNHdpeuYFFFSPbUn5Q2Ea3cgva5W0re1cJkHN2e\",\"calls\":[{\"from\":\"0x1111111254eeb25477b68fb85ed929f73a960582\",\"gas\":177450,\"gasUsed\":3000,\"to\":\"0x0000000000000000000000000000000000000001\",\"input\":\"60q000d2l65gUUVI9tSflDYRrdyC9rlbSt7VwmQc3Z4AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAG8TPjmKCMKoYutujG5XnM4vQhXiBZmkNO+jwwea6WJ3yag5PxIPN9R7J5KhPAVF06uuqw3L7iLqNLJ7z755ogOU=\",\"output\":\"AAAAAAAAAAAAAAAAQ2lccVZOV+LG5Ot+tmP+Kcb1KgA=\"},{\"from\":\"0x1111111254eeb25477b68fb85ed929f73a960582\",\"gas\":171702,\"gasUsed\":14229,\"to\":\"0x807cf9a772d5a3f9cefbc1192e939d62f0d9bd38\",\"input\":\"Fia6futKtNNHdpeuYFFFSPbUn5Q2Ea3cgva5W0re1cJkHN2eAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQcTPjmKCMKoYutujG5XnM4vQhXiBZmkNO+jwwea6WJ3yag5PxIPN9R7J5KhPAVF06uuqw3L7iLqNLJ7z755ogOUb\",\"output\":\"Fia6fgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=\",\"calls\":[{\"from\":\"0x807cf9a772d5a3f9cefbc1192e939d62f0d9bd38\",\"gas\":161924,\"gasUsed\":6980,\"to\":\"0xec3c808e408026a346f2752c990854aca61ad838\",\"input\":\"Fia6futKtNNHdpeuYFFFSPbUn5Q2Ea3cgva5W0re1cJkHN2eAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQcTPjmKCMKoYutujG5XnM4vQhXiBZmkNO+jwwea6WJ3yag5PxIPN9R7J5KhPAVF06uuqw3L7iLqNLJ7z755ogOUb\",\"output\":\"Fia6fgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=\",\"calls\":[{\"from\":\"0x807cf9a772d5a3f9cefbc1192e939d62f0d9bd38\",\"gas\":155932,\"gasUsed\":3000,\"to\":\"0x0000000000000000000000000000000000000001\",\"input\":\"TetFfRRAP0BEwVbRODTjxkN13J95gqAMJ7Ss79MoOgMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAG8TPjmKCMKoYutujG5XnM4vQhXiBZmkNO+jwwea6WJ3yag5PxIPN9R7J5KhPAVF06uuqw3L7iLqNLJ7z755ogOU=\",\"output\":\"AAAAAAAAAAAAAAAA6fGRwJZ5p0Wh16fyBlhoGxjZoxo=\"}],\"value\":0}]},{\"from\":\"0x1111111254eeb25477b68fb85ed929f73a960582\",\"gas\":132030,\"gasUsed\":15025,\"to\":\"0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2\",\"input\":\"I7hy3QAAAAAAAAAAAAAAAIB8+ady1aP5zvvBGS6TnWLw2b04AAAAAAAAAAAAAAAAX1FfbFJLGMow93g/tY3Uvi6ZBOwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABJ3ku4945AA==\",\"output\":\"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=\",\"logs\":[{\"address\":\"0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2\",\"topics\":[\"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef\",\"0x000000000000000000000000807cf9a772d5a3f9cefbc1192e939d62f0d9bd38\",\"0x0000000000000000000000005f515f6c524b18ca30f7783fb58dd4be2e9904ec\"],\"data\":\"0x0000000000000000000000000000000000000000000000000127792ee3de3900\",\"blockNumber\":\"0x0\",\"transactionHash\":\"0x702beb5ccd0c96ba12de36fda799d95a08ec4d174b6f09bb8bb7bcdcef7f5f0f\",\"transactionIndex\":\"0x0\",\"blockHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"logIndex\":\"0x0\",\"removed\":false}],\"value\":0,\"contract_call\":{\"name\":\"transferFrom\",\"params\":[{\"name\":\"src\",\"value\":\"0x807cf9a772d5a3f9cefbc1192e939d62f0d9bd38\",\"type\":\"address\"},{\"name\":\"dst\",\"value\":\"0x5f515f6c524b18ca30f7783fb58dd4be2e9904ec\",\"type\":\"address\"},{\"name\":\"wad\",\"value\":83168360428091648,\"type\":\"uint256\"}]}},{\"from\":\"0x1111111254eeb25477b68fb85ed929f73a960582\",\"gas\":116753,\"gasUsed\":43831,\"to\":\"0x744d70fdbe2ba4cf95131626614a1763df805b9e\",\"input\":\"I7hy3QAAAAAAAAAAAAAAAF9RX2xSSxjKMPd4P7WN1L4umQTsAAAAAAAAAAAAAAAAgHz5p3LVo/nO+8EZLpOdYvDZvTgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAkxUG19wIuagAA==\",\"output\":\"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=\",\"calls\":[{\"from\":\"0x744d70fdbe2ba4cf95131626614a1763df805b9e\",\"gas\":109899,\"gasUsed\":669,\"to\":\"0x52ae2b53c847327f95a5084a7c38c0adb12fd302\",\"input\":\"SjkxSQAAAAAAAAAAAAAAAF9RX2xSSxjKMPd4P7WN1L4umQTsAAAAAAAAAAAAAAAAgHz5p3LVo/nO+8EZLpOdYvDZvTgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAkxUG19wIuagAA==\",\"output\":\"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=\",\"value\":0}],\"logs\":[{\"address\":\"0x744d70fdbe2ba4cf95131626614a1763df805b9e\",\"topics\":[\"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef\",\"0x0000000000000000000000005f515f6c524b18ca30f7783fb58dd4be2e9904ec\",\"0x000000000000000000000000807cf9a772d5a3f9cefbc1192e939d62f0d9bd38\"],\"data\":\"0x00000000000000000000000000000000000000000000024c541b5f7022e6a000\",\"blockNumber\":\"0x0\",\"transactionHash\":\"0x702beb5ccd0c96ba12de36fda799d95a08ec4d174b6f09bb8bb7bcdcef7f5f0f\",\"transactionIndex\":\"0x0\",\"blockHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"logIndex\":\"0x0\",\"removed\":false}],\"value\":0,\"contract_call\":{\"name\":\"transferFrom\",\"params\":[{\"name\":\"src\",\"value\":\"0x5f515f6c524b18ca30f7783fb58dd4be2e9904ec\",\"type\":\"address\"},{\"name\":\"dst\",\"value\":\"0x807cf9a772d5a3f9cefbc1192e939d62f0d9bd38\",\"type\":\"address\"},{\"name\":\"wad\",\"value\":10852746058000000000000,\"type\":\"uint256\"}]}}],\"logs\":[{\"address\":\"0x1111111254eeb25477b68fb85ed929f73a960582\",\"topics\":[\"0xc3b639f02b125bfa160e50739b8c44eb2d1b6908e2b6d5925c6d770f2ca78127\"],\"data\":\"0xeb4ab4d3477697ae60514548f6d49f943611addc82f6b95b4aded5c2641cdd9e0000000000000000000000000000000000000000000000000127792ee3de3900\",\"blockNumber\":\"0x0\",\"transactionHash\":\"0x702beb5ccd0c96ba12de36fda799d95a08ec4d174b6f09bb8bb7bcdcef7f5f0f\",\"transactionIndex\":\"0x0\",\"blockHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"logIndex\":\"0x0\",\"removed\":false}],\"value\":0}"
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

		parse, err := p.ParseWithCallFrame(tradelogs_type.ConvertCallFrame(&callFrame), *eventLog, 0)
		require.NoError(t, err)
		require.Equal(t, expectedTradeLog, parse)
	}
}
