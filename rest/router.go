package rest

import (
	"github.com/go-chi/chi/v5"
)

type Router struct {
	R *chi.Mux
}

func NewRouter() *Router {
	rr := chi.NewRouter()
	return &Router{
		R: rr,
	}
}
