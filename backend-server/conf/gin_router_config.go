package conf

import (
	"log"
	"net/http"
)

func InitServer(environmentConfig *Config, registerHandlers func(mux *http.ServeMux)) {
	//gin.SetMode

	// router configuration
	mux := http.DefaultServeMux
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// endpoints
	registerHandlers(mux)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
