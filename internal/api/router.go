package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router struct {
	Mux *chi.Mux
}

func NewRouter() *Router {
	return &Router{Mux: chi.NewRouter()}
}

func (r *Router) UrlPairRoutes(h UrlPairHandlers) {
	r.Mux.Mount("/debug", middleware.Profiler())
	r.Mux.Route("/shorten", func(r chi.Router) {

		r.Post("/", h.addPair)
		r.Post("/{url}", h.addPairHashParam)

		r.Get("/{hash}", h.getFullUrl)
	})
}
