package models

import (
	"time"

	"github.com/google/uuid"
)

type OrderDetails struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	OrderID    int       `gorm:"column:order_id"`
	MenuID     uuid.UUID `gorm:"column:menu_id"`
	Quantity   int
	TotalPrice int       `gorm:"column:total_price"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}
