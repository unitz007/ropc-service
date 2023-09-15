package services

import (
	"errors"
	"ropc-service/mocks"
	"ropc-service/model/entities"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_CreateClientTest(t *testing.T) {
	t.Run("should panic if create client fails", func(t *testing.T) {
		clientRepository := new(mocks.ClientRepository)
		client := &entities.Client{ClientId: uuid.NewString()}
		clientRepository.On("CreateClient", client).Return(errors.New("could not create client"))
		clientService := NewClientService(clientRepository)

		exec := func() {
			clientService.CreateClient(client)
		}

		assert.PanicsWithError(t, "could not create client", exec)
		clientRepository.AssertCalled(t, "CreateClient", client)
	})

	t.Run("client id should not be empty", func(t *testing.T) {
		client := &entities.Client{ClientSecret: "secret"}
		clientRepository := new(mocks.ClientRepository)
		clientService := NewClientService(clientRepository)

		clientRepository.On("CreateClient", client).Return(nil)

		exec := func() {
			clientService.CreateClient(client)
		}

		assert.PanicsWithError(t, "client id is required", exec)
		clientRepository.AssertNotCalled(t, "CreateClient", client)
	})
}
