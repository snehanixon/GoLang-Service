package main

import (
	"fmt"
	"net/http"
)

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List of Users")
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User Details")
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User Created Successfully")
}

func main() {
	http.HandleFunc("/users", getUsersHandler)
	http.HandleFunc("/user", getUserHandler)
	http.HandleFunc("/create-user", createUserHandler)

	fmt.Println("User Service running on http://localhost:8082")

	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}