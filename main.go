package main

import (
	"car-catalog-backend/database"
	"car-catalog-backend/middleware" // подключаем middleware
	"car-catalog-backend/routes"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	// Подключаемся к базе данных
	database.Connect()

	// Инициализация маршрутов
	r := gin.Default()

	// Использование middleware
	r.Use(middleware.AuthRequired()) // Используем созданную middleware для авторизации

	routes.SetupRoutes(r)

	// Запуск сервера
	if err := r.Run(":8080"); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s\n", err)
	}
}
