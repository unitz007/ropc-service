package services

import (
	"ropc-service/model/dto"
	"ropc-service/model/entities"
)

type AuthenticatorContract interface {
	Authenticate(user *entities.User, client *entities.Client) (*dto.Token, error)
}

type Authenticator struct {
	userAuthenticator   UserAuthenticatorContract
	clientAuthenticator ClientAuthenticatorContract
}

func InstantiateAuthenticator(uA UserAuthenticatorContract, cA ClientAuthenticatorContract) *Authenticator {
	return &Authenticator{
		userAuthenticator:   uA,
		clientAuthenticator: cA,
	}
}

func (selfC Authenticator) Authenticate(user *entities.User, client *entities.Client) (*dto.Token, error) {

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

	var u2 *entities.User
	var c2 *entities.Client

	for val := range channel {
		if err, ok := val.(error); ok && err != nil {
			return nil, err
		} else {
			if u, ok := val.(*entities.User); ok {
				u2 = u
			}

			if c, ok := val.(*entities.Client); ok {
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

	token := &dto.Token{AccessToken: accessToken}

	return token, nil
}
