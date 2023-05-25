package services

import (
	"log"
	"ropc-service/model"
)

type AuthenticationService interface {
	Authenticate(user *model.User, client *model.Client) (*model.Token, error)
}

type AuthenticationServiceImpl struct {
	userAuthenticator   UserAuthenticatorContract
	clientAuthenticator ClientAuthenticatorContract
}

var tokenUtil = InstantiateTokenUtil()

func InstantiateAuthenticator(userAuthenticator UserAuthenticatorContract, clientAuthenticator ClientAuthenticatorContract) *AuthenticationServiceImpl {
	return &AuthenticationServiceImpl{
		userAuthenticator:   userAuthenticator,
		clientAuthenticator: clientAuthenticator,
	}
}

func (selfC AuthenticationServiceImpl) Authenticate(user *model.User, client *model.Client) (*model.Token, error) {

	channel := make(chan any)

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

	var token *model.Token

	for val := range channel {
		if err, ok := val.(error); ok && err != nil {
			log.Println("Got here")
			return nil, err
		} else {
			s, err := tokenUtil.GenerateToken(&model.User{}, &model.Client{})
			if err != nil {
				return nil, err
			} else {
				token = &model.Token{AccessToken: s}
				break
			}
		}
	}

	return token, nil
}
