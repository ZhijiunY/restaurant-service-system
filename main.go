package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

	db := initDB()
}

func initDB() *sql.DB {
	conn := connectToDB()
	if conn == nil {
		log.Panic("can't connect to database")
	}
	return conn
}

func connectToDB() *sql.DB {
	counts := 0

	dsn := os.Getenv("DSN") // 設定環境變數，獲取dsn字串，來自os.Getenv調用環境變數

	for {
		// connection 和 err 來自調用尚不存在的開放資料庫
		connection, err := openDB(dsn)

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
