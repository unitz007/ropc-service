package services

import (
	"errors"
	"ropc-service/model/entities"
	"ropc-service/repositories"
)

type ClientService interface {
	GetClients() []entities.Client
	CreateClient(client *entities.Client)
}

type clientService struct {
	clientRepository repositories.ClientRepository
}

func (s *clientService) CreateClient(client *entities.Client) {

	if client.ClientId == "" {
		panic(errors.New("client id is required"))
	}

	err := s.clientRepository.CreateClient(client)

	if err != nil {
		panic(err)
	}
}

func NewClientService(repository repositories.ClientRepository) ClientService {
	return &clientService{clientRepository: repository}
}

func (s *clientService) GetClients() []entities.Client {
	return s.clientRepository.GetClients()
}
