package controllers

import (
	"car-catalog-backend/database"
	"car-catalog-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Получить все автомобили
func GetCars(c *gin.Context) {
	var cars []models.Car
	if err := database.DB.Preload("Brand").Find(&cars).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch cars"})
		return
	}
	c.JSON(http.StatusOK, cars)
}

// Получить автомобиль по ID
func GetCarByID(c *gin.Context) {
	id := c.Param("id")
	var car models.Car
	if err := database.DB.Preload("Brand").First(&car, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
		return
	}
	c.JSON(http.StatusOK, car)
}

// Создать новый автомобиль
func CreateCar(c *gin.Context) {
	var car models.Car
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&car).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create car"})
		return
	}

	// Получаем бренд после создания
	database.DB.Preload("Brand").First(&car, car.ID)

	c.JSON(http.StatusCreated, car)
}

// Обновить информацию о автомобиле
func UpdateCar(c *gin.Context) {
	id := c.Param("id")
	var car models.Car
	if err := database.DB.First(&car, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
		return
	}

	// Обновление данных
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&car).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update car"})
		return
	}

	// Подгружаем связанный бренд
	database.DB.Preload("Brand").First(&car, car.ID)

	c.JSON(http.StatusOK, car)
}

// Удалить автомобиль
func DeleteCar(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Car{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Car deleted successfully"})
}
