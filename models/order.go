package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID            string      `gorm:"primaryKey" json:"id"`
	OrderDate     time.Time   `gorm:"not null" json:"order_date"`
	TotalPrice    int64       `gorm:"not null" json:"total_price"`
	TotalQuantity int64       `gorm:"not null" json:"total_quantity"`
	OrderItems    []OrderItem `json:"order_items"`
	User          User        `gorm:"foreignKey:UserID" json:"user"`
	Table         Table       `gorm:"foreignKey:TableID" json:"table"`
	Created_at    time.Time   `json:"created_at"`
	Updated_at    time.Time   `json:"updated_at"`
}
