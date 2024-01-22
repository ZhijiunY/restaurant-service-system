package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// var (
// 	ctx = context.Background()
// )

// type OrderItem struct {
// 	Name     string `json:"name"`
// 	Quantity string `json:"quantity"`
// }

// Index page
func GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Home website",
	})
}

// Menu page
func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.HTML(http.StatusOK, "menu.tmpl", gin.H{
			"title": "Menu website",
		})
	}
}
