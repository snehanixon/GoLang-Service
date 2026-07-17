package main

import (
	"fmt"
	"net/http"
)

func products(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Product List")
}

func product(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Product Details")
}

func addProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Product Added Successfully")
}

func main() {
	http.HandleFunc("/products", products)
	http.HandleFunc("/product", product)
	http.HandleFunc("/add-product", addProduct)

	fmt.Println("Product Service running on http://localhost:8084")
	http.ListenAndServe(":8084", nil)
}

