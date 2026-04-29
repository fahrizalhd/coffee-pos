package product

import (
	"time"

	"coffee-pos-api/internal/modules/category"
)

type Product struct {
	ID         int64   `gorm:"primaryKey" json:"id"`
	Name       string  `json:"name"`
	CategoryID int64   `json:"category_id"`
	BasePrice  float64 `json:"base_price"`
	IsActive   bool    `json:"is_active"`

	Category category.Category `gorm:"foreignKey:CategoryID" json:"category"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
