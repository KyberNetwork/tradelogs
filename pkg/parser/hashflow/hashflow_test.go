package hashflow

import (
	"context"
	"encoding/json"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"

	tradelogstype "github.com/KyberNetwork/tradelogs/pkg/types"
	tradingTypes "github.com/KyberNetwork/tradinglib/pkg/types"
)

const rpcURL = ""

func TestFetchEvent(t *testing.T) {
	t.Skip()
	p := MustNewParser()
	require.Equal(t, p.abi.Events[TradeEvent].ID, common.HexToHash("0x8cf3dec1929508e5677d7db003124e74802bfba7250a572205a9986d86ca9f1e"))
	client, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)
	logs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: big.NewInt(17363812),
		ToBlock:   big.NewInt(17363812),
		Addresses: nil,
		Topics:    [][]common.Hash{{common.HexToHash("0x8cf3dec1929508e5677d7db003124e74802bfba7250a572205a9986d86ca9f1e")}},
	})
	require.NoError(t, err)
	d, _ := json.Marshal(logs)
	t.Log(string(d))
}

func TestParseEvent(t *testing.T) {
	eventRaw := `{
		"address":"0x4306aefe6c0e605f22207a43203ae7d7e4fb42c9",
		"topics":[
			"0x8cf3dec1929508e5677d7db003124e74802bfba7250a572205a9986d86ca9f1e"
			],
		"data":"0x0000000000000000000000009863470bd3549563d5499c179ede3c0fcbcf3eb40067616e64616c6674686562726f776e67786d786e6900190878e8b3c77600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec7000000000000000000000000000000000000000000000000002386f26fc10000000000000000000000000000000000000000000000000000000000000121a0f6",
		"blockNumber":"0x108f364",
		"transactionHash":"0x2345ecbe3c5f79d11b632bbd606ed74d7ada809992945bb4ed67f826660229af",
		"transactionIndex":"0x20",
		"blockHash":"0x1057d2b4e14ff582aa1571c4d26c28f662825f2cfb01cfed4009dd125248e5fa",
		"logIndex":"0xe0",
		"removed":false
	}`
	event := types.Log{}
	err := json.Unmarshal([]byte(eventRaw), &event)
	require.NoError(t, err)
	p := MustNewParser()
	log, err := p.Parse(event, uint64(time.Now().Unix()))
	require.NoError(t, err)
	require.Equal(t, log.EventHash, p.eventHash)
	t.Log(log)
}

func TestParseWithCallFrame(t *testing.T) {
	t.Skip("Need to add the rpc url that enables the trace call JSON-RPC")
	p := MustNewParser()
	client, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)

	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash("0x68cea489a8d0cafb5521c7a019897ad2060d9a8f3627cc5bfd2e733d4962d62d"))
	require.NoError(t, err)

	rawData := "{\n    \"from\": \"0xdef1c0ded9bec7f1a1670819833240f027b25eff\",\n    \"gas\": 158955,\n    \"gasUsed\": 62139,\n    \"to\": \"0x5ebac8dbfbba22168471b0f914131d1976536a25\",\n    \"input\": \"2sdI1AAAAAAAAAAAAAAAAFgthyobCU/Ej13jHTtz8tm+R97xAAAAAAAAAAAAAAAAwCqqObIj/o0KDlxPJ+rZCDx1bMIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABCNqFOAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAky06/E/qtAAAAAAAAAAAAAAAA/4uk0fw3YvYVTMlCzPMASaKgzsYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMOh1PDiSphW8tuEl2PP3O1VmnmxAAAAAGZ84SQAAAAAAAAAAAAAAAAAAAAAAAAAAGZ84MoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAcpqmkjALSKGVFRSEhnRjDzvBurefvbEMh83NYQ9ZNf0h0m592fRDLqXc3P4V15F2f7CRHJU60qpowCifEXdXfKAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAky06/E/qt\",\n    \"output\": \"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACTLTr8T+q0AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABCNqFOA==\",\n    \"calls\": [\n        {\n            \"from\": \"0xdef1c0ded9bec7f1a1670819833240f027b25eff\",\n            \"gas\": 150382,\n            \"gasUsed\": 3000,\n            \"to\": \"0x0000000000000000000000000000000000000001\",\n            \"input\": \"l+OB+kmS5T6BnGGHB6PMq/Qaq0INZVBDo8Vq37p8nUgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHKappIwC0ihlRUUhIZ0Yw87wbq3n72xDIfNzWEPWTX9IdJufdn0Qy6l3Nz+FdeRdn+wkRyVOtKqaMAonxF3V3yg=\",\n            \"output\": \"AAAAAAAAAAAAAAAA/4uk0fw3YvYVTMlCzPMASaKgzsY=\"\n        },\n        {\n            \"from\": \"0xdef1c0ded9bec7f1a1670819833240f027b25eff\",\n            \"gas\": 143508,\n            \"gasUsed\": 8225,\n            \"to\": \"0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2\",\n            \"input\": \"I7hy3QAAAAAAAAAAAAAAACL53PRkcITWwxsnZfaRDNhcF4wYAAAAAAAAAAAAAAAA/4uk0fw3YvYVTMlCzPMASaKgzsYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAJMtOvxP6rQ==\",\n            \"output\": \"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=\",\n            \"logs\": [\n                {\n                    \"address\": \"0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2\",\n                    \"topics\": [\n                        \"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef\",\n                        \"0x00000000000000000000000022f9dcf4647084d6c31b2765f6910cd85c178c18\",\n                        \"0x000000000000000000000000ff8ba4d1fc3762f6154cc942ccf30049a2a0cec6\"\n                    ],\n                    \"data\": \"0x0000000000000000000000000000000000000000000000000024cb4ebf13faad\",\n                    \"blockNumber\": \"0x0\",\n                    \"transactionHash\": \"0x8212c857001fc294c889faad28afad3cd92287a7a3dc580d27f56f221881072a\",\n                    \"transactionIndex\": \"0x0\",\n                    \"blockHash\": \"0x0000000000000000000000000000000000000000000000000000000000000000\",\n                    \"logIndex\": \"0x0\",\n                    \"removed\": false\n                }\n            ],\n            \"value\": 0,\n            \"contract_call\": {\n                \"name\": \"transferFrom\",\n                \"params\": [\n                    {\n                        \"name\": \"src\",\n                        \"value\": \"0x22f9dcf4647084d6c31b2765f6910cd85c178c18\",\n                        \"type\": \"address\"\n                    },\n                    {\n                        \"name\": \"dst\",\n                        \"value\": \"0xff8ba4d1fc3762f6154cc942ccf30049a2a0cec6\",\n                        \"type\": \"address\"\n                    },\n                    {\n                        \"name\": \"wad\",\n                        \"value\": 10356638235228845,\n                        \"type\": \"uint256\"\n                    }\n                ]\n            }\n        },\n        {\n            \"from\": \"0xdef1c0ded9bec7f1a1670819833240f027b25eff\",\n            \"gas\": 135048,\n            \"gasUsed\": 37307,\n            \"to\": \"0x582d872a1b094fc48f5de31d3b73f2d9be47def1\",\n            \"input\": \"I7hy3QAAAAAAAAAAAAAAAP+LpNH8N2L2FUzJQszzAEmioM7GAAAAAAAAAAAAAAAAIvnc9GRwhNbDGydl9pEM2FwXjBgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABCNqFOA==\",\n            \"output\": \"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=\",\n            \"logs\": [\n                {\n                    \"address\": \"0x582d872a1b094fc48f5de31d3b73f2d9be47def1\",\n                    \"topics\": [\n                        \"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef\",\n                        \"0x000000000000000000000000ff8ba4d1fc3762f6154cc942ccf30049a2a0cec6\",\n                        \"0x00000000000000000000000022f9dcf4647084d6c31b2765f6910cd85c178c18\"\n                    ],\n                    \"data\": \"0x0000000000000000000000000000000000000000000000000000000108da8538\",\n                    \"blockNumber\": \"0x0\",\n                    \"transactionHash\": \"0x8212c857001fc294c889faad28afad3cd92287a7a3dc580d27f56f221881072a\",\n                    \"transactionIndex\": \"0x0\",\n                    \"blockHash\": \"0x0000000000000000000000000000000000000000000000000000000000000000\",\n                    \"logIndex\": \"0x0\",\n                    \"removed\": false\n                },\n                {\n                    \"address\": \"0x582d872a1b094fc48f5de31d3b73f2d9be47def1\",\n                    \"topics\": [\n                        \"0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925\",\n                        \"0x000000000000000000000000ff8ba4d1fc3762f6154cc942ccf30049a2a0cec6\",\n                        \"0x000000000000000000000000def1c0ded9bec7f1a1670819833240f027b25eff\"\n                    ],\n                    \"data\": \"0xfffffffffffffffffffffffffffffffffffffffffffffffffffe50f44bb63acc\",\n                    \"blockNumber\": \"0x0\",\n                    \"transactionHash\": \"0x8212c857001fc294c889faad28afad3cd92287a7a3dc580d27f56f221881072a\",\n                    \"transactionIndex\": \"0x0\",\n                    \"blockHash\": \"0x0000000000000000000000000000000000000000000000000000000000000000\",\n                    \"logIndex\": \"0x0\",\n                    \"removed\": false\n                }\n            ],\n            \"value\": 0,\n            \"contract_call\": {\n                \"name\": \"transferFrom\",\n                \"params\": [\n                    {\n                        \"name\": \"src\",\n                        \"value\": \"0xff8ba4d1fc3762f6154cc942ccf30049a2a0cec6\",\n                        \"type\": \"address\"\n                    },\n                    {\n                        \"name\": \"dst\",\n                        \"value\": \"0x22f9dcf4647084d6c31b2765f6910cd85c178c18\",\n                        \"type\": \"address\"\n                    },\n                    {\n                        \"name\": \"wad\",\n                        \"value\": 4443505976,\n                        \"type\": \"uint256\"\n                    }\n                ]\n            }\n        }\n    ],\n    \"logs\": [\n        {\n            \"address\": \"0xdef1c0ded9bec7f1a1670819833240f027b25eff\",\n            \"topics\": [\n                \"0xac75f773e3a92f1a02b12134d65e1f47f8a14eabe4eaf1e24624918e6a8b269f\"\n            ],\n            \"data\": \"0xe25553e667f4d3053b4ef37683048b05c29e966879880976357a4dd8a03eedb2000000000000000000000000ff8ba4d1fc3762f6154cc942ccf30049a2a0cec600000000000000000000000022f9dcf4647084d6c31b2765f6910cd85c178c18000000000000000000000000582d872a1b094fc48f5de31d3b73f2d9be47def1000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000000000000000000000000000000000000108da85380000000000000000000000000000000000000000000000000024cb4ebf13faad\",\n            \"blockNumber\": \"0x0\",\n            \"transactionHash\": \"0x8212c857001fc294c889faad28afad3cd92287a7a3dc580d27f56f221881072a\",\n            \"transactionIndex\": \"0x0\",\n            \"blockHash\": \"0x0000000000000000000000000000000000000000000000000000000000000000\",\n            \"logIndex\": \"0x0\",\n            \"removed\": false\n        }\n    ],\n    \"value\": 0\n}"
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

		parse, err := p.ParseWithCallFrame(tradelogstype.ConvertCallFrame(&callFrame), *eventLog, uint64(time.Now().Unix()))
		require.NoError(t, err)
		t.Log(parse)
	}
}
