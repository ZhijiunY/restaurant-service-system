package main

import (
	"log"
	"time"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"github.com/gin-gonic/gin"

	// "github.com/glebarez/sqlite"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// router.Use(middleware.CorsMiddleware())

	// routes.PublicRoutes(router)

	// router.Use(middleware.Authorization())
	// routes.PrivateRoutes(router)

	log.Fatal(router.Run())

	// setup database
	db := initDB()
}

// init database
func initDB() *gorm.DB {
	conn := connectToDB()
	if conn == nil {
		log.Panic("can't connect to database")
	}
	return conn
}

// connect to the database
func connectToDB() *gorm.DB {
	counts := 0

	// dsn := os.Getenv("DSN") // 設定環境變數，獲取dsn字串，來自os.Getenv調用環境變數
	dsn := "gorm.db"

	for {
		// connection 和 err 來自調用尚不存在的開放資料庫
		connection, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})

		// 確認連接
		if err != nil {
			log.Println("postgres not yet ready...")
		} else {
			log.Print("connected to database!")
			return connection
		}

		// 如果遇到錯誤，再適十次
		if counts > 10 {
			return nil
		}

		// 否則
		log.Print("Backing off for 1 seconds")
		time.Sleep(1 * time.Second)
		// 增加 counts++
		counts++
		continue
	}
}

// open database
func openDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
