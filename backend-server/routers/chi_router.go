package routers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ChiRouter struct {
	router *chi.Mux
}

func NewChiRouter(router *chi.Mux) Router {
	return &ChiRouter{router: router}
}

func (mux *ChiRouter) Get(path string, handlerFunc func(w http.ResponseWriter, r *http.Request)) {
	mux.router.Get(path, handlerFunc)
}

func (mux *ChiRouter) Serve(addr string) error {
	return http.ListenAndServe(addr, mux.router)
}

func (mux *ChiRouter) Post(path string, handlerFunc func(w http.ResponseWriter, r *http.Request)) {
	mux.router.Post(path, handlerFunc)
}
