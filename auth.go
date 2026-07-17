package main

import (
	"fmt"
	"net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login Successful")
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User Registered Successfully")
}

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/register", registerHandler)

	fmt.Println("Auth Service running on http://localhost:8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}