package handler

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func search(w http.ResponseWriter, r *http.Request) error {
	request := chi.URLParam(r, "req")
}
