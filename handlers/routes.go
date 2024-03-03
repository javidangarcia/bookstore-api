package handlers

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Post("/books", CreateBook)
	r.Get("/books/{isbn}", GetBookByISBN)

	return r
}