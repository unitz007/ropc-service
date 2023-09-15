package conf

import (
	"log"
	"ropc-service/routers"
	"time"
)

func InitServer(router routers.Router, registerHandlers func(router routers.Router)) {

	// endpoints
	registerHandlers(router)

	PORT := EnvironmentConfig.ServerPort()

	go func() {
		time.Sleep(time.Millisecond * 10)
		log.Println("Server started on", PORT)
	}()

	log.Fatal(router.Serve(":" + PORT))
}
