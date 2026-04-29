package category

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup) {
	category := r.Group("/categories")

	category.POST("", Create)
	category.GET("", FindAll)
	category.GET("/:id", FindByID)
	category.PUT("/:id", Update)
	category.DELETE("/:id", Delete)
}
