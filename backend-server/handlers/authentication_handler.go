package handlers

import (
	"net/http"
	"ropc-service/authenticators"
	"ropc-service/model/dto"
	"ropc-service/model/entities"
	"ropc-service/repositories"
	"ropc-service/services"
)

const authenticationSuccessMsg = "Authentication successful"

func Authentication(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		http.ServeFile(w, r, "./assets/html/login.html")
		return

	case http.MethodPost:
		var loginRequest *dto.LoginRequest

		err := JsonToStruct(r.Body, &loginRequest)
		if err != nil {
			panic(err)
		}

		userRepository := repositories.NewUserRepository()
		clientRepository := repositories.NewClientRepository()

		userAuthenticator := authenticators.InstantiateUserAuthenticator(userRepository)
		clientAuthenticator := authenticators.InstantiateClientAuthenticator(clientRepository)

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
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
