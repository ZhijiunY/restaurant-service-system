package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Table struct {
	gorm.Model
	ID               uuid.UUID `bson:"_id"`
	Number_of_guests *int      `json:"number_of_guests" validate:"required"`
	Table_number     *int      `json:"table_number" validate:"required"`
	Created_at       time.Time `json:"created_at"`
	Updated_at       time.Time `json:"updated_at"`
	Table_id         string    `json:"table_id"`
}

// type Table struct {
// 	gorm.Model
// 	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
// 	User        []User
// 	TableNumber int       `gorm:"size:255"`
// 	CreatedAt   time.Time `json:"created_at"`
// 	UpdatedAt   time.Time `json:"updated_at"`
// }
