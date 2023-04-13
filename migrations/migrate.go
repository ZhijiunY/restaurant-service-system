package migrations

import (
	"fmt"

	"github.com/ZhijiunY/restaurant-service-system/models"
	"github.com/ZhijiunY/restaurant-service-system/utils"
)

func Migrate() {
	utils.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	utils.DB.AutoMigrate(&models.User{}, &models.Menu{}, &models.Table{}, &models.Order{}, &models.OrderItem{})
	fmt.Println("Migration complete")

	// DB.Migrator().CreateConstraint(&models.User{}, "Order")
	// DB.Migrator().CreateConstraint(&models.User{}, "fk_users_order")

	// DB.Migrator().CreateConstraint(&models.Table{}, "User")
	// DB.Migrator().CreateConstraint(&models.Table{}, "fk_table_users")

	// DB.Migrator().CreateConstraint(&models.Order{}, "OrderDetails")
	// DB.Migrator().CreateConstraint(&models.Order{}, "fk_order_orderDetails")

	// DB.Migrator().CreateConstraint(&models.Menus{}, "OrderDetails")
	// DB.Migrator().CreateConstraint(&models.Menus{}, "fk_menus_orderDetails")

	// fmt.Println("Foreign Keys created successfully")

}
