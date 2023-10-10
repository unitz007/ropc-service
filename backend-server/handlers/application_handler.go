package handlers

import (
	"backend-server/model"
	"backend-server/repositories"
	"errors"
	"net/http"
)

type ApplicationHandler interface {
	CreateApplication(w http.ResponseWriter, r *http.Request)
}

type applicationHandler struct {
	applicationRepository repositories.ApplicationRepository
}

func NewApplicationHandler(applicationRepository repositories.ApplicationRepository) ApplicationHandler {
	return &applicationHandler{
		applicationRepository: applicationRepository,
	}
}

func (c *applicationHandler) CreateApplication(w http.ResponseWriter, r *http.Request) {

	var request *model.CreateApplication

	err := JsonToStruct(r.Body, &request)
	if err != nil {
		panic(errors.New("invalid prepareRequest body"))
	}

	app, err := model.NewApplication(request.ClientId)
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
