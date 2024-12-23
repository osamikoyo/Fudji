package handler

import (
	"fudji/internal/loger"
	"net/http"
)

type handlerFunc func(w http.ResponseWriter, r *http.Request) error

func getErrorHandlerFunc(h handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			loger.New().Error().Err(err)
		}
	}
}
