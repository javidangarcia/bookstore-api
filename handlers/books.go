package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/javidangarcia/bookstore-api/data"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := data.Book{}

	decodeErr := json.NewDecoder(r.Body).Decode(&book)

	if decodeErr != nil {
		respondWithError(w, "Invalid request body.", http.StatusBadRequest)
		return
	}

	createBookErr := data.CreateBook(book)

	if createBookErr != nil {
		respondWithError(w, "This book already exists.", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Book has been successfully created."))
}