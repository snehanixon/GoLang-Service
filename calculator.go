package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func addHandler(w http.ResponseWriter, r *http.Request) {
	a, _ := strconv.Atoi(r.URL.Query().Get("a"))
	b, _ := strconv.Atoi(r.URL.Query().Get("b"))

	result := a + b
	fmt.Fprintf(w, "Addition Result: %d", result)
}

func subHandler(w http.ResponseWriter, r *http.Request) {
	a, _ := strconv.Atoi(r.URL.Query().Get("a"))
	b, _ := strconv.Atoi(r.URL.Query().Get("b"))

	result := a - b
	fmt.Fprintf(w, "Subtraction Result: %d", result)
}

func mulHandler(w http.ResponseWriter, r *http.Request) {
	a, _ := strconv.Atoi(r.URL.Query().Get("a"))
	b, _ := strconv.Atoi(r.URL.Query().Get("b"))

	result := a * b
	fmt.Fprintf(w, "Multiplication Result: %d", result)
}

func divHandler(w http.ResponseWriter, r *http.Request) {
	a, _ := strconv.Atoi(r.URL.Query().Get("a"))
	b, _ := strconv.Atoi(r.URL.Query().Get("b"))

	if b == 0 {
		fmt.Fprintf(w, "Error: Division by zero is not allowed")
		return
	}

	result := float64(a) / float64(b)
	fmt.Fprintf(w, "Division Result: %.2f", result)
}

func main() {
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/sub", subHandler)
	http.HandleFunc("/mul", mulHandler)
	http.HandleFunc("/div", divHandler)

	fmt.Println("Calculator Service running on http://localhost:8083")

	err := http.ListenAndServe(":8083", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}