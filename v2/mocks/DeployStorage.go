// Code generated by mockery v2.46.2. DO NOT EDIT.

package mocks

import (
	zerox_deployment "github.com/KyberNetwork/tradelogs/v2/pkg/storage/zerox_deployment"
	mock "github.com/stretchr/testify/mock"
)

// MockDeployStorage is an autogenerated mock type for the IStorage type
type MockDeployStorage struct {
	mock.Mock
}

// Get provides a mock function with given fields:
func (_m *MockDeployStorage) Get() ([]zerox_deployment.Deployment, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 []zerox_deployment.Deployment
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]zerox_deployment.Deployment, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []zerox_deployment.Deployment); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]zerox_deployment.Deployment)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: d
func (_m *MockDeployStorage) Insert(d zerox_deployment.Deployment) error {
	ret := _m.Called(d)

	if len(ret) == 0 {
		panic("no return value specified for Insert")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(zerox_deployment.Deployment) error); ok {
		r0 = rf(d)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockDeployStorage creates a new instance of MockDeployStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDeployStorage(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDeployStorage {
	mock := &MockDeployStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}