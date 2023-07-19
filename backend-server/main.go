package main

import (
	"net/http"
	"ropc-service/conf"
	"ropc-service/handlers"
	"ropc-service/middlewares"
)

func main() {
	environmentConfig := conf.LoadConfig() // load environment configuration
	conf.InitGormConfig(environmentConfig) // initialize database

	// set up request handlers
	requestHandlers := func(mux *http.ServeMux) {

		// authenticate
		mux.HandleFunc("/login", middlewares.PanicRecovery(handlers.Authentication))

		// user
		mux.HandleFunc("/users", middlewares.PanicRecovery(handlers.CreateUser))

		//secured resource
		mux.HandleFunc("/user_details", middlewares.PanicRecovery(middlewares.Security(handlers.UserDetailsHandler)))
	}

	// load gin context
	conf.InitServer(environmentConfig, requestHandlers)
}
