package services

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"ropc-service/mocks"
	"ropc-service/model"
	"testing"
)

var (
	userAuthenticatorMock, clientAuthenticatorMock, authenticatorServiceTest = resetMocks()
)

func Test_AuthenticationFailure(t *testing.T) {

	// assertions for user authentication
	userAuthenticatorMock.On("Authenticate", wrongUsername, wrongPassword).Return(nil, errors.New("invalid user credentials"))
	token, err := authenticatorServiceTest.Authenticate(&wrongTestUser, &model.Client{})
	assert.EqualError(t, err, "invalid user credentials")
	assert.Nil(t, token, "token should be nil")

	userAuthenticatorMock.AssertCalled(t, "Authenticate", wrongUsername, wrongPassword)
	clientAuthenticatorMock.AssertNotCalled(t, "Authenticate", mock.Anything, mock.Anything)

	// asserts that client authenticator is called when user authentication is successful
	userAuthenticatorMock, clientAuthenticatorMock, authenticatorServiceTest = resetMocks()

	userAuthenticatorMock.On("Authenticate", rightUsername, rightPassword).Return(&rightTestUser, nil)
	clientAuthenticatorMock.On("Authenticate", wrongClientId, wrongClientSecret).Return(nil, errors.New("invalid client credentials"))
	token, err = authenticatorServiceTest.Authenticate(&rightTestUser, &wrongClient)

	assert.Nil(t, token)
	assert.EqualError(t, err, "invalid client credentials")
}

func resetMocks() (*mocks.UserAuthenticatorContract, *mocks.ClientAuthenticatorContract, *AuthenticationServiceImpl) {

	userAuthenticatorMock := new(mocks.UserAuthenticatorContract)
	clientAuthenticatorMock := new(mocks.ClientAuthenticatorContract)
	authenticatorServiceTest := InstantiateAuthenticator(userAuthenticatorMock, clientAuthenticatorMock)

	return userAuthenticatorMock, clientAuthenticatorMock, authenticatorServiceTest
}
