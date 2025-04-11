package main

import (
	"fmt"
	"log"
	"method_and_interface/internal/product"
	"method_and_interface/pkg/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Tải dữ liệu từ file khi khởi động
	product.LoadData()

	r := mux.NewRouter()

	// Định nghĩa các endpoint
	r.HandleFunc("/products", product.GetProducts).Methods("GET")
	r.HandleFunc("/products", product.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{id:[0-9]+}", product.GetProduct).Methods("GET")
	r.HandleFunc("/products/{id:[0-9]+}", product.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id:[0-9]+}", product.DeleteProduct).Methods("DELETE")
	r.Use(utils.LoggingMiddleware)

	fmt.Println("Server is running on port 8888")
	log.Fatal(http.ListenAndServe(":8888", r))
}
