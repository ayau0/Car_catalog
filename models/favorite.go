package models

type Favorite struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	UserID    uint   `json:"user_id"`    // Ссылка на пользователя
	CarID     uint   `json:"car_id"`     // Ссылка на автомобиль
	CreatedAt string `json:"created_at"` // Дата добавления
}
