package conf

import (
	"fmt"
	"ropc-service/logger"
	"ropc-service/routers"
	"time"
)

func InitServer(router routers.Router, registerHandlers func(router routers.Router)) {

	// endpoints
	registerHandlers(router)

	PORT := EnvironmentConfig.ServerPort()

	go func() {
		time.Sleep(time.Millisecond * 10)
		msg := fmt.Sprintf("Server started on port %s, with %s", PORT, router.Name())
		logger.Info(msg)
	}()

	err := router.Serve(":" + PORT)

	if err != nil {
		logger.Error(err.Error(), true)
	}
}
