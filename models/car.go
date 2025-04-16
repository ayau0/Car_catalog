package models

type Car struct {
	ID       uint    `json:"id" gorm:"primaryKey"`
	Name     string  `json:"name"`
	BrandID  uint    `json:"brand_id"`
	Brand    Brand   `json:"Brand"` // Вложенная структура бренда
	Year     int     `json:"year"`
	Price    float64 `json:"price"`
	ImageURL string  `json:"image_url"`
}
