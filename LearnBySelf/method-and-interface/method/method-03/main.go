package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
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

// Struct để trả về thông tin kết quả bao gồm danh sách sản phẩm và thông tin phân trang
type PaginatedResponse struct {
	Total      int       `json:"total"`
	Page       int       `json:"page"`
	Limit      int       `json:"limit"`
	TotalPages int       `json:"total_pages"`
	Products   []Product `json:"products"`
}

var (
	products   = make(map[int]Product) // Lưu trữ danh sách sản phẩm
	nextID     = 1                     // ID tự động tăng
	productsMu sync.Mutex              // Bảo vệ dữ liệu khi truy cập đồng thời
)

// Hàm phân trang và lọc
func paginateAndFilter(products map[int]Product, page, limit int, nameFilter string) ([]Product, int) {
	var filteredProducts []Product

	// Lọc theo tên sản phẩm nếu có
	for _, product := range products {
		if nameFilter != "" && !strings.Contains(strings.ToLower(product.Name), strings.ToLower(nameFilter)) {
			continue
		}
		filteredProducts = append(filteredProducts, product)
	}

	// Tính tổng số sản phẩm sau khi lọc
	total := len(filteredProducts)

	// Phân trang
	start := (page - 1) * limit
	end := start + limit
	if start > total {
		return []Product{}, total
	}
	if end > total {
		end = total
	}

	return filteredProducts[start:end], total
}

// Handler lấy danh sách sản phẩm với phân trang và lọc
func getProducts(w http.ResponseWriter, r *http.Request) {
	productsMu.Lock()
	defer productsMu.Unlock()

	// Lấy tham số truy vấn
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	limit, err2 := strconv.Atoi(r.URL.Query().Get("limit"))

	fmt.Println("page:", page, "limit:", limit)

	// Nếu không có tham số page hoặc limit, trả về tất cả sản phẩm
	if err != nil || page <= 0 || err2 != nil || limit <= 0 {
		// Trả về tất cả sản phẩm
		var allProducts []Product
		for _, product := range products {
			allProducts = append(allProducts, product)
		}

		// Nếu không có sản phẩm nào
		if len(allProducts) == 0 {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(allProducts)
		return
	}

	// Nếu có tham số page và limit, thực hiện phân trang và lọc
	nameFilter := r.URL.Query().Get("name")
	paginatedProducts, total := paginateAndFilter(products, page, limit, nameFilter)

	// Tính tổng số trang
	totalPages := (total + limit - 1) / limit

	// Tạo response chứa thông tin phân trang và sản phẩm
	response := PaginatedResponse{
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
		Products:   paginatedProducts,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Handler thêm một sản phẩm mới
func createProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct Product

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newProduct); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

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
