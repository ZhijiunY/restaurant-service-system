package controllers

import (
	"net/http"

	"github.com/ZhijiunY/restaurant-service-system/middleware"
	"github.com/ZhijiunY/restaurant-service-system/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetLoginPage
func LoginPage(c *gin.Context) {
	// name := c.Param("name")
	c.HTML(
		http.StatusOK, "login.tmpl", gin.H{},
	)
}

// GetSignupPage
func SignupPage(c *gin.Context) {
	c.HTML(
		http.StatusOK, "signup.tmpl", gin.H{},
	)
}

// PostLogin
func Login(c *gin.Context) {
	name := c.Request.FormValue("name")
	password := c.Request.FormValue("password")

	if hasSession := middleware.HasSession(c); hasSession {
		c.String(200, "already Logged in")
		return
	}

	user := models.UserDetailByName(name)

	if err := middleware.Compare(user.Password, password); err != nil {
		c.String(401, "password mismatch")
		return
	}

	middleware.SaveAuthSession(c, user.ID)

	c.String(200, "login successful")
}

// PostLogout
func Logout(c *gin.Context) {
	if hasSession := middleware.HasSession(c); hasSession {
		c.String(401, "用戶未登入")
		return
	}
	middleware.ClearAuthSession(c)
	c.String(200, "退出成功")
}

// PostSignup
func Signup(c *gin.Context) {
	var user models.User
	user.Name = c.Request.FormValue("name")
	user.Email = c.Request.FormValue("email")

	if hasSession := middleware.HasSession(c); hasSession {
		c.String(200, "用户已登陆")
		return
	}

	if existUser := models.UserDetailByName(user.Name); existUser.ID != uuid.Nil {
		c.String(200, "用户名已存在")
		return
	}

	if c.Request.FormValue("password") != c.Request.FormValue("password_confirmation") {
		c.String(200, "密码不一致")
		return
	}

	if pwd, err := middleware.Encrypt(c.Request.FormValue("password")); err == nil {
		user.Password = pwd
	}

	models.AddUser(&user)

	middleware.SaveAuthSession(c, user.ID)

	c.String(200, "注册成功")
}

// func Signup(c *gin.Context) {
// 	email := c.PostForm("email")
// 	password := c.PostForm("password")
// 	confirmPassword := c.PostForm("confirm_password")
// }
