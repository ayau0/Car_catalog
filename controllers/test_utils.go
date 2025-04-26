package controllers

import (
	"car-catalog-backend/database"
	"car-catalog-backend/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

// Инициализация тестовой базы данных
func InitTestDB() {
	database.Connect() // Подключаемся к базе данных

	// Очистка данных с учетом внешнего ключа
	database.DB.Exec("DELETE FROM cars")   // Сначала удаляем машины
	database.DB.Exec("DELETE FROM brands") // Затем удаляем бренды

	// Миграции
	database.DB.AutoMigrate(&models.Brand{}, &models.Car{})
}

func SeedTestData() {
	// Сначала добавляем бренды
	tesla := models.Brand{Name: "Tesla", Country: "USA", Description: "Electric cars", LogoURL: "url"}
	mercedes := models.Brand{Name: "Mercedes", Country: "Germany", Description: "Luxury cars", LogoURL: "url"}
	audi := models.Brand{Name: "Audi", Country: "Germany", Description: "Premium cars", LogoURL: "url"}

	// Добавляем бренды в базу данных
	database.DB.Create(&tesla)
	database.DB.Create(&mercedes)
	database.DB.Create(&audi)

	// Проверяем, что бренды были добавлены
	database.DB.First(&tesla, "name = ?", "Tesla")
	database.DB.First(&mercedes, "name = ?", "Mercedes")
	database.DB.First(&audi, "name = ?", "Audi")

	// Логирование для отладки
	fmt.Println("Added brands:", tesla, mercedes, audi)

	// Теперь добавляем машины, ссылаясь на правильный brand_id
	teslaCar := models.Car{BrandID: tesla.ID, Name: "Tesla Model X", Year: 2022, Price: 80000}
	mercedesCar := models.Car{BrandID: mercedes.ID, Name: "Mercedes-Benz S-Class", Year: 2021, Price: 90000}
	audiCar := models.Car{BrandID: audi.ID, Name: "Audi A6", Year: 2022, Price: 50000}

	// Создаем машины с правильными brand_id
	database.DB.Create(&teslaCar)
	database.DB.Create(&mercedesCar)
	database.DB.Create(&audiCar)

	// Логирование для отладки
	fmt.Println("Added cars:", teslaCar, mercedesCar, audiCar)
}

// Настроим тестовый роутер
func SetupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/brands", GetBrands)
	router.GET("/brands/:id", GetBrandByID)
	router.GET("/brands/cars", GetBrandWithCars)
	router.POST("/brands", CreateBrand)
	router.PUT("/brands/:id", UpdateBrand)
	router.DELETE("/brands/:id", DeleteBrand)
	router.GET("/cars", GetCars)
	router.GET("/cars/:id", GetCarByID)
	router.POST("/cars", CreateCar)
	router.PUT("/cars/:id", UpdateCar)
	router.DELETE("/cars/:id", DeleteCar)
	return router
}
