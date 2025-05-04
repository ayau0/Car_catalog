package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"net/http"
)

func CheckUserExistence(c *gin.Context) {
	userID := c.Param("id")

	client := resty.New()
	resp, err := client.R().
		Get("http://user-service:8081/users/" + userID)

	if err != nil || resp.StatusCode() != http.StatusOK {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found in user-service"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User exists in user-service"})
}
