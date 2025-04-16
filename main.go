package main

import (
	"car-catalog-backend/database"
	"car-catalog-backend/models"
	"car-catalog-backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	database.DB.AutoMigrate(&models.Car{}, &models.Brand{}, &models.User{})

	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
