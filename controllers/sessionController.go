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
			// c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			// 	"message:": "need to login!!!!!!!",
			// })
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
		// sc.AuthRequired()

		// get form value
		name := c.PostForm("name")
		email := c.PostForm("email")
		password := c.PostForm("password")
		fmt.Println("fet value error")

		// if hasSession := middleware.HasSession(c); hasSession {
		// 	c.String(200, "already logged in")
		// 	return
		// }
		// if existUser := models.UserDetailByName(name); existUser.ID != uuid.Nil {
		// 	c.String(200, "user already exists")
		// 	return
		// }
		// if pwd, err := middleware.Encrypt(password); err == nil {
		// 	password = pwd
		// }

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
			c.String(200, "user already exists")
			return
		}

		// Hash password
		if pwd, err := middleware.Encrypt(password); err == nil {
			password = pwd
		}

		// create new user
		newUser := &models.User{
			ID:         uuid.New(),
			Name:       name,
			Password:   password,
			Email:      email,
			Created_at: time.Now(),
			Updated_at: time.Now(),
		}

		// Store user in database
		err := utils.DB.Create(newUser).Error
		if err != nil {
			// c.AbortWithError(http.StatusInternalServerError, err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to create user"})
			return
		}

		// 在session中儲存用戶資訊
		// session := sessions.Default(c)
		// session.Set(userkey, newUser.ID)
		// session.Set(userkey, email)
		// err = session.Save()
		// if err != nil {
		// 	c.AbortWithError(http.StatusInternalServerError, err)
		// 	fmt.Println("store session error 400")
		// 	return
		// }
		middleware.SaveAuthSession(c, newUser.ID)

		// Redirect to login page
		c.Redirect(http.StatusSeeOther, "/user/login")
		//c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})

	}
}

func (sc *SessionController) LoginGet() gin.HandlerFunc {
	return func(c *gin.Context) {

		middleware.ClearAuthSession(c)

		session := sessions.Default(c)
		user := session.Get(userkey)
		if user != nil {
			c.HTML(http.StatusBadRequest, "login.tmpl", gin.H{
				"content": "Please logout first",
				"user":    user,
			})
			fmt.Println("please louout first 400")
			return
		}
		c.HTML(http.StatusOK, "login.tmpl", gin.H{
			"content": "",
			"user":    user,
		})

		fmt.Println("login user 200")

		middleware.ClearAuthSession(c)

	}
}

func (sc *SessionController) LoginPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		middleware.ClearAuthSession(c)

		sc.AuthRequired()

		hasSession := middleware.HasSession(c)
		if hasSession {
			c.Redirect(http.StatusSeeOther, "/menu")
			fmt.Println("Authentication login")
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

			// Check if user is already logged in
			if hasSession := middleware.HasSession(c); hasSession {
				c.String(200, "already logged")
				fmt.Println("already logged")
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

			hashedPwd := user.Password

			// Compare password
			if match := middleware.Compare(password, hashedPwd); !match {
				c.HTML(http.StatusBadRequest, "login.tmpl", gin.H{
					"err": "incorrect password",
				})
				fmt.Println("invalid password")
				return
			}

			// Save user ID to session
			middleware.SaveAuthSession(c, user.ID)

			// Create cookie
			cookie := &http.Cookie{
				Name:  "auth_token",
				Value: user.ID.String(),
				Path:  "/",
			}
			// Set cookie
			http.SetCookie(c.Writer, cookie)
			fmt.Println("set cookie")

		}
		// Redirect to home page
		c.Redirect(http.StatusSeeOther, "/menu")
	}

}

func (sc *SessionController) LogoutPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		sc.AuthRequired()

		// clear session
		if err := middleware.ClearAuthSession(c); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			fmt.Println("Failed to delete session")
			return
		}

		// 從用戶中注銷，並將用戶重定向回主頁
		// Logout from the user and redirect the user back to the homepage.
		middleware.ClearAuthSession(c)
		c.Redirect(http.StatusMovedPermanently, "/")
	}

	// 	session := sessions.Default(c)
	// 	user := session.Get(userkey)

	// 	if user == nil {
	// 		return
	// 	}
	// 	session.Delete(userkey)
	// 	if err := session.Save(); err != nil {
	// 		return
	// 	}

	// 	c.Redirect(http.StatusMovedPermanently, "/")
	// }

}
