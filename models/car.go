package models

import "gorm.io/gorm"

// Car представляет данные об автомобиле
type Car struct {
	gorm.Model
	Name         string  `json:"name"`
	BrandID      uint    `json:"brand_id"`
	Year         int     `json:"year"`
	Price        float64 `json:"price"`
	BodyType     string  `json:"body_type"`
	EngineType   string  `json:"engine_type"`
	Horsepower   int     `json:"horsepower"`
	Transmission string  `json:"transmission"`
	ImageURL     string  `json:"image_url"`
	Brand        Brand   `gorm:"foreignKey:BrandID"` // Связь с брендом
}
