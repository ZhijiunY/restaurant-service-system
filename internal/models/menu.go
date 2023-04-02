package models

import (
	"time"

	"github.com/google/uuid"
)

type Menus struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name         string    `gorm:"column:name"`
	OrderDetails []OrderDetails
	Price        float64   `gorm:"column:price,omitempty"`
	CreatedAt    time.Time `gorm:"column:created_at,omitempty"`
	UpdatedAt    time.Time `gorm:"column:updated_at,omitempty"`
}
