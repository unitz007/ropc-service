package services

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"ropc-service/mocks"
	"ropc-service/model/entities"
	"testing"
)

func Test_UserAuthenticationFailure(t *testing.T) {

	// should fail if user does not exist
	userAuthenticator, userRepositoryMock := resetUserAuthenticatorMock()
	userRepositoryMock.On("GetUser", mock.Anything).Return(nil, errors.New("invalid user"))
	user, err := userAuthenticator.Authenticate(wrongUsername, mock.Anything)
	assert.Equal(t, err.Error(), "invalid user")
	assert.Nil(t, user)
	userRepositoryMock.AssertCalled(t, "GetUser", mock.Anything)

	// should fail with wrong password combination
	userAuthenticator, userRepositoryMock = resetUserAuthenticatorMock()
	userRepositoryMock.On("GetUser", rightUsername).Return(&entities.User{Username: rightUsername, Password: hashedRightPassword}, nil)
	user, err = userAuthenticator.Authenticate(rightUsername, wrongPassword)
	assert.EqualError(t, err, "invalid user credentials")
	assert.Nil(t, user)
	userRepositoryMock.AssertCalled(t, "GetUser", rightUsername)

	// test with right encryption but still wrong password nonetheless, should fail
	userAuthenticator, userRepositoryMock = resetUserAuthenticatorMock()
	userRepositoryMock.On("GetUser", rightUsername).Return(&entities.User{Username: rightUsername, Password: hashedRightPassword}, nil)
	user, err = userAuthenticator.Authenticate(rightUsername, wrongPassword)
	assert.EqualError(t, err, "invalid user credentials")
	assert.Nil(t, user)
	userRepositoryMock.AssertCalled(t, "GetUser", rightUsername)

}

func Test_UserAuthenticationSuccess(t *testing.T) {

	// successful authentication
	userAuthenticator, userRepositoryMock := resetUserAuthenticatorMock()

	userRepositoryMock.On("GetUser", rightUsername).Return(&entities.User{Username: rightUsername, Password: hashedRightPassword}, nil)
	user, err := userAuthenticator.Authenticate(rightUsername, rightPassword)
	assert.NotNil(t, user)
	assert.Nil(t, err)
	userRepositoryMock.AssertCalled(t, "GetUser", rightUsername)
}

func resetUserAuthenticatorMock() (*UserAuthenticator, *mocks.UserRepository) {

	userRepositoryMock := new(mocks.UserRepository)
	userAuthenticatorMock := &UserAuthenticator{
		userRepository: userRepositoryMock,
	}

	return userAuthenticatorMock, userRepositoryMock
}
