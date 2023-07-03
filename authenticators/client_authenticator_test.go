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
	clientRepositoryMock.On("GetClient", mock.Anything).Return(nil, errors.New(InvalidClientMessage))
	user, err := clientAuthenticator.Authenticate(WrongUsername, mock.Anything)
	assert.Equal(t, err.Error(), InvalidClientMessage)
	assert.Nil(t, user)
	clientRepositoryMock.AssertCalled(t, "GetClient", mock.Anything)

	// should fail with unencrypted client secret in database
	clientAuthenticator, clientRepositoryMock = resetClientAuthenticatorMock()
	clientRepositoryMock.On("GetClient", rightClientId).Return(&entities.Client{ClientId: rightClientId, ClientSecret: rightClientSecret}, nil)
	user, err = clientAuthenticator.Authenticate(rightClientId, WrongClientSecret)
	assert.EqualError(t, err, InvalidClientMessage)
	assert.Nil(t, user)
	clientRepositoryMock.AssertCalled(t, "GetClient", rightClientId)

	// should fail with wrong client secret
	clientAuthenticator, clientRepositoryMock = resetClientAuthenticatorMock()
	clientRepositoryMock.On("GetClient", rightClientId).Return(&entities.Client{ClientId: rightClientId, ClientSecret: hashedRightClientSecret}, nil)
	user, err = clientAuthenticator.Authenticate(rightClientId, WrongClientSecret)
	assert.EqualError(t, err, InvalidClientMessage)
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

func Test_ThirdPartyValidation(t *testing.T) {

	thirdPartyClient := "mbb_85893922"
	var clientRepositoryMock *mocks.ClientRepository
	var thirdPartyAuthenticatorMock *mocks.ThirdPartyClientAuthenticator

	t.Run("Should call third party validation", func(t *testing.T) {
		clientRepositoryMock = new(mocks.ClientRepository)
		thirdPartyAuthenticatorMock = new(mocks.ThirdPartyClientAuthenticator)

		clientRepositoryMock.On("GetClient", thirdPartyClient).Return(&entities.Client{}, nil)
		thirdPartyAuthenticatorMock.On("Authenticate", thirdPartyClient, mock.Anything).Return(false, nil)

		clientAuthenticator := ClientAuthenticator{clientRepositoryMock, thirdPartyAuthenticatorMock}

		_, _ = clientAuthenticator.Authenticate(thirdPartyClient, mock.Anything)

		clientRepositoryMock.AssertNotCalled(t, "GetClient", thirdPartyClient)

		thirdPartyAuthenticatorMock.AssertCalled(t, "Authenticate", thirdPartyClient, mock.Anything)

	})

	t.Run("Third party failed authentication", func(t *testing.T) {
		thirdPartyAuthenticatorMock = new(mocks.ThirdPartyClientAuthenticator)
		thirdPartyAuthenticatorMock.On("Authenticate", thirdPartyClient, mock.Anything).Return(false, nil)

		clientAuthenticator := ClientAuthenticator{clientRepositoryMock, thirdPartyAuthenticatorMock}

		_, err := clientAuthenticator.Authenticate(thirdPartyClient, mock.Anything)

		if err == nil {
			t.Fatal("Error is expected but got nil")
		}
	})

}
