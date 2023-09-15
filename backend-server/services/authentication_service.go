package services

import (
	"ropc-service/authenticators"
	"ropc-service/model/dto"
	"ropc-service/model/entities"
	"ropc-service/utils"
)

type Authenticator interface {
	Authenticate(user *entities.User, client *entities.Client) (*dto.Token, error)
}

type authenticator struct {
	userAuthenticator   authenticators.UserAuthenticator
	clientAuthenticator authenticators.ClientAuthenticator
}

func InstantiateAuthenticator(uA authenticators.UserAuthenticator, cA authenticators.ClientAuthenticator) *authenticator {
	return &authenticator{
		userAuthenticator:   uA,
		clientAuthenticator: cA,
	}
}

func (selfC authenticator) Authenticate(user *entities.User, client *entities.Client) (*dto.Token, error) {

	u, err := selfC.userAuthenticator.Authenticate(user.Username, user.Password)
	if err != nil {
		return nil, err
	}

	c, err := selfC.clientAuthenticator.Authenticate(client.ClientId, client.ClientSecret)
	if err != nil {
		return nil, err
	}

	accessToken, err := utils.GenerateToken(u, c)
	if err != nil {
		return nil, err
	}

	token := &dto.Token{AccessToken: accessToken}

	return token, nil
}
