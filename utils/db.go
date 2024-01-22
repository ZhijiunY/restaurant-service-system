package utils

import (
	"log"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	// redisCache *redis.Client
	redisClient *redis.Client
)

func Connect() {

	var err error

	// gorm, postgres initialization
	db, err := gorm.Open(postgres.Open("postgres://postgres:password@localhost:5432/restaurant_service"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to postgres")
	}
	DB = db

	// Redis initialization
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // 或者是你的 Redis 地址
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	log.Println("Connected to PostgreSQL database")

}
