package services

import (
	"ropc-service/model/entities"
)

type ClientService interface {
	GetClients() []entities.Client
}
