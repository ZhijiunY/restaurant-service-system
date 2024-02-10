package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	ID          uuid.UUID `gorm:"primaryKey" json:"id"`
	Quantity    int       `json:"quantity" gorm:"not null"  form:"quantity"`
	Name        string    `json:"name" gorm:"type:varchar(100);not null" form:"name"`
	Description string    `gorm:"type:text" form:"description"`
	Price       float64   `json:"price" gorm:"type:decimal(10,2);not null" form:"price"`
	Menu        *Menu     `gorm:"foreignKey:MenuID" json:"menu"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}
