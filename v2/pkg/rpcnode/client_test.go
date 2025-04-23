package rpcnode_test

import (
	"context"
	"testing"

	"github.com/KyberNetwork/tradelogs/v2/pkg/rpcnode"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestGetBlockByTxHash(t *testing.T) {
	t.Skip("Need to add the rpc url that enables the trace call JSON-RPC")
	l := zap.S()
	ethClient, err := ethclient.Dial("rpcNode")
	assert.NoError(t, err)

	client := rpcnode.NewClient(l, ethClient)
	blockNumber, err := client.BlockByTxHash(context.Background(), "0x918c85c01c531264a2f20347644c6b136487004a3161ef2609644a5249057e9a")
	assert.NoError(t, err)
	t.Log(blockNumber)
}
