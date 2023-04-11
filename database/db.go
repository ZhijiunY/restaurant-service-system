package database

import (
	"github.com/ZhijiunY/restaurant-service-system/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:password@localhost:5432/restaurant_service"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db

	err = DB.AutoMigrate(&models.User{}, &models.Menu{}, &models.Table{}, &models.Order{}, &models.OrderItem{})
	if err != nil {
		panic("Failed to create tables!")
	}

	// DB.Migrator().CreateConstraint(&models.User{}, "Order")
	// DB.Migrator().CreateConstraint(&models.User{}, "fk_users_order")

	// DB.Migrator().CreateConstraint(&models.Table{}, "User")
	// DB.Migrator().CreateConstraint(&models.Table{}, "fk_table_users")

	// DB.Migrator().CreateConstraint(&models.Order{}, "OrderDetails")
	// DB.Migrator().CreateConstraint(&models.Order{}, "fk_order_orderDetails")

	// DB.Migrator().CreateConstraint(&models.Menus{}, "OrderDetails")
	// DB.Migrator().CreateConstraint(&models.Menus{}, "fk_menus_orderDetails")

}
