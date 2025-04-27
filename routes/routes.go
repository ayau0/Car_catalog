package routes

import (
	"car-catalog-backend/controllers"
	"car-catalog-backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Роуты для авторизации
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Защищенные роуты
	authorized := r.Group("/")
	authorized.Use(middleware.AuthRequired()) // Middleware для проверки JWT
	{
		// Роуты для автомобилей
		authorized.GET("/cars", controllers.GetCars)
		authorized.GET("/cars/:id", controllers.GetCarByID)
		authorized.POST("/cars", controllers.CreateCar)
		authorized.PUT("/cars/:id", controllers.UpdateCar)
		authorized.DELETE("/cars/:id", controllers.DeleteCar)

		// Роуты для брендов
		authorized.GET("/brands", controllers.GetBrands)
		authorized.GET("/brands/:id", controllers.GetBrandByID)
		authorized.GET("/brands/cars", controllers.GetBrandWithCars)
		authorized.POST("/brands", controllers.CreateBrand)
		authorized.PUT("/brands/:id", controllers.UpdateBrand)
		authorized.DELETE("/brands/:id", controllers.DeleteBrand)

		// Роуты для категорий
		authorized.GET("/categories", controllers.GetCategories)
		authorized.POST("/categories", controllers.CreateCategory)
		authorized.PUT("/categories/:id", controllers.UpdateCategory)
		authorized.DELETE("/categories/:id", controllers.DeleteCategory)

		// Роуты для избранного
		authorized.POST("/favorites", controllers.AddToFavorites)
		authorized.GET("/favorites/:user_id", controllers.GetFavorites)
		authorized.DELETE("/favorites/:id", controllers.RemoveFromFavorites)
	}
}
