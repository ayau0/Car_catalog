package main

import (
	"car-catalog-backend/database"
	"car-catalog-backend/models"
	"car-catalog-backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Подключаемся к базе данных
	database.Connect()

	// Автоматическая миграция
	database.DB.AutoMigrate(&models.Car{}, &models.Brand{})

	// Запуск сервера
	r := gin.Default()
	routes.SetupRoutes(r) // Настройка маршрутов
	r.Run(":8080")        // Запуск на порту 8080
}
