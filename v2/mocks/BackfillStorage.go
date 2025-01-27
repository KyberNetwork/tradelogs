// Code generated by mockery v2.46.2. DO NOT EDIT.

package mocks

import (
	backfill "github.com/KyberNetwork/tradelogs/v2/pkg/storage/backfill"
	mock "github.com/stretchr/testify/mock"
)

// MockBackfillStorage is an autogenerated mock type for the IStorage type
type MockBackfillStorage struct {
	mock.Mock
}

// CreateTask provides a mock function with given fields: task
func (_m *MockBackfillStorage) CreateTask(task backfill.Task) (int, error) {
	ret := _m.Called(task)

	if len(ret) == 0 {
		panic("no return value specified for CreateTask")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(backfill.Task) (int, error)); ok {
		return rf(task)
	}
	if rf, ok := ret.Get(0).(func(backfill.Task) int); ok {
		r0 = rf(task)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(backfill.Task) error); ok {
		r1 = rf(task)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTask provides a mock function with given fields: taskID
func (_m *MockBackfillStorage) DeleteTask(taskID int) error {
	ret := _m.Called(taskID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTask")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(taskID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields:
func (_m *MockBackfillStorage) Get() ([]backfill.State, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 []backfill.State
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]backfill.State, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []backfill.State); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]backfill.State)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRunningTaskNumber provides a mock function with given fields:
func (_m *MockBackfillStorage) GetRunningTaskNumber() (int, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetRunningTaskNumber")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func() (int, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTask provides a mock function with given fields:
func (_m *MockBackfillStorage) GetTask() ([]backfill.Task, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetTask")
	}

	var r0 []backfill.Task
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]backfill.Task, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []backfill.Task); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]backfill.Task)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTaskByID provides a mock function with given fields: taskID
func (_m *MockBackfillStorage) GetTaskByID(taskID int) (backfill.Task, error) {
	ret := _m.Called(taskID)

	if len(ret) == 0 {
		panic("no return value specified for GetTaskByID")
	}

	var r0 backfill.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (backfill.Task, error)); ok {
		return rf(taskID)
	}
	if rf, ok := ret.Get(0).(func(int) backfill.Task); ok {
		r0 = rf(taskID)
	} else {
		r0 = ret.Get(0).(backfill.Task)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(taskID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: backfilled
func (_m *MockBackfillStorage) Update(backfilled uint64) error {
	ret := _m.Called(backfilled)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(backfilled)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTask provides a mock function with given fields: id, block, status
func (_m *MockBackfillStorage) UpdateTask(id int, block *uint64, status string) error {
	ret := _m.Called(id, block, status)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTask")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int, *uint64, string) error); ok {
		r0 = rf(id, block, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockBackfillStorage creates a new instance of MockBackfillStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockBackfillStorage(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockBackfillStorage {
	mock := &MockBackfillStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
