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

func (r *Router) PathMetaRoutes(h PathMetaHandlers) {
	r.Mux.Mount("/debug", middleware.Profiler())
	r.Mux.Route("/shorten", func(r chi.Router) {
		r.Post("/", h.addPath)

		r.Get("/{hash}", h.getFullUrl)
		r.Get("/{hash}/visits", h.listVisits)
	})
}
