package authenticators

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
	user, err := userAuthenticator.Authenticate(WrongUsername, mock.Anything)
	assert.Equal(t, err.Error(), "invalid user")
	assert.Nil(t, user)
	userRepositoryMock.AssertCalled(t, "GetUser", mock.Anything)

	// should fail with wrong password combination
	userAuthenticator, userRepositoryMock = resetUserAuthenticatorMock()
	userRepositoryMock.On("GetUser", RightUsername).Return(&entities.User{Username: RightUsername, Password: hashedRightPassword}, nil)
	user, err = userAuthenticator.Authenticate(RightUsername, WrongPassword)
	assert.EqualError(t, err, "invalid user credentials")
	assert.Nil(t, user)
	userRepositoryMock.AssertCalled(t, "GetUser", RightUsername)

	// test with right encryption but still wrong password nonetheless, should fail
	userAuthenticator, userRepositoryMock = resetUserAuthenticatorMock()
	userRepositoryMock.On("GetUser", RightUsername).Return(&entities.User{Username: RightUsername, Password: hashedRightPassword}, nil)
	user, err = userAuthenticator.Authenticate(RightUsername, WrongPassword)
	assert.EqualError(t, err, "invalid user credentials")
	assert.Nil(t, user)
	userRepositoryMock.AssertCalled(t, "GetUser", RightUsername)

}

func Test_UserAuthenticationSuccess(t *testing.T) {

	// successful authentication
	userAuthenticator, userRepositoryMock := resetUserAuthenticatorMock()

	userRepositoryMock.On("GetUser", RightUsername).Return(&entities.User{Username: RightUsername, Password: hashedRightPassword}, nil)
	user, err := userAuthenticator.Authenticate(RightUsername, RightPassword)
	assert.NotNil(t, user)
	assert.Nil(t, err)
	userRepositoryMock.AssertCalled(t, "GetUser", RightUsername)
}

func resetUserAuthenticatorMock() (*UserAuthenticator, *mocks.UserRepository) {

	userRepositoryMock := new(mocks.UserRepository)
	userAuthenticatorMock := &UserAuthenticator{
		userRepository: userRepositoryMock,
	}

	return userAuthenticatorMock, userRepositoryMock
}
