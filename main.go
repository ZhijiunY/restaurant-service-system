package main

import (
	"log"

	"github.com/ZhijiunY/restaurant-service-system/migrations"
	"github.com/ZhijiunY/restaurant-service-system/routes"
	"github.com/ZhijiunY/restaurant-service-system/utils"
	"github.com/go-redis/redis/v8"
	_ "github.com/joho/godotenv/autoload"

	// "github.com/kardianos/govendor/migrate"
	_ "github.com/lib/pq"
)

var redisClient *redis.Client

// func setupLogging() {
// 	f, _ := os.Create("gin.log")
// 	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
// }

// var foodCollection *mongo.Collection = utils.OpenCollection(utils.Client, "food")

func main() {
	// setupLogging()

	utils.ConnectToDb()                                      // connect to database
	migrations.Migrate()                                     // migrate migrations
	router := routes.InitRouter(utils.RedisClient, utils.DB) // Initialize Router
	log.Println("Server started on port 8080")
	router.Run(":8080")
}
