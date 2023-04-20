package controllers

import (
	"net/http"

	"github.com/ZhijiunY/restaurant-service-system/models"
	"github.com/gin-gonic/gin"
)

// Index page
func GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Home website",
	})
}

// Menu page
func GetMenu(c *gin.Context) {
	user := &models.User{Name: "simba"}
	c.HTML(http.StatusOK, "menu.tmpl", gin.H{
		"title": "Menu website",
		"Name":  user.Name,
	})
}

// Manager page
func GetManager(c *gin.Context) {
	c.HTML(http.StatusOK, "order.tmpl", gin.H{
		"title": "Manager website",
	})
}
