package models

import "gorm.io/gorm"

// Модель автомобиля
type Car struct {
	gorm.Model
	Name    string  `json:"name"`
	Year    int     `json:"year"`
	Price   float64 `json:"price"`
	BrandID uint    `json:"brand_id"`
	Brand   Brand   `gorm:"foreignKey:BrandID" json:"brand"`

	// Добавление связи с пользователем
	UserID uint `json:"user_id"`
}
