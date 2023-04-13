package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         uuid.UUID `gorm:"primaryKey" json:"id"`
	Name       string    `gorm:"not null" json:"name"`
	Password   string    `gorm:"not null" json:"password"`
	Email      string    `gorm:"not null" json:"email"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
