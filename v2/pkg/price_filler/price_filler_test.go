package pricefiller

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/KyberNetwork/tradelogs/v2/pkg/mtm"
	"github.com/KyberNetwork/tradelogs/v2/pkg/storage/tradelogs/types"
	"github.com/test-go/testify/require"
	"go.uber.org/zap"
)

// go test -v -timeout 30s -run ^TestFillPrice$ github.com/KyberNetwork/tradelogs/pkg/pricefiller
func TestFillPrice(t *testing.T) {
	t.Skip("Need to add mtm url")
	httpClient := &http.Client{}
	mtmClient := mtm.NewMtmClient("", httpClient)
	filler, err := NewPriceFiller(zap.S(), nil, mtmClient, nil)
	require.NoError(t, err)

	tradeLogs := []types.TradeLog{
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

	m, _ := json.Marshal(tradeLogs)
	t.Log(string(m))
}

func TestFillBebopPrice(t *testing.T) {
	t.Skip("Need to add mtm url")
	httpClient := &http.Client{}
	mtmClient := mtm.NewMtmClient("", httpClient)
	filler, err := NewPriceFiller(zap.S(), nil, mtmClient, nil)
	require.NoError(t, err)

	rawTradelogs := []byte(`[{"exchange":"","order_hash":"52344126666663492359326433927181107201","maker":"0x51c72848c68a965f66fa7a88855f9f7784502a7f","taker":"0x589c86cc1f6043f99222843d397ea4e770841cae","maker_token":"0xba100000625a3754423978a60c9317c58a424e3d","taker_token":"0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2","maker_token_amount":"418270652861628178","taker_token_amount":"3500000000000000","contract_address":"0xbbbbbBB520d69a9775E85b458C58c648259FAD5F","block_number":20153388,"tx_hash":"0xff46ac555ec7da7aa484864dc0df90217b7f46dfa51627c20ef6e25451c64b15","log_index":133,"timestamp":1730258461000,"event_hash":"0xadd7095becdaa725f0f33243630938c861b0bba83dfd217d4055701aa768ec2e","expiration_date":1719133686,"maker_token_price":0,"taker_token_price":0,"maker_usd_amount":0,"taker_usd_amount":0,"state":""},{"exchange":"","order_hash":"52344126666663492359326433927181107201","maker":"0x51c72848c68a965f66fa7a88855f9f7784502a7f","taker":"0x589c86cc1f6043f99222843d397ea4e770841cae","maker_token":"0x6b175474e89094c44da98b954eedeac495271d0f","taker_token":"0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2","maker_token_amount":"1320463992019695220","taker_token_amount":"3500000000000000","contract_address":"0xbbbbbBB520d69a9775E85b458C58c648259FAD5F","block_number":20153388,"tx_hash":"0xff46ac555ec7da7aa484864dc0df90217b7f46dfa51627c20ef6e25451c64b15","log_index":133,"trade_index":1,"timestamp":1730258461000,"event_hash":"0xadd7095becdaa725f0f33243630938c861b0bba83dfd217d4055701aa768ec2e","expiration_date":1719133686,"maker_token_price":0,"taker_token_price":0,"maker_usd_amount":0,"taker_usd_amount":0,"state":""},{"exchange":"","order_hash":"52344126666663492359326433927181107201","maker":"0x51c72848c68a965f66fa7a88855f9f7784502a7f","taker":"0x589c86cc1f6043f99222843d397ea4e770841cae","maker_token":"0xc00e94cb662c3520282e6f5717214004a7f26888","taker_token":"0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2","maker_token_amount":"26747477356835299","taker_token_amount":"3500000000000000","contract_address":"0xbbbbbBB520d69a9775E85b458C58c648259FAD5F","block_number":20153388,"tx_hash":"0xff46ac555ec7da7aa484864dc0df90217b7f46dfa51627c20ef6e25451c64b15","log_index":133,"trade_index":2,"timestamp":1730258461000,"event_hash":"0xadd7095becdaa725f0f33243630938c861b0bba83dfd217d4055701aa768ec2e","expiration_date":1719133686,"maker_token_price":0,"taker_token_price":0,"maker_usd_amount":0,"taker_usd_amount":0,"state":""},{"exchange":"","order_hash":"52344126666663492359326433927181107201","maker":"0x51c72848c68a965f66fa7a88855f9f7784502a7f","taker":"0x589c86cc1f6043f99222843d397ea4e770841cae","maker_token":"0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48","taker_token":"0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2","maker_token_amount":"1320537","taker_token_amount":"3500000000000000","contract_address":"0xbbbbbBB520d69a9775E85b458C58c648259FAD5F","block_number":20153388,"tx_hash":"0xff46ac555ec7da7aa484864dc0df90217b7f46dfa51627c20ef6e25451c64b15","log_index":133,"trade_index":3,"timestamp":1730258461000,"event_hash":"0xadd7095becdaa725f0f33243630938c861b0bba83dfd217d4055701aa768ec2e","expiration_date":1719133686,"maker_token_price":0,"taker_token_price":0,"maker_usd_amount":0,"taker_usd_amount":0,"state":""},{"exchange":"","order_hash":"52344126666663492359326433927181107201","maker":"0x51c72848c68a965f66fa7a88855f9f7784502a7f","taker":"0x589c86cc1f6043f99222843d397ea4e770841cae","maker_token":"0x6b3595068778dd592e39a122f4f5a5cf09c90fe2","taker_token":"0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2","maker_token_amount":"1614402915976888452","taker_token_amount":"3500000000000000","contract_address":"0xbbbbbBB520d69a9775E85b458C58c648259FAD5F","block_number":20153388,"tx_hash":"0xff46ac555ec7da7aa484864dc0df90217b7f46dfa51627c20ef6e25451c64b15","log_index":133,"trade_index":4,"timestamp":1730258461000,"event_hash":"0xadd7095becdaa725f0f33243630938c861b0bba83dfd217d4055701aa768ec2e","expiration_date":1719133686,"maker_token_price":0,"taker_token_price":0,"maker_usd_amount":0,"taker_usd_amount":0,"state":""}]`)

	var tradeLogs []types.TradeLog
	_ = json.Unmarshal(rawTradelogs, &tradeLogs)
	tradeLogs = []types.TradeLog{
		{
			Exchange:         "bebop",
			Taker:            "0x589c86cc1f6043f99222843d397ea4e770841cae",
			MakerToken:       "0xba100000625a3754423978a60c9317c58a424e3d,0x6b175474e89094c44da98b954eedeac495271d0f,0xc00e94cb662c3520282e6f5717214004a7f26888,0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48,0x6b3595068778dd592e39a122f4f5a5cf09c90fe2",
			TakerToken:       "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
			MakerTokenAmount: "418270652861628178,1320463992019695220,26747477356835299,1320537,1614402915976888452",
			TakerTokenAmount: "3500000000000000",
			ContractAddress:  "0xbbbbbBB520d69a9775E85b458C58c648259FAD5F",
			BlockNumber:      20153388,
			TxHash:           "0xff46ac555ec7da7aa484864dc0df90217b7f46dfa51627c20ef6e25451c64b15",
			LogIndex:         133,
			Timestamp:        1730344560000,
			EventHash:        "0xadd7095becdaa725f0f33243630938c861b0bba83dfd217d4055701aa768ec2e",
		},
	}
	filler.FullFillTradeLogs(tradeLogs)

	m, _ := json.Marshal(tradeLogs)
	t.Log(string(m))
}
