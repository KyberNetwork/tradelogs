package kyberswaprfq

import (
	"context"
	"encoding/json"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	tradingTypes "github.com/KyberNetwork/tradinglib/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

const rpcURL = ""

func TestParseWithCallFrame(t *testing.T) {
	t.Skip("Need to add the rpc url that enables the trace call JSON-RPC")
	expectedTradelog := storage.TradeLog{
		OrderHash:        "0x7ccd6f83d5bee5508e9cef5fcad12f06b2c57d0bacff28a44e57d3119c5939bf",
		Maker:            "0x807cF9A772d5a3f9CeFBc1192e939D62f0D9bD38",
		Taker:            "0xf081470f5C6FBCCF48cC4e5B82Dd926409DcdD67",
		MakerToken:       "0xdAC17F958D2ee523a2206206994597C13D831ec7",
		TakerToken:       "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
		MakerTokenAmount: "10416972177",
		TakerTokenAmount: "3072875000000000000",
		ContractAddress:  "0x7A819Fa46734a49D0112796f9377E024c350FB26",
		BlockNumber:      20182193,
		TxHash:           "0x1759c2ea3e51c8bb746008db60c234eb17be587d0d68b2fa5d02443759f88bd4",
		LogIndex:         313,
		Timestamp:        0,
		EventHash:        "0x71cea972ff7e532f1478ece362ee7d7e5be0654905843c97ecca781a3a0c9d4a",
		MakerTraits:      "",
		Expiry:           1719481285,
	}
	p := MustNewParser()
	client, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)

	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash("0x1759c2ea3e51c8bb746008db60c234eb17be587d0d68b2fa5d02443759f88bd4"))
	require.NoError(t, err)

	rawData := "{\"from\":\"0xf081470f5c6fbccf48cc4e5b82dd926409dcdd67\",\"gas\":206466,\"gasUsed\":86201,\"to\":\"0x7a819fa46734a49d0112796f9377e024c350fb26\",\"input\":\"urpYVQAAAAAAAAAAAAAAAAAAAAAAAAAAZn0zxcpbRyUxrViuAAAAAAAAAAAAAAAA2sF/lY0u5SOiIGIGmUWXwT2DHscAAAAAAAAAAAAAAADAKqo5siP+jQoOXE8n6tkIPHVswgAAAAAAAAAAAAAAAIB8+ady1aP5zvvBGS6TnWLw2b04AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACbOZhkQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACqlC4geLbAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAWAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACqlC4geLbAAAAAAAAAAAAAAAAAA8IFHD1xvvM9IzE5bgt2SZAnc3WcAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQcdBlOkb32uWOlPf2UN/q0GyDc4rIz2ENdYJ58l8+T2WMgAAOxgtzYHyIH883laibr8K5PFCv0iSRY11x2Vc8SAcAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==\",\"output\":\"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAmzmYZEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAqpQuIHi2wAA==\",\"calls\":[{\"from\":\"0x7a819fa46734a49d0112796f9377e024c350fb26\",\"gas\":178228,\"gasUsed\":3000,\"to\":\"0x0000000000000000000000000000000000000001\",\"input\":\"fM1vg9W+5VCOnO9fytEvBrLFfQus/yikTlfTEZxZOb8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHMdBlOkb32uWOlPf2UN/q0GyDc4rIz2ENdYJ58l8+T2WMgAAOxgtzYHyIH883laibr8K5PFCv0iSRY11x2Vc8SA=\",\"output\":\"AAAAAAAAAAAAAAAAVFoxXCETx5OSCh0/Bg0bYu7dpLQ=\"},{\"from\":\"0x7a819fa46734a49d0112796f9377e024c350fb26\",\"gas\":171194,\"gasUsed\":14255,\"to\":\"0x807cf9a772d5a3f9cefbc1192e939d62f0d9bd38\",\"input\":\"Fia6fnzNb4PVvuVQjpzvX8rRLwayxX0LrP8opE5X0xGcWTm/AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQcdBlOkb32uWOlPf2UN/q0GyDc4rIz2ENdYJ58l8+T2WMgAAOxgtzYHyIH883laibr8K5PFCv0iSRY11x2Vc8SAcAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==\",\"output\":\"Fia6fgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=\",\"calls\":[{\"from\":\"0x807cf9a772d5a3f9cefbc1192e939d62f0d9bd38\",\"gas\":161418,\"gasUsed\":7000,\"to\":\"0xec3c808e408026a346f2752c990854aca61ad838\",\"input\":\"Fia6fnzNb4PVvuVQjpzvX8rRLwayxX0LrP8opE5X0xGcWTm/AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQcdBlOkb32uWOlPf2UN/q0GyDc4rIz2ENdYJ58l8+T2WMgAAOxgtzYHyIH883laibr8K5PFCv0iSRY11x2Vc8SAcAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==\",\"output\":\"Fia6fgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=\",\"calls\":[{\"from\":\"0x807cf9a772d5a3f9cefbc1192e939d62f0d9bd38\",\"gas\":155415,\"gasUsed\":3000,\"to\":\"0x0000000000000000000000000000000000000001\",\"input\":\"hSaao3QlitjmLOlr9dKd6aEIKWYRF/q7pLz1hOH3NzMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHMdBlOkb32uWOlPf2UN/q0GyDc4rIz2ENdYJ58l8+T2WMgAAOxgtzYHyIH883laibr8K5PFCv0iSRY11x2Vc8SA=\",\"output\":\"AAAAAAAAAAAAAAAA6fGRwJZ5p0Wh16fyBlhoGxjZoxo=\"}],\"value\":0}]},{\"from\":\"0x7a819fa46734a49d0112796f9377e024c350fb26\",\"gas\":155719,\"gasUsed\":24530,\"to\":\"0xdac17f958d2ee523a2206206994597c13d831ec7\",\"input\":\"I7hy3QAAAAAAAAAAAAAAAIB8+ady1aP5zvvBGS6TnWLw2b04AAAAAAAAAAAAAAAA8IFHD1xvvM9IzE5bgt2SZAnc3WcAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACbOZhkQ==\",\"logs\":[{\"address\":\"0xdac17f958d2ee523a2206206994597c13d831ec7\",\"topics\":[\"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef\",\"0x000000000000000000000000807cf9a772d5a3f9cefbc1192e939d62f0d9bd38\",\"0x000000000000000000000000f081470f5c6fbccf48cc4e5b82dd926409dcdd67\"],\"data\":\"0x000000000000000000000000000000000000000000000000000000026ce66191\",\"blockNumber\":\"0x0\",\"transactionHash\":\"0x1759c2ea3e51c8bb746008db60c234eb17be587d0d68b2fa5d02443759f88bd4\",\"transactionIndex\":\"0x0\",\"blockHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"logIndex\":\"0x0\",\"removed\":false}],\"value\":0,\"contract_call\":{\"name\":\"transferFrom\",\"params\":[{\"name\":\"src\",\"value\":\"0x807cf9a772d5a3f9cefbc1192e939d62f0d9bd38\",\"type\":\"address\"},{\"name\":\"dst\",\"value\":\"0xf081470f5c6fbccf48cc4e5b82dd926409dcdd67\",\"type\":\"address\"},{\"name\":\"wad\",\"value\":10416972177,\"type\":\"uint256\"}]}},{\"from\":\"0x7a819fa46734a49d0112796f9377e024c350fb26\",\"gas\":130054,\"gasUsed\":8225,\"to\":\"0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2\",\"input\":\"I7hy3QAAAAAAAAAAAAAAAPCBRw9cb7zPSMxOW4LdkmQJ3N1nAAAAAAAAAAAAAAAAgHz5p3LVo/nO+8EZLpOdYvDZvTgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAqpQuIHi2wAA==\",\"output\":\"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=\",\"logs\":[{\"address\":\"0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2\",\"topics\":[\"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef\",\"0x000000000000000000000000f081470f5c6fbccf48cc4e5b82dd926409dcdd67\",\"0x000000000000000000000000807cf9a772d5a3f9cefbc1192e939d62f0d9bd38\"],\"data\":\"0x0000000000000000000000000000000000000000000000002aa50b881e2db000\",\"blockNumber\":\"0x0\",\"transactionHash\":\"0x1759c2ea3e51c8bb746008db60c234eb17be587d0d68b2fa5d02443759f88bd4\",\"transactionIndex\":\"0x0\",\"blockHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"logIndex\":\"0x0\",\"removed\":false}],\"value\":0,\"contract_call\":{\"name\":\"transferFrom\",\"params\":[{\"name\":\"src\",\"value\":\"0xf081470f5c6fbccf48cc4e5b82dd926409dcdd67\",\"type\":\"address\"},{\"name\":\"dst\",\"value\":\"0x807cf9a772d5a3f9cefbc1192e939d62f0d9bd38\",\"type\":\"address\"},{\"name\":\"wad\",\"value\":3072875000000000000,\"type\":\"uint256\"}]}}],\"logs\":[{\"address\":\"0x7a819fa46734a49d0112796f9377e024c350fb26\",\"topics\":[\"0x71cea972ff7e532f1478ece362ee7d7e5be0654905843c97ecca781a3a0c9d4a\",\"0x000000000000000000000000807cf9a772d5a3f9cefbc1192e939d62f0d9bd38\",\"0x000000000000000000000000f081470f5c6fbccf48cc4e5b82dd926409dcdd67\"],\"data\":\"0x7ccd6f83d5bee5508e9cef5fcad12f06b2c57d0bacff28a44e57d3119c5939bf000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec7000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000000000000000000000000000000000026ce661910000000000000000000000000000000000000000000000002aa50b881e2db000\",\"blockNumber\":\"0x0\",\"transactionHash\":\"0x1759c2ea3e51c8bb746008db60c234eb17be587d0d68b2fa5d02443759f88bd4\",\"transactionIndex\":\"0x0\",\"blockHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"logIndex\":\"0x0\",\"removed\":false}],\"value\":0}"
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
		require.Equal(t, expectedTradelog, parse)
	}
}
