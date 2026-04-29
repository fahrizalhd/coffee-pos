package product

import (
	"coffee-pos-api/internal/database"
	"coffee-pos-api/internal/modules/category"
	"coffee-pos-api/internal/response"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var req ProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "Invalid request")
		return
	}

	var cat category.Category

	if err := database.DB.First(&cat, req.CategoryID).Error; err != nil {
		response.Error(c, 404, "Category not found")
		return
	}

	product := Product{
		Name:       req.Name,
		CategoryID: req.CategoryID,
		BasePrice:  req.BasePrice,
		IsActive:   req.IsActive,
	}

	if err := database.DB.Create(&product).Error; err != nil {
		response.Error(c, 500, "Failed to create product")
		return
	}

	database.DB.Preload("Category").First(&product, product.ID)

	response.Success(c, 201, "Product created", product)
}

func FindAll(c *gin.Context) {
	var products []Product

	if err := database.DB.
		Preload("Category").
		Find(&products).Error; err != nil {

		response.Error(c, 500, "Failed to fetch products")
		return
	}

	response.Success(c, 200, "Products fetched", products)
}

func FindByID(c *gin.Context) {
	id := c.Param("id")

	var product Product

	if err := database.DB.
		Preload("Category").
		First(&product, id).Error; err != nil {

		response.Error(c, 404, "Product not found")
		return
	}

	response.Success(c, 200, "Product fetched", product)
}

func Update(c *gin.Context) {
	id := c.Param("id")

	var req ProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "Invalid request")
		return
	}

	var product Product

	if err := database.DB.First(&product, id).Error; err != nil {
		response.Error(c, 404, "Product not found")
		return
	}

	var cat category.Category

	if err := database.DB.First(&cat, req.CategoryID).Error; err != nil {
		response.Error(c, 404, "Category not found")
		return
	}

	product.Name = req.Name
	product.CategoryID = req.CategoryID
	product.BasePrice = req.BasePrice
	product.IsActive = req.IsActive

	if err := database.DB.Save(&product).Error; err != nil {
		response.Error(c, 500, "Failed to update product")
		return
	}

	database.DB.Preload("Category").First(&product, product.ID)

	response.Success(c, 200, "Product updated", product)
}

func Delete(c *gin.Context) {
	id := c.Param("id")

	var product Product

	if err := database.DB.First(&product, id).Error; err != nil {
		response.Error(c, 404, "Product not found")
		return
	}

	if err := database.DB.Delete(&product).Error; err != nil {
		response.Error(c, 500, "Failed to delete product")
		return
	}

	response.Success(c, 200, "Product deleted", nil)
}
