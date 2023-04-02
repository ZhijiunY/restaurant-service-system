package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	UserID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	OrderDetails []OrderDetails
	TableID      uuid.UUID `gorm:"column:table_id,omitempty"`
	Quantity     int
	TotalPrice   float64
	CreatedAt    time.Time `gorm:"column:created_at,omitempty"`
	UpdatedAt    time.Time `gorm:"column:updated_at,omitempty"`
}
