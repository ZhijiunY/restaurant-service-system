package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderDetails struct {
	gorm.Model
	ID            uuid.UUID `bson:"_id"`
	Quantity      *string   `json:"quantity" validate:"required,eq=S|eq=M|eq=L"`
	Unit_price    *float64  `json:"unit_price" validate:"required"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
	Food_id       *string   `json:"food_id" validate:"required"`
	Order_item_id string    `json:"order_item_id"`
	Order_id      string    `json:"order_id" validate:"required"`
}

// type OrderDetails struct {
// 	gorm.Model
// 	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
// 	OrderID    int       `json:"order_id"`
// 	MenuID     uuid.UUID `json:"menu_id"`
// 	Quantity   int
// 	TotalPrice int       `json:"total_price"`
// 	CreatedAt  time.Time `json:"created_at"`
// 	UpdatedAt  time.Time `json:"updated_at"`
// }
