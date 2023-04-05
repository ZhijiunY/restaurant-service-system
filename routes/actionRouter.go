package routes

import (
	"github.com/ZhijiunY/restaurant-service-system/controllers"
	"github.com/gin-gonic/gin"
)

func ActionRoutes(router *gin.Engine) {
	router.GET("/home", controllers.GetHome)
	router.GET("/about", controllers.GetAbout)
	router.GET("/menu", controllers.GetMenu)
	router.GET("/login", controllers.GetLogin)
	router.GET("/signup", controllers.GetSignup)
}
