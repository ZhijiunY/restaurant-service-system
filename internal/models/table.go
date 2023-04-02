package models

import (
	"time"

	"github.com/google/uuid"
)

type Table struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	User        []User
	TableNumber int       `gorm:"size:255"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}
