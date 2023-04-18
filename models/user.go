package models

import (
	"time"

	"github.com/ZhijiunY/restaurant-service-system/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         uuid.UUID `gorm:"primaryKey" json:"id"`
	Name       string    `gorm:"not null" json:"name"`
	Password   string    `gorm:"not null" json:"password"`
	Email      string    `gorm:"not null" json:"email"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

func AddUser(user *User) {
	utils.DB.Create(&user)
}

func UserDetailByName(name string) (user User) {
	utils.DB.Where("name = ?", name).First(&user)
	return
}

func UserDetailByEmail(email string) (user User) {
	utils.DB.Where("email = ?", email).First(&user)
	return
}

func UserDetail(id uuid.UUID) (user User) {
	utils.DB.Where("id = ?", id).First(&user)
	return
}
