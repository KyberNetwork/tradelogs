package pancakeswap

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/KyberNetwork/tradelogs/pkg/rpcnode"
	"github.com/KyberNetwork/tradelogs/pkg/tracecall"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

var rpcURL = os.Getenv("TEST_RPC_URL")

func TestFetchEvent(t *testing.T) {
	t.Skip("Need to add the rpc url that enables the trace call JSON-RPC")
	ethClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		panic(err)
	}
	traceCalls := tracecall.NewCache(rpcnode.NewClient(zap.S(), ethClient))
	p := MustNewParser(traceCalls)
	assert.Equal(t, p.abi.Events[FilledEvent].ID, common.HexToHash("0x78ad7ec0e9f89e74012afa58738b6b661c024cb0fd185ee2f616c0a28924bd66"))
	logs, err := ethClient.FilterLogs(context.Background(), ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: big.NewInt(20827954),
		ToBlock:   big.NewInt(20827954),
		Addresses: nil,
		Topics: [][]common.Hash{
			{
				p.abi.Events[FilledEvent].ID,
			},
		},
	})
	assert.NoError(t, err)
	d, err := json.Marshal(logs)
	assert.NoError(t, err)
	t.Log(string(d))
}

func TestParseEvent(t *testing.T) {
	t.Skip("Need to add the rpc url that enables the trace call JSON-RPC")
	event := types.Log{}
	eventRaw := `{"address":"0x35db01d1425685789dcc9228d47c7a5c049388d8","topics":["0x78ad7ec0e9f89e74012afa58738b6b661c024cb0fd185ee2f616c0a28924bd66","0x7cd33c73177739d79731891669df1b5f27319d0c9e7f6a24d4903968dbde68c9","0x000000000000000000000000f3771ddf312bba0588c50944519b306fb0eeae16","0x0000000000000000000000009d24d495f7380ba80dc114d8c2cf1a54a68e25a4"],"data":"0x000468328d3deca81b41ae49eac5f79318c11f665cc8ac1a5f5dbbadefbc8b30","blockNumber":"0x13dcf32","transactionHash":"0x3db95e8b739ddfb2f2d967e3664a176c7ce13b2c9eeb8feaa4efb0a837245045","transactionIndex":"0xa3","blockHash":"0xb7e22e4a7d9800bce6f87ad4f1ea83aa7191f322aa7314ea732dffd12e29dcd7","logIndex":"0x204","removed":false}`
	err := json.Unmarshal([]byte(eventRaw), &event)
	assert.NoError(t, err)
	ethClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		panic(err)
	}
	traceCalls := tracecall.NewCache(rpcnode.NewClient(zap.S(), ethClient))
	p := MustNewParser(traceCalls)
	log, err := p.Parse(event, 1727271335)
	assert.NoError(t, err)
	fmt.Printf("%+v\n", log)
	assert.Equal(t, log.EventHash, p.eventHash)
	assert.Equal(t, log.Maker, "0xF3771dDF312bBA0588c50944519B306fb0EeaE16")
	assert.Equal(t, log.Taker, "0x9D24d495F7380BA80dC114D8C2cF1a54a68e25A4")
	assert.Equal(t, log.MakerTokenAmount, "40809476822677611")
	assert.Equal(t, log.MakerToken, "0x0000000000000000000000000000000000000000")
	assert.Equal(t, log.TakerTokenAmount, "118603267")
	assert.Equal(t, log.TakerToken, "0xdAC17F958D2ee523a2206206994597C13D831ec7")
	assert.Equal(t, log.Expiry, uint64(1727271435))
	t.Log(log)
}
