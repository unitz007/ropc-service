package authenticators

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"ropc-service/model/entities"
	"ropc-service/repositories"
	"strings"
)

const InvalidClientMessage = "invalid client credentials"
const ConnectionErrorMessage = "could not authenticate client"

type ClientAuthenticatorContract interface {
	Authenticate(clientId, clientSecret string) (*entities.Client, error)
}

type ClientAuthenticator struct {
	repository                    repositories.ClientRepository
	thirdPartyClientAuthenticator ThirdPartyClientAuthenticator
}

func InstantiateClientAuthenticator(repository repositories.ClientRepository) *ClientAuthenticator {
	return &ClientAuthenticator{
		repository: repository,
	}
}

func (selfC ClientAuthenticator) Authenticate(clientId, clientSecret string) (*entities.Client, error) {

	if strings.HasPrefix(clientId, "mbb_") {
		ok, err := selfC.thirdPartyClientAuthenticator.Authenticate(clientId, clientSecret)

		if err != nil {
			return nil, errors.New(ConnectionErrorMessage)
		}

		if !*ok {
			return nil, errors.New(InvalidClientMessage)

		}

		return &entities.Client{ClientId: clientId, ClientSecret: clientSecret}, nil
	}

	client, err := selfC.repository.GetClient(clientId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(client.ClientSecret), []byte(clientSecret)); err != nil || err == bcrypt.ErrMismatchedHashAndPassword {
		log.Println(err)
		return nil, errors.New(InvalidClientMessage)
	}

	return client, nil
}
