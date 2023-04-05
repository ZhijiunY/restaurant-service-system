package routes

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// connect to template
	router.LoadHTMLGlob("./static/templates/*")
	router.Static("/static", "./static")

	UserRoutes(router)
	ActionRoutes(router)

	// // set session middleware
	// store := cookie.NewStore([]byte("loginuser"))
	// router.Use(sessions.Sessions("mysession", store))

	// {
	// 	// register
	// 	router.GET("/signup.tmpl", controllers.RegisterGet)
	// 	router.POST("/signup.tmpl", controllers.RegisterPost)

	// 	// login
	// 	router.GET("/login", controllers.LoginGet)
	// 	router.POST("/login", controllers.LoginPost)

	// 	// home
	// 	router.GET("/home.tmpl", controllers.HomeGet)

	// 	// exit
	// 	router.GET("/exit", controllers.ExitGet)

	// }
	return router
}
