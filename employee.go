package main

import (
	"encoding/json"
	"net/http"
)

type Employee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

func employeeHandler(w http.ResponseWriter, r *http.Request) {
	employee := Employee{
		ID:   1001,
		Name: "John",
		Role: "Developer",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employee)
}

func main() {
	http.HandleFunc("/employee", employeeHandler)
	println("Employee Service running on :8091")
	http.ListenAndServe(":8091", nil)
}