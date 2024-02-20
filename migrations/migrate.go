package migrations

import (
	"fmt"

	"github.com/ZhijiunY/restaurant-service-system/models"
	"github.com/ZhijiunY/restaurant-service-system/utils"
)

func Migrate() {
	utils.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	utils.DB.AutoMigrate(
		&models.User{}, &models.Menu{}, &models.Table{},
		&models.Order{}, &models.OrderItem{},
	)

	fmt.Println("Migration complete")
}
