package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	ID         uuid.UUID `bson:"_id"`
	Text       string    `json:"text"`
	Title      string    `json:"title"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Note_id    string    `json:"note_id"`
}
