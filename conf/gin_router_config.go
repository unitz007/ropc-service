package conf

import (
	"log"
	"net/http"
)

func InitGinContext(environmentConfig *Config, registerHandlers func(mux *http.ServeMux)) {
	//gin.SetMode(environmentConfig.GinMode)

	// router configuration
	//router := gin.Default()
	mux := http.DefaultServeMux

	//router.Use(gin.Recovery())

	// serve static files
	//router.Static("/assets", "./assets")
	//router.LoadHTMLGlob("assets/**/*")

	// endpoints
	registerHandlers(mux)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
