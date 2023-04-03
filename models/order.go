package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID         uuid.UUID `bson:"_id"`
	Order_Date time.Time `json:"order_date" validate:"required"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Order_id   string    `json:"order_id"`
	Table_id   *string   `json:"table_id" validate:"required"`
}

// type Order struct {
// 	gorm.Model
// 	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
// 	UserID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
// 	OrderDetails []OrderDetails
// 	TableID      uuid.UUID `gorm:"omitempty" json:"table_id"`
// 	Quantity     int       `json:"quantity"`
// 	TotalPrice   float64   `Json:"total_price"`
// 	CreatedAt    time.Time `gorm:"omitempty" json:"created_at"`
// 	UpdatedAt    time.Time `gorm:"omitempty" json:"updated_at"`
// }
