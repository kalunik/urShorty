package controller

import (
	"github.com/go-chi/chi/v5"
)

type Router struct {
	Mux *chi.Mux
}

func NewRouter() *Router {
	r := chi.NewRouter()
	//middleware

	r.Route("/shorty", func(r chi.Router) {
		r.Post("/", addLink)
		r.Post("/", addFriendlyLink)
		r.Get("/{key}", searchLink)
	})
	return &Router{Mux: r}
}
