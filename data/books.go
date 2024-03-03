package data

import (
	"context"
)

type Book struct {
	ISBN string `json:"isbn"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	Author string `json:"author"`
	Genre string `json:"genre"`
	Publisher string `json:"publisher"`
	YearPublished int `json:"yearPublished"`
	CopiesSold int `json:"copiesSold"`
}

func CreateBook(book Book) error {
	tag, err := pool.Exec(context.Background(), `
			INSERT INTO books (isbn, name, description, price, author, 
			genre, publisher, year_published, copies_sold) 
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`, 
			book.ISBN, book.Name, book.Description, book.Price, book.Author, 
			book.Genre, book.Publisher, book.YearPublished, book.CopiesSold)
			
	if err != nil && !tag.Insert() {
		return err
	}

	return nil
}