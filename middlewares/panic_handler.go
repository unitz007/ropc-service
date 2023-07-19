package middlewares

import (
	"log"
	"net/http"
	"ropc-service/handlers"
	"ropc-service/model/dto"
)

func PanicRecovery(h func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				errorMsg := "Something went wrong"
				if e, ok := err.(error); ok {
					errorMsg = e.Error()
				}
				_ = handlers.PrintResponse(http.StatusBadRequest, w, dto.NewResponse(errorMsg, nil))
			}
		}()

		h(w, r)
	}
}
