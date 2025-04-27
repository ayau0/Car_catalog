package models

type Brand struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Country     string `json:"country"`
	Description string `json:"description"`
	LogoURL     string `json:"logo_url"`
	Cars        []Car  `json:"cars,omitempty"` // Связанные машины
}
