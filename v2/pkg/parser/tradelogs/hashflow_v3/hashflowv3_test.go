package hashflowv3

import (
	"context"
	"encoding/json"
	"math/big"
	"os"
	"strings"
	"testing"
	"time"

	tradelogstypes "github.com/KyberNetwork/tradelogs/v2/pkg/types"
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
	require.Equal(t, p.abi.Events[TradeEvent].ID, common.HexToHash("0x34f57786fb01682fb4eec88d340387ef01a168fe345ea5b76f709d4e560c10eb"))
	client, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)
	logs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: big.NewInt(18267792),
		ToBlock:   big.NewInt(18268792),
		Addresses: nil,
		Topics:    [][]common.Hash{{common.HexToHash("0x34f57786fb01682fb4eec88d340387ef01a168fe345ea5b76f709d4e560c10eb")}},
	})
	require.NoError(t, err)
	d, _ := json.Marshal(logs)
	t.Log(string(d))
}

func TestParseEvent(t *testing.T) {
	eventRaw := `{
		"address":"0x403a1a99095d95501f57269d0f137e5790d40406",
		"topics":[
			"0x34f57786fb01682fb4eec88d340387ef01a168fe345ea5b76f709d4e560c10eb"
			],
		"data":"0x00000000000000000000000013414b047539298d5aed429722211681eaab43b700000000000000000000000013414b047539298d5aed429722211681eaab43b7100100000000000000001748090b00ffffffffffffff001b95ca55a556b000000000000000000000000000006b175474e89094c44da98b954eedeac495271d0f000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000000000000000000000000000004563918244f40000000000000000000000000000000000000000000000000000000aa888b55c36e5",
		"blockNumber":"0x116c013",
		"transactionHash":"0x092681c0f57a996361aa0e6634a5f8ccd0c59cafeb9ceb746f2c5ba0692b39da",
		"transactionIndex":"0x41",
		"blockHash":"0xb77e1bdfe3f03464b4a70cb81e40b6b89704e773d4daac8328ac9bbb000735a0",
		"logIndex":"0x5d",
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

func TestParseWithCallFrame(t *testing.T) {
	t.Skip("Need to add the rpc url that enables the trace call JSON-RPC")
	p := MustNewParser()
	client, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)

	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash("0x092681c0f57a996361aa0e6634a5f8ccd0c59cafeb9ceb746f2c5ba0692b39da"))
	require.NoError(t, err)

	rawData := "{\n                                \"from\": \"0x403a1a99095d95501f57269d0f137e5790d40406\",\n                                \"gas\": \"0xd90f\",\n                                \"gasUsed\": \"0xd90f\",\n                                \"to\": \"0xca310b1b942a30ff4b40a5e1b69ab4607ec79bc1\",\n                                \"input\": \"0xc52ac7200000000000000000000000000000000000000000000000000000000000000020000000000000000000000000403a1a99095d95501f57269d0f137e5790d404060000000000000000000000002008b6c3d07b061a84f790c035c2f6dc11a0be7000000000000000000000000013414b047539298d5aed429722211681eaab43b700000000000000000000000013414b047539298d5aed429722211681eaab43b70000000000000000000000006b175474e89094c44da98b954eedeac495271d0f000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000000000000000000000000000004563918244f400000000000000000000000000000000000000000000000000004563918244f40000000000000000000000000000000000000000000000000000000aa888b55c36e500000000000000000000000000000000000000000000000000000000651bb2a60000000000000000000000000000000000000000000000000000018af4313d9c100100000000000000001748090b00ffffffffffffff001b95ca55a556b0000000000000000000000000000000000000000000000000000000000000000001a0000000000000000000000000000000000000000000000000000000000000004179f3a1fb7ba08bce99f689d9ecd1e4ddfa671e7228726b5a5380f01179590e455585279ae912ae6513fcc6edbdfac8df98e6df08e0b7186f47809fb0429748551c00000000000000000000000000000000000000000000000000000000000000\",\n                                \"calls\": [\n                                    {\n                                        \"from\": \"0x403a1a99095d95501f57269d0f137e5790d40406\",\n                                        \"gas\": \"0xb82a\",\n                                        \"gasUsed\": \"0xbb8\",\n                                        \"to\": \"0x0000000000000000000000000000000000000001\",\n                                        \"input\": \"0xf7f7f3acab04130f7075f57a8d2a29ab8758a8fbda70dfac55cd10e04f6b75dd000000000000000000000000000000000000000000000000000000000000001c79f3a1fb7ba08bce99f689d9ecd1e4ddfa671e7228726b5a5380f01179590e455585279ae912ae6513fcc6edbdfac8df98e6df08e0b7186f47809fb042974855\",\n                                        \"output\": \"0x0000000000000000000000002008b6c3d07b061a84f790c035c2f6dc11a0be70\",\n                                        \"type\": \"STATICCALL\"\n                                    },\n                                    {\n                                        \"from\": \"0x403a1a99095d95501f57269d0f137e5790d40406\",\n                                        \"gas\": \"0x3c00\",\n                                        \"gasUsed\": \"0x3ab1\",\n                                        \"to\": \"0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2\",\n                                        \"input\": \"0x23b872dd0000000000000000000000002008b6c3d07b061a84f790c035c2f6dc11a0be7000000000000000000000000013414b047539298d5aed429722211681eaab43b7000000000000000000000000000000000000000000000000000aa888b55c36e5\",\n                                        \"output\": \"0x0000000000000000000000000000000000000000000000000000000000000001\",\n                                        \"logs\": [\n                                            {\n                                                \"address\": \"0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2\",\n                                                \"topics\": [\n                                                    \"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef\",\n                                                    \"0x0000000000000000000000002008b6c3d07b061a84f790c035c2f6dc11a0be70\",\n                                                    \"0x00000000000000000000000013414b047539298d5aed429722211681eaab43b7\"\n                                                ],\n                                                \"data\": \"0x000000000000000000000000000000000000000000000000000aa888b55c36e5\",\n                                                \"position\": \"0x0\"\n                                            }\n                                        ],\n                                        \"value\": \"0x0\",\n                                        \"type\": \"CALL\"\n                                    }\n                                ],\n                                \"logs\": [\n                                    {\n                                        \"address\": \"0x403a1a99095d95501f57269d0f137e5790d40406\",\n                                        \"topics\": [\n                                            \"0x34f57786fb01682fb4eec88d340387ef01a168fe345ea5b76f709d4e560c10eb\"\n                                        ],\n                                        \"data\": \"0x00000000000000000000000013414b047539298d5aed429722211681eaab43b700000000000000000000000013414b047539298d5aed429722211681eaab43b7100100000000000000001748090b00ffffffffffffff001b95ca55a556b000000000000000000000000000006b175474e89094c44da98b954eedeac495271d0f000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000000000000000000000000000004563918244f40000000000000000000000000000000000000000000000000000000aa888b55c36e5\",\n                                        \"position\": \"0x1\"\n                                    }\n                                ],\n                                \"value\": \"0x0\",\n                                \"type\": \"DELEGATECALL\"\n                            }"
	var callFrame tradelogstypes.CallFrame
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

		parse, err := p.ParseWithCallFrame(callFrame, *eventLog, uint64(time.Now().Unix()))
		require.NoError(t, err)
		t.Log(parse)
	}
}
