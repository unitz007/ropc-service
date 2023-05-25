package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"ropc-service/model"
	"ropc-service/repositories"
)

type ClientAuthenticatorContract interface {
	Authenticate(clientId, clientSecret string) (*model.Client, error)
}

type ClientAuthenticator struct {
	repository repositories.ClientRepository
}

func InstantiateClientAuthenticator() *ClientAuthenticator {
	return &ClientAuthenticator{
		repository: repositories.NewClientRepository(),
	}
}

func (selfC ClientAuthenticator) Authenticate(clientId, clientSecret string) (*model.Client, error) {

	client, err := selfC.repository.GetClient(clientId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(client.ClientSecret), []byte(clientSecret)); err != nil || err == bcrypt.ErrMismatchedHashAndPassword {
		log.Println(err)
		return nil, errors.New("invalid client credentials")
	}

	return client, nil
}
