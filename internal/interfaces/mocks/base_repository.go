// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	interfaces "github.com/SpringCare/sh-go-workshop/internal/interfaces"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// BaseRepository is an autogenerated mock type for the BaseRepository type
type BaseRepository[M interfaces.Model, V interfaces.ViewModel] struct {
	mock.Mock
}

// Create provides a mock function with given fields: model
func (_m *BaseRepository[M, V]) Create(model M) (*V, error) {
	ret := _m.Called(model)

	var r0 *V
	var r1 error
	if rf, ok := ret.Get(0).(func(M) (*V, error)); ok {
		return rf(model)
	}
	if rf, ok := ret.Get(0).(func(M) *V); ok {
		r0 = rf(model)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*V)
		}
	}

	if rf, ok := ret.Get(1).(func(M) error); ok {
		r1 = rf(model)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *BaseRepository[M, V]) Delete(id uuid.UUID) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *BaseRepository[M, V]) GetAll() ([]V, error) {
	ret := _m.Called()

	var r0 []V
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]V, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []V); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]V)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: id
func (_m *BaseRepository[M, V]) GetById(id uuid.UUID) (*V, error) {
	ret := _m.Called(id)

	var r0 *V
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (*V, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) *V); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*V)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: model, id
func (_m *BaseRepository[M, V]) Update(model M, id uuid.UUID) (*V, error) {
	ret := _m.Called(model, id)

	var r0 *V
	var r1 error
	if rf, ok := ret.Get(0).(func(M, uuid.UUID) (*V, error)); ok {
		return rf(model, id)
	}
	if rf, ok := ret.Get(0).(func(M, uuid.UUID) *V); ok {
		r0 = rf(model, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*V)
		}
	}

	if rf, ok := ret.Get(1).(func(M, uuid.UUID) error); ok {
		r1 = rf(model, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewBaseRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewBaseRepository creates a new instance of BaseRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBaseRepository[M interfaces.Model, V interfaces.ViewModel](t mockConstructorTestingTNewBaseRepository) *BaseRepository[M, V] {
	mock := &BaseRepository[M, V]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}