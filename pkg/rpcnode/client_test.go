package rpcnode

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
)

func TestGetSender(t *testing.T) {
	t.Skip("need rpc url")
	rpcURL := ""
	ethClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		panic(err)
	}
	rpcnode := NewClient(nil, ethClient)
	from, err := rpcnode.GetTxOriginByTxHash(context.Background(), "0x9570e65ab98a007f69e485a8c90ec7256cffedb9d7983cf8d2f0a566ea3c46a3")
	if err != nil {
		t.Log(err)
	}
	t.Log(from)
}
