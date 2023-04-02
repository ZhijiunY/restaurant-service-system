package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {

	//this represents the customer
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	FirstName  string    `gorm:"type:varchar(255);not null" column:"first_name"`
	SecondName string    `gorm:"type:varchar(255);not null" column:"second_name"`
	Email      string    `gorm:"uniqueIndex" column:"email"`
	Order      []Order
	TableID    uuid.UUID `gorm:"column:table_id,omitempty"`
	CreatedAt  time.Time `gorm:"column:created_at,omitempty"`
	UpdatedAt  time.Time `gorm:"column:updated_at,omitempty"`
}
