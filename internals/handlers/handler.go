package handlers

import (
	"github/metasoftonic/coffee-api/internals/repository"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler(store repository.Store) *Handler {
	h := &Handler{
		Mux:  chi.NewMux(),
		repo: store,
	}

	h.Route("/coffees", func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Get("/", h.CoffeList())
		r.Post("/", h.CreateCoffee())
	})

	return h
}

type Handler struct {
	*chi.Mux
	repo repository.Store
}
