package main

import (
	"github.com/javidangarcia/bookstore-api/data"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	data.ConnectDatabase()
	defer data.CloseDatabase()
}