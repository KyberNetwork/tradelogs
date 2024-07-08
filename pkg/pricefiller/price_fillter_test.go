package pricefiller

import (
	"testing"

	"github.com/KyberNetwork/go-binance/v2"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"github.com/test-go/testify/require"
)

// go test -v -timeout 30s -run ^TestFillPrice$ github.com/KyberNetwork/tradelogs/pkg/pricefiller
func TestFillPrice(t *testing.T) {
	t.Skip("Need to add Binance credentials")
	bClient := binance.NewClient("", "")
	filler, err := NewPriceFiller(bClient, nil)
	if err != nil {
		require.NoError(t, err)
	}

	tradeLogs := []storage.TradeLog{
		{
			Taker:            "0x807cf9a772d5a3f9cefbc1192e939d62f0d9bd38",
			MakerToken:       "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
			TakerToken:       "0x320623b8e4ff03373931769a31fc52a4e78b5d70",
			MakerTokenAmount: "503568522079108960",
			TakerTokenAmount: "336970435721651800000000",
			ContractAddress:  "0x6131b5fae19ea4f9d964eac0408e4408b66337b5",
			BlockNumber:      20230391,
			TxHash:           "0x0aad2e38b90390d6d060a5886ddd20d312c6beb3a305b2360bd6d25593ddf058",
			LogIndex:         57,
			Timestamp:        1720062815000,
			EventHash:        "0xd6d4f5681c246c9f42c203e287975af1601f8df8035a9251f79aab5c8f09e2f8",
		},
		{
			Taker:            "0x807cf9a772d5a3f9cefbc1192e939d62f0d9bd38",
			MakerToken:       "0x7fc66500c84a76ad7e9c93437bfc5ac33e2ddae9",
			TakerToken:       "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
			MakerTokenAmount: "179003122862979007021",
			TakerTokenAmount: "4688506567482331000",
			ContractAddress:  "0x6131b5fae19ea4f9d964eac0408e4408b66337b5",
			BlockNumber:      20230305,
			TxHash:           "0x21b8f97e43ff6debbe2ab323f079f0a936c14c16bdb6693587b48a5d66dcc37c",
			LogIndex:         136,
			Timestamp:        1720061783000,
			EventHash:        "0xd6d4f5681c246c9f42c203e287975af1601f8df8035a9251f79aab5c8f09e2f8",
		},
		{
			Taker:            "0x807cf9a772d5a3f9cefbc1192e939d62f0d9bd38",
			MakerToken:       "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
			TakerToken:       "0x514910771af9ca656af840dff83e8264ecf986ca",
			MakerTokenAmount: "783414511682884466",
			TakerTokenAmount: "190359937738916760000",
			ContractAddress:  "0x6131b5fae19ea4f9d964eac0408e4408b66337b5",
			BlockNumber:      20230291,
			TxHash:           "0x2ff00bcfa69c85fdbd0c7fc9b193717009751a32cf661c31815d9745deb68552",
			LogIndex:         136,
			Timestamp:        1720061615000,
			EventHash:        "0xd6d4f5681c246c9f42c203e287975af1601f8df8035a9251f79aab5c8f09e2f8",
		},
	}
	filler.FullFillTradeLogs(tradeLogs)

	t.Log(tradeLogs)
}
