package data

import "context"

type Author struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Biography string `json:"biography"`
	Publisher string `json:"publisher"`
}

func CreateAuthor(author Author) error {
	tag, err := pool.Exec(context.Background(), `
			INSERT INTO authors (first_name, last_name, biography, publisher)
			VALUES ($1, $2, $3, $4)`, 
			author.FirstName, author.LastName, author.Biography, author.Publisher)
			
	if err != nil && !tag.Insert() {
		return err
	}

	return nil
}