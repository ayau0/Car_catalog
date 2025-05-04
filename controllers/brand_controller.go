package controllers

import (
	"car-catalog-backend/database"
	"car-catalog-backend/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
func GetBrandWithCars(c *gin.Context) {
	var brand models.Brand

	// Получаем brand_id из query, а не из маршрута
	brandID := c.Query("brand_id")
	if brandID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "brand_id is required"})
		return
	}

	// Пагинация
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	offset := (pageInt - 1) * limitInt

	// Получаем бренд
	if err := database.DB.First(&brand, brandID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Brand not found"})
		return
	}

	// Логируем информацию о бренде
	fmt.Printf("Brand found: %+v\n", brand)

	// Получаем машины этого бренда с пагинацией
	var cars []models.Car
	if err := database.DB.
		Where("brand_id = ?", brandID).
		Limit(limitInt).
		Offset(offset).
		Preload("Brand").
		Find(&cars).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load cars"})
		return
	}

	// Логируем количество найденных машин
	fmt.Printf("Found %d cars for brand_id: %s\n", len(cars), brandID)

	// Ответ
	c.JSON(http.StatusOK, gin.H{
		"brand": brand,
		"page":  pageInt,
		"limit": limitInt,
		"cars":  cars,
	})
}
