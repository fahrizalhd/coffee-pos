package order

import (
	"time"

	"coffee-pos-api/internal/modules/product"
	"coffee-pos-api/internal/modules/user"
)

type Order struct {
	ID            int64       `gorm:"primaryKey" json:"id"`
	InvoiceNumber string      `json:"invoice_number"`
	CashierID     int64       `json:"cashier_id"`
	TotalAmount   float64     `json:"total_amount"`
	Status        OrderStatus `json:"status"`

	Cashier user.User `gorm:"foreignKey:CashierID" json:"cashier"`

	OrderItems []OrderItem `gorm:"foreignKey:OrderID" json:"order_items"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrderItem struct {
	ID        int64 `gorm:"primaryKey" json:"id"`
	OrderID   int64 `json:"order_id"`
	ProductID int64 `json:"product_id"`

	Qty      int     `json:"qty"`
	Price    float64 `json:"price"`
	Subtotal float64 `json:"subtotal"`

	Product product.Product `gorm:"foreignKey:ProductID" json:"product"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
