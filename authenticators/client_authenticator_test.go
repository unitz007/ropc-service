package authenticators

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"ropc-service/mocks"
	"ropc-service/model/entities"
	"testing"
)

func Test_ClientAuthenticationFailure(t *testing.T) {

	// should fail if client does not exist
	clientAuthenticator, clientRepositoryMock := resetClientAuthenticatorMock()
	clientRepositoryMock.On("GetClient", mock.Anything).Return(nil, errors.New("invalid client"))
	user, err := clientAuthenticator.Authenticate(WrongUsername, mock.Anything)
	assert.Equal(t, err.Error(), "invalid client")
	assert.Nil(t, user)
	clientRepositoryMock.AssertCalled(t, "GetClient", mock.Anything)

	// should fail with unencrypted client secret in database
	clientAuthenticator, clientRepositoryMock = resetClientAuthenticatorMock()
	clientRepositoryMock.On("GetClient", rightClientId).Return(&entities.Client{ClientId: rightClientId, ClientSecret: rightClientSecret}, nil)
	user, err = clientAuthenticator.Authenticate(rightClientId, WrongClientSecret)
	assert.EqualError(t, err, "invalid client credentials")
	assert.Nil(t, user)
	clientRepositoryMock.AssertCalled(t, "GetClient", rightClientId)

	// should fail with wrong client secret
	clientAuthenticator, clientRepositoryMock = resetClientAuthenticatorMock()
	clientRepositoryMock.On("GetClient", rightClientId).Return(&entities.Client{ClientId: rightClientId, ClientSecret: hashedRightClientSecret}, nil)
	user, err = clientAuthenticator.Authenticate(rightClientId, WrongClientSecret)
	assert.EqualError(t, err, "invalid client credentials")
	assert.Nil(t, user)
	clientRepositoryMock.AssertCalled(t, "GetClient", rightClientId)
}

func Test_ClientAuthenticationSuccess(t *testing.T) {

	// successful authentication
	clientAuthenticator, clientRepositoryMock := resetClientAuthenticatorMock()

	clientRepositoryMock.On("GetClient", rightClientId).Return(&entities.Client{ClientId: rightClientId, ClientSecret: hashedRightClientSecret}, nil)
	client, err := clientAuthenticator.Authenticate(rightClientId, rightClientSecret)
	assert.NotNil(t, client)
	assert.Nil(t, err)
	clientRepositoryMock.AssertCalled(t, "GetClient", rightClientId)
}

func resetClientAuthenticatorMock() (*ClientAuthenticator, *mocks.ClientRepository) {

	clientRepositoryMock := new(mocks.ClientRepository)
	clientAuthenticatorMock := &ClientAuthenticator{
		repository: clientRepositoryMock,
	}

	return clientAuthenticatorMock, clientRepositoryMock
}
