package paraswap

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

	storageTypes "github.com/KyberNetwork/tradelogs/v2/pkg/storage/tradelogs/types"
	tradelogstype "github.com/KyberNetwork/tradelogs/v2/pkg/types"
	tradingTypes "github.com/KyberNetwork/tradinglib/pkg/types"
)

var rpcURL = os.Getenv("TEST_RPC_URL")

func TestFetchEvent(t *testing.T) {
	t.Skip()
	p := MustNewParser()
	require.Equal(t, p.abi.Events[OrderFilledEvent].ID, common.HexToHash("0x6621486d9c28838df4a87d2cca5007bc2aaf6a5b5de083b1db8faf709302c473"))
	client, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)
	logs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: big.NewInt(20182223),
		ToBlock:   big.NewInt(20182223),
		Addresses: nil,
		Topics: [][]common.Hash{
			{
				common.HexToHash("0x6621486d9c28838df4a87d2cca5007bc2aaf6a5b5de083b1db8faf709302c473"),
			},
		},
	})
	require.NoError(t, err)
	d, _ := json.Marshal(logs)
	t.Log(string(d))
}

func TestParseEvent(t *testing.T) {
	eventRaw := `{
		"address":"0xe92b586627cca7a83dc919cc7127196d70f55a06",
		"topics":["0x6621486d9c28838df4a87d2cca5007bc2aaf6a5b5de083b1db8faf709302c473","0xd3c7eed1fe6c2c0684383c9f4a4f87d7136a96e60e9a5db381d324e9198133f0","0x000000000000000000000000bb289bc97591f70d8216462df40ed713011b968a","0x000000000000000000000000def171fe48cf0115b1d80b88dc8eab59176fee57"],
		"data":"0x000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000000000000000000000000000006f285be082e6dc410000000000000000000000004d224452801aced8b2f0aebe155379bb5d5943810000000000000000000000000000000000000000000000bad91fdd52aac724d2",
		"blockNumber":"0x10174f2",
		"transactionHash":"0xbed8812c5b929036c8b6c20c834482afcee3bc5a7100e2926a90604413c35e5f",
		"transactionIndex":"0x57",
		"blockHash":"0x66f5cae466d1a24aab266717d06df7234b9b826b2fe2891bdfa343ac42c8605f",
		"logIndex":"0xe7",
		"removed":false
		}`
	event := types.Log{}
	err := json.Unmarshal([]byte(eventRaw), &event)
	require.NoError(t, err)
	p := MustNewParser()
	log, err := p.Parse(event, uint64(time.Now().Unix()))
	require.NoError(t, err)
	require.Equal(t, log[0].EventHash, p.eventHash)
	t.Log(log)
}

func TestGetExpiry(t *testing.T) {
	rawData := "{\"from\":\"0xe92b586627cca7a83dc919cc7127196d70f55a06\",\"gas\":439558,\"gasUsed\":81856,\"to\":\"0xe92b586627cca7a83dc919cc7127196d70f55a06\",\"input\":\"2muErwAAAAAAAAAAAAAAALXQS3WlFEMj2OnNU5e0WbY+qfg8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGZ865gAAAAAAAAAAAAAAADAKqo5siP+jQoOXE8n6tkIPHVswgAAAAAAAAAAAAAAANrBf5WNLuUjoiBiBplFl8E9gx7HAAAAAAAAAAAAAAAA/4uk0fw3YvYVTMlCzPMASaKgzsYAAAAAAAAAAAAAAADe8XH+SM8BFbHYC4jcjqtZF2/uVwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAxsXYc+42y0AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAALSaIYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACy0F4AAAAAAAAAAAAAAAAA3vFx/kjPARWx2AuI3I6rWRdv7lcAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAJAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEFF2GlTCyjfKqxeSnhRpctD5CKOnjsi0oSYQ7n+1Z6U5gt9WmOhjBCEsdfdI+AICIoYp8H7s2IG2RnNZteMaCMIGwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA\",\"output\":\"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADEzgZ32hzno=\",\"calls\":[{\"from\":\"0xe92b586627cca7a83dc919cc7127196d70f55a06\",\"gas\":425708,\"gasUsed\":3000,\"to\":\"0x0000000000000000000000000000000000000001\",\"input\":\"aRK0qv1rumvbPPqx8/Zl8kwCBN4ml9J7BgvU55P5dVoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAG0XYaVMLKN8qrF5KeFGly0PkIo6eOyLShJhDuf7VnpTmC31aY6GMEISx190j4AgIihinwfuzYgbZGc1m14xoIwg=\",\"output\":\"AAAAAAAAAAAAAAAA/4uk0fw3YvYVTMlCzPMASaKgzsY=\"},{\"from\":\"0xe92b586627cca7a83dc919cc7127196d70f55a06\",\"gas\":398772,\"gasUsed\":32125,\"to\":\"0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2\",\"input\":\"I7hy3QAAAAAAAAAAAAAAAP+LpNH8N2L2FUzJQszzAEmioM7GAAAAAAAAAAAAAAAA3vFx/kjPARWx2AuI3I6rWRdv7lcAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMTOBnfaHOeg==\",\"output\":\"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=\",\"logs\":[{\"address\":\"0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2\",\"topics\":[\"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef\",\"0x000000000000000000000000ff8ba4d1fc3762f6154cc942ccf30049a2a0cec6\",\"0x000000000000000000000000def171fe48cf0115b1d80b88dc8eab59176fee57\"],\"data\":\"0x0000000000000000000000000000000000000000000000000c4ce0677da1ce7a\",\"blockNumber\":\"0x0\",\"transactionHash\":\"0x62f60a83309ddff6ecdad1e4ebb7bc1c756115f44ef899d122ade9ccbb874c20\",\"transactionIndex\":\"0x0\",\"blockHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"logIndex\":\"0x0\",\"removed\":false}],\"value\":0,\"contract_call\":{\"name\":\"transferFrom\",\"params\":[{\"name\":\"src\",\"value\":\"0xff8ba4d1fc3762f6154cc942ccf30049a2a0cec6\",\"type\":\"address\"},{\"name\":\"dst\",\"value\":\"0xdef171fe48cf0115b1d80b88dc8eab59176fee57\",\"type\":\"address\"},{\"name\":\"wad\",\"value\":886329961779154554,\"type\":\"uint256\"}]}},{\"from\":\"0xe92b586627cca7a83dc919cc7127196d70f55a06\",\"gas\":365390,\"gasUsed\":9730,\"to\":\"0xdac17f958d2ee523a2206206994597c13d831ec7\",\"input\":\"I7hy3QAAAAAAAAAAAAAAAN7xcf5IzwEVsdgLiNyOq1kXb+5XAAAAAAAAAAAAAAAA/4uk0fw3YvYVTMlCzPMASaKgzsYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAstBeAA==\",\"logs\":[{\"address\":\"0xdac17f958d2ee523a2206206994597c13d831ec7\",\"topics\":[\"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef\",\"0x000000000000000000000000def171fe48cf0115b1d80b88dc8eab59176fee57\",\"0x000000000000000000000000ff8ba4d1fc3762f6154cc942ccf30049a2a0cec6\"],\"data\":\"0x00000000000000000000000000000000000000000000000000000000b2d05e00\",\"blockNumber\":\"0x0\",\"transactionHash\":\"0x62f60a83309ddff6ecdad1e4ebb7bc1c756115f44ef899d122ade9ccbb874c20\",\"transactionIndex\":\"0x0\",\"blockHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"logIndex\":\"0x0\",\"removed\":false}],\"value\":0,\"contract_call\":{\"name\":\"transferFrom\",\"params\":[{\"name\":\"src\",\"value\":\"0xdef171fe48cf0115b1d80b88dc8eab59176fee57\",\"type\":\"address\"},{\"name\":\"dst\",\"value\":\"0xff8ba4d1fc3762f6154cc942ccf30049a2a0cec6\",\"type\":\"address\"},{\"name\":\"wad\",\"value\":3000000000,\"type\":\"uint256\"}]}}],\"logs\":[{\"address\":\"0xe92b586627cca7a83dc919cc7127196d70f55a06\",\"topics\":[\"0x6621486d9c28838df4a87d2cca5007bc2aaf6a5b5de083b1db8faf709302c473\",\"0x6912b4aafd6bba6bdb3cfab1f3f665f24c0204de2697d27b060bd4e793f9755a\",\"0x000000000000000000000000ff8ba4d1fc3762f6154cc942ccf30049a2a0cec6\",\"0x000000000000000000000000def171fe48cf0115b1d80b88dc8eab59176fee57\"],\"data\":\"0x000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000000000000000000000000000000c4ce0677da1ce7a000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec700000000000000000000000000000000000000000000000000000000b2d05e00\",\"blockNumber\":\"0x0\",\"transactionHash\":\"0x62f60a83309ddff6ecdad1e4ebb7bc1c756115f44ef899d122ade9ccbb874c20\",\"transactionIndex\":\"0x0\",\"blockHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"logIndex\":\"0x0\",\"removed\":false}],\"value\":0}"
	var callFrame tradingTypes.CallFrame
	p := MustNewParser()
	err := json.Unmarshal([]byte(rawData), &callFrame)
	require.NoError(t, err)
	rfqOrderParam, err := p.getRFQOrderParams(tradelogstype.ConvertCallFrame(&callFrame))
	require.NoError(t, err)
	require.NotNil(t, rfqOrderParam)
	t.Log(rfqOrderParam)
	require.Equal(t, uint64(1719462808), uint64(rfqOrderParam.Expiry))
}

func TestParseWithCallFrame(t *testing.T) {
	t.Skip("Need to add the rpc url that enables the trace call JSON-RPC")
	expectedTradeLog := []storageTypes.TradeLog{
		{
			OrderHash:              "0x14e3f8874f48ded712f9f3f823636e65cef85e8b3d6df079382812af1fa76d47",
			Maker:                  "0xbAd9ADa0E9633ED27Fa183dBdEceef76712a6Ee1",
			Taker:                  "0xDEF171Fe48CF0115B1d80b88dc8eAB59176FEe57",
			MakerToken:             "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
			TakerToken:             "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
			MakerTokenAmount:       "321000000",
			TakerTokenAmount:       "94793992268096036",
			MakerTokenOriginAmount: "321064203",
			TakerTokenOriginAmount: "94812951952474816",
			ContractAddress:        "0xe92b586627ccA7a83dC919cc7127196d70f55a06",
			BlockNumber:            20182223,
			TxHash:                 "0x12b704d281025c6c7c3507573d05914d80ff8816117a8a10755eb7515d4a6cd6",
			LogIndex:               236,
			Timestamp:              0,
			EventHash:              "0x6621486d9c28838df4a87d2cca5007bc2aaf6a5b5de083b1db8faf709302c473",
			Expiry:                 1719481757,
		},
	}
	p := MustNewParser()

	rawEvent := `{"address":"0xe92b586627cca7a83dc919cc7127196d70f55a06","topics":["0x6621486d9c28838df4a87d2cca5007bc2aaf6a5b5de083b1db8faf709302c473","0x14e3f8874f48ded712f9f3f823636e65cef85e8b3d6df079382812af1fa76d47","0x000000000000000000000000bad9ada0e9633ed27fa183dbdeceef76712a6ee1","0x000000000000000000000000def171fe48cf0115b1d80b88dc8eab59176fee57"],"data":"0x000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb480000000000000000000000000000000000000000000000000000000013221240000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000000000000000000000000000000150c6a23bfb2224","blockNumber":"0x133f4cf","transactionHash":"0x12b704d281025c6c7c3507573d05914d80ff8816117a8a10755eb7515d4a6cd6","transactionIndex":"0x6d","blockHash":"0x850d2c22592da0bb43acfb22111d32a9956d5ac475a03403453159bdbff33caa","logIndex":"0xec","removed":false}`
	var event types.Log
	err := json.Unmarshal([]byte(rawEvent), &event)
	require.NoError(t, err)

	rawData := "{\"from\":\"0xe92b586627cca7a83dc919cc7127196d70f55a06\",\"gas\":241607,\"gasUsed\":110524,\"to\":\"0xe92b586627cca7a83dc919cc7127196d70f55a06\",\"input\":\"2muErwAAAAAAFO0I+q9/rRjDC+6A+cPCQQXRRM96N3lK/6MpAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGZ9NZ0AAAAAAAAAAAAAAACguGmRxiGLNsHRnUounrDONgbrSAAAAAAAAAAAAAAAAMAqqjmyI/6NCg5cTyfq2Qg8dWzCAAAAAAAAAAAAAAAAutmtoOljPtJ/oYPb3s7vdnEqbuEAAAAAAAAAAAAAAADe8XH+SM8BFbHYC4jcjqtZF2/uVwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAATIw0LAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAVDX4KE7wsAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFQxqI7+yIkAAAAAAAAAAAAAAAA3vFx/kjPARWx2AuI3I6rWRdv7lcAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAJAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEEbhY6eEvJIlNVbwU+8jJPNg4uYI5+7AiaHPiKHtxU7TgNJn66yyHNWL8YkA4MLRmOtgz5BsL7zmj+2ZqL3RRnM1QAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA\",\"output\":\"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABMiEkA=\",\"calls\":[{\"from\":\"0xe92b586627cca7a83dc919cc7127196d70f55a06\",\"gas\":227050,\"gasUsed\":10599,\"to\":\"0xbad9ada0e9633ed27fa183dbdeceef76712a6ee1\",\"input\":\"Fia6fhTj+IdPSN7XEvnz+CNjbmXO+F6LPW3weTgoEq8fp21HAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQRuFjp4S8kiU1VvBT7yMk82Di5gjn7sCJoc+Ioe3FTtOA0mfrrLIc1YvxiQDgwtGY62DPkGwvvOaP7ZmovdFGczVAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==\",\"output\":\"Fia6fgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=\",\"calls\":[{\"from\":\"0xbad9ada0e9633ed27fa183dbdeceef76712a6ee1\",\"gas\":217976,\"gasUsed\":4790,\"to\":\"0xdb432af6f3bcb74242f507dd1c7f4aca5f2d63c8\",\"input\":\"Fia6fhTj+IdPSN7XEvnz+CNjbmXO+F6LPW3weTgoEq8fp21HAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQRuFjp4S8kiU1VvBT7yMk82Di5gjn7sCJoc+Ioe3FTtOA0mfrrLIc1YvxiQDgwtGY62DPkGwvvOaP7ZmovdFGczVAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==\",\"output\":\"Fia6fgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=\",\"calls\":[{\"from\":\"0xdb432af6f3bcb74242f507dd1c7f4aca5f2d63c8\",\"gas\":213014,\"gasUsed\":3000,\"to\":\"0x0000000000000000000000000000000000000001\",\"input\":\"kcgcbRJM2uKiJy0N6hLlV0PcBCvfNNs/R5xAigv+W0kAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAG4WOnhLySJTVW8FPvIyTzYOLmCOfuwImhz4ih7cVO04DSZ+usshzVi/GJAODC0ZjrYM+QbC+85o/tmai90UZzNU=\",\"output\":\"AAAAAAAAAAAAAAAAajLMBE3WNZwnu2bnsC3ObdD9okc=\"}]}]},{\"from\":\"0xe92b586627cca7a83dc919cc7127196d70f55a06\",\"gas\":192574,\"gasUsed\":48549,\"to\":\"0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48\",\"input\":\"I7hy3QAAAAAAAAAAAAAAALrZraDpYz7Sf6GD297O73ZxKm7hAAAAAAAAAAAAAAAA3vFx/kjPARWx2AuI3I6rWRdv7lcAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEyISQA==\",\"output\":\"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=\",\"calls\":[{\"from\":\"0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48\",\"gas\":182428,\"gasUsed\":41254,\"to\":\"0x43506849d7c04f9138d1a2050bbf3a0c054402dd\",\"input\":\"I7hy3QAAAAAAAAAAAAAAALrZraDpYz7Sf6GD297O73ZxKm7hAAAAAAAAAAAAAAAA3vFx/kjPARWx2AuI3I6rWRdv7lcAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEyISQA==\",\"output\":\"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=\",\"logs\":[{\"address\":\"0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48\",\"topics\":[\"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef\",\"0x000000000000000000000000bad9ada0e9633ed27fa183dbdeceef76712a6ee1\",\"0x000000000000000000000000def171fe48cf0115b1d80b88dc8eab59176fee57\"],\"data\":\"0x0000000000000000000000000000000000000000000000000000000013221240\",\"blockNumber\":\"0x0\",\"transactionHash\":\"0x12b704d281025c6c7c3507573d05914d80ff8816117a8a10755eb7515d4a6cd6\",\"transactionIndex\":\"0x0\",\"blockHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"logIndex\":\"0x0\",\"removed\":false}],\"value\":0,\"contract_call\":{\"name\":\"transferFrom\",\"params\":[{\"name\":\"src\",\"value\":\"0xbad9ada0e9633ed27fa183dbdeceef76712a6ee1\",\"type\":\"address\"},{\"name\":\"dst\",\"value\":\"0xdef171fe48cf0115b1d80b88dc8eab59176fee57\",\"type\":\"address\"},{\"name\":\"wad\",\"value\":321000000,\"type\":\"uint256\"}]}}],\"value\":0,\"contract_call\":{\"name\":\"transferFrom\",\"params\":[{\"name\":\"src\",\"value\":\"0xbad9ada0e9633ed27fa183dbdeceef76712a6ee1\",\"type\":\"address\"},{\"name\":\"dst\",\"value\":\"0xdef171fe48cf0115b1d80b88dc8eab59176fee57\",\"type\":\"address\"},{\"name\":\"wad\",\"value\":321000000,\"type\":\"uint256\"}]}},{\"from\":\"0xe92b586627cca7a83dc919cc7127196d70f55a06\",\"gas\":143023,\"gasUsed\":10225,\"to\":\"0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2\",\"input\":\"I7hy3QAAAAAAAAAAAAAAAN7xcf5IzwEVsdgLiNyOq1kXb+5XAAAAAAAAAAAAAAAAutmtoOljPtJ/oYPb3s7vdnEqbuEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABUMaiO/siJA==\",\"output\":\"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=\",\"logs\":[{\"address\":\"0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2\",\"topics\":[\"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef\",\"0x000000000000000000000000def171fe48cf0115b1d80b88dc8eab59176fee57\",\"0x000000000000000000000000bad9ada0e9633ed27fa183dbdeceef76712a6ee1\"],\"data\":\"0x0000000000000000000000000000000000000000000000000150c6a23bfb2224\",\"blockNumber\":\"0x0\",\"transactionHash\":\"0x12b704d281025c6c7c3507573d05914d80ff8816117a8a10755eb7515d4a6cd6\",\"transactionIndex\":\"0x0\",\"blockHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"logIndex\":\"0x0\",\"removed\":false}],\"value\":0,\"contract_call\":{\"name\":\"transferFrom\",\"params\":[{\"name\":\"src\",\"value\":\"0xdef171fe48cf0115b1d80b88dc8eab59176fee57\",\"type\":\"address\"},{\"name\":\"dst\",\"value\":\"0xbad9ada0e9633ed27fa183dbdeceef76712a6ee1\",\"type\":\"address\"},{\"name\":\"wad\",\"value\":94793992268096036,\"type\":\"uint256\"}]}}],\"logs\":[{\"address\":\"0xe92b586627cca7a83dc919cc7127196d70f55a06\",\"topics\":[\"0x6621486d9c28838df4a87d2cca5007bc2aaf6a5b5de083b1db8faf709302c473\",\"0x14e3f8874f48ded712f9f3f823636e65cef85e8b3d6df079382812af1fa76d47\",\"0x000000000000000000000000bad9ada0e9633ed27fa183dbdeceef76712a6ee1\",\"0x000000000000000000000000def171fe48cf0115b1d80b88dc8eab59176fee57\"],\"data\":\"0x000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb480000000000000000000000000000000000000000000000000000000013221240000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000000000000000000000000000000150c6a23bfb2224\",\"blockNumber\":\"0x0\",\"transactionHash\":\"0x12b704d281025c6c7c3507573d05914d80ff8816117a8a10755eb7515d4a6cd6\",\"transactionIndex\":\"0x0\",\"blockHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"logIndex\":\"0x0\",\"removed\":false}],\"value\":0}"
	var callFrame tradingTypes.CallFrame
	err = json.Unmarshal([]byte(rawData), &callFrame)
	require.NoError(t, err)

	parse, err := p.ParseWithCallFrame(tradelogstype.ConvertCallFrame(&callFrame), event, 0)
	require.NoError(t, err)
	t.Log(parse)
	require.Equal(t, expectedTradeLog, parse)

}
