package pstaker //nolint: testpackage

import (
	"context"
	"encoding/json"
	"math/big"
	"testing"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/test-go/testify/require"
)

func TestFetchEvent(t *testing.T) {
	t.Skip()
	p := MustNewParser()
	require.Equal(t, p.abi.Events[SwapEvent].ID, common.HexToHash("0xe00361d207b252a464323eb23d45d42583e391f2031acdd2e9fa36efddd43cb0"))
	client, err := ethclient.Dial("")
	require.NoError(t, err)
	logs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: big.NewInt(17538476),
		ToBlock:   big.NewInt(17538476),
		Addresses: nil,
		Topics:    [][]common.Hash{{p.abi.Events[SwapEvent].ID}},
	})
	require.NoError(t, err)
	d, err := json.Marshal(logs)
	require.NoError(t, err)
	t.Log(string(d))
}

func TestParseEvent(t *testing.T) {
	eventRaw := `{
		"address":"0xdef171fe48cf0115b1d80b88dc8eab59176fee57","topics":["0xe00361d207b252a464323eb23d45d42583e391f2031acdd2e9fa36efddd43cb0","0x0000000000000000000000003a9b02c14d967333ce68e353a64653e258da2baa","0x000000000000000000000000eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee","0x0000000000000000000000007f39c581f595b53c5cb19bd0b3f8da6c935e2ca0"],"data":"0x32598393fe124476a72c13ece9ad2a7d00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000040000000000000000000000000003a9b02c14d967333ce68e353a64653e258da2baa00000000000000000000000000000000000000000000006318989ad746100000000000000000000000000000000000000000000000000057bf1e61ab00fb472b000000000000000000000000000000000000000000000057bf1e61ab00fb472b","blockNumber":"0x10b9dac","transactionHash":"0x6ff0fee62b3dc8a59e64a2df0ff810e2e5b594f593cbba98c12f621c9d594e3a","transactionIndex":"0x6a","blockHash":"0x9bdfdf6247f1969fb9c9096352bb004c0d5c66b537818def499d744eb9a54748","logIndex":"0xe4","removed":false
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

func TestFetchDirectEvent(t *testing.T) {
	t.Skip()
	p := MustNewParser()
	require.Equal(t, p.abi.Events[SwapDirect].ID, common.HexToHash("0xd2d73da2b5fd52cd654d8fd1b514ad57355bad741de639e3a1c3a20dd9f17347"))
	client, err := ethclient.Dial("")
	require.NoError(t, err)
	logs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: big.NewInt(18743024),
		ToBlock:   big.NewInt(18743024),
		Addresses: nil,
		Topics:    [][]common.Hash{{p.abi.Events[SwapDirect].ID}},
	})
	require.NoError(t, err)
	d, err := json.Marshal(logs)
	require.NoError(t, err)
	t.Log(string(d))
}

func TestParseDirectEvent(t *testing.T) {
	eventRaw := `{
		"address":"0xdef171fe48cf0115b1d80b88dc8eab59176fee57","topics":["0xd2d73da2b5fd52cd654d8fd1b514ad57355bad741de639e3a1c3a20dd9f17347","0x0000000000000000000000003a9b02c14d967333ce68e353a64653e258da2baa","0x0000000000000000000000007f39c581f595b53c5cb19bd0b3f8da6c935e2ca0","0x000000000000000000000000eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"],"data":"0x915958345a264710bbd844f9e5d98a8d00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000040000000000000000000000000003a9b02c14d967333ce68e353a64653e258da2baa0000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000015df8c7137e65d000000000000000000000000000000000000000000000000001922b86271d97b33a400000000000000000000000000000000000000000000001922b86271d97b33a4","blockNumber":"0x11dfef0","transactionHash":"0x819dc882c6f853673ea7f8d1c194df885e0dfc22aa6f1e47897b1c9e3faa1392","transactionIndex":"0x81","blockHash":"0x2f6799f7438ebb22dffaa55401de03ce86cffba156f636fad46a4a9b135573ce","logIndex":"0xce","removed":false
	}`
	event := types.Log{}
	err := json.Unmarshal([]byte(eventRaw), &event)
	require.NoError(t, err)
	p := MustNewParser()
	log, err := p.Parse(event, uint64(time.Now().Unix()))
	require.NoError(t, err)
	require.Equal(t, log.EventHash, p.directEventHash)
	t.Log(log)
}
