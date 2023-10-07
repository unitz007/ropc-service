package main

import (
	"ropc-service/authenticators"
	"ropc-service/conf"
	"ropc-service/handlers"
	"ropc-service/middlewares"
	"ropc-service/repositories"
	"ropc-service/routers"

	"github.com/gorilla/mux"
)

const (
	loginPath  = "/token"
	userPath   = "/users"
	clientPath = "/apps"
)

func main() {

	router := routers.NewMuxMultiplexer(mux.NewRouter())

	// set up request handlers
	requestHandlers := func(router routers.Router) {

		config := conf.EnvironmentConfig
		DB := conf.NewDataBase(config)

		// Repositories
		clientRepository := repositories.NewClientRepository(DB)
		//userRepository := repositories.NewUserRepository(DB)

		clientAuthenticator := authenticators.NewClientAuthenticator(clientRepository)
		//userAuthenticator := authenticators.NewUserAuthenticator(userRepository)

		//authenticator := services.NewAuthenticator(clientAuthenticator)

		authenticationHandler := handlers.NewAuthenticationHandler(router, clientAuthenticator)
		userHandler := handlers.NewUserAuthenticatorHandler()

		clientHandler := handlers.NewApplicationHandler(clientRepository)

		// authenticate
		router.Post(loginPath, middlewares.PanicRecovery(authenticationHandler.Authenticate))
		router.Get(loginPath, middlewares.PanicRecovery(authenticationHandler.LoginPage))

		// user
		router.Post(userPath, middlewares.RequestLogger(middlewares.PanicRecovery(userHandler.CreateUser)))

		// client
		router.Post(clientPath, middlewares.PanicRecovery(clientHandler.CreateApplication))

		//secured resource
		router.Get(userPath, middlewares.PanicRecovery(middlewares.Security(userHandler.GetUserDetails)))
	}

	conf.InitServer(router, requestHandlers)

}
