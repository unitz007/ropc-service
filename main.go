package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"ropc-service/conf"
	"ropc-service/handlers"
)

func main() {

	globalConfig := conf.LoadConfig()
	conf.InitGormConfig(&globalConfig)

	gin.SetMode(globalConfig.GinMode)

	router := gin.Default()
	router.Use(gin.Recovery())

	router.POST("/login", handlers.Authentication)

	err := router.Run(fmt.Sprintf(":%s", globalConfig.ServerPort))
	if err != nil {
		log.Fatal(err)
	}
}
