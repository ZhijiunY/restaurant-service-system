package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID            string      `gorm:"primaryKey" json:"id"`
	UserID        int64       `gorm:"not null" json:"user_id"`
	TableID       int64       `gorm:"not null" json:"table_id"`
	OrderDate     time.Time   `gorm:"not null" json:"order_date"`
	TotalPrice    float64     `gorm:"not null" json:"total_price"`
	TotalQuantity int64       `gorm:"not null" json:"total_quantity"`
	OrderItems    []OrderItem `json:"order_item"`
	User          User        `gorm:"foreignKey:UserID" json:"user"`
	Table         Table       `gorm:"foreignKey:TableID" json:"table"`
	Created_at    time.Time   `json:"created_at"`
	Updated_at    time.Time   `json:"updated_at"`
}

// GetTotalCount
func (order *Order) GetTotalQuantity() int64 {
	var totalQuantity int64
	for _, v := range order.OrderItems {
		totalQuantity = totalQuantity + v.Quantity
	}
	return totalQuantity
}

// GetTotalPrice
func (order *Order) GetTotalPrice() float64 {
	var totalPrice float64

	for _, v := range order.OrderItems {
		totalPrice = totalPrice + (float64(v.Quantity) * v.Menu.Price)
	}
	return totalPrice
}
