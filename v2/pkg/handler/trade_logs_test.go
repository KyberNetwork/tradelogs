package handler

import (
	"context"
	"github.com/KyberNetwork/tradelogs/v2/mocks"
	"github.com/KyberNetwork/tradelogs/v2/pkg/parser"
	zxotc2 "github.com/KyberNetwork/tradelogs/v2/pkg/parser/zxotc"
	"github.com/KyberNetwork/tradelogs/v2/pkg/rpcnode"
	"github.com/KyberNetwork/tradelogs/v2/pkg/storage"
	"github.com/KyberNetwork/tradelogs/v2/pkg/storage/types"
	types2 "github.com/KyberNetwork/tradelogs/v2/pkg/types"
	types3 "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	"os"
	"testing"
)

var rpcURL = os.Getenv("TEST_RPC_URL")

func TestTradeLogHandler_ProcessBlock(t *testing.T) {
	t.Skip("Need to set the rpc url env that enables the trace call JSON-RPC")

	ethClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		t.Fatal(err)
	}
	client := rpcnode.NewClient(zap.S(), ethClient)

	mockStorage := &mocks.MockStorage{}
	mockStorage.On("Type").Return("zerox").
		On("Insert", mock.Anything).Return(nil)
	s := storage.NewManager(zap.S(), []types.Storage{mockStorage})

	p := zxotc2.MustNewParser()

	mockKafka := &mocks.MockPublisher{}
	mockKafka.On("Publish", mock.Anything, mock.Anything).Return(nil)

	h := NewTradeLogHandler(zap.S(), client, s, []parser.Parser{p}, "test", mockKafka)

	err = h.ProcessBlock("0x04b65fabd0eaaa00eae00782128a8add39e30098552738c305610259f14ea048", 20181990, 1725436442)
	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, mockStorage.AssertNumberOfCalls(t, "Insert", 1))
	assert.True(t, mockKafka.AssertNumberOfCalls(t, "Publish", 1))
}

func TestAssignLogIndexes(t *testing.T) {
	t.Skip("Need to set the rpc url env that enables the trace call JSON-RPC")

	ethClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		t.Fatal(err)
	}
	client := rpcnode.NewClient(zap.S(), ethClient)

	blockHash := "0x70c955e0b8fb254bf160445a8db0afd25b4429b513d3f6f8189eadca24ecacb4"

	traceCalls, err := client.FetchTraceCalls(context.Background(), blockHash)
	assert.Nil(t, err)

	logs, err := client.FetchLogsByBlockHash(context.Background(), blockHash)
	assert.Nil(t, err)

	id := 0
	callLogs := make([]types2.CallLog, 0)
	for _, call := range traceCalls {
		id = assignLogIndexes(&call.CallFrame, id)
		callLogs = append(callLogs, getLogs(call.CallFrame)...)
	}

	assert.Equal(t, len(logs), len(callLogs))

	logsMap := make(map[int]types3.Log, len(logs))
	for _, log := range logs {
		logsMap[int(log.Index)] = log
	}

	callLogsMap := make(map[int]types2.CallLog, len(callLogs))
	for _, log := range callLogs {
		callLogsMap[log.Index] = log
	}

	for i := range logsMap {
		assert.Equal(t, logsMap[i].Index, uint(callLogsMap[i].Index))
		assert.Equal(t, logsMap[i].Address, callLogsMap[i].Address)
		assert.Equal(t, logsMap[i].Topics, callLogsMap[i].Topics)
	}
}

func getLogs(call types2.CallFrame) []types2.CallLog {
	result := make([]types2.CallLog, 0)
	for _, subCall := range call.Calls {
		result = append(result, getLogs(subCall)...)
	}
	result = append(result, call.Logs...)
	return result
}