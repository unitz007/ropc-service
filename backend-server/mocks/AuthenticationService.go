// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	model "ropc-service/model/dto"
	"ropc-service/model/entities"

	mock "github.com/stretchr/testify/mock"
)

// AuthenticationService is an autogenerated mock type for the AuthenticationService type
type AuthenticationService struct {
	mock.Mock
}

// Authenticate provides a mock function with given fields: user, client
func (_m *AuthenticationService) Authenticate(user *entities.User, client *entities.Client) (*model.Token, error) {
	ret := _m.Called(user, client)

	var r0 *model.Token
	var r1 error
	if rf, ok := ret.Get(0).(func(*entities.User, *entities.Client) (*model.Token, error)); ok {
		return rf(user, client)
	}
	if rf, ok := ret.Get(0).(func(*entities.User, *entities.Client) *model.Token); ok {
		r0 = rf(user, client)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Token)
		}
	}

	if rf, ok := ret.Get(1).(func(*entities.User, *entities.Client) error); ok {
		r1 = rf(user, client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAuthenticationService interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthenticationService creates a new instance of AuthenticationService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthenticationService(t mockConstructorTestingTNewAuthenticationService) *AuthenticationService {
	mock := &AuthenticationService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}