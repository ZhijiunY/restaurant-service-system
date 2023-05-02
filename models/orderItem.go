package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	ID       uuid.UUID `gorm:"primaryKey" json:"id"`
	OrderID  int64     `gorm:"not null" json:"order_id"`
	MenuID   int64     `gorm:"not null" json:"menu_id"`
	Quantity int64     `gorm:"not null" json:"quantity"`
	//Order      Order     `gorm:"foreignKey:OrderID" json:"order"`
	Menu       *Menu     `gorm:"foreignKey:MenuID" json:"menu"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

// GetNotePrice  金額小計
func (OrderItem *OrderItem) GetNotePrice() float64 {
	price := OrderItem.Menu.Price
	return float64(OrderItem.Quantity) * price
}
