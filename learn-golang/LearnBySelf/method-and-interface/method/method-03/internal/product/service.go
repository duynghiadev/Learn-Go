package product

import (
	"strings"
)

// Hàm phân trang và lọc
func PaginateAndFilter(products map[int]Product, page, limit int, nameFilter string) ([]Product, int) {
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
