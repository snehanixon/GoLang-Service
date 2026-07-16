package main

import (
	"fmt"
	"net/http"
)

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")

	if city == "" {
		city = "Unknown"
	}

	fmt.Fprintf(w, "Weather in %s: Sunny, 30°C", city)
}

func main() {
	http.HandleFunc("/weather", weatherHandler)

	fmt.Println("Weather Service running on http://localhost:8086")
	http.ListenAndServe(":8086", nil)
}