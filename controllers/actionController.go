package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHome(c *gin.Context) {
	c.HTML(http.StatusOK, "home.tmpl", gin.H{
		"title": "Home website",
	})
}

func GetLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", gin.H{
		"title": "Login website",
	})
}

func GetSignup(c *gin.Context) {

}

func GetAbout(c *gin.Context) {
	c.HTML(http.StatusOK, "about.tmpl", gin.H{
		"title": "About website",
	})

}

func GetMenu(c *gin.Context) {
	c.HTML(http.StatusOK, "menu.tmpl", gin.H{
		"title": "Menu website",
	})

}
