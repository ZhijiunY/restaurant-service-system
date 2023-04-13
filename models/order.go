package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID         uuid.UUID `gorm:"primaryKey" json:"id"`
	UserID     int       `gorm:"not null" json:"user_id"`
	TableID    int       `gorm:"not null" json:"table_id"`
	OrderDate  time.Time `gorm:"not null" json:"order_date"`
	TotalPrice float64   `gorm:"not null" json:"total_price"`
	User       User      `gorm:"foreignKey:UserID" json:"user"`
	Table      Table     `gorm:"foreignKey:TableID" json:"table"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
