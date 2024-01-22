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

// func NewSessionController(store cookie.Store) {
// 	panic("unimplemented")
// }

func InitRouter(redisClient *redis.Client) *gin.Engine {
	router := gin.New()

	router.Use(middleware.Logger())
	router.Use(gin.Recovery())

	// BasicAuth
	router.Use(gin.BasicAuth(gin.Accounts{"Simba": "1234"}))

	// session
	middleware.EnableCookieSession()
	store := cookie.NewStore(secret)
	router.Use(sessions.Sessions(sessionName, store))
	sessionController := controllers.NewSessionController(store)
	orderController := controllers.NewOrderController(redisClient, context.Background())
	router.Use(sessionController.LoadAndSave())

	// connect to template | Static file
	router.LoadHTMLGlob("./templates/**/*")
	router.Static("/static", "./static")

	// Grouping routes
	// need authentication
	MainRoutes := router.Group("/")
	{ // 需要通過 middleware.AuthSessionMiddle() 才能進入後面的路由
		MainRoutes.GET("/", controllers.GetIndex)
		MainRoutes.GET("/menu", controllers.NewSessionController(store).AuthRequired(), controllers.GetMenu())
		MainRoutes.GET("/order", controllers.NewSessionController(store).AuthRequired(), orderController.GetOrder())
		MainRoutes.POST("/submit-order", controllers.NewSessionController(store).AuthRequired(), orderController.SubmitOrder())
	}

	// auth
	AuthRoutes := router.Group("/auth")
	{
		AuthRoutes.GET("/getlogin", controllers.NewSessionController(store).LoginGet())
		AuthRoutes.GET("/signup", controllers.NewSessionController(store).SignupGet())
		AuthRoutes.POST("/login", controllers.NewSessionController(store).LoginPost())
		AuthRoutes.POST("/logout", controllers.NewSessionController(store).LogoutPost())
		AuthRoutes.POST("/signup", controllers.NewSessionController(store).SignupPost())

		AuthRoutes.Static("/static", "./static")
	}

	// user
	// UserRoutes := router.Group("/user")
	{
		// 	// UserRoutes.GET("/", controllers.GetUser)
		// 	// UserRoutes.POST("/", controllers.CreateUser)
		// 	// UserRoutes.PUT("/:id", controllers.UpdateUser)
		// 	// UserRoutes.DELETE("/:id", controllers.DeleteUser)

		// UserRoutes.GET("/users", controllers.NewSessionController(store).GetUsers())
		// UserRoutes.GET("/users/:user_id", controllers.NewSessionController(store).GetUser())
		// UserRoutes.POST("/users/signup", controllers.NewSessionController(store).SignUp())
		// UserRoutes.POST("/users/login", controllers.NewSessionController(store).Login())
	}

	return router
}
