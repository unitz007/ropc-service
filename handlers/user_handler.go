package handlers

import (
	"net/http"
	"ropc-service/model/dto"
	"ropc-service/model/entities"
	"ropc-service/repositories"
	"ropc-service/services"
)

const (
	UserCreated = "User created Successfully"
)

func CreateUser(response http.ResponseWriter, request *http.Request) {

	switch request.Method {
	case http.MethodPost:
		var user *entities.User

		err := JsonToStruct(request.Body, &user)
		if err != nil {
			panic(err)
		}

		userRepository := repositories.NewUserRepository()
		userService := services.NewUserService(userRepository)

		_, err = userService.CreateUser(user)
		if err != nil {
			panic(err)

		} else {
			_ = PrintResponse(http.StatusCreated, response, dto.NewResponse(UserCreated, nil))
		}

		return
	default:
		response.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
