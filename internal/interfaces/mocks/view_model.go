// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ViewModel is an autogenerated mock type for the ViewModel type
type ViewModel struct {
	mock.Mock
}

type mockConstructorTestingTNewViewModel interface {
	mock.TestingT
	Cleanup(func())
}

// NewViewModel creates a new instance of ViewModel. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewViewModel(t mockConstructorTestingTNewViewModel) *ViewModel {
	mock := &ViewModel{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
