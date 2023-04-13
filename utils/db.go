package utils

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	// redisCache *redis.Client
)

func Connect() {

	var err error

	db, err := gorm.Open(postgres.Open("postgres://postgres:password@localhost:5432/restaurant_service"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to postgres")
	}
	// redisCache = redis.NewClient(&redis.Options{
	// 	Addr: "localhost:6379",
	// })
	// _, err = redisCache.Ping().Result()
	// if err != nil {
	// 	panic("failed to connect Redis")
	// }

	DB = db

}
