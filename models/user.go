package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID            uuid.UUID `bson:"_id"`
	First_name    *string   `json:"first_name" validate:"required,min=2,max=100"`
	Last_name     *string   `json:"last_name" validate:"required,min=2,max=100"`
	Password      *string   `json:"Password" validate:"required,min=6"`
	Email         *string   `json:"email" validate:"email,required"`
	Avatar        *string   `json:"avatar"`
	Phone         *string   `json:"phone" validate:"required"`
	Token         *string   `json:"token"`
	Refresh_Token *string   `json:"refresh_token"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
	User_id       string    `json:"user_id"`
}

// type User struct {
// 	gorm.Model
// 	//this represents the customer
// 	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
// 	FirstName  string    `gorm:"type:varchar(255);not null" json:"first_name"`
// 	SecondName string    `gorm:"type:varchar(255);not null" json:"second_name"`
// 	Email      string    `gorm:"uniqueIndex" json:"email"`
// 	Order      []Order
// 	TableID    uuid.UUID `json:"table_id,omitempty"`
// 	CreatedAt  time.Time `json:"table_created_at,omitempty"`
// 	UpdatedAt  time.Time `json:"table_updated_at,omitempty"`
// }
