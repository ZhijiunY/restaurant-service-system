package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const userkey = "user"

type SessionController struct {
	store sessions.Store
}

func NewSessionController(store sessions.Store) *SessionController {
	return &SessionController{store}
}

func (sc *SessionController) LoadAndSave() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		defer session.Save()

		c.Set("session", session)
		c.Next()
	}
}

// 驗證是否已登入
func (sc *SessionController) AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(userkey)
		if user == nil {
			c.Redirect(http.StatusMovedPermanently, "/user/login")
			// c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			// 	"message:": "need to login!!!!!!!",
			// })
			return
		}
		c.Next()
	}
}

func (sc *SessionController) SignupGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(userkey)
		if user != nil {
			c.HTML(http.StatusBadRequest, "sighup.tmpl",
				gin.H{
					"content": "Please sighup first",
					"user":    user,
				})
			return
		}
		c.HTML(http.StatusOK, "signup.tmpl", gin.H{
			"content": "",
			"user":    user,
		})
	}

	// return func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "signup.tmpl", gin.H{
	// 		"content": "",
	// 		"user":    nil,
	// 	})
	// }
}

func (sc *SessionController) SignupPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		email := c.PostForm("email")
		password := c.PostForm("password")

		// Validate email and password
		if email == "" || password == "" {
			c.HTML(http.StatusBadRequest, "signup.tmpl", gin.H{
				"content": "Email and password cannot be empty",
				"user":    nil,
			})
			fmt.Println("validate error")
			return
		}

		// Store user in session
		session.Set(userkey, email)
		err := session.Save()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.HTML(http.StatusOK, "signup.tmpl", gin.H{
			"content": "Sign up success",
			"user":    email,
		})
	}
}

func (sc *SessionController) LoginGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(userkey)
		if user != nil {
			c.HTML(http.StatusBadRequest, "login.tmpl",
				gin.H{
					"content": "Please logout first",
					"user":    user,
				})
			return
		}
		c.HTML(http.StatusOK, "login.tmpl", gin.H{
			"content": "",
			"user":    user,
		})
	}
}

func (sc *SessionController) LoginPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(userkey)
		if user != nil {
			c.HTML(http.StatusBadRequest, "login.tmpl", gin.H{"content": "Please logout first"})
			return
		}

		// username := c.PostForm("username")
		// password := c.PostForm("password")

		// if EmptyUserPass(username, password) {
		// 	c.HTML(http.StatusBadRequest, "login.html", gin.H{"content": "Parameters can't be empty"})
		// 	return
		// }

		// if !CheckUserPass(username, password) {
		// 	c.HTML(http.StatusUnauthorized, "login.html", gin.H{"content": "Incorrect username or password"})
		// 	return
		// }

		// SaveSession(c, username)
		// c.Redirect(http.StatusMovedPermanently, "/dashboard")
	}
}

func (sc *SessionController) LogoutGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(userkey)
		if user == nil {
			return
		}
		session.Delete(userkey)
		if err := session.Save(); err != nil {
			return
		}

		c.Redirect(http.StatusMovedPermanently, "/")
	}
}

func (sc *SessionController) IndexGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(userkey)
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"content": "This is an index page...",
			"user":    user,
		})
	}
}

func (sc *SessionController) DashboardGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(userkey)
		c.HTML(http.StatusOK, "dashboard.tmpl", gin.H{
			"content": "This is a dashboard page...",
			"user":    user,
		})
	}
}

// // GetLoginPage
// func LoginPage(c *gin.Context) {
// 	// name := c.Param("name")
// 	c.HTML(
// 		http.StatusOK, "login.tmpl", gin.H{},
// 	)
// }

// // GetSignupPage
// func SignupPage(c *gin.Context) {
// 	c.HTML(
// 		http.StatusOK, "signup.tmpl", gin.H{},
// 	)
// }

// // PostLogin
// func Login(c *gin.Context) {
// 	name := c.Request.FormValue("name")
// 	password := c.Request.FormValue("password")

// 	if hasSession := middleware.HasSession(c); hasSession {
// 		c.String(200, "already Logged in")
// 		return
// 	}

// 	user := models.UserDetailByName(name)

// 	if err := middleware.Compare(user.Password, password); err != nil {
// 		c.String(401, "password mismatch")
// 		return
// 	}

// 	middleware.SaveAuthSession(c, user.ID)

// 	c.String(200, "login successful")
// }

// // PostLogout
// func Logout(c *gin.Context) {
// 	if hasSession := middleware.HasSession(c); hasSession {
// 		c.String(401, "用戶未登入")
// 		return
// 	}
// 	middleware.ClearAuthSession(c)
// 	c.String(200, "退出成功")
// }

// // PostSignup
// func Signup(c *gin.Context) {
// 	var user models.User
// 	user.Name = c.Request.FormValue("name")
// 	user.Password = c.Request.FormValue("password")
// 	user.Email = c.Request.FormValue("email")

// 	if hasSession := middleware.HasSession(c); hasSession {
// 		c.String(200, "用户已登陆")
// 		return
// 	}

// 	if existUser := models.UserDetailByName(user.Name); existUser.ID != uuid.Nil {
// 		c.String(200, "用户名已存在")
// 		return
// 	}

// 	if c.Request.FormValue("password") != c.Request.FormValue("password_confirmation") {
// 		c.String(200, "密码不一致")
// 		return
// 	}

// 	if pwd, err := middleware.Encrypt(c.Request.FormValue("password")); err == nil {
// 		user.Password = pwd
// 	}

// 	models.AddUser(&user)

// 	middleware.SaveAuthSession(c, user.ID)

// 	c.String(200, "注册成功")
// }

// func Me(c *gin.Context) {
// 	currentUser := c.MustGet("userId").(uint)
// 	c.JSON(http.StatusOK, gin.H{
// 		"code": 1,
// 		"data": currentUser,
// 	})
// }

// func Signup(c *gin.Context) {
// 	email := c.PostForm("email")
// 	password := c.PostForm("password")
// 	confirmPassword := c.PostForm("confirm_password")
// }
