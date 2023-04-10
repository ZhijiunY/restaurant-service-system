package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `bson:"_id"`
	UserName *string   `json:"user_name" validate:"required,min=2,max=100"`
	Password *string   `json:"password" validate:"required,min=6"`
	Email    *string   `json:"email" validate:"email,required"`
}

//  Token         *string   `json:"token"`
// 	Refresh_Token *string   `json:"refresh_token"`
// 	Created_at    time.Time `json:"created_at"`
// 	Updated_at    time.Time `json:"updated_at"`

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
