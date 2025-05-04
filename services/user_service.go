package services

import (
	"github.com/go-resty/resty/v2"
	"log"
)

var client = resty.New()

func GetUserData(userID string) ([]byte, error) {
	client := resty.New()
	resp, err := client.R().
		SetPathParams(map[string]string{
			"id": string(savedCar.UserID),
		}).
		Get("http://localhost:8081/users/{id}")

	if err != nil {
		log.Println("Error calling user-service:", err)
		return nil, err
	}
	return resp.Body(), nil
}
