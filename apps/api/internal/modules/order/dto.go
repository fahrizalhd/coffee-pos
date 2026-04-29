package order

type CreateOrderRequest struct {
	Items []CreateOrderItemRequest `json:"items"`
}

type CreateOrderItemRequest struct {
	ProductID int64 `json:"product_id"`
	Qty       int   `json:"qty"`
}
