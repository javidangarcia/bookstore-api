package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Post("/books", CreateBook)

	return r
}

type ApiError struct {
	Error string `json:"error"`
}

func respondWithError(w http.ResponseWriter, msg string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ApiError{Error: msg})
}