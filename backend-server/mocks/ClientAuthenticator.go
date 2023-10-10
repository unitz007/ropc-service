// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	model "backend-server/model"

	mock "github.com/stretchr/testify/mock"
)

// ClientAuthenticator is an autogenerated mock type for the ClientAuthenticator type
type ClientAuthenticator struct {
	mock.Mock
}

// Authenticate provides a mock function with given fields: clientId, clientSecret
func (_m *ClientAuthenticator) Authenticate(clientId string, clientSecret string) (*model.Token, error) {
	ret := _m.Called(clientId, clientSecret)

	var r0 *model.Token
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*model.Token, error)); ok {
		return rf(clientId, clientSecret)
	}
	if rf, ok := ret.Get(0).(func(string, string) *model.Token); ok {
		r0 = rf(clientId, clientSecret)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Token)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(clientId, clientSecret)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewClientAuthenticator interface {
	mock.TestingT
	Cleanup(func())
}

// NewClientAuthenticator creates a new instance of ClientAuthenticator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClientAuthenticator(t mockConstructorTestingTNewClientAuthenticator) *ClientAuthenticator {
	mock := &ClientAuthenticator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
