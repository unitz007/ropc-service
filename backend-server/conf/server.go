package conf

import (
	"log"
	"ropc-service/routers"
)

func InitServer(router routers.Router, registerHandlers func(router routers.Router)) {

	// endpoints
	registerHandlers(router)

	// serve
	log.Fatal(router.Serve(":" + EnvironmentConfig.ServerPort()))
}
