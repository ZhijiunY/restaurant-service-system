package migrations

import (
	"fmt"
	"time"

	"github.com/ZhijiunY/restaurant-service-system/models"
	"github.com/ZhijiunY/restaurant-service-system/utils"
	"github.com/google/uuid"
)

func Migrate() {
	utils.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	utils.DB.AutoMigrate(&models.User{}, &models.Menu{}, &models.Table{}, &models.Order{}, &models.OrderItem{})
	fmt.Println("Migration complete")

	// create menus
	menus := []models.Menu{
		{ID: uuid.New(), Name: "Hambuger", Description: "good", Price: 10, Created_at: time.Now(), Updated_at: time.Now()},
		{ID: uuid.New(), Name: "Apple", Description: "good", Price: 20, Created_at: time.Now(), Updated_at: time.Now()},
		{ID: uuid.New(), Name: "Banana", Description: "good", Price: 30, Created_at: time.Now(), Updated_at: time.Now()},
		{ID: uuid.New(), Name: "Pizza", Description: "good", Price: 40, Created_at: time.Now(), Updated_at: time.Now()},
		{ID: uuid.New(), Name: "Salads", Description: "good", Price: 50, Created_at: time.Now(), Updated_at: time.Now()},
		{ID: uuid.New(), Name: "Cake", Description: "good", Price: 60, Created_at: time.Now(), Updated_at: time.Now()},
	}

	// store database in postgres
	for _, menu := range menus {
		result := utils.DB.Create(&menu)
		if result.Error != nil {
			fmt.Println(result.Error)
		}
	}
}
