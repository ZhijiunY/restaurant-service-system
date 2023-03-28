package main

import (
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"github.com/ZhijiunY/restaurant-service-system/cmd/web/initializers"
	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	initializers.ConnectToDB(&config)
}

func main() {
	// // create sessions to connect to redis
	// session := initSession()

	// // create waitGroup
	// wg := sync.WaitGroup{}

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

	log.Fatal(router.Run())
}

// ----------------------------------------------------------------

// Session
func initSession() *scs.SessionManager {
	// gob.Register(data.User{})

	// set up session
	session := scs.New()

	session.Store = redisstore.New(initRedis())
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteDefaultMode
	session.Cookie.Secure = true

	return session
}

// Redis
func initRedis() *redis.Pool {
	redisPool := &redis.Pool{
		MaxIdle: 10, Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", os.Getenv("REDIS"))
		},
	}

	return redisPool
}
