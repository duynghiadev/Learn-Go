package product

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
