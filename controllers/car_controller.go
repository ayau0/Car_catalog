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
	if err := database.DB.Preload("Brand").First(&car, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
		return
	}
	c.JSON(http.StatusOK, car)
}

// Создание нового автомобиля
func CreateCar(c *gin.Context) {
	var car models.Car
	// Пробуем связать данные из запроса с моделью автомобиля
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Логируем полученные данные
	log.Printf("Received car data: %+v", car)

	// Проверка существования пользователя через внешний сервис
	client := resty.New()
	resp, err := client.R().
		SetPathParams(map[string]string{
			"id": fmt.Sprintf("%d", car.UserID),
		}).
		Get("http://user-service:8081/users/{id}") // Внутренний адрес сервиса

	// Если ошибка запроса или неправильный статус ответа, возвращаем ошибку
	if err != nil {
		log.Printf("Error during user check: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "User service unavailable"})
		return
	}

	// Логируем ответ от внешнего сервиса
	log.Printf("Response from user service: %s", resp.String())

	// Если статус ответа не 200, значит пользователь не найден
	if resp.StatusCode() != http.StatusOK {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	// Создание автомобиля и добавление связи с брендом
	if err := database.DB.Preload("Brand").Create(&car).Error; err != nil {
		log.Printf("Error creating car: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create car"})
		return
	}

	// Логируем успешное создание автомобиля
	log.Printf("Car created successfully: %+v", car)
	c.JSON(http.StatusCreated, car)
}

// Обновить информацию о автомобиле
func UpdateCar(c *gin.Context) {
	id := c.Param("id")
	var car models.Car
	if err := database.DB.First(&car, "id = ?", id).Error; err != nil {
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
	if err := database.DB.Delete(&models.Car{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Car deleted successfully"})
}
