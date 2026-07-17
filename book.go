package main

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func bookHandler(w http.ResponseWriter, r *http.Request) {
	book := Book{
		ID:     1,
		Title:  "Go Programming",
		Author: "Alan Donovan",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func main() {
	http.HandleFunc("/book", bookHandler)
	println("Book Service running on :8092")
	http.ListenAndServe(":8092", nil)
}