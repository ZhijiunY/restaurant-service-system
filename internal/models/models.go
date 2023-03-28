package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	//this represents the customer
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	FirstName  string    `gorm:"type:varchar(255);not null"`
	SecondName string    `gorm:"type:varchar(255);not null"`
	Email      string    `gorm:"uniqueIndex;not null"`
	Order      []Order
	TableID    int `gorm:"column:table_id;not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Menus struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name         string    `gorm:"column:name"`
	OrderDetails []OrderDetails
	Price        float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Table struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	User        []User
	TableNumber int `gorm:"size:255"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Order struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	UserID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	OrderDetails []OrderDetails
	TableID      int `gorm:"column:table_id"`
	Quantity     int
	TotalPrice   float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type OrderDetails struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	OrderID    int       `gorm:"column:order_id"`
	MenuID     int       `gorm:"column:menu_id"`
	Quantity   int
	TotalPrice int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
