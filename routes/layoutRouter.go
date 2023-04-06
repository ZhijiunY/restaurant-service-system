package routes

import (
	"github.com/ZhijiunY/restaurant-service-system/controllers"
	"github.com/gin-gonic/gin"
)

func LayoutRoutes(router *gin.Engine) {
	router.GET("/home", controllers.GetHome)
	router.GET("/menu", controllers.GetMenu)
	router.GET("/manager", controllers.GetManager)
}
