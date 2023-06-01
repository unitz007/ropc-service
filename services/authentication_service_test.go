package services

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"ropc-service/authenticators"
	"ropc-service/mocks"
	"ropc-service/model/entities"
	"testing"
)

var (
	userAuthenticatorMock, clientAuthenticatorMock, authenticatorServiceTest = resetMocks()
)

func Test_AuthenticationFailure(t *testing.T) {

	t.Skip("Skipping authentication test")

	// assertions for user authentication
	userAuthenticatorMock.On("Authenticate", authenticators.WrongUsername, authenticators.WrongPassword).Return(nil, errors.New("invalid user credentials"))
	token, err := authenticatorServiceTest.Authenticate(&authenticators.WrongTestUser, &entities.Client{})
	assert.EqualError(t, err, "invalid user credentials")
	assert.Nil(t, token, "token should be nil")

	userAuthenticatorMock.AssertCalled(t, "Authenticate", authenticators.WrongUsername, authenticators.WrongPassword)
	clientAuthenticatorMock.AssertNotCalled(t, "Authenticate", mock.Anything, mock.Anything)

	// asserts that client authenticator is called when user authentication is successful
	userAuthenticatorMock, clientAuthenticatorMock, authenticatorServiceTest = resetMocks()

	userAuthenticatorMock.On("Authenticate", authenticators.RightUsername, authenticators.RightPassword).Return(&authenticators.RightTestUser, nil)
	clientAuthenticatorMock.On("Authenticate", authenticators.WrongClientId, authenticators.WrongClientSecret).Return(nil, errors.New("invalid client credentials"))
	token, err = authenticatorServiceTest.Authenticate(&authenticators.RightTestUser, &authenticators.WrongClient)

	assert.Nil(t, token)
	assert.EqualError(t, err, "invalid client credentials")
}

func resetMocks() (*mocks.UserAuthenticatorContract, *mocks.ClientAuthenticatorContract, *Authenticator) {

	userAuthenticatorMock := new(mocks.UserAuthenticatorContract)
	clientAuthenticatorMock := new(mocks.ClientAuthenticatorContract)
	authenticatorServiceTest := InstantiateAuthenticator(userAuthenticatorMock, clientAuthenticatorMock)

	return userAuthenticatorMock, clientAuthenticatorMock, authenticatorServiceTest
}
