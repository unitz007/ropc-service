package handlers

import (
	"errors"
	"net/http"
	"ropc-service/model"
	"ropc-service/repositories"
	"ropc-service/routers"

	"github.com/gorilla/mux"
)

type ApplicationHandler interface {
	CreateApplication(w http.ResponseWriter, r *http.Request)
}

type applicationHandler[T any] struct {
	router                routers.Router
	applicationRepository repositories.ApplicationRepository
}

func NewApplicationHandler(applicationRepository repositories.ApplicationRepository) ApplicationHandler {
	return &applicationHandler[*mux.Route]{
		applicationRepository: applicationRepository,
	}
}

func (c *applicationHandler[T]) CreateApplication(w http.ResponseWriter, r *http.Request) {

	var request *model.CreateApplication

	err := JsonToStruct(r.Body, &request)
	if err != nil {
		panic(errors.New("invalid prepareRequest body"))
	}

	app, err := model.NewApplication(request.ClientId, request.ClientSecret)
	if err != nil {
		panic(err)
	}

	err = c.applicationRepository.Create(app)
	if err != nil {
		panic(err)
	}

	response := model.NewResponse[*model.ApplicationDto]("application created successfully", app.ToDTO())

	_ = PrintResponse[*model.Response[*model.ApplicationDto]](http.StatusCreated, w, response)

}
