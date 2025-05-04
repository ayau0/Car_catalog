package controllers

import (
	"car-catalog-backend/database"
	"car-catalog-backend/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"net/http"
)

// POST /saved-cars
func CreateSavedCar(c *gin.Context) {
	var savedCar models.SavedCar
	if err := c.ShouldBindJSON(&savedCar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Проверка пользователя через внешний сервис
	client := resty.New()
	resp, err := client.R().
		SetPathParam("id", fmt.Sprintf("%d", savedCar.UserID)).
		Get("http://localhost:8081/users/{id}")
	if err != nil || resp.StatusCode() != http.StatusOK {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	// Проверка существования автомобиля
	var car models.Car
	if err := database.DB.First(&car, savedCar.CarID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Car not found"})
		return
	}

	// Создание сохранённой записи
	if err := database.DB.Create(&savedCar).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save car"})
		return
	}

	c.JSON(http.StatusCreated, savedCar)
}

// GET /saved-cars/:user_id
func GetSavedCarsByUser(c *gin.Context) {
	userID := c.Param("user_id")
	var savedCars []models.SavedCar
	if err := database.DB.Preload("Car.Brand").Where("user_id = ?", userID).Find(&savedCars).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch saved cars"})
		return
	}
	c.JSON(http.StatusOK, savedCars)
}
