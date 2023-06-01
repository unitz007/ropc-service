package services

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"ropc-service/mocks"
	model "ropc-service/model/entities"
	"testing"
)

func TestRegisterUser(t *testing.T) {

	var testUser *model.User
	var userMockRepository *mocks.UserRepository

	t.Run("user without username", func(t *testing.T) {

		userMockRepository = new(mocks.UserRepository)
		testUser = &model.User{
			Username: "",
		}

		userMockRepository.On("CreateUser", testUser).Return(mock.Anything)
		_, err := NewUserService(userMockRepository).CreateUser(testUser)
		assert.EqualError(t, err, "username is required")
		userMockRepository.AssertNotCalled(t, "CreateUser", testUser)
	})

	t.Run("user without password", func(t *testing.T) {
		userMockRepository = new(mocks.UserRepository)
		testUser = &model.User{
			Username: "User",
			Password: "",
		}

		userMockRepository.On("CreateUser", testUser).Return(mock.Anything)
		_, err := NewUserService(userMockRepository).CreateUser(testUser)
		assert.EqualError(t, err, "password is required")
		userMockRepository.AssertNotCalled(t, "CreateUser", testUser)
	})

	t.Run("user without email", func(t *testing.T) {
		userMockRepository = new(mocks.UserRepository)
		testUser = &model.User{
			Username: "User",
			Password: "Password",
			Email:    "",
		}

		userMockRepository.On("CreateUser", testUser).Return(mock.Anything)
		_, err := NewUserService(userMockRepository).CreateUser(testUser)
		assert.EqualError(t, err, "email is required")
		userMockRepository.AssertNotCalled(t, "CreateUser", testUser)

	})

	t.Run("successful registration", func(t *testing.T) {
		userMockRepository = new(mocks.UserRepository)

		password := "Password"

		testUser = &model.User{
			Username: "User",
			Password: password,
			Email:    "test@gmail.com",
		}

		userMockRepository.On("CreateUser", testUser).Return(testUser, nil)
		newUser, err := NewUserService(userMockRepository).CreateUser(testUser)
		assert.NoError(t, err, "error is supposed to be nil")
		assert.NotNil(t, newUser, "newUser should be not be nil")
		userMockRepository.AssertCalled(t, "CreateUser", testUser)

		encErr := bcrypt.CompareHashAndPassword([]byte(newUser.Password), []byte(password))
		assert.NoError(t, encErr, "user password was not encrypted")
	})
}
