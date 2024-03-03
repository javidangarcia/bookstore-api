package data

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func ConnectDatabase() {
	var err error
	pool, err = pgxpool.New(context.Background(), os.Getenv("CONNECTION_STRING"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
}

func CloseDatabase() {
	pool.Close()
}