package data

import (
	"context"
)

type Book struct {
	ISBN string `json:"isbn"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	AuthorID int `json:"authorId"`
	Genre string `json:"genre"`
	Publisher string `json:"publisher"`
	YearPublished int `json:"yearPublished"`
	CopiesSold int `json:"copiesSold"`
}

func CreateBook(book Book) error {
	tag, err := pool.Exec(context.Background(), `
			INSERT INTO books (isbn, name, description, price, author_id, 
			genre, publisher, year_published, copies_sold) 
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`, 
			book.ISBN, book.Name, book.Description, book.Price, book.AuthorID, 
			book.Genre, book.Publisher, book.YearPublished, book.CopiesSold)
			
	if err != nil && !tag.Insert() {
		return err
	}

	return nil
}

func FetchBookByISBN(isbn string) (Book, error) {
	book := &Book{}

	row := pool.QueryRow(context.Background(), `
		SELECT isbn, name, description, price, author_id, genre, publisher, year_published, copies_sold 
		FROM books 
		WHERE isbn = $1`, isbn)

	err := row.Scan(&book.ISBN, &book.Name, &book.Description, &book.Price, &book.AuthorID, &book.Genre, 
					&book.Publisher, &book.YearPublished, &book.CopiesSold)

	if err != nil {
		return Book{}, err
	}

	return *book, nil
}

func BookExists(isbn string) (bool, error) {
	var count int

	row := pool.QueryRow(context.Background(), `SELECT COUNT(*) FROM books WHERE isbn = $1`, isbn)

	err := row.Scan(&count)

	if err != nil {
		return false, err
	}

	return count == 1, nil
}

func FetchBooksByAuthorID(authorID int) ([]Book, error) {
	var books []Book

	rows, err := pool.Query(context.Background(), `
		SELECT isbn, name, description, price, author_id, genre, publisher, year_published, copies_sold
		FROM books 
		WHERE author_id = $1`, authorID)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var book Book

		rows.Scan(&book.ISBN, &book.Name, &book.Description, &book.Price, &book.AuthorID, &book.Genre, 
				&book.Publisher, &book.YearPublished, &book.CopiesSold)

		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, err
}