package product

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup) {
	product := r.Group("/products")

	product.POST("", Create)
	product.GET("", FindAll)
	product.GET("/:id", FindByID)
	product.PUT("/:id", Update)
	product.DELETE("/:id", Delete)
}
