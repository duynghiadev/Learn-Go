package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
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
	dataFile   = "data.json"           // File lưu trữ dữ liệu
)

var fileMu sync.Mutex // Tạo mutex cho ghi file

// Hàm tải dữ liệu từ file JSON
func loadData() {
	file, err := os.Open(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("data.json not found, starting with empty data.")
			return
		}
		log.Fatalf("Failed to open data file: %v", err)
	}
	defer file.Close()

	// Đọc toàn bộ nội dung tệp
	var loadedProducts []Product
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&loadedProducts); err != nil {
		// Xử lý khi file rỗng hoặc không đúng định dạng
		if err.Error() == "EOF" {
			fmt.Println("data.json is empty, starting with empty data.")
			return
		}
		log.Fatalf("Failed to decode data file: %v", err)
	}

	productsMu.Lock()
	defer productsMu.Unlock()

	for _, product := range loadedProducts {
		products[product.ID] = product
		if product.ID >= nextID {
			nextID = product.ID + 1
		}
	}
	fmt.Println("Data loaded successfully.")
}

// Hàm lưu dữ liệu mới vào file (chỉ ghi sản phẩm mới)
func appendData(newProduct Product) {
	fileMu.Lock()
	defer fileMu.Unlock()

	// Mở file với chế độ "append" (nối vào cuối file)
	file, err := os.OpenFile(dataFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Printf("Failed to open data file: %v", err)
		return
	}
	defer file.Close()

	// Đọc dữ liệu hiện tại trong file
	var productsInFile []Product
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&productsInFile); err != nil && err.Error() != "EOF" {
		log.Printf("Failed to decode existing file data: %v", err)
	}

	// Thêm sản phẩm mới vào danh sách
	productsInFile = append(productsInFile, newProduct)

	// Quay lại đầu file và ghi lại toàn bộ danh sách
	file.Truncate(0) // Xóa nội dung cũ trong file
	file.Seek(0, 0)  // Quay lại đầu file

	// Ghi toàn bộ danh sách sản phẩm vào file
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(productsInFile); err != nil {
		log.Printf("Failed to encode products to file: %v", err)
	} else {
		fmt.Println("New product saved to file.")
	}
}

func saveData() {
	fileMu.Lock()
	defer fileMu.Unlock()

	file, err := os.Create(dataFile)
	if err != nil {
		log.Printf("Failed to create data file: %v", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	productsMu.Lock()
	defer productsMu.Unlock()
	if err := encoder.Encode(products); err != nil {
		log.Printf("Failed to encode data to file: %v", err)
	}
}

// Hàm phân trang và lọc
func paginateAndFilter(products map[int]Product, page, limit int, nameFilter string) ([]Product, int) {
	var filteredProducts []Product

	// Lọc theo tên sản phẩm nếu có
	for _, product := range products {
		productName := strings.ToLower(product.Name)
		filterName := strings.ToLower(nameFilter)

		if nameFilter != "" && !strings.Contains(productName, filterName) {
			continue
		}
		filteredProducts = append(filteredProducts, product)
	}

	// Tính tổng số sản phẩm sau khi lọc
	total := len(filteredProducts)
	fmt.Println("total:", total)

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

	// Lấy tham số truy vấn (query parameters)
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
	fmt.Println("totalPages:", totalPages)

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

// Hàm xử lý tạo sản phẩm mới
func createProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request to create product")

	var newProduct Product

	// Giải mã dữ liệu từ request body
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newProduct); err != nil {
		fmt.Println("Failed to decode request body:", err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	productsMu.Lock()
	defer productsMu.Unlock()

	// Gán ID cho sản phẩm mới và tăng nextID
	newProduct.ID = nextID
	nextID++
	products[newProduct.ID] = newProduct

	// Lưu sản phẩm mới vào file (append dữ liệu mới)
	fmt.Println("Product created successfully, saving to file...")
	appendData(newProduct)

	// Trả về sản phẩm đã tạo
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newProduct)
	fmt.Println("Response sent to client")
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

// middleware
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request: %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Tải dữ liệu từ file khi khởi động
	loadData()

	r := mux.NewRouter()

	// Định nghĩa các endpoint
	r.HandleFunc("/products", getProducts).Methods("GET")
	r.HandleFunc("/products", createProduct).Methods("POST")
	r.HandleFunc("/products/{id:[0-9]+}", getProduct).Methods("GET")
	r.HandleFunc("/products/{id:[0-9]+}", updateProduct).Methods("PUT")
	r.HandleFunc("/products/{id:[0-9]+}", deleteProduct).Methods("DELETE")
	r.Use(loggingMiddleware)

	fmt.Println("Server is running on port 8888")
	log.Fatal(http.ListenAndServe(":8888", r))
}
