package routes

import (
	"car-catalog-backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Роуты для автомобилей
	r.GET("/cars", controllers.GetCars)
	r.GET("/cars/:id", controllers.GetCarByID)
	r.POST("/cars", controllers.CreateCar)
	r.PUT("/cars/:id", controllers.UpdateCar)
	r.DELETE("/cars/:id", controllers.DeleteCar)

	// Роуты для брендов
	r.GET("/brands", controllers.GetBrands)
	r.GET("/brands/:id", controllers.GetBrandByID)
	r.GET("/brands/:id/cars", controllers.GetBrandWithCars)
	r.POST("/brands", controllers.CreateBrand)
	r.PUT("/brands/:id", controllers.UpdateBrand)
	r.DELETE("/brands/:id", controllers.DeleteBrand)
}
