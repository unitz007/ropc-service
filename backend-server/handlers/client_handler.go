package handlers

import (
	"errors"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"ropc-service/model/dto"
	"ropc-service/model/entities"
	"ropc-service/routers"
	"ropc-service/services"
)

type ClientHandler interface {
	CreateClient(w http.ResponseWriter, r *http.Request)
}

type clientHandler[T any] struct {
	router        routers.Router
	clientService services.ClientService
}

func NewClientHandler(clientService services.ClientService) ClientHandler {
	return &clientHandler[*mux.Route]{
		clientService: clientService,
	}
}

func (c *clientHandler[T]) CreateClient(_ http.ResponseWriter, r *http.Request) {

	var clientRequest *dto.Client

	err := JsonToStruct(r.Body, &clientRequest)
	if err != nil {
		panic(errors.New("invalid request body"))
	}

	client, err := entities.NewClient(clientRequest.ClientId, uuid.New().String())
	if err != nil {
		panic(err)
	}

	//clientRepository := repositories.NewClientRepository(conf.NewDataBase(conf.EnvironmentConfig))
	//
	//clientService := services.NewClientService(clientRepository)

	c.clientService.CreateClient(client)

}
