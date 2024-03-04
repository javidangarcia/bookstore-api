package handlers

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Post("/books", handleCreateBook)
	r.Get("/books/{isbn}", handleGetBookByISBN)
	r.Post("/authors", handleCreateAuthor)
	r.Get("/authors/{authorID}/books", handleGetBooksByAuthorID)

	return r
}