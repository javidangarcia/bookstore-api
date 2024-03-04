CREATE TABLE IF NOT EXISTS authors (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    biography TEXT NOT NULL,
    publisher VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS books (
    isbn VARCHAR(13) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    author_id INT,
    genre VARCHAR(255),
    publisher VARCHAR(255),
    year_published INT,
    copies_sold INT,
    FOREIGN KEY (author_id) REFERENCES authors(id)
);