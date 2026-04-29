package order

import (
	"coffee-pos-api/internal/database"
	"coffee-pos-api/internal/modules/product"
	"coffee-pos-api/internal/response"
	"coffee-pos-api/internal/utils"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var req CreateOrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "Invalid request")
		return
	}

	userID := c.GetInt64("user_id")

	var totalAmount float64

	order := Order{
		InvoiceNumber: utils.GenerateInvoiceNumber(),
		CashierID:     userID,
		Status:        OrderStatusPending,
	}

	if err := database.DB.Create(&order).Error; err != nil {
		response.Error(c, 500, "Failed to create order")
		return
	}

	for _, item := range req.Items {
		var productData product.Product

		if err := database.DB.
			First(&productData, item.ProductID).Error; err != nil {

			response.Error(c, 404, "Product not found")
			return
		}

		subtotal := productData.BasePrice * float64(item.Qty)

		orderItem := OrderItem{
			OrderID:   order.ID,
			ProductID: productData.ID,
			Qty:       item.Qty,
			Price:     productData.BasePrice,
			Subtotal:  subtotal,
		}

		if err := database.DB.Create(&orderItem).Error; err != nil {
			response.Error(c, 500, "Failed to create order item")
			return
		}

		totalAmount += subtotal
	}

	order.TotalAmount = totalAmount

	database.DB.Save(&order)

	database.DB.
		Preload("Cashier").
		Preload("OrderItems").
		Preload("OrderItems.Product").
		Preload("OrderItems.Product.Category").
		First(&order, order.ID)

	response.Success(
		c,
		201,
		"Order created",
		ToOrderResponse(order),
	)
}

func FindAll(c *gin.Context) {
	var orders []Order

	if err := database.DB.
		Preload("Cashier").
		Preload("OrderItems").
		Preload("OrderItems.Product").
		Find(&orders).Error; err != nil {

		response.Error(c, 500, "Failed to fetch orders")
		return
	}

	var responses []OrderResponse

	for _, order := range orders {
		responses = append(
			responses,
			ToOrderResponse(order),
		)
	}

	response.Success(c, 200, "Orders fetched", responses)
}

func FindByID(c *gin.Context) {
	id := c.Param("id")

	var order Order

	if err := database.DB.
		Preload("Cashier").
		Preload("OrderItems").
		Preload("OrderItems.Product").
		First(&order, id).Error; err != nil {

		response.Error(c, 404, "Order not found")
		return
	}

	response.Success(
		c,
		200,
		"Order fetched",
		ToOrderResponse(order),
	)
}
