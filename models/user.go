package models

// Модель для пользователя
type User struct {
	ID       uint   `json:"ID"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}
