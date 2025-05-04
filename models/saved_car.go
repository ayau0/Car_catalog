package models

import "gorm.io/gorm"

type SavedCar struct {
	gorm.Model
	UserID uint `json:"user_id"` // ID пользователя из внешнего сервиса
	CarID  uint `json:"car_id"`
	Car    Car  `gorm:"foreignKey:CarID" json:"car"`
}
