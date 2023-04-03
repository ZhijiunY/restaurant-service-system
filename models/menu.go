package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Menus struct {
	gorm.Model
	ID         uuid.UUID  `bson:"_id"`
	Name       string     `json:"name" validate:"required"`
	Category   string     `json:"category" validate:"required"`
	Start_Date *time.Time `json:"start_date"`
	End_Date   *time.Time `json:"end_date"`
	Created_at time.Time  `json:"created_at"`
	Updated_at time.Time  `json:"updated_at"`
	Menu_id    string     `json:"food_id"`
}

// type Menus struct {
// 	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
// 	Name         string    `json:"name"`
// 	OrderDetails []OrderDetails
// 	Price        float64   `gorm:"omitempty" json:"price"`
// 	CreatedAt    time.Time `gorm:"omitempty" json:"created_at"`
// 	UpdatedAt    time.Time `gorm:"omitempty" json:"updated_at"`
// }

// type Menu struct {
// 	gorm.Model
// 	ID         uuid.UUID  `bson:"_id"`
// 	Name       string     `json:"name" validate:"required"`
// 	Category   string     `json:"category" validate:"required"`
// 	Start_Date *time.Time `json:"start_date"`
// 	End_Date   *time.Time `json:"end_date"`
// 	Created_at time.Time  `json:"created_at"`
// 	Updated_at time.Time  `json:"updated_at"`
// 	Menu_id    string     `json:"food_id"`
// }
