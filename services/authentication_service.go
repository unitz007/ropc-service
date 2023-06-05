package services

import (
	"log"
	"ropc-service/authenticators"
	"ropc-service/model/dto"
	"ropc-service/model/entities"
	"ropc-service/utils"
	"sync"
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

	var wg sync.WaitGroup
	wg.Add(2)
	var error2 error
	var verifiedUser *entities.User
	var verifiedClient *entities.Client

	go func() {
		defer wg.Done()
		u, err := selfC.userAuthenticator.Authenticate(user.Username, user.Password)
		if err != nil {
			error2 = err
			return
		}
		verifiedUser = u

	}()

	go func() {
		defer wg.Done()
		c, err := selfC.clientAuthenticator.Authenticate(client.ClientId, client.ClientSecret)
		if err != nil {
			error2 = err
			return
		}
		verifiedClient = c
	}()

	wg.Wait()

	log.Println("er", error2)

	if error2 != nil {
		return nil, error2
	}

	accessToken, err := utils.GenerateToken(verifiedUser, verifiedClient)
	if err != nil {
		return nil, err
	}

	token := &dto.Token{AccessToken: accessToken}

	return token, nil
}
