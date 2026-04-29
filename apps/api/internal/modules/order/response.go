package order

type OrderResponse struct {
	ID            int64               `json:"id"`
	InvoiceNumber string              `json:"invoice_number"`
	TotalAmount   float64             `json:"total_amount"`
	Status        string              `json:"status"`
	Cashier       CashierResponse     `json:"cashier"`
	Items         []OrderItemResponse `json:"items"`
	CreatedAt     string              `json:"created_at"`
}

type CashierResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type OrderItemResponse struct {
	ProductID   int64   `json:"product_id"`
	ProductName string  `json:"product_name"`
	Category    string  `json:"category"`
	Qty         int     `json:"qty"`
	Price       float64 `json:"price"`
	Subtotal    float64 `json:"subtotal"`
}
