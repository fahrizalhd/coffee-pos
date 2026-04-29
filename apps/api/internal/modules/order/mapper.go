package order

func ToOrderResponse(order Order) OrderResponse {
	var items []OrderItemResponse

	for _, item := range order.OrderItems {
		items = append(items, OrderItemResponse{
			ProductID:   item.ProductID,
			ProductName: item.Product.Name,
			Category:    item.Product.Category.Name,
			Qty:         item.Qty,
			Price:       item.Price,
			Subtotal:    item.Subtotal,
		})
	}

	return OrderResponse{
		ID:            order.ID,
		InvoiceNumber: order.InvoiceNumber,
		TotalAmount:   order.TotalAmount,
		Status:        string(order.Status),
		Cashier: CashierResponse{
			ID:   order.Cashier.ID,
			Name: order.Cashier.Name,
		},
		Items:     items,
		CreatedAt: order.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
