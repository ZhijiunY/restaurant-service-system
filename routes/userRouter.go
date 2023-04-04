package routes

import (
	"github.com/ZhijiunY/restaurant-service-system/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/index.tmpl", controllers.GetUsers)
	router.POST("/", controllers.CreateUsers)
	router.DELETE("/:id", controllers.DeleteUsers)
	router.PUT("/:id", controllers.UpdateUsers)
}
