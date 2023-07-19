// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	"ropc-service/model/entities"

	mock "github.com/stretchr/testify/mock"
)

// TokenUtil is an autogenerated mock type for the TokenUtil type
type TokenUtil struct {
	mock.Mock
}

// GenerateToken provides a mock function with given fields: user, client
func (_m *TokenUtil) GenerateToken(user *entities.User, client *entities.Client) (string, error) {
	ret := _m.Called(user, client)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*entities.User, *entities.Client) (string, error)); ok {
		return rf(user, client)
	}
	if rf, ok := ret.Get(0).(func(*entities.User, *entities.Client) string); ok {
		r0 = rf(user, client)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*entities.User, *entities.Client) error); ok {
		r1 = rf(user, client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTokenUtil interface {
	mock.TestingT
	Cleanup(func())
}

// NewTokenUtil creates a new instance of TokenUtil. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTokenUtil(t mockConstructorTestingTNewTokenUtil) *TokenUtil {
	mock := &TokenUtil{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}