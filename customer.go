package main

import (
	"encoding/json"
	"net/http"
)

type Customer struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func customerHandler(w http.ResponseWriter, r *http.Request) {
	customer := Customer{
		ID:    1,
		Name:  "Sneha",
		Email: "sneha@example.com",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

func main() {
	http.HandleFunc("/customer", customerHandler)
	println("Customer Service running on :8090")
	http.ListenAndServe(":8090", nil)
}