package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ZhijiunY/restaurant-service-system/middleware"
	"github.com/ZhijiunY/restaurant-service-system/models"
	"github.com/ZhijiunY/restaurant-service-system/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	userkey  = "user"
	emailkey = "email"
)

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

// Check if user is already logged in
func (sc *SessionController) AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(userkey)
		if user == nil {
			c.Redirect(http.StatusMovedPermanently, "/user/login")

			return
		}
		c.Next()
	}

}

// 註冊頁面
func (sc *SessionController) SignupGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(userkey)
		if user != nil {
			c.HTML(http.StatusBadRequest, "login.tmpl", gin.H{
				"content": "Please logout first",
				"user":    user,
			})
			return
		}
		c.HTML(http.StatusOK, "signup.tmpl", gin.H{
			"content": "",
			"user":    user,
			"auth":    user,
		})
	}
}

func (sc *SessionController) SignupPost() gin.HandlerFunc {
	return func(c *gin.Context) {

		// get form value
		name := c.PostForm("name")
		email := c.PostForm("email")
		password := c.PostForm("password")
		fmt.Println("fet value error")

		if hasSession := middleware.HasSession(c); hasSession {
			c.String(200, "already logged in")
			return
		}
		if existUser := models.UserDetailByName(name); existUser.ID != uuid.Nil {
			c.String(200, "user already exists")
			return
		}

		// Validate email and password
		if email == "" || password == "" || name == "" {
			c.HTML(http.StatusBadRequest, "signup.tmpl", gin.H{
				"content": "Email, password, or name cannot be empty",
				"user":    nil,
			})
			return
		}

		// Check if user already exists
		if existUser := models.UserDetailByName(name); existUser.ID != uuid.Nil {
			c.String(http.StatusBadRequest, "user already exists")
			return
		}

		// Hash password
		hashedPwd, err := middleware.Encrypt(password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Password encryption failed"})
			return
		}

		// create new user
		newUser := &models.User{
			ID:         uuid.New(),
			Name:       name,
			Password:   hashedPwd,
			Email:      email,
			Created_at: time.Now(),
			Updated_at: time.Now(),
		}

		// Store user in database
		if err := utils.DB.Create(newUser).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create user"})
			return
		}

		middleware.SaveAuthSession(c, newUser.ID)

		// Redirect to login page
		c.Redirect(http.StatusSeeOther, "/auth/login")

	}
}

func (sc *SessionController) LoginGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(userkey)
		if user != nil {
			c.HTML(http.StatusBadRequest, "login.tmpl", gin.H{
				"content": "Please logout first",
			})
			return
		}
		c.HTML(http.StatusOK, "login.tmpl", gin.H{
			"content": "",
		})
		middleware.ClearAuthSession(c)

	}
}

func RenderNavigation(c *gin.Context) {
	session := sessions.Default(c)
	userName := session.Get("Name")

	c.HTML(http.StatusOK, "navigation.tmpl", gin.H{
		"Name": userName,
	})
}

func (sc *SessionController) LoginPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		middleware.ClearAuthSession(c)

		hasSession := middleware.HasSession(c)
		if hasSession {
			c.Redirect(http.StatusSeeOther, "/menu")
		} else {

			// Get form values
			email := c.PostForm("email")
			password := c.PostForm("password")

			// Validate email and password
			if email == "" || password == "" {
				c.HTML(http.StatusBadRequest, "login.tmpl", gin.H{
					"content": "Email and password cannot be empty",
					"user":    nil,
				})
				fmt.Println("Email and password cannot be empty")
				return
			}

			// Check if user exists in database
			// Verify user credentials
			err := utils.DB.Where("email = ?", email).First(&user).Error
			if err != nil {
				c.HTML(http.StatusBadRequest, "login.tmpl", gin.H{
					"content": "Invalid email or password",
					"user":    nil,
				})
				fmt.Println("Invalid email or password")
				return
			}

			// Compare password
			hashedPwd := user.Password
			if match := middleware.Compare(password, hashedPwd); !match {
				c.HTML(http.StatusBadRequest, "login.tmpl", gin.H{
					"err": "incorrect password",
				})
				fmt.Println("invalid password")
				return
			}

			// Save user ID to session
			middleware.SaveAuthSession(c, user.ID)

			// 在用户登入成功後保存用户ID到session
			session := sessions.Default(c)
			session.Set(userkey, user.ID)
			session.Set("Name", user.Name) // 這裡假設 user 結構體有一个 Name 字段
			session.Save()

		}
		// Redirect to menu page
		c.Redirect(http.StatusSeeOther, "/menu")

	}
}

func (sc *SessionController) LogoutPost() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Clear the user's session
		session := sessions.Default(c)
		session.Clear()
		err := session.Save()
		if err != nil {
			// Handle the error, perhaps by logging it and redirecting with an error message
			c.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
				"content": "Error clearing the session",
			})
			return
		}

		// Redirect to the login page or home page
		c.Redirect(http.StatusSeeOther, "/auth/login")
	}
}

// 登入後出現使用者名字
// 點餐頁面可以出現順暢的總計小計
