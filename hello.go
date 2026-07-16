package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", hello)

	fmt.Println("Hello World Service Started...")
	fmt.Println("Open: http://localhost:8080/")

	http.ListenAndServe(":8080", nil)
}