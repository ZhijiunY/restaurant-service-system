package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", gin.H{
		"title": "Login website",
	})
}

func Signup(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.tmpl", gin.H{
		"title": "Signup website",
	})
}
