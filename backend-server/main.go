package main

import (
	"backend-server/conf"
	"backend-server/handlers"
	"backend-server/repositories"
	"backend-server/routers"
	"backend-server/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	loginPath = "/token"
	appPath   = "/apps"
)

func main() {

	// properties
	mul := mux.NewRouter()
	router := routers.NewRouter(mul)
	config := conf.NewConfig()
	DB := conf.NewDataBase(config)

	// Repositories
	applicationRepository := repositories.NewApplicationRepository(DB)

	// services
	authenticatorService := services.NewAuthenticatorService(applicationRepository)

	// Handlers
	authenticationHandler := handlers.NewAuthenticationHandler(authenticatorService)
	applicationHandler := handlers.NewApplicationHandler(applicationRepository)

	// Server
	server := NewServer(router)
	server.RegisterHandler(appPath, http.MethodPost, applicationHandler.CreateApplication)
	server.RegisterHandler(loginPath, http.MethodPost, authenticationHandler.Authenticate)

	log.Fatal(server.Start(":" + config.ServerPort()))
}
