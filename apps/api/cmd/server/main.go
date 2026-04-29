package main

import (
	"log"

	"coffee-pos-api/internal/config"
	"coffee-pos-api/internal/database"
	"coffee-pos-api/internal/database/seeders"
	"coffee-pos-api/internal/middleware"
	"coffee-pos-api/internal/modules/category"
	"coffee-pos-api/internal/modules/order"
	"coffee-pos-api/internal/modules/product"
	"coffee-pos-api/internal/modules/user"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	db := database.Connect(cfg)

	if err := seeders.SeedAdmin(db); err != nil {
		log.Fatal("Failed to seed admin")
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// PUBLIC ROUTES
	user.RegisterPublicRoutes(r)

	// PROTECTED ROUTES
	auth := r.Group("/")

	auth.Use(middleware.AuthMiddleware(cfg))

	user.RegisterProtectedRoutes(auth)
	category.RegisterRoutes(auth)
	product.RegisterRoutes(auth)
	order.RegisterRoutes(auth)

	log.Println("Server running on port", cfg.AppPort)

	r.Run(":" + cfg.AppPort)
}
