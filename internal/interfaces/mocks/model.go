// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Model is an autogenerated mock type for the Model type
type Model struct {
	mock.Mock
}

type mockConstructorTestingTNewModel interface {
	mock.TestingT
	Cleanup(func())
}

// NewModel creates a new instance of Model. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewModel(t mockConstructorTestingTNewModel) *Model {
	mock := &Model{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
