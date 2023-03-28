package initializers

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

// ConnectToDB connects to the database
func ConnectToDB(config *Config) (*gorm.DB, error) {
	counts := 0

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)
	dsn := os.Getenv("DSN")

	for {
		// connection 和 err 來自調用尚不存在的開放資料庫
		connection, err := openDB(dsn)

		// 確認連接
		if err != nil {
			log.Println("postgres not yet ready...")
		} else {
			log.Print("connected to database!")
			return connection, nil
		}

		// 如果遇到錯誤，再試十次
		if counts > 10 {
			return nil, err
		}

		log.Print("Backing off for 1 seconds")
		time.Sleep(1 * time.Second)
		// 增加 counts++
		counts++
		continue
	}
}

// open database
func openDB(dsn string) (*gorm.DB, error) {
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
		return nil, err
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return nil, err
	}

	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected Successfully to the Database")
	return DB, nil
}
