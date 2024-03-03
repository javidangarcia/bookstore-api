package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/javidangarcia/bookstore-api/data"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := data.Book{}

	decodeErr := json.NewDecoder(r.Body).Decode(&book)

	if decodeErr != nil || bookHasMissingFields(book) {
		respondWithError(w, "Some required fields are missing in the request.", http.StatusBadRequest)
		return
	}

	bookExists, bookExistsErr := data.BookExists(book.ISBN)

	if bookExistsErr != nil {
		respondWithError(w, "Oops! Something went wrong on our end. Please try again later.", http.StatusInternalServerError)
		return
	}

	if bookExists {
		respondWithError(w, "This book already exists in our database.", http.StatusConflict)
		return
	}

	createBookErr := data.CreateBook(book)

	if createBookErr != nil {
		respondWithError(w, "Oops! Something went wrong on our end. Please try again later.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Book successfully created."))
}

func GetBookByISBN(w http.ResponseWriter, r *http.Request) {
	isbn := chi.URLParam(r, "isbn")
	book, err := data.GetBookByISBN(isbn)

	if err != nil {
		respondWithError(w, "The book with the provided ISBN was not found.", http.StatusNotFound)
		return
	}

	bookJSON, err := json.Marshal(book)

	if err != nil {
		respondWithError(w, "Oops! Something went wrong on our end. Please try again later.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bookJSON)
}