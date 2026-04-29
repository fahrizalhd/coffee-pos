package order

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup) {
	order := r.Group("/orders")

	order.POST("", Create)
	order.GET("", FindAll)
	order.GET("/:id", FindByID)
}
