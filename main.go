package main

import (
	"log"
	"sync"
	"time"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"github.com/gin-gonic/gin"

	// "github.com/glebarez/sqlite"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}

	// setup database
	db := initDB()

	// create sessions to connect to redis
	session := initSession()

	// create waitGroup
	wg := sync.WaitGroup{}

	// // create loggers
	// infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// // set up the application config
	// app := Config{
	// 	Session:       session,
	// 	DB:            db,
	// 	InfoLog:       infoLog,
	// 	ErrorLog:      errorLog,
	// 	Wait:          &wg,
	// 	Models:        data.New(db),
	// 	ErrorChan:     make(chan error),
	// 	ErrorChanDone: make(chan bool),
	// }

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// router.Use(middleware.CorsMiddleware())

	// routes.PublicRoutes(router)

	// router.Use(middleware.Authorization())
	// routes.PrivateRoutes(router)

	log.Fatal(router.Run())
}

// ----------------------------------------------------------------
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

	dsn := "gorm.db"

	for {
		// connection 和 err 來自調用尚不存在的開放資料庫
		// connection, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
		connection, err := openDB(dsn)

		// 確認連接
		if err != nil {
			log.Println("postgres not yet ready...")
		} else {
			log.Print("connected to database!")
			return connection
		}

		// 如果遇到錯誤，再試十次
		if counts > 10 {
			return nil
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

// ----------------------------------------------------------------

// // Session
// func initSession() *scs.SessionManager {
// 	// gob.Register(data.User{})

// 	// set up session
// 	session := scs.New()

// 	session.Store = redisstore.New(initRedis())
// 	session.Lifetime = 24 * time.Hour
// 	session.Cookie.Persist = true
// 	session.Cookie.SameSite = http.SameSiteDefaultMode
// 	session.Cookie.Secure = true

// 	return session
// }

// // Redis
// func initRedis() *redis.Pool {
// 	redisPool := &redis.Pool{
// 		MaxIdle: 10, Dial: func() (redis.Conn, error) {
// 			return redis.Dial("tcp", os.Getenv("REDIS"))
// 		},
// 	}

// 	return redisPool
// }
