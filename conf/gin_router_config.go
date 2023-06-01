package conf

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func InitGinContext(environmentConfig *Config, registerHandlers func(engine *gin.Engine)) {
	gin.SetMode(environmentConfig.GinMode)

	// router configuration
	router := gin.Default()
	router.Use(gin.Recovery())

	// serve static files
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("assets/**/*")

	// endpoints
	registerHandlers(router)

	err := router.Run(fmt.Sprintf(":%s", environmentConfig.ServerPort))
	if err != nil {
		log.Fatal(err)
	}
}
