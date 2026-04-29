package product

type ProductRequest struct {
	Name       string  `json:"name"`
	CategoryID int64   `json:"category_id"`
	BasePrice  float64 `json:"base_price"`
	IsActive   bool    `json:"is_active"`
}
