package utils

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB          *gorm.DB
	RedisClient *redis.Client
)

func ConnectToDb() {

	var err error
	ctx := context.Background()

	// gorm, postgres initialization
	db, err := gorm.Open(postgres.Open("postgres://postgres:password@localhost:5432/restaurant_service"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to postgres")
	}
	DB = db

	// Redis initialization
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // 或者是你的 Redis 地址
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	// 檢查 Redis 連接
	_, err = RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	log.Println("Connected to PostgreSQL and Redis")

}
