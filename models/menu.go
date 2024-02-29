package models

import (
	"time"

	"gorm.io/gorm"
)

// type Menu struct {
// 	gorm.Model
// 	ID          uuid.UUID `gorm:"primaryKey" json:"id"`
// 	FoodType    string    `gorm:"foodType" json:"foodType"`
// 	Name        string    `gorm:"not null" json:"name"`
// 	Description string    `gorm:"not null" json:"description"`
// 	Price       float64   `gorm:"not null" json:"price"`
// 	Count       int       `json:"count"`
// 	Created_at  time.Time `json:"created_at"`
// 	Updated_at  time.Time `json:"updated_at"`
// }

type Menu struct {
	gorm.Model
	// ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	FoodType    string    `gorm:"type:string" json:"foodType"`
	Name        string    `gorm:"type:string" json:"name"`
	Description string    `gorm:"type:string" json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
