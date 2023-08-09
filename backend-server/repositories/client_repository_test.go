package repositories

import (
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"ropc-service/conf"
	"ropc-service/mocks"
	"ropc-service/model/entities"
	"testing"
)

func TestCreateClient(t *testing.T) {
	var db conf.Database[gorm.DB]

	t.Run("should return error if client_id already exists", func(t *testing.T) {
		db = mocks.NewDatabaseMock(t)
		repo := NewClientRepository(db)

		clientId := "client_id"
		client := &entities.Client{
			ClientId:     clientId,
			ClientSecret: mock.Anything,
		}

		err := repo.CreateClient(client)
		require.NoError(t, err, "should not return error")

		err = repo.CreateClient(client)
		require.Errorf(t, err, "should return error")

		expectedErrorMsg := "client Id already exists"

		if err.Error() != expectedErrorMsg {
			t.Errorf("expected %s, got %s", expectedErrorMsg, err.Error())
		}
	})

	t.Run("should fail if client secret or client id is empty", func(t *testing.T) {
		db = mocks.NewDatabaseMock(t)
		repository := NewClientRepository(db)
		err := repository.CreateClient(&entities.Client{ClientId: "clientId"})
		require.Errorf(t, err, "should return error cause client secret is empty")
	})

}
