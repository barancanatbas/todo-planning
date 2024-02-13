// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	entity "task-manager/internal/model/entity"

	mock "github.com/stretchr/testify/mock"

	provider "task-manager/pkg/provider"

	url "net/url"
)

// IProviderClient is an autogenerated mock type for the IProviderClient type
type IProviderClient struct {
	mock.Mock
}

// Get provides a mock function with given fields: _a0, providerModel
func (_m *IProviderClient) Get(_a0 *url.URL, providerModel provider.ITask) ([]entity.Task, error) {
	ret := _m.Called(_a0, providerModel)

	var r0 []entity.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(*url.URL, provider.ITask) ([]entity.Task, error)); ok {
		return rf(_a0, providerModel)
	}
	if rf, ok := ret.Get(0).(func(*url.URL, provider.ITask) []entity.Task); ok {
		r0 = rf(_a0, providerModel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(*url.URL, provider.ITask) error); ok {
		r1 = rf(_a0, providerModel)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIProviderClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewIProviderClient creates a new instance of IProviderClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIProviderClient(t mockConstructorTestingTNewIProviderClient) *IProviderClient {
	mock := &IProviderClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}