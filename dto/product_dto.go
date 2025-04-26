package dto

// CreateProductInput dùng để bind JSON khi tạo mới product
type CreateProductInput struct {
    Name          string  `json:"name" binding:"required"`
    Description   string  `json:"description"`
    Price         float64 `json:"price" binding:"required"`
    StockQuantity int     `json:"stock_quantity" binding:"required"`
    Status        string  `json:"status"`
	CategoryIDs    []uint  `json:"categories"` // Danh sách ID của các category
}

// UpdateProductInput dùng để bind JSON khi cập nhật product
type UpdateProductInput struct {
    Name          *string  `json:"name"`
    Description   *string  `json:"description"`
    Price         *float64 `json:"price"`
    StockQuantity *int     `json:"stock_quantity"`
    Status        *string  `json:"status"`
}