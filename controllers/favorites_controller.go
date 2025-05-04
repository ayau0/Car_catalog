package controllers

import (
	"car-catalog-backend/database"
	"car-catalog-backend/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"log"
	"net/http"
)

func AddToFavorites(c *gin.Context) {
	var favorite models.Favorite
	if err := c.ShouldBindJSON(&favorite); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Проверка пользователя через Resty
	client := resty.New()
	resp, err := client.R().Get(fmt.Sprintf("http://localhost:8081/users/%d", favorite.UserID))
	if err != nil || resp.StatusCode() != http.StatusOK {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	// Проверка машины
	var car models.Car
	if err := database.DB.First(&car, favorite.CarID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Car not found"})
		return
	}

	// Сохранение в базу данных
	if err := database.DB.Create(&favorite).Error; err != nil {
		log.Printf("Error saving favorite: %v\n", err) // Дополнительное логирование ошибки
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not add to favorites"})
		return
	}

	c.JSON(http.StatusCreated, favorite)
}

// Получение списка избранных автомобилей пользователя
func GetFavorites(c *gin.Context) {
	userID := c.DefaultQuery("user_id", "") // Получаем user_id из query параметра

	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	var favorites []models.Favorite
	if err := database.DB.Where("user_id = ?", userID).Find(&favorites).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch favorites"})
		return
	}

	c.JSON(http.StatusOK, favorites)
}

// Удалить из избранного
func RemoveFromFavorites(c *gin.Context) {
	id := c.Param("id") // Получаем id из параметра

	// Удаляем элемент из избранного
	if err := database.DB.Delete(&models.Favorite{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Favorite not found"})
		return
	}

	// Сообщение об успешном удалении
	c.JSON(http.StatusOK, gin.H{"message": "Removed from favorites successfully"})
}
