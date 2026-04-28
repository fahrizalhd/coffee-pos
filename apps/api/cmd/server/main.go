package main

import (
	"log"

	"coffee-pos-api/internal/config"
	"coffee-pos-api/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	db := database.Connect(cfg)

	_ = db

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	log.Println("server running on port", cfg.AppPort)

	r.Run(":" + cfg.AppPort)
}
