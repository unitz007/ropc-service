package handlers

import (
	"net/http"
	"ropc-service/conf"
	"ropc-service/model/dto"
	"ropc-service/model/entities"
	"ropc-service/repositories"
	"ropc-service/services"
	"ropc-service/utils"
)

const (
	userCreated = "User created Successfully"
)

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUserDetails(w http.ResponseWriter, r *http.Request)
}

type userHandler struct{}

func NewUserAuthenticatorHandler() UserHandler {
	return &userHandler{}
}

func (u *userHandler) CreateUser(response http.ResponseWriter, request *http.Request) {
	var user *entities.User

	err := JsonToStruct(request.Body, &user)
	if err != nil {
		panic(err)
	}

	database := conf.NewDataBase(conf.EnvironmentConfig)

	userRepository := repositories.NewUserRepository(database)
	userService := services.NewUserService(userRepository)

	_, err = userService.CreateUser(user)
	if err != nil {
		panic(err)

	}
	_ = PrintResponse(http.StatusCreated, response, dto.NewResponse[any](userCreated, nil))
}

func (u *userHandler) GetUserDetails(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		accessToken := r.Header.Get("Authorization")
		claims, err := utils.ValidateToken(accessToken)
		if err != nil {
			panic(err)
		}

		userDetails := dto.UserDetails{
			Username: claims["username"].(string),
			Email:    claims["email"].(string),
			ClientId: claims["client_id"].(string),
		}

		_ = PrintResponse(http.StatusOK, w, dto.NewResponse("User details fetched successfully",
			userDetails))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
