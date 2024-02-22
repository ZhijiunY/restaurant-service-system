package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Index page
func GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Home website",
	})
}

// Menu page
func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userName := session.Get("Name")
		if userName == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User Name not found"})
			return
		}

		c.HTML(http.StatusOK, "menu.tmpl", gin.H{
			"title":    "Menu website",
			"userName": userName,
		})
	}
}
