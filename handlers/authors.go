package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/javidangarcia/bookstore-api/data"
)

func handleCreateAuthor(w http.ResponseWriter, r *http.Request) {
	author := data.Author{}

	decodeErr := json.NewDecoder(r.Body).Decode(&author)

	if decodeErr != nil || authorHasMissingFields(author) {
		respondWithError(w, "Some required fields are missing in the request.", http.StatusBadRequest)
		return
	}

	createAuthorErr := data.CreateAuthor(author)

	if createAuthorErr != nil {
		respondWithError(w, "Oops! Something went wrong on our end. Please try again later.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Author successfully created."))
}