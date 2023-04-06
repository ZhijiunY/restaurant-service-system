package routes

import (
	"github.com/ZhijiunY/restaurant-service-system/controllers"
	"github.com/gin-gonic/gin"
)

func SessionRoutes(router *gin.Engine) {
	router.GET("/login", controllers.Login)
	router.GET("/signup", controllers.Signup)
}
