package handlers

import (
	"errors"
	"net/http"
	"ropc-service/authenticators"
	"ropc-service/model"
	"ropc-service/routers"

	"github.com/gorilla/mux"
)

const authenticationSuccessMsg = "Authentication successful"

type AuthenticationHandler interface {
	GetMux() *mux.Router
	Authenticate(w http.ResponseWriter, r *http.Request)
	LoginPage(w http.ResponseWriter, r *http.Request)
}

type authenticationHandler[T any] struct {
	router        routers.Router
	authenticator authenticators.ClientAuthenticator
}

func NewAuthenticationHandler(router routers.Router, authenticator authenticators.ClientAuthenticator) AuthenticationHandler {
	return &authenticationHandler[*mux.Router]{router, authenticator}
}

func (a *authenticationHandler[T]) Authenticate(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		panic(errors.New("invalid content-type"))
	}

	clientId := r.FormValue("client_id")
	if clientId == "" {
		panic(errors.New("client id is required"))
	}

	clientSecret := r.FormValue("client_secret")
	if clientSecret == "" {
		panic(errors.New("client secret is required"))
	}

	grantType := r.FormValue("grant_type")
	if grantType == "" {
		panic(errors.New("grant type is required"))
	}

	token, err := a.authenticator.Authenticate(clientId, clientSecret)
	if err != nil {
		_ = PrintResponse(http.StatusUnauthorized, w, &model.Response[string]{Message: err.Error()})
		return
	}

	_ = PrintResponse[any](http.StatusOK, w, token)
}

func (a *authenticationHandler[T]) LoginPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./assets/html/login.html")
}

func (a *authenticationHandler[T]) GetMux() *mux.Router {
	return &mux.Router{}
}
