// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	entity "task-manager/internal/model/entity"

	mock "github.com/stretchr/testify/mock"
)

// ITaskRepository is an autogenerated mock type for the ITaskRepository type
type ITaskRepository struct {
	mock.Mock
}

// GetNumberOfWeek provides a mock function with given fields:
func (_m *ITaskRepository) GetNumberOfWeek() (entity.Week, error) {
	ret := _m.Called()

	var r0 entity.Week
	var r1 error
	if rf, ok := ret.Get(0).(func() (entity.Week, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() entity.Week); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(entity.Week)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: tasks
func (_m *ITaskRepository) Save(tasks []entity.Task) error {
	ret := _m.Called(tasks)

	var r0 error
	if rf, ok := ret.Get(0).(func([]entity.Task) error); ok {
		r0 = rf(tasks)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveNumberOfWeek provides a mock function with given fields: numberOfWeek
func (_m *ITaskRepository) SaveNumberOfWeek(numberOfWeek int) error {
	ret := _m.Called(numberOfWeek)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(numberOfWeek)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewITaskRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewITaskRepository creates a new instance of ITaskRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewITaskRepository(t mockConstructorTestingTNewITaskRepository) *ITaskRepository {
	mock := &ITaskRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
