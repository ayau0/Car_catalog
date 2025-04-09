package controllers

import (
	"car-catalog-backend/database"
	"car-catalog-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Получить все бренды
func GetBrands(c *gin.Context) {
	var brands []models.Brand
	if err := database.DB.Find(&brands).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, brands)
}

// Получить бренд по ID
func GetBrandByID(c *gin.Context) {
	var brand models.Brand
	id := c.Param("id")
	if err := database.DB.First(&brand, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Brand not found"})
		return
	}
	c.JSON(http.StatusOK, brand)
}

// Создать новый бренд
func CreateBrand(c *gin.Context) {
	var brand models.Brand
	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&brand).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, brand)
}

// Обновить бренд по ID
func UpdateBrand(c *gin.Context) {
	var brand models.Brand
	id := c.Param("id")
	if err := database.DB.First(&brand, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Brand not found"})
		return
	}

	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&brand).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, brand)
}

// Удалить бренд по ID
func DeleteBrand(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Brand{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Brand not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Brand deleted successfully"})
}

// Получить бренд по ID с машинами
func GetBrandWithCars(c *gin.Context) {
	var brand models.Brand
	id := c.Param("id")

	if err := database.DB.Preload("Cars").First(&brand, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Brand not found"})
		return
	}

	c.JSON(http.StatusOK, brand)
}
