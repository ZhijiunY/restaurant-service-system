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
		MainRoutes.GET("/", controllers.GetIndex)
		MainRoutes.GET("/menu", controllers.GetMenu)
		MainRoutes.GET("/order", controllers.GetManager)

	}

	AuthRoutes := router.Group("/")
	{
		AuthRoutes.GET("/login", controllers.LoginPage)
		AuthRoutes.GET("/signup", controllers.SignupPage)
		// SessionRoutes.POST("login", controllers.Login)
		// SessionRoutes.POST("signup", controllers.Signup)
	}

	UserRoutes := router.Group("/user")
	{
		UserRoutes.GET("/", controllers.GetUser)
		UserRoutes.POST("/", controllers.CreateUser)
		UserRoutes.PUT("/:id", controllers.UpdateUser)
		UserRoutes.DELETE("/:id", controllers.DeleteUser)
	}

	return router
}
