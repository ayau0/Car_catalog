package main

import (
	"car-catalog-backend/database"
	"car-catalog-backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
