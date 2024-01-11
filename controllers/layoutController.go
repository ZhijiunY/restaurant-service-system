package controllers

import (
	"net/http"

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
		// tmpl := template.Must(template.ParseFiles("layout.html"))
		// userName := middleware.GetUserSession(c)
		c.HTML(http.StatusOK, "menu.tmpl", gin.H{
			"title": "Menu website",
			// "Name":  userName,
		})
	}
}
