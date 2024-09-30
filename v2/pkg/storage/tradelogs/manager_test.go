package tradelogs

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/KyberNetwork/tradelogs/v2/mocks"
	"github.com/KyberNetwork/tradelogs/v2/pkg/storage/tradelogs/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

func TestSimple(t *testing.T) {
	l := zap.S()
	mockStorage := &mocks.MockStorage{}
	mockStorage.On("Exchange").Return("test").
		On("Insert", mock.Anything).Return(nil).
		On("Get", mock.Anything).Return(nil, nil)
	s := NewManager(l, []types.Storage{mockStorage})
	var tradeLogs []types.TradeLog
	byteValue, err := os.Open("test.json")
	assert.NoError(t, err)
	assert.NoError(t, json.NewDecoder(byteValue).Decode(&tradeLogs), "failed to parse tradelogs")
	for i := range tradeLogs {
		tradeLogs[i].Exchange = "test"
	}
	assert.NoError(t, s.Insert(tradeLogs), "failed to insert tradelogs")
	assert.True(t, mockStorage.AssertNumberOfCalls(t, "Insert", 1))

	_, err = s.Get(types.TradeLogsQuery{})
	assert.NoError(t, err, "failed to query tradelogs")
	assert.True(t, mockStorage.AssertNumberOfCalls(t, "Get", 1))
}
