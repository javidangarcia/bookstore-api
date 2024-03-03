CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    isbn VARCHAR(13) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    author VARCHAR(255) NOT NULL,
    genre VARCHAR(255),
    publisher VARCHAR(255),
    year_published INT,
    copies_sold INT
);