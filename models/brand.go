package models

import "gorm.io/gorm"

// Brand представляет данные о бренде автомобиля
type Brand struct {
	gorm.Model
	Name        string `json:"name"`
	Country     string `json:"country"`
	Description string `json:"description"`
	LogoURL     string `json:"logo_url"`
	Cars        []Car  `gorm:"foreignKey:BrandID"`
}
