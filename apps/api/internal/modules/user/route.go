package user

import "github.com/gin-gonic/gin"

func RegisterPublicRoutes(r *gin.Engine) {
	users := r.Group("/users")

	users.POST("/login", Login)
}

func RegisterProtectedRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.GET("/profile", Profile)
}
