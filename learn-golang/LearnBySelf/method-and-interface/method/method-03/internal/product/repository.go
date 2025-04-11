package product

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
)

var (
	fileMu   sync.Mutex
	dataFile = "data.json"
)

// Hàm tải dữ liệu từ file JSON
func LoadData() {
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
func AppendData(newProduct Product) {
	// Ghi vào file trong một goroutine riêng để tránh blocking
	go func() {
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
	}()
}

// Hàm lưu dữ liệu mới vào file (cập nhật toàn bộ sản phẩm)
func SaveData() {
	// Ghi vào file trong một goroutine riêng để tránh blocking
	go func() {
		fileMu.Lock()
		defer fileMu.Unlock()

		// Mở hoặc tạo file data.json
		file, err := os.Create(dataFile)
		if err != nil {
			log.Printf("Failed to create data file: %v", err)
			return
		}
		defer file.Close()

		// Lấy danh sách các sản phẩm
		var allProducts []Product
		productsMu.Lock()
		defer productsMu.Unlock()

		for _, product := range products {
			allProducts = append(allProducts, product)
		}

		// Ghi danh sách sản phẩm vào file
		encoder := json.NewEncoder(file)
		if err := encoder.Encode(allProducts); err != nil {
			log.Printf("Failed to encode data to file: %v", err)
		} else {
			fmt.Println("Data saved successfully.")
		}
	}()
}
