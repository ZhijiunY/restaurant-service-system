package routes

import (
	"github.com/ZhijiunY/restaurant-service-system/controllers"
	"github.com/ZhijiunY/restaurant-service-system/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const (
	sessionName = "my-session"
	userkey     = "user"
)

var secret = []byte("secret")

func InitRouter() *gin.Engine {
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
	router.Use(sessionController.LoadAndSave())
	// sessionMiddleware := middleware.EnableCookieSession()
	// router.Use(sessionMiddleware)

	// connect to template
	// Static file
	router.LoadHTMLGlob("./templates/**/*")
	router.Static("/static", "./static")

	// Grouping routes
	MainRoutes := router.Group("/")
	{ // 需要通過 middleware.AuthSessionMiddle() 才能進入後面的路由
		MainRoutes.GET("/", controllers.GetIndex)
		MainRoutes.GET("/menu", controllers.NewSessionController(store).AuthRequired(), controllers.GetMenu)
		MainRoutes.GET("/order", controllers.GetManager)
	}

	AuthRoutes := router.Group("/user")
	{
		AuthRoutes.GET("/login", controllers.NewSessionController(store).LoginGet())
		AuthRoutes.GET("/signup", controllers.NewSessionController(store).SignupGet())
		AuthRoutes.POST("/login", controllers.NewSessionController(store).LoginPost())
		AuthRoutes.POST("/logout", controllers.NewSessionController(store).LogoutPost())
		AuthRoutes.POST("/signup", controllers.NewSessionController(store).SignupPost())

		AuthRoutes.Static("/static", "./static")
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

func NewSessionController(store cookie.Store) {
	panic("unimplemented")
}
