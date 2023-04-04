package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	ID               uuid.UUID `bson:"_id" primaryKey:"_id"`
	Order_id         string    `json:"order_id"`
	Payment_method   *string   `json:"payment_method" validate:"eq=CARD|eq=CASH|eq="`
	Payment_status   *string   `json:"payment_status" validate:"required,eq=PENDING|eq=PAID"`
	Payment_due_date time.Time `json:"Payment_due_date"`
	Created_at       time.Time `json:"created_at"`
	Updated_at       time.Time `json:"updated_at"`
}
