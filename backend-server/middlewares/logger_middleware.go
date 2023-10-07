package middlewares

import (
	"fmt"
	"net/http"
	"ropc-service/logger"
)

func RequestLogger(h func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		s := fmt.Sprintf("%v, %v", r.Method, r.URL.Path)
		logger.Info(s)
		h(w, r)

	}

}
