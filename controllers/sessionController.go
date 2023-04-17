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

		sc.AuthRequired()
		session := sessions.Default(c)
		user := session.Get(userkey)

		// // if user != nil, 則表示用戶已經登錄
		// if user != nil {
		// 	c.HTML(http.StatusBadRequest, "signup.tmpl",
		// 		gin.H{
		// 			"content": "already logged in",
		// 			"user":    user,
		// 		})
		// 	fmt.Println("already logged in")
		// 	return
		// }

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
		sc.AuthRequired()
		// // 檢查是否已有有效的 session。
		// if middleware.CheckSession(c) {
		// 	// 如果有，則將用戶重定向到主頁面或其他頁面。
		// 	c.Redirect(http.StatusMovedPermanently, "/")
		// 	return
		// }

		// get form value
		name := c.PostForm("name")
		email := c.PostForm("email")
		password := c.PostForm("password")

		fmt.Println("fet value error")

		// Validate email and password
		if email == "" || password == "" || name == "" {
			c.HTML(http.StatusBadRequest, "signup.tmpl", gin.H{
				"content": "Email and password cannot be empty",
				"user":    nil,
			})
			fmt.Println("Email, Password or name are empty ")
			return
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
			fmt.Println("500")
			return
		}

		// 在session中儲存用戶資訊
		session := sessions.Default(c)
		session.Set(userkey, newUser.ID)
		session.Set(userkey, email)
		err = session.Save()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			fmt.Println("store session error 400")
			return
		}

		// Redirect to login page
		c.Redirect(http.StatusSeeOther, "/user/login")
		c.JSON(http.StatusOK, gin.H{
			"auth": "signup",
		})
		//c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})

	}
}

// func signupPOST(c *gin.Context) {
// 	// 檢查是否已有有效的 session。
// 	if middleware.CheckSession(c) {
// 		// 如果有，則將用戶重定向到主頁面或其他頁面。
// 		c.Redirect(http.StatusMovedPermanently, "/")
// 		return
// 	}

// 	// 設置 cookie-based session 中間件，以便在該請求中處理 session。
// 	sessionMiddleware := middleware.EnableCookieSession()

// 	// 解析表單中提交的用戶註冊信息。
// 	// ...

// 	// 創建新的用戶，並將用戶信息寫入數據庫。
// 	// ...

// 	// 保存用戶的 session。
// 	middleware.SaveAuthSession(c, userID)

// 	// 將用戶重定向到註冊成功頁面。
// 	// ...
// }

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
			fmt.Println("please louout first 400")
			return
		}
		c.HTML(http.StatusOK, "login.tmpl", gin.H{
			"content": "",
			"user":    user,
			"auth":    "login",
		})
		fmt.Println("login user 200")
	}
}

func (sc *SessionController) LoginPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		// Get form values
		email := c.PostForm("email")
		password := c.PostForm("password")
		//auth := c.PostForm("auth")

		// Validate email and password
		if email == "" || password == "" {
			c.HTML(http.StatusBadRequest, "login.tmpl", gin.H{
				"content": "Email and password cannot be empty",
				"user":    nil,
			})
			fmt.Println("empty error")
			return
		}

		// Check if user exists in database
		// Verify user credentials
		// 驗證用戶憑據
		var user models.User
		err := utils.DB.Where("email = ?", email).First(&user).Error
		if err != nil || user.Password != password {
			c.HTML(http.StatusBadRequest, "login.tmpl", gin.H{
				"content": "Invalid email or password",
				"user":    nil,
			})
			fmt.Println("user error")
			return
		}
		c.HTML(http.StatusOK, "menu.tmpl", gin.H{
			"content": "login successfully",
			"user":    user,
			//"auth":    auth,
		})
		fmt.Println("Status OK")

		// // Check password
		// err = middleware.Compare(user.Password, password)
		// if err != nil {
		// 	c.HTML(http.StatusBadRequest, "login.tmpl", gin.H{
		// 		"content": "無效的郵件或密碼",
		// 		"user":    nil,
		// 	})
		// 	fmt.Println("password error")
		// 	return
		// }

		// 為用戶創建新會話
		session.Set(userkey, user.ID)
		session.Set(emailkey, email)
		err = session.Save()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			fmt.Println("Failed to create session")
			return
		}

		// 重定向到主頁
		c.Redirect(http.StatusSeeOther, "/menu")
	}
}

func (sc *SessionController) LogoutPost() gin.HandlerFunc {
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
