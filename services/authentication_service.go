package services

import (
	"ropc-service/model"
)

type AuthenticationService interface {
	Authenticate(user *model.User, client *model.Client) (*model.Token, error)
}

type AuthenticationServiceImpl struct {
	userAuthenticator   UserAuthenticatorContract
	clientAuthenticator ClientAuthenticatorContract
}

func InstantiateAuthenticator(userAuthenticator UserAuthenticatorContract, clientAuthenticator ClientAuthenticatorContract) *AuthenticationServiceImpl {
	return &AuthenticationServiceImpl{
		userAuthenticator:   userAuthenticator,
		clientAuthenticator: clientAuthenticator,
	}
}

func (selfC AuthenticationServiceImpl) Authenticate(user *model.User, client *model.Client) (*model.Token, error) {

	channel := make(chan any, 2)

	go func() {
		u, err := selfC.userAuthenticator.Authenticate(user.Username, user.Password)
		if err != nil {
			channel <- err
			return
		}
		channel <- u
	}()

	go func() {
		c, err := selfC.clientAuthenticator.Authenticate(client.ClientId, client.ClientSecret)
		if err != nil {
			channel <- err
			return
		}
		channel <- c
	}()

	var u2 *model.User
	var c2 *model.Client

	for val := range channel {
		if err, ok := val.(error); ok && err != nil {
			return nil, err
		} else {
			if u, ok := val.(*model.User); ok {
				u2 = u
			}

			if c, ok := val.(*model.Client); ok {
				c2 = c
			}
		}

		if u2 != nil && c2 != nil {
			break
		}
	}

	accessToken, err := GenerateToken(u2, c2)
	if err != nil {
		return nil, err
	}

	token := &model.Token{AccessToken: accessToken}

	return token, nil
}
