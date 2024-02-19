package routes

import (
	"context"

	"github.com/ZhijiunY/restaurant-service-system/controllers"
	"github.com/ZhijiunY/restaurant-service-system/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

const (
	sessionName = "my-session"
	userkey     = "user"
)

var secret = []byte("secret")

func InitRouter(redisClient *redis.Client) *gin.Engine {
	router := gin.Default()

	router.Use(middleware.Logger())
	router.Use(gin.Recovery())

	// BasicAuth
	router.Use(gin.BasicAuth(gin.Accounts{"Simba": "1234"}))

	// session
	middleware.EnableCookieSession()
	store := cookie.NewStore(secret)
	sessionController := controllers.NewSessionController(store)
	orderController := controllers.NewOrderController(redisClient, context.Background())

	router.Use(sessions.Sessions(sessionName, store))
	router.Use(sessionController.LoadAndSave())

	// connect to template | Static file
	router.LoadHTMLGlob("./templates/**/*")
	router.Static("/static", "./static")

	// Grouping routes
	MainRoutes := router.Group("/")
	{ // 需要通過 middleware.AuthSessionMiddle() 才能進入後面的路由
		MainRoutes.GET("/", controllers.GetIndex)
		MainRoutes.GET("/menu", sessionController.AuthRequired(), controllers.GetMenu())
		MainRoutes.GET("/order", sessionController.AuthRequired(), orderController.GetOrder())
		MainRoutes.POST("/submit-order", sessionController.AuthRequired(), orderController.SubmitOrder())
		MainRoutes.GET("/show-orders", sessionController.AuthRequired(), orderController.ShowOrders())
	}

	// auth
	AuthRoutes := router.Group("/auth")
	{
		AuthRoutes.GET("/getlogin", sessionController.LoginGet())
		AuthRoutes.GET("/signup", sessionController.SignupGet())
		AuthRoutes.POST("/login", sessionController.LoginPost())
		AuthRoutes.POST("/logout", sessionController.LogoutPost())
		AuthRoutes.POST("/signup", sessionController.SignupPost())

		AuthRoutes.Static("/static", "./static")
	}

	return router
}
