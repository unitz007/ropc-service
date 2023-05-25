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

var tokenUtil = InstantiateTokenUtil()

func InstantiateAuthenticator(userAuthenticator UserAuthenticatorContract, clientAuthenticator ClientAuthenticatorContract) *AuthenticationServiceImpl {
	return &AuthenticationServiceImpl{
		userAuthenticator:   userAuthenticator,
		clientAuthenticator: clientAuthenticator,
	}
}

func (selfC AuthenticationServiceImpl) Authenticate(user *model.User, client *model.Client) (*model.Token, error) {

	u, err := selfC.userAuthenticator.Authenticate(user.Username, user.Password)
	if err != nil {
		return nil, err
	}

	c, err := selfC.clientAuthenticator.Authenticate(client.ClientId, client.ClientSecret)
	if err != nil {
		return nil, err
	}

	token, err := tokenUtil.GenerateToken(u, c)
	if err != nil {
		return nil, err
	}

	return &model.Token{AccessToken: token}, nil
}
