package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	//http.HandleFunc("/api/search", SearchHandler)
	http.HandleFunc("/api/details", DetailHandler)
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
