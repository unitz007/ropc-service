package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"ropc-service/authenticators"
	"ropc-service/conf"
	"ropc-service/model/dto"
	"ropc-service/model/entities"
	"ropc-service/repositories"
	"ropc-service/routers"
	"ropc-service/services"
)

const authenticationSuccessMsg = "Authentication successful"

type AuthenticationHandler interface {
	GetMux() *mux.Router
	Login(w http.ResponseWriter, r *http.Request)
	LoginPage(w http.ResponseWriter, r *http.Request)
}

type authenticationHandler[T any] struct {
	router routers.Router
}

func NewAuthenticationHandler(router routers.Router) AuthenticationHandler {
	return &authenticationHandler[*mux.Router]{router}
}

func (a *authenticationHandler[T]) Login(w http.ResponseWriter, r *http.Request) {

	var loginRequest *dto.LoginRequest

	err := JsonToStruct(r.Body, &loginRequest)
	if err != nil {
		panic(err)
	}

	database := conf.NewDataBase(conf.EnvironmentConfig)

	userRepository := repositories.NewUserRepository(database)
	clientRepository := repositories.NewClientRepository(database)

	userAuthenticator := authenticators.NewUserAuthenticator(userRepository)
	clientAuthenticator := authenticators.NewClientAuthenticator(clientRepository)

	user := &entities.User{
		Username: loginRequest.Username,
		Password: loginRequest.Password,
	}

	client := &entities.Client{
		ClientId:     loginRequest.ClientId,
		ClientSecret: loginRequest.ClientSecret,
		GrantType:    loginRequest.GrantType,
	}

	token, err := services.InstantiateAuthenticator(userAuthenticator, clientAuthenticator).Authenticate(user, client)
	if err != nil {
		_ = PrintResponse(http.StatusUnauthorized, w, &dto.Response[string]{Message: err.Error()})
		return
	}

	_ = PrintResponse(http.StatusOK, w, dto.NewResponse(authenticationSuccessMsg, token))
}

func (a *authenticationHandler[T]) LoginPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./assets/html/login.html")
}

func (a *authenticationHandler[T]) GetMux() *mux.Router {
	return &mux.Router{}
}
