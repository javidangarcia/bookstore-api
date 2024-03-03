package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/javidangarcia/bookstore-api/data"
)

type ApiError struct {
	Error string `json:"error"`
}

func respondWithError(w http.ResponseWriter, msg string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ApiError{Error: msg})
}

func bookHasMissingFields(book data.Book) bool {
	return book.ISBN == "" || book.Name == "" || book.Description == "" || book.Price == 0 || 
	book.Author == "" || book.Genre == "" || book.Publisher == "" || book.YearPublished == 0 || 
	book.CopiesSold == 0
}