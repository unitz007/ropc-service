package routers

import (
	"net/http"
)

type Router interface {
	Get(path string, handlerFunc func(w http.ResponseWriter, r *http.Request))
	Serve(addr string) error
	Post(path string, handler func(w http.ResponseWriter, r *http.Request))
}
