package main

import (
	"fmt"
	"net/http"

	"github.com/javidangarcia/bookstore-api/data"
	"github.com/javidangarcia/bookstore-api/handlers"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	data.ConnectDatabase()
	defer data.CloseDatabase()

	r := handlers.Routes()

	fmt.Println("Server is running on port 3000...")

	http.ListenAndServe(":3000", r)
}