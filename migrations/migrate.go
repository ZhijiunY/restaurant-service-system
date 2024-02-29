package migrations

import (
	"fmt"

	"github.com/ZhijiunY/restaurant-service-system/models"
	"github.com/ZhijiunY/restaurant-service-system/utils"
)

func Migrate() {
	utils.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	utils.DB.AutoMigrate(
		&models.User{}, &models.Menu{},
		&models.Order{}, &models.OrderItem{},
	)

	// Create a new menu
	// // batch insert from `[]map[string]interface{}{}`
	// utils.DB.Model(&models.Menu{}).Create([]map[string]interface{}{
	// 	{"FoodType": "主食", "Name": "北京烤鴨", "Description": "**", "Price": 380},
	// 	{"FoodType": "主食", "Name": "意式千層麵", "Description": "*", "Price": 240},
	// 	{"FoodType": "主食", "Name": "日式壽司", "Description": "**", "Price": 185},
	// 	{"FoodType": "點心", "Name": "提拉米蘇", "Description": "**", "Price": 80},
	// 	{"FoodType": "點心", "Name": "馬卡龍", "Description": "*", "Price": 90},
	// 	{"FoodType": "點心", "Name": "芝士蛋糕", "Description": "**", "Price": 100},
	// 	{"FoodType": "飲料", "Name": "珍珠奶茶", "Description": "****", "Price": 50},
	// 	{"FoodType": "飲料", "Name": "抹茶拿鐵", "Description": "*****", "Price": 70},
	// 	{"FoodType": "飲料", "Name": "鮮榨果汁", "Description": "*****", "Price": 60},
	// })

	fmt.Println("Migration complete")
}

// utils.DB.AutoMigrate(
// 	&models.User{}, &models.Menu{}, &models.Table{},
// 	&models.Order{}, &models.OrderItem{},
// )
