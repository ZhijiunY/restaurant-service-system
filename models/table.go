package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Table struct {
	gorm.Model
	ID              uuid.UUID `gorm:"primaryKey" json:"id"`
	SeatingCapacity uint      `gorm:"not null" json:"capacity"`
	Created_at      time.Time `json:"created_at"`
	Updated_at      time.Time `json:"updated_at"`
}
