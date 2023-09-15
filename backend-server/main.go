package main

import (
	"ropc-service/conf"
	"ropc-service/handlers"
	"ropc-service/middlewares"
	"ropc-service/repositories"
	"ropc-service/routers"
	"ropc-service/services"

	"github.com/gorilla/mux"
)

const (
	loginPath  = "/login"
	userPath   = "/users"
	clientPath = "/clients"
)

func main() {

	router := routers.NewMuxMultiplexer(mux.NewRouter())

	// set up request handlers
	requestHandlers := func(router routers.Router) {

		authenticationHandler := handlers.NewAuthenticationHandler(router)
		userHandler := handlers.NewUserAuthenticatorHandler()
		config := conf.EnvironmentConfig
		db := conf.NewDataBase(config)

		clientRepository := repositories.NewClientRepository(db)
		clientService := services.NewClientService(clientRepository)
		clientHandler := handlers.NewClientHandler(clientService)

		// authenticate
		router.Post(loginPath, middlewares.PanicRecovery(authenticationHandler.Login))
		router.Get(loginPath, middlewares.PanicRecovery(authenticationHandler.LoginPage))

		// user
		router.Post(userPath, middlewares.PanicRecovery(userHandler.CreateUser))

		// client
		router.Post(clientPath, middlewares.PanicRecovery(clientHandler.CreateClient))

		//secured resource
		router.Get(userPath, middlewares.PanicRecovery(middlewares.Security(userHandler.GetUserDetails)))
	}

	conf.InitServer(router, requestHandlers)

}
