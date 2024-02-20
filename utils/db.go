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
	connectToPostgres()
	connectToRedis()
}

func connectToPostgres() {
	var err error
	// PostgreSQL 連接配置
	dsn := "postgres://postgres:password@localhost:5432/restaurant_service"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	log.Println("Successfully connected to PostgreSQL")
}

func connectToRedis() {
	var err error
	ctx := context.Background()
	// Redis 連接配置
	redisOptions := &redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	}
	RedisClient = redis.NewClient(redisOptions)
	_, err = RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("Successfully connected to Redis")
}
