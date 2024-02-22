package routes

import (
	"context"

	"github.com/ZhijiunY/restaurant-service-system/controllers"
	"github.com/ZhijiunY/restaurant-service-system/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/postgres"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

const (
	sessionName = "my-session"
	secretKey   = "secret"
)

func InitRouter(redisClient *redis.Client, db *gorm.DB) *gin.Engine {
	router := gin.Default()

	configureMiddlewares(router, db)
	configureRoutes(router, redisClient, db)

	return router
}

func configureMiddlewares(router *gin.Engine, db *gorm.DB) {
	router.Use(
		middleware.Logger(), gin.Recovery(),
		gin.BasicAuth(gin.Accounts{"Simba": "1234"}),
	)

	// 配置會話存儲
	store := configureSessionStore(db)
	router.Use(sessions.Sessions(sessionName, store))
}

func configureSessionStore(db *gorm.DB) sessions.Store {
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get sql.DB from gorm.DB: " + err.Error())
	}

	store, err := postgres.NewStore(sqlDB, []byte("secret"))
	if err != nil {
		panic("failed to create session store: " + err.Error())
	}
	return store
}

func configureRoutes(router *gin.Engine, redisClient *redis.Client, db *gorm.DB) {
	store := configureSessionStore(db)
	sessionController := controllers.NewSessionController(store)
	orderController := controllers.NewOrderController(redisClient, context.Background())

	router.LoadHTMLGlob("./templates/**/*")
	router.Static("/static", "./static")

	setupMainRoutes(router, sessionController)
	setupAuthRoutes(router, sessionController)
	setupOrderRoutes(router, sessionController, orderController)
}

func setupMainRoutes(router *gin.Engine, sessionController *controllers.SessionController) {
	router.GET("/", controllers.GetIndex)
	router.GET("/menu", sessionController.AuthRequired(), controllers.GetMenu())
}

func setupOrderRoutes(router *gin.Engine, sessionController *controllers.SessionController, orderController *controllers.OrderController) {
	orderRoutes := router.Group("/order")
	orderRoutes.GET("/", orderController.GetOrder())
	orderRoutes.POST("/submit-order", sessionController.AuthRequired(), orderController.SubmitOrder())
	orderRoutes.GET("/show-orders", sessionController.AuthRequired(), orderController.ShowOrders())
	orderRoutes.GET("/generate-qr", sessionController.AuthRequired(), orderController.GenerateOrderQRCode())
	orderRoutes.Static("/static", "./static")
}

func setupAuthRoutes(router *gin.Engine, sessionController *controllers.SessionController) {
	authRoutes := router.Group("/auth")
	authRoutes.GET("/getlogin", sessionController.LoginGet())
	authRoutes.GET("/signup", sessionController.SignupGet())
	authRoutes.POST("/login", sessionController.LoginPost())
	authRoutes.POST("/logout", sessionController.LogoutPost())
	authRoutes.POST("/signup", sessionController.SignupPost())
	authRoutes.Static("/static", "./static")
}
