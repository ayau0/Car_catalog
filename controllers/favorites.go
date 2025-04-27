package controllers

import (
	"car-catalog-backend/database"
	"car-catalog-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Добавить в избранное
func AddToFavorites(c *gin.Context) {
	var favorite models.Favorite
	// Привязываем JSON запрос к структуре Favorite
	if err := c.ShouldBindJSON(&favorite); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Проверяем, существует ли пользователь
	var user models.User
	if err := database.DB.First(&user, favorite.UserID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	// Проверяем, существует ли автомобиль
	var car models.Car
	if err := database.DB.First(&car, favorite.CarID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Car not found"})
		return
	}

	// Проверим, добавлен ли уже этот автомобиль в избранное
	var existingFavorite models.Favorite
	if err := database.DB.Where("user_id = ? AND car_id = ?", favorite.UserID, favorite.CarID).First(&existingFavorite).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"message": "This car is already in favorites"})
		return
	}

	// Добавление в избранное
	if err := database.DB.Create(&favorite).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add to favorites"})
		return
	}

	// Возвращаем добавленный элемент в ответе
	c.JSON(http.StatusCreated, favorite)
}

// Получить все избранные элементы пользователя
func GetFavorites(c *gin.Context) {
	userID := c.Param("user_id")
	var favorites []models.Favorite

	// Получаем все избранные элементы для пользователя и подгружаем связанные автомобили
	if err := database.DB.Where("user_id = ?", userID).Preload("Car").Find(&favorites).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get favorites"})
		return
	}

	// Возвращаем избранные элементы
	c.JSON(http.StatusOK, favorites)
}

// Убрать из избранного
func RemoveFromFavorites(c *gin.Context) {
	id := c.Param("id")
	// Удаляем элемент из избранного
	if err := database.DB.Delete(&models.Favorite{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Favorite not found"})
		return
	}

	// Сообщение об успешном удалении
	c.JSON(http.StatusOK, gin.H{"message": "Removed from favorites successfully"})
}
