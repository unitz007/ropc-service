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

	var clientRepositoryMock *mocks.ClientRepository

	t.Run("Client does not exist", func(t *testing.T) {

		clientRepositoryMock = new(mocks.ClientRepository)
		clientAuthenticator := &clientAuthenticator{
			repository: clientRepositoryMock,
		}
		clientRepositoryMock.On("GetClient", mock.Anything).Return(nil, errors.New(InvalidClientMessage))
		user, err := clientAuthenticator.Authenticate(WrongUsername, mock.Anything)
		assert.Equal(t, err.Error(), InvalidClientMessage)
		assert.Nil(t, user)
		clientRepositoryMock.AssertCalled(t, "GetClient", mock.Anything)

	})

	t.Run("Client with plain password on Database", func(t *testing.T) {

		clientRepositoryMock = new(mocks.ClientRepository)
		clientAuthenticator := &clientAuthenticator{
			repository: clientRepositoryMock,
		}

		// should fail with unencrypted client secret in database
		clientRepositoryMock.On("GetClient", rightClientId).Return(&entities.Client{ClientId: rightClientId, ClientSecret: rightClientSecret}, nil)
		user, err := clientAuthenticator.Authenticate(rightClientId, WrongClientSecret)
		assert.EqualError(t, err, InvalidClientMessage)
		assert.Nil(t, user)
		clientRepositoryMock.AssertCalled(t, "GetClient", rightClientId)

	})

	t.Run("Client with wrong secret", func(t *testing.T) {
		clientRepositoryMock = new(mocks.ClientRepository)
		clientAuthenticator := &clientAuthenticator{
			repository: clientRepositoryMock,
		}

		// should fail with wrong client secret
		clientRepositoryMock.On("GetClient", rightClientId).Return(&entities.Client{ClientId: rightClientId, ClientSecret: hashedRightClientSecret}, nil)
		user, err := clientAuthenticator.Authenticate(rightClientId, WrongClientSecret)
		assert.EqualError(t, err, InvalidClientMessage)
		assert.Nil(t, user)
		clientRepositoryMock.AssertCalled(t, "GetClient", rightClientId)
	})

}

func Test_ClientAuthenticationSuccess(t *testing.T) {

	// successful authentication
	clientRepositoryMock := new(mocks.ClientRepository)
	clientAuthenticator := &clientAuthenticator{
		repository: clientRepositoryMock,
	}

	clientRepositoryMock.On("GetClient", rightClientId).Return(&entities.Client{ClientId: rightClientId, ClientSecret: hashedRightClientSecret}, nil)
	client, err := clientAuthenticator.Authenticate(rightClientId, rightClientSecret)
	assert.NotNil(t, client)
	assert.Nil(t, err)
	clientRepositoryMock.AssertCalled(t, "GetClient", rightClientId)
}

func Test_ThirdPartyValidation(t *testing.T) {

	thirdPartyClient := "mbb_85893922"
	var clientRepositoryMock *mocks.ClientRepository
	var thirdPartyAuthenticatorMock *mocks.ThirdPartyClientAuthenticator
	failed := false

	t.Run("Should call third party validation", func(t *testing.T) {
		clientRepositoryMock = new(mocks.ClientRepository)
		thirdPartyAuthenticatorMock = new(mocks.ThirdPartyClientAuthenticator)

		clientRepositoryMock.On("GetClient", thirdPartyClient).Return(&entities.Client{}, nil)
		thirdPartyAuthenticatorMock.On("Authenticate", thirdPartyClient, mock.Anything).Return(&failed, nil)

		clientAuthenticator := clientAuthenticator{clientRepositoryMock, thirdPartyAuthenticatorMock}

		_, _ = clientAuthenticator.Authenticate(thirdPartyClient, mock.Anything)

		clientRepositoryMock.AssertNotCalled(t, "GetClient", thirdPartyClient)

		thirdPartyAuthenticatorMock.AssertCalled(t, "Authenticate", thirdPartyClient, mock.Anything)

	})

	t.Run("Third party failed authentication", func(t *testing.T) {
		thirdPartyAuthenticatorMock = new(mocks.ThirdPartyClientAuthenticator)
		thirdPartyAuthenticatorMock.On("Authenticate", thirdPartyClient, mock.Anything).Return(&failed, nil)

		clientAuthenticator := clientAuthenticator{
			repository:                    clientRepositoryMock,
			thirdPartyClientAuthenticator: thirdPartyAuthenticatorMock,
		}

		_, err := clientAuthenticator.Authenticate(thirdPartyClient, mock.Anything)

		if err == nil {
			t.Fatal("Error is expected but got nil")
		}

		if err.Error() != InvalidClientMessage {
			t.Errorf("Expected %s but got %s", InvalidClientMessage, err.Error())
		}
	})

	t.Run("Request to third party failed (timeout)", func(t *testing.T) {
		thirdPartyAuthenticatorMock = new(mocks.ThirdPartyClientAuthenticator)
		thirdPartyAuthenticatorMock.On("Authenticate", thirdPartyClient, mock.Anything).Return(nil, errors.New("request to third party failed"))

		clientAuthenticator := clientAuthenticator{clientRepositoryMock, thirdPartyAuthenticatorMock}

		_, err := clientAuthenticator.Authenticate(thirdPartyClient, mock.Anything)
		if err == nil {
			t.Fatal("Error is expected but got nil")
		}

		if err.Error() != ConnectionErrorMessage {
			t.Errorf("Expected %s but got %s", ConnectionErrorMessage, err.Error())
		}
	})

	t.Run("Successful client third part authentication", func(t *testing.T) {

		success := true

		thirdPartyAuthenticatorMock = new(mocks.ThirdPartyClientAuthenticator)
		thirdPartyAuthenticatorMock.On("Authenticate", thirdPartyClient, mock.Anything).Return(&success, nil)

		clientAuthenticator := clientAuthenticator{clientRepositoryMock, thirdPartyAuthenticatorMock}
		client, _ := clientAuthenticator.Authenticate(thirdPartyClient, mock.Anything)

		if client == nil {
			t.Fatal("Expected client but got nil")
		}

	})

}
