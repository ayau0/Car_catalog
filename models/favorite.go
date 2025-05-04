package models

type Favorite struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	UserID uint `json:"user_id"`
	CarID  uint `json:"car_id"`
	Car    Car  `json:"car" gorm:"foreignKey:CarID"`
}
