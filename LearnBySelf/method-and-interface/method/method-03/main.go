package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

// Struct Product đại diện cho một sản phẩm
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

var (
	products   = make(map[int]Product) // Lưu trữ danh sách sản phẩm
	nextID     = 1                     // ID tự động tăng
	productsMu sync.Mutex              // Bảo vệ dữ liệu khi truy cập đồng thời
)

// Handler lấy danh sách sản phẩm
func getProducts(w http.ResponseWriter, r *http.Request) {
	productsMu.Lock()
	defer productsMu.Unlock()

	var productList []Product
	for _, product := range products {
		productList = append(productList, product)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(productList)
}

// Handler thêm một sản phẩm mới
func createProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct Product

	// type 1
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newProduct); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// type 2
	// if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
	// 	http.Error(w, "Invalid input", http.StatusBadRequest)
	// 	return
	// }

	productsMu.Lock()
	defer productsMu.Unlock()

	newProduct.ID = nextID
	nextID++
	products[newProduct.ID] = newProduct

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newProduct)
}

// Handler lấy thông tin chi tiết của một sản phẩm
func getProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	productsMu.Lock()
	defer productsMu.Unlock()

	product, exists := products[id]
	if !exists {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// Handler cập nhật thông tin sản phẩm
func updateProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedProduct Product
	if err := json.NewDecoder(r.Body).Decode(&updatedProduct); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	productsMu.Lock()
	defer productsMu.Unlock()

	_, exists := products[id]
	if !exists {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	updatedProduct.ID = id
	products[id] = updatedProduct

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedProduct)
}

// Handler xóa một sản phẩm
func deleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	productsMu.Lock()
	defer productsMu.Unlock()

	_, exists := products[id]
	if !exists {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	delete(products, id)

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	r := mux.NewRouter()

	// Định nghĩa các endpoint
	r.HandleFunc("/products", getProducts).Methods("GET")
	r.HandleFunc("/products", createProduct).Methods("POST")
	r.HandleFunc("/products/{id:[0-9]+}", getProduct).Methods("GET")
	r.HandleFunc("/products/{id:[0-9]+}", updateProduct).Methods("PUT")
	r.HandleFunc("/products/{id:[0-9]+}", deleteProduct).Methods("DELETE")

	fmt.Println("Server is running on port 8888")
	log.Fatal(http.ListenAndServe(":8888", r))
}
