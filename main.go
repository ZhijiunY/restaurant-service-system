package main

import (
	"log"

	"github.com/ZhijiunY/restaurant-service-system/database"
	"github.com/ZhijiunY/restaurant-service-system/migrations"
	"github.com/ZhijiunY/restaurant-service-system/routes"
	_ "github.com/joho/godotenv/autoload"

	// "github.com/kardianos/govendor/migrate"
	_ "github.com/lib/pq"
)

// var db *gorm.DB

// func init() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	dbHost := os.Getenv("DB_HOST")
// 	dbPort := os.Getenv("DB_PORT")
// 	dbUser := os.Getenv("DB_USER")
// 	dbName := os.Getenv("DB_NAME")
// 	dbPassword := os.Getenv("DB_PASSWORD")

// 	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPassword)

// 	// var err error
// 	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Println("Connected to PostgreSQL database")
// }

// func setupLogging() {
// 	f, _ := os.Create("gin.log")
// 	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
// }

func main() {
	// connect to PostgreSQL database
	database.Connect()

	migrations.Migrate()

	// Initialize Router
	router := routes.InitRouter()

	log.Println("Server started!")
	router.Run(":8080")
}
