package zxotc

import (
	"context"
	"encoding/json"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	tradelogstype "github.com/KyberNetwork/tradelogs/pkg/types"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"

	tradingTypes "github.com/KyberNetwork/tradinglib/pkg/types"
)

const rpcURL = ""

func TestFetchEvent(t *testing.T) {
	t.Skip()
	p := MustNewParser()
	require.Equal(t, p.abi.Events[OtcOrderFilledEvent].ID, common.HexToHash("0xac75f773e3a92f1a02b12134d65e1f47f8a14eabe4eaf1e24624918e6a8b269f"))
	client, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)
	logs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: big.NewInt(16282125),
		ToBlock:   big.NewInt(16282255),
		Addresses: nil,
		Topics: [][]common.Hash{
			{
				common.HexToHash("0xac75f773e3a92f1a02b12134d65e1f47f8a14eabe4eaf1e24624918e6a8b269f"),
			},
		},
	})
	require.NoError(t, err)
	d, _ := json.Marshal(logs)
	t.Log(string(d))
}

func TestParseEvent(t *testing.T) {
	eventRaw := `{
        "address": "0xdef1c0ded9bec7f1a1670819833240f027b25eff",
        "topics": [
            "0xac75f773e3a92f1a02b12134d65e1f47f8a14eabe4eaf1e24624918e6a8b269f"
        ],
        "data": "0xb1db529b00f4a98a4c0dbf8fc7c888be9cd20da5b95136c0887aef7a6a96637d000000000000000000000000af0b0000f0210d0f421f0009c72406703b50506b00000000000000000000000074de5d4fcbf63e00296fd95d33236b9794016631000000000000000000000000c944e90c64b2c07662a292be6244bdf05cda44a7000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb4800000000000000000000000000000000000000000000124272b1b960ca00000000000000000000000000000000000000000000000000000000000001276a5f90",
        "blockNumber": "0xf87219",
        "transactionHash": "0x4153c12bbef7fcae36bed6a1b27b42f72b15302afea70e2c5aa52022f92a7eca",
        "transactionIndex": "0x15",
        "blockHash": "0x15db0c1f7d0fcf132f6598e4420bf45de265f4b9d8b168383a049f62bdcd0485",
        "logIndex": "0x3b",
        "removed": false
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

func TestParseTxEvent(t *testing.T) {
	t.Skip()
	txHash := common.HexToHash("0xd436854ddf0ecb3c087f0de1a73b62a61aae1c144eee989efa467f259d6f372b")
	ethClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		panic(err)
	}
	p := MustNewParser()

	tx, err := ethClient.TransactionReceipt(context.Background(), txHash)
	require.NoError(t, err)
	for _, log := range tx.Logs {
		if log.Topics[0].Hex() == p.eventHash {
			tradeLog, err := p.Parse(*log, tx.BlockNumber.Uint64())
			require.NoError(t, err)
			t.Log(tradeLog)
		}
	}
}

func TestGetExpiry(t *testing.T) {
	rawData := "{\n    \"from\": \"0xdef1c0ded9bec7f1a1670819833240f027b25eff\",\n    \"gas\": 158955,\n    \"gasUsed\": 62139,\n    \"to\": \"0x5ebac8dbfbba22168471b0f914131d1976536a25\",\n    \"input\": \"2sdI1AAAAAAAAAAAAAAAAFgthyobCU/Ej13jHTtz8tm+R97xAAAAAAAAAAAAAAAAwCqqObIj/o0KDlxPJ+rZCDx1bMIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABCNqFOAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAky06/E/qtAAAAAAAAAAAAAAAA/4uk0fw3YvYVTMlCzPMASaKgzsYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMOh1PDiSphW8tuEl2PP3O1VmnmxAAAAAGZ84SQAAAAAAAAAAAAAAAAAAAAAAAAAAGZ84MoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAcpqmkjALSKGVFRSEhnRjDzvBurefvbEMh83NYQ9ZNf0h0m592fRDLqXc3P4V15F2f7CRHJU60qpowCifEXdXfKAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAky06/E/qt\",\n    \"output\": \"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACTLTr8T+q0AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABCNqFOA==\",\n    \"calls\": [\n        {\n            \"from\": \"0xdef1c0ded9bec7f1a1670819833240f027b25eff\",\n            \"gas\": 150382,\n            \"gasUsed\": 3000,\n            \"to\": \"0x0000000000000000000000000000000000000001\",\n            \"input\": \"l+OB+kmS5T6BnGGHB6PMq/Qaq0INZVBDo8Vq37p8nUgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHKappIwC0ihlRUUhIZ0Yw87wbq3n72xDIfNzWEPWTX9IdJufdn0Qy6l3Nz+FdeRdn+wkRyVOtKqaMAonxF3V3yg=\",\n            \"output\": \"AAAAAAAAAAAAAAAA/4uk0fw3YvYVTMlCzPMASaKgzsY=\"\n        },\n        {\n            \"from\": \"0xdef1c0ded9bec7f1a1670819833240f027b25eff\",\n            \"gas\": 143508,\n            \"gasUsed\": 8225,\n            \"to\": \"0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2\",\n            \"input\": \"I7hy3QAAAAAAAAAAAAAAACL53PRkcITWwxsnZfaRDNhcF4wYAAAAAAAAAAAAAAAA/4uk0fw3YvYVTMlCzPMASaKgzsYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAJMtOvxP6rQ==\",\n            \"output\": \"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=\",\n            \"logs\": [\n                {\n                    \"address\": \"0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2\",\n                    \"topics\": [\n                        \"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef\",\n                        \"0x00000000000000000000000022f9dcf4647084d6c31b2765f6910cd85c178c18\",\n                        \"0x000000000000000000000000ff8ba4d1fc3762f6154cc942ccf30049a2a0cec6\"\n                    ],\n                    \"data\": \"0x0000000000000000000000000000000000000000000000000024cb4ebf13faad\",\n                    \"blockNumber\": \"0x0\",\n                    \"transactionHash\": \"0x8212c857001fc294c889faad28afad3cd92287a7a3dc580d27f56f221881072a\",\n                    \"transactionIndex\": \"0x0\",\n                    \"blockHash\": \"0x0000000000000000000000000000000000000000000000000000000000000000\",\n                    \"logIndex\": \"0x0\",\n                    \"removed\": false\n                }\n            ],\n            \"value\": 0,\n            \"contract_call\": {\n                \"name\": \"transferFrom\",\n                \"params\": [\n                    {\n                        \"name\": \"src\",\n                        \"value\": \"0x22f9dcf4647084d6c31b2765f6910cd85c178c18\",\n                        \"type\": \"address\"\n                    },\n                    {\n                        \"name\": \"dst\",\n                        \"value\": \"0xff8ba4d1fc3762f6154cc942ccf30049a2a0cec6\",\n                        \"type\": \"address\"\n                    },\n                    {\n                        \"name\": \"wad\",\n                        \"value\": 10356638235228845,\n                        \"type\": \"uint256\"\n                    }\n                ]\n            }\n        },\n        {\n            \"from\": \"0xdef1c0ded9bec7f1a1670819833240f027b25eff\",\n            \"gas\": 135048,\n            \"gasUsed\": 37307,\n            \"to\": \"0x582d872a1b094fc48f5de31d3b73f2d9be47def1\",\n            \"input\": \"I7hy3QAAAAAAAAAAAAAAAP+LpNH8N2L2FUzJQszzAEmioM7GAAAAAAAAAAAAAAAAIvnc9GRwhNbDGydl9pEM2FwXjBgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABCNqFOA==\",\n            \"output\": \"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=\",\n            \"logs\": [\n                {\n                    \"address\": \"0x582d872a1b094fc48f5de31d3b73f2d9be47def1\",\n                    \"topics\": [\n                        \"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef\",\n                        \"0x000000000000000000000000ff8ba4d1fc3762f6154cc942ccf30049a2a0cec6\",\n                        \"0x00000000000000000000000022f9dcf4647084d6c31b2765f6910cd85c178c18\"\n                    ],\n                    \"data\": \"0x0000000000000000000000000000000000000000000000000000000108da8538\",\n                    \"blockNumber\": \"0x0\",\n                    \"transactionHash\": \"0x8212c857001fc294c889faad28afad3cd92287a7a3dc580d27f56f221881072a\",\n                    \"transactionIndex\": \"0x0\",\n                    \"blockHash\": \"0x0000000000000000000000000000000000000000000000000000000000000000\",\n                    \"logIndex\": \"0x0\",\n                    \"removed\": false\n                },\n                {\n                    \"address\": \"0x582d872a1b094fc48f5de31d3b73f2d9be47def1\",\n                    \"topics\": [\n                        \"0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925\",\n                        \"0x000000000000000000000000ff8ba4d1fc3762f6154cc942ccf30049a2a0cec6\",\n                        \"0x000000000000000000000000def1c0ded9bec7f1a1670819833240f027b25eff\"\n                    ],\n                    \"data\": \"0xfffffffffffffffffffffffffffffffffffffffffffffffffffe50f44bb63acc\",\n                    \"blockNumber\": \"0x0\",\n                    \"transactionHash\": \"0x8212c857001fc294c889faad28afad3cd92287a7a3dc580d27f56f221881072a\",\n                    \"transactionIndex\": \"0x0\",\n                    \"blockHash\": \"0x0000000000000000000000000000000000000000000000000000000000000000\",\n                    \"logIndex\": \"0x0\",\n                    \"removed\": false\n                }\n            ],\n            \"value\": 0,\n            \"contract_call\": {\n                \"name\": \"transferFrom\",\n                \"params\": [\n                    {\n                        \"name\": \"src\",\n                        \"value\": \"0xff8ba4d1fc3762f6154cc942ccf30049a2a0cec6\",\n                        \"type\": \"address\"\n                    },\n                    {\n                        \"name\": \"dst\",\n                        \"value\": \"0x22f9dcf4647084d6c31b2765f6910cd85c178c18\",\n                        \"type\": \"address\"\n                    },\n                    {\n                        \"name\": \"wad\",\n                        \"value\": 4443505976,\n                        \"type\": \"uint256\"\n                    }\n                ]\n            }\n        }\n    ],\n    \"logs\": [\n        {\n            \"address\": \"0xdef1c0ded9bec7f1a1670819833240f027b25eff\",\n            \"topics\": [\n                \"0xac75f773e3a92f1a02b12134d65e1f47f8a14eabe4eaf1e24624918e6a8b269f\"\n            ],\n            \"data\": \"0xe25553e667f4d3053b4ef37683048b05c29e966879880976357a4dd8a03eedb2000000000000000000000000ff8ba4d1fc3762f6154cc942ccf30049a2a0cec600000000000000000000000022f9dcf4647084d6c31b2765f6910cd85c178c18000000000000000000000000582d872a1b094fc48f5de31d3b73f2d9be47def1000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000000000000000000000000000000000000108da85380000000000000000000000000000000000000000000000000024cb4ebf13faad\",\n            \"blockNumber\": \"0x0\",\n            \"transactionHash\": \"0x8212c857001fc294c889faad28afad3cd92287a7a3dc580d27f56f221881072a\",\n            \"transactionIndex\": \"0x0\",\n            \"blockHash\": \"0x0000000000000000000000000000000000000000000000000000000000000000\",\n            \"logIndex\": \"0x0\",\n            \"removed\": false\n        }\n    ],\n    \"value\": 0\n}"
	var callFrame tradingTypes.CallFrame
	p := MustNewParser()
	err := json.Unmarshal([]byte(rawData), &callFrame)
	require.NoError(t, err)
	rfqOrderParam, err := p.getRFQOrderParams(tradelogstype.ConvertCallFrame(&callFrame))
	require.NoError(t, err)
	require.NotNil(t, rfqOrderParam)
	require.Equal(t, uint64(1719460132), rfqOrderParam.GetExpiry())
}

func TestParseWithCallFrame(t *testing.T) {
	t.Skip("Need to add the rpc url that enables the trace call JSON-RPC")
	expectedTradeLog := storage.TradeLog{
		OrderHash:        "0xceefc26698e9c77d3bff738449a28f2a29c72ffc93392c4d8a156b56f8d03a22",
		Maker:            "0xff8Ba4D1fC3762f6154cc942CCF30049A2A0cEC6",
		Taker:            "0x22F9dCF4647084d6C31b2765F6910cd85C178C18",
		MakerToken:       "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
		TakerToken:       "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
		MakerTokenAmount: "28332999",
		TakerTokenAmount: "8453527433100000",
		ContractAddress:  "0xDef1C0ded9bec7F1a1670819833240f027b25EfF",
		BlockNumber:      20181990,
		TxHash:           "0x2ae57498f98fec3d5d053ade3d9dfdd6d5ec6c66e9c0a18fad836b1c9a2dfb3a",
		LogIndex:         342,
		Timestamp:        0,
		EventHash:        "0xac75f773e3a92f1a02b12134d65e1f47f8a14eabe4eaf1e24624918e6a8b269f",
		Expiry:           1719478903,
	}
	p := MustNewParser()
	client, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)

	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash("0x2ae57498f98fec3d5d053ade3d9dfdd6d5ec6c66e9c0a18fad836b1c9a2dfb3a"))
	require.NoError(t, err)

	rawData := "{\"from\":\"0xdef1c0ded9bec7f1a1670819833240f027b25eff\",\"gas\":136080,\"gasUsed\":66881,\"to\":\"0x5ebac8dbfbba22168471b0f914131d1976536a25\",\"input\":\"2sdI1AAAAAAAAAAAAAAAAKC4aZHGIYs2wdGdSi6esM42ButIAAAAAAAAAAAAAAAAwCqqObIj/o0KDlxPJ+rZCDx1bMIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAbBTxwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAeCHA7v7rgAAAAAAAAAAAAAAAA/4uk0fw3YvYVTMlCzPMASaKgzsYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACeVG/rr0m8S3M1MpuroUFGr0HYWAAAAAGZ9KncAAAAAAAAAAAAAAAAAAAAAAAAAAGZ9Kh0AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAcNVhjiKaMuCcp2KNj2r7VV7jYuz1awBo8HbxWpcVyPsdMNI/JKnbYiSrwT9s1mIhew8tyPtiUrxbPx7uzXi5WegAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAeCHA7v7rg\",\"output\":\"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAB4IcDu/uuAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAbBTxw==\",\"calls\":[{\"from\":\"0xdef1c0ded9bec7f1a1670819833240f027b25eff\",\"gas\":127865,\"gasUsed\":3000,\"to\":\"0x0000000000000000000000000000000000000001\",\"input\":\"IX5aqefvovx38NwqMliCKhLUABcMpoucCzMs/GLd7FUAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHDVYY4imjLgnKdijY9q+1Ve42Ls9WsAaPB28VqXFcj7HTDSPySp22Ikq8E/bNZiIXsPLcj7YlK8Wz8e7s14uVno=\",\"output\":\"AAAAAAAAAAAAAAAA/4uk0fw3YvYVTMlCzPMASaKgzsY=\"},{\"from\":\"0xdef1c0ded9bec7f1a1670819833240f027b25eff\",\"gas\":120990,\"gasUsed\":8225,\"to\":\"0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2\",\"input\":\"I7hy3QAAAAAAAAAAAAAAACL53PRkcITWwxsnZfaRDNhcF4wYAAAAAAAAAAAAAAAA/4uk0fw3YvYVTMlCzPMASaKgzsYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHghwO7+64A==\",\"output\":\"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=\",\"logs\":[{\"address\":\"0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2\",\"topics\":[\"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef\",\"0x00000000000000000000000022f9dcf4647084d6c31b2765f6910cd85c178c18\",\"0x000000000000000000000000ff8ba4d1fc3762f6154cc942ccf30049a2a0cec6\"],\"data\":\"0x000000000000000000000000000000000000000000000000001e08703bbfbae0\",\"blockNumber\":\"0x0\",\"transactionHash\":\"0x2ae57498f98fec3d5d053ade3d9dfdd6d5ec6c66e9c0a18fad836b1c9a2dfb3a\",\"transactionIndex\":\"0x0\",\"blockHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"logIndex\":\"0x0\",\"removed\":false}],\"value\":0,\"contract_call\":{\"name\":\"transferFrom\",\"params\":[{\"name\":\"src\",\"value\":\"0x22f9dcf4647084d6c31b2765f6910cd85c178c18\",\"type\":\"address\"},{\"name\":\"dst\",\"value\":\"0xff8ba4d1fc3762f6154cc942ccf30049a2a0cec6\",\"type\":\"address\"},{\"name\":\"wad\",\"value\":8453527433100000,\"type\":\"uint256\"}]}},{\"from\":\"0xdef1c0ded9bec7f1a1670819833240f027b25eff\",\"gas\":112530,\"gasUsed\":42049,\"to\":\"0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48\",\"input\":\"I7hy3QAAAAAAAAAAAAAAAP+LpNH8N2L2FUzJQszzAEmioM7GAAAAAAAAAAAAAAAAIvnc9GRwhNbDGydl9pEM2FwXjBgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAbBTxw==\",\"output\":\"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=\",\"calls\":[{\"from\":\"0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48\",\"gas\":110033,\"gasUsed\":41254,\"to\":\"0x43506849d7c04f9138d1a2050bbf3a0c054402dd\",\"input\":\"I7hy3QAAAAAAAAAAAAAAAP+LpNH8N2L2FUzJQszzAEmioM7GAAAAAAAAAAAAAAAAIvnc9GRwhNbDGydl9pEM2FwXjBgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAbBTxw==\",\"output\":\"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=\",\"logs\":[{\"address\":\"0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48\",\"topics\":[\"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef\",\"0x000000000000000000000000ff8ba4d1fc3762f6154cc942ccf30049a2a0cec6\",\"0x00000000000000000000000022f9dcf4647084d6c31b2765f6910cd85c178c18\"],\"data\":\"0x0000000000000000000000000000000000000000000000000000000001b053c7\",\"blockNumber\":\"0x0\",\"transactionHash\":\"0x2ae57498f98fec3d5d053ade3d9dfdd6d5ec6c66e9c0a18fad836b1c9a2dfb3a\",\"transactionIndex\":\"0x0\",\"blockHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"logIndex\":\"0x0\",\"removed\":false}],\"value\":0,\"contract_call\":{\"name\":\"transferFrom\",\"params\":[{\"name\":\"src\",\"value\":\"0xff8ba4d1fc3762f6154cc942ccf30049a2a0cec6\",\"type\":\"address\"},{\"name\":\"dst\",\"value\":\"0x22f9dcf4647084d6c31b2765f6910cd85c178c18\",\"type\":\"address\"},{\"name\":\"wad\",\"value\":28332999,\"type\":\"uint256\"}]}}],\"value\":0,\"contract_call\":{\"name\":\"transferFrom\",\"params\":[{\"name\":\"src\",\"value\":\"0xff8ba4d1fc3762f6154cc942ccf30049a2a0cec6\",\"type\":\"address\"},{\"name\":\"dst\",\"value\":\"0x22f9dcf4647084d6c31b2765f6910cd85c178c18\",\"type\":\"address\"},{\"name\":\"wad\",\"value\":28332999,\"type\":\"uint256\"}]}}],\"logs\":[{\"address\":\"0xdef1c0ded9bec7f1a1670819833240f027b25eff\",\"topics\":[\"0xac75f773e3a92f1a02b12134d65e1f47f8a14eabe4eaf1e24624918e6a8b269f\"],\"data\":\"0xceefc26698e9c77d3bff738449a28f2a29c72ffc93392c4d8a156b56f8d03a22000000000000000000000000ff8ba4d1fc3762f6154cc942ccf30049a2a0cec600000000000000000000000022f9dcf4647084d6c31b2765f6910cd85c178c18000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000000000000000000000000000000000000001b053c7000000000000000000000000000000000000000000000000001e08703bbfbae0\",\"blockNumber\":\"0x0\",\"transactionHash\":\"0x2ae57498f98fec3d5d053ade3d9dfdd6d5ec6c66e9c0a18fad836b1c9a2dfb3a\",\"transactionIndex\":\"0x0\",\"blockHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"logIndex\":\"0x0\",\"removed\":false}],\"value\":0}"
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

		parse, err := p.ParseWithCallFrame(tradelogstype.ConvertCallFrame(&callFrame), *eventLog, 0)
		require.NoError(t, err)
		t.Log(parse)
		require.Equal(t, expectedTradeLog, parse)
	}
}
