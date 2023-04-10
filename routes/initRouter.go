package routes

import (
	"github.com/ZhijiunY/restaurant-service-system/controllers"
	"github.com/ZhijiunY/restaurant-service-system/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(middleware.Logger())
	router.Use(gin.Recovery())
	router.Use(gin.BasicAuth(gin.Accounts{"Simba": "1234"}))

	// connect to template
	// Static file
	router.LoadHTMLGlob("./templates/**/*")
	router.Static("/static", "./static")

	// Grouping routes
	MainRoutes := router.Group("/")
	{
		MainRoutes.GET("/", controllers.GetHome)
		MainRoutes.GET("/menu", controllers.GetMenu)
		MainRoutes.GET("/order", controllers.GetManager)

	}

	SessionRoutes := router.Group("/")
	{
		SessionRoutes.GET("/login", controllers.LoginPage)
		SessionRoutes.GET("/signup", controllers.SignupPage)
		// SessionRoutes.POST("login", controllers.Login)
		// SessionRoutes.POST("signup", controllers.Signup)
	}

	UserRoutes := router.Group("/users")
	{
		UserRoutes.GET("/", controllers.GetUser)
		UserRoutes.POST("/", controllers.CreateUsers)
		UserRoutes.PUT("/:id", controllers.UpdateUsers)
		UserRoutes.DELETE("/:id", controllers.DeleteUsers)
	}

	return router
}
