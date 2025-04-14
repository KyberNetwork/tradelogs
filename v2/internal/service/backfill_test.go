package service

import (
	"testing"

	"github.com/KyberNetwork/tradelogs/v2/internal/worker"
	"github.com/KyberNetwork/tradelogs/v2/mocks"
	"github.com/KyberNetwork/tradelogs/v2/pkg/parser"
	"github.com/KyberNetwork/tradelogs/v2/pkg/storage/backfill"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

func TestBackfill_NewBackfill(t *testing.T) {
	backfillStorage := &mocks.MockBackfillStorage{}
	stateStorage := &mocks.MockState{}
	rpcClient := &mocks.MockRPCClient{}

	backfillStorage.On("GetTask").Return([]backfill.Task{}, nil)

	w := worker.NewBackFiller(nil, nil, nil, backfillStorage, stateStorage, zap.S(), rpcClient, nil)
	_, err := NewBackfillService(backfillStorage, zap.S(), w)
	assert.NoError(t, err)

	backfillStorage.AssertCalled(t, "GetTask")
}

func TestBackfill_NewBackfillTask(t *testing.T) {
	backfillStorage := &mocks.MockBackfillStorage{}
	stateStorage := &mocks.MockState{}
	rpcClient := &mocks.MockRPCClient{}
	mockParser := &mocks.MockParser{}

	backfillStorage.On("GetTask").Return([]backfill.Task{}, nil).
		On("CreateTask", mock.Anything).Return(0, nil).
		On("Get").Return(
		[]backfill.State{
			{
				Exchange:        "test",
				DeployBlock:     20962650,
				BackFilledBlock: 20962657,
			},
		}, nil).
		On("GetRunningTaskNumber").Return(1, nil).
		On("UpdateTask", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	rpcClient.On("FetchLogs", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return([]types.Log{}, nil)

	mockParser.On("Exchange").Return("test").
		On("Address").Return("").
		On("Topics").Return([]string{})

	w := worker.NewBackFiller(nil, nil, nil, backfillStorage, stateStorage, zap.S(), rpcClient, []parser.Parser{mockParser})
	srv, err := NewBackfillService(backfillStorage, zap.S(), w)
	assert.NoError(t, err)

	testcases := []struct {
		name         string
		from         uint64
		to           uint64
		exchange     string
		expectedFrom uint64
		expectedTo   uint64
		hasError     bool
	}{
		{
			name:     "fully covered",
			from:     20962650,
			to:       20962657,
			hasError: true,
		},
		{
			name:     "invalid from",
			from:     20962645,
			to:       20962657,
			hasError: true,
		},
		{
			name:         "partially covered",
			from:         20962655,
			to:           20962659,
			exchange:     "test",
			expectedFrom: 20962657,
			expectedTo:   20962659,
			hasError:     false,
		},
		{
			name:         "normal case",
			from:         20962660,
			to:           20962665,
			exchange:     "test",
			expectedFrom: 20962660,
			expectedTo:   20962665,
			hasError:     false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			_, _, err = srv.NewBackfillTask(tc.from, tc.to, tc.exchange)
			if tc.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				task := backfill.Task{
					Exchange:  tc.exchange,
					FromBlock: tc.expectedFrom,
					ToBlock:   tc.expectedTo,
				}
				backfillStorage.AssertCalled(t, "CreateTask", task)
			}
		})
	}
}

func TestBackfill_ListTask(t *testing.T) {
	backfillStorage := &mocks.MockBackfillStorage{}
	stateStorage := &mocks.MockState{}
	rpcClient := &mocks.MockRPCClient{}

	backfillStorage.On("GetTask").Return(
		[]backfill.Task{
			{
				ID:             1,
				Exchange:       "test",
				FromBlock:      20962660,
				ToBlock:        20962667,
				ProcessedBlock: 20962665,
				Status:         backfill.StatusTypeDone,
			},
		}, nil).
		On("CreateTask", mock.Anything).Return(0, nil).
		On("Get").Return(
		[]backfill.State{
			{
				Exchange:        "test",
				DeployBlock:     20962650,
				BackFilledBlock: 20962657,
			},
		}, nil)

	rpcClient.On("FetchLogs", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return([]types.Log{}, nil)

	w := worker.NewBackFiller(nil, nil, nil, backfillStorage, stateStorage, zap.S(), rpcClient, nil)
	srv, err := NewBackfillService(backfillStorage, zap.S(), w)
	assert.NoError(t, err)

	tasks, err := srv.ListTask()
	assert.NoError(t, err)
	assert.Len(t, tasks, 2)

	expectedTasks := []backfill.Task{
		{
			ID:             -1,
			Exchange:       "",
			FromBlock:      20962650,
			ToBlock:        20962657,
			ProcessedBlock: 20962657,
			Status:         backfill.StatusTypeProcessing,
		},
		{
			ID:             1,
			Exchange:       "test",
			FromBlock:      20962660,
			ToBlock:        20962667,
			ProcessedBlock: 20962665,
			Status:         backfill.StatusTypeDone,
		},
	}
	assert.Equal(t, expectedTasks, tasks)
}
