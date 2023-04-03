package main

import (
	"fmt"

	"github.com/ZhijiunY/restaurant-service-system/database"
	"github.com/ZhijiunY/restaurant-service-system/models"
)

// func init() {
// 	config, err := initializers.LoadConfig(".")
// 	if err != nil {
// 		log.Fatal("Could not load environment variables", err)
// 	}

// 	initializers.ConnectDB(&config)
// }

func Migrate() {
	database.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	database.DB.AutoMigrate(&models.User{}, &models.Menus{}, &models.Table{}, &models.Order{}, &models.OrderDetails{})
	fmt.Println("Migration complete")

	// initializers.DB.Model(&models.User{}).
	// 	AddForeignKey("table_id", "tables(id)", "CASCADE", "RESTRICT")
	// fmt.Println("cann't add User{} foreign key")

	// initializers.DB.Model(&models.Order{}).
	// 	AddForeignKey("user_id", "users(id)", "CASCADE", "RESTRICT").
	// 	AddForeignKey("table_id", "tables(id)", "CASCADE", "RESTRICT")
	// fmt.Println("cann't add Order{} foreign key")

	// initializers.DB.Model(&models.OrderDetails{}).
	// 	AddForeignKey("order_id", "orders(id)", "CASCADE", "RESTRICT").
	// 	AddForeignKey("menu_id", "menus(id)", "CASCADE", "RESTRICT")
	// fmt.Println("cann't add OrderDetails{} foreign key")

	// fmt.Println("Foreign Keys created successfully")

}
