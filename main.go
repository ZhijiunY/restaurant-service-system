package main

import (
	"io"
	"log"
	"os"

	"github.com/ZhijiunY/restaurant-service-system/migrations"
	"github.com/ZhijiunY/restaurant-service-system/routes"
	"github.com/ZhijiunY/restaurant-service-system/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	_ "github.com/joho/godotenv/autoload"

	// "github.com/kardianos/govendor/migrate"
	_ "github.com/lib/pq"
)

var redisClient *redis.Client

func setupLogging() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

// var foodCollection *mongo.Collection = utils.OpenCollection(utils.Client, "food")

func main() {
	setupLogging()

	// connect to PostgreSQL database
	utils.Connect()

	// Migrations
	migrations.Migrate()

	// Initialize Router
	router := routes.InitRouter(redisClient)
	log.Println("Server started!")
	router.Run(":8080")
}
