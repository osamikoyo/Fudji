package handler

import "github.com/go-chi/chi/v5"

func RegisterRoutes(r *chi.Mux) {
	r.Get("/search/{req}", getErrorHandlerFunc(search))
}
