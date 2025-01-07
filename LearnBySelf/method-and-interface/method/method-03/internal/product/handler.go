package product

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

var products = make(map[int]Product)
var nextID = 1
var productsMu sync.Mutex

// Handler lấy danh sách sản phẩm với phân trang và lọc
func GetProducts(w http.ResponseWriter, r *http.Request) {
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
	paginatedProducts, total := PaginateAndFilter(products, page, limit, nameFilter)

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
func CreateProduct(w http.ResponseWriter, r *http.Request) {
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
	AppendData(newProduct)

	// Trả về sản phẩm đã tạo
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newProduct)
	fmt.Println("Response sent to client")
}

// Handler lấy thông tin chi tiết của một sản phẩm
func GetProduct(w http.ResponseWriter, r *http.Request) {
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
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
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

	// Cập nhật thông tin sản phẩm
	updatedProduct.ID = id
	products[id] = updatedProduct

	// Lưu dữ liệu vào file sau khi cập nhật
	go SaveData()

	// Trả về sản phẩm đã cập nhật
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedProduct)
}

// Hàm xóa sản phẩm và cập nhật lại file
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
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

	// Xóa sản phẩm khỏi bộ nhớ
	delete(products, id)

	// Cập nhật lại file sau khi xóa sản phẩm
	go SaveData()

	w.WriteHeader(http.StatusNoContent)
}
