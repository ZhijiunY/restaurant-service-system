package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Home website",
	})
}

func GetMenu(c *gin.Context) {
	c.HTML(http.StatusOK, "menu.tmpl", gin.H{
		"title": "Menu website",
	})

}

func GetManager(c *gin.Context) {
	c.HTML(http.StatusOK, "order.tmpl", gin.H{
		"title": "Manager website",
	})
}
