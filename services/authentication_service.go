package services

import (
	"ropc-service/authenticators"
	"ropc-service/model/dto"
	"ropc-service/model/entities"
	"ropc-service/utils"
)

type AuthenticatorContract interface {
	Authenticate(user *entities.User, client *entities.Client) (*dto.Token, error)
}

type Authenticator struct {
	userAuthenticator   authenticators.UserAuthenticatorContract
	clientAuthenticator authenticators.ClientAuthenticatorContract
}

func InstantiateAuthenticator(uA authenticators.UserAuthenticatorContract, cA authenticators.ClientAuthenticatorContract) *Authenticator {
	return &Authenticator{
		userAuthenticator:   uA,
		clientAuthenticator: cA,
	}
}

func (selfC Authenticator) Authenticate(user *entities.User, client *entities.Client) (*dto.Token, error) {

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
