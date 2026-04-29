package category

import (
	"coffee-pos-api/internal/database"
	"coffee-pos-api/internal/response"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var category Category

	if err := c.ShouldBindJSON(&category); err != nil {
		response.Error(c, 400, "Invalid request")
		return
	}

	if err := database.DB.Create(&category).Error; err != nil {
		response.Error(c, 500, "Failed to create category")
		return
	}

	response.Success(c, 201, "Category created", category)
}

func FindAll(c *gin.Context) {
	var categories []Category

	if err := database.DB.Find(&categories).Error; err != nil {
		response.Error(c, 500, "Failed to fetch categories")
		return
	}

	response.Success(c, 200, "Categories fetched", categories)
}

func FindByID(c *gin.Context) {
	id := c.Param("id")

	var category Category

	if err := database.DB.First(&category, id).Error; err != nil {
		response.Error(c, 404, "Category not found")
		return
	}

	response.Success(c, 200, "Category fetched", category)
}

func Update(c *gin.Context) {
	id := c.Param("id")

	var category Category

	if err := database.DB.First(&category, id).Error; err != nil {
		response.Error(c, 404, "Category not found")
		return
	}

	if err := c.ShouldBindJSON(&category); err != nil {
		response.Error(c, 400, "Invalid request")
		return
	}

	if err := database.DB.Save(&category).Error; err != nil {
		response.Error(c, 500, "Failed to update category")
		return
	}

	response.Success(c, 200, "Category updated", category)
}

func Delete(c *gin.Context) {
	id := c.Param("id")

	var category Category

	if err := database.DB.First(&category, id).Error; err != nil {
		response.Error(c, 404, "Category not found")
		return
	}

	if err := database.DB.Delete(&category).Error; err != nil {
		response.Error(c, 500, "Failed to delete category")
		return
	}

	response.Success(c, 200, "Category deleted", nil)
}
