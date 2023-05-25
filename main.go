package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"ropc-service/handlers"
	"ropc-service/repositories"
)

func main() {

	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	gin.SetMode(os.Getenv("GIN_MODE"))

	router := gin.Default()
	router.Use(gin.Recovery())
	repositories.InitGormConfig()
	repositories.DatabaseConfig.PreLoadData()

	router.POST("/login", handlers.Authentication)

	err = router.Run(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
	if err != nil {
		log.Fatal(err)
	}
}
