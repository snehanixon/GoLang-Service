package main

import (
	"encoding/json"
	"net/http"
)

type Student struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Course string `json:"course"`
}

func studentHandler(w http.ResponseWriter, r *http.Request) {
	student := Student{
		ID:     101,
		Name:   "Sneha",
		Course: "Computer Science",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func main() {
	http.HandleFunc("/student", studentHandler)
	println("Student Service running on :8093")
	http.ListenAndServe(":8093", nil)
}