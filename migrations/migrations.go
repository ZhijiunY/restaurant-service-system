package migrations

import (
	"github.com/ZhijiunY/restaurant-service-system/cmd/web/initializers"
	"github.com/ZhijiunY/restaurant-service-system/internal/models"
	//"github.com/jinzhu/gorm"
)

// func Migrate() {
// 	initializers.DB.AutoMigrate(&models.User{}, &models.Menus{}, &models.Table{}, &models.Order{})
// 	initializers.DB.Model(&models.User{}).AddForeignKey("table_id", "table(id)", "CASCADE", "RESTRICT")
// }

func AutoMigrate() {
	// Auto-migrate all models
	initializers.DB.AutoMigrate(&models.User{}, &models.Menus{}, &models.Table{}, &models.Order{}, &models.OrderDetails{})
}

func AddForeignKey() {
	// Add foreign key constraint for User table
	initializers.DB.Model(&models.User{}).
		AddForeignKey("table_id", "tables(id)", "CASCADE", "RESTRICT")

	// Add foreign key constraint for Order table
	initializers.DB.Model(&models.Order{}).
		AddForeignKey("user_id", "users(id)", "CASCADE", "RESTRICT").
		AddForeignKey("table_id", "tables(id)", "CASCADE", "RESTRICT")

	// Add foreign key constraint for OrderDetails table
	initializers.DB.Model(&models.OrderDetails{}).
		AddForeignKey("order_id", "orders(id)", "CASCADE", "RESTRICT").
		AddForeignKey("menu_id", "menus(id)", "CASCADE", "RESTRICT")
}
