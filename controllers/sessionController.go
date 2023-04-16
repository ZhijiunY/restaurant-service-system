package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ZhijiunY/restaurant-service-system/models"
	"github.com/ZhijiunY/restaurant-service-system/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

// 註冊頁面
func (sc *SessionController) SignupGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(userkey)

		// if user != nil, 則表示用戶已經登錄
		if user != nil {
			c.HTML(http.StatusBadRequest, "signup.tmpl",
				gin.H{
					"content": "already logged in",
					"user":    user,
				})
			fmt.Println("already logged in")
			return
		}

		// 如果 user 變數為 nil，則表示用戶還沒有登錄
		c.HTML(http.StatusOK, "signup.tmpl", gin.H{
			"content": "",
			"user":    user,
		})
		fmt.Println("please sign up first")
	}
}

func (sc *SessionController) SignupPost() gin.HandlerFunc {
	return func(c *gin.Context) {

		// // get form value
		session := sessions.Default(c)
		name := c.PostForm("name")
		email := c.PostForm("email")
		password := c.PostForm("password")
		fmt.Println("fet value error")

		// // Validate email and password
		// if email == "" || password == "" || name == "" {
		// 	c.HTML(http.StatusBadRequest, "signup.tmpl", gin.H{
		// 		"content": "Email and password cannot be empty",
		// 		"user":    nil,
		// 	})
		// 	fmt.Println("Email ro Password error ")
		// 	return
		// }

		// // Store user in session
		// session.Set(userkey, email)
		// err := session.Save()
		// if err != nil {
		// 	c.AbortWithError(http.StatusInternalServerError, err)
		// 	fmt.Println("store session error 400")
		// 	return
		// }

		// // Store user in database
		// user := models.User{Email: email, Password: password}
		// err = utils.DB.Create(&user).Error
		// if err != nil {
		// 	c.AbortWithError(http.StatusInternalServerError, err)
		// 	fmt.Println("500")
		// 	return
		// }

		// // 創建新用戶
		// user.ID = uuid.New()
		// user.Created_at = time.Now()
		// user.Updated_at = time.Now()

		// // Redirect to login page
		// c.Redirect(http.StatusSeeOther, "/user/login")

		// 編寫處理 POST 請求的代碼
		// name := c.PostForm("name")
		// email := c.PostForm("email")
		// password := c.PostForm("password")

		// 創建新用戶
		newUser := &models.User{
			ID:         uuid.New(),
			Name:       name,
			Password:   password,
			Email:      email,
			Created_at: time.Now(),
			Updated_at: time.Now(),
		}

		// 將用戶儲存到數據庫
		err := utils.DB.Create(newUser).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to create user"})
			return
		}

		// 在session中儲存用戶資訊
		//session := sessions.Default(c)
		session.Set(userkey, newUser.ID)
		session.Save()

		// Redirect to login page
		c.Redirect(http.StatusSeeOther, "/user/login")

		//c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})

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
			fmt.Println("please louout error")
			return
		}
		c.HTML(http.StatusOK, "login.tmpl", gin.H{
			"content": "",
			"user":    user,
		})
		fmt.Println("login user error")
	}
}

func (sc *SessionController) LoginPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userkey := session.Get(userkey)

		email := c.PostForm("email")
		password := c.PostForm("password")
		if userkey != nil {
			c.HTML(http.StatusBadRequest, "login.tmpl", gin.H{"content": "Please logout first"})
			fmt.Println("please log out firsts")
			return

		}

		// Validate email and password
		if email == "" || password == "" {
			c.HTML(http.StatusBadRequest, "login.tmpl", gin.H{
				"content": "Email and password cannot be empty",
				"user":    nil,
			})
			return
		}

		// Check if user exists in database
		var user models.User
		err := utils.DB.Where("email = ?", email).First(&user).Error
		if err != nil {
			c.HTML(http.StatusBadRequest, "login.tmpl", gin.H{
				"content": "Invalid email or password",
				"user":    nil,
			})
			return
		}

		// if userkey.password != password {
		// 	c.HTML(http.StatusBadRequest, "login.tmpl", gin.H{
		// 		"content": "Invalid email or password",
		// 		"user":    nil,
		// 	})
		// 	return
		// }

		// Store user in session
		session.Set("user", email)
		err = session.Save()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.Redirect(http.StatusSeeOther, "/menu")

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
