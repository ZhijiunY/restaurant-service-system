package controllers

import (
	"fmt"
	"log"
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
	ID       = "ID"
	userName = "Name"
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

		// 檢查Session中是否存在用戶ID
		userID := session.Get(ID)
		if userID == nil {
			// 如果用戶ID不存在，則認為用戶未登入，重定向到登入頁面
			log.Println("User is not logged in, redirected to login page")
			c.Redirect(http.StatusSeeOther, "/auth/getlogin")
			c.Abort() // 確保不再繼續處理後續的請求處理函數
			return
		}

		// 如果用戶ID存在，則認為用戶已登入，繼續處理後續的請求處理函數
		log.Println("The user is logged in, continue processing the request")
		c.Next()
	}
}

// 註冊頁面
func (sc *SessionController) SignupGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		// 檢查Session中是否已有用戶ID，來判斷用戶是否已登入
		userID := session.Get("ID")
		if userID != nil {
			// 如果用戶已經登入，則重定向到主頁面或菜單頁面，並提示用戶先登出
			c.Redirect(http.StatusSeeOther, "/")
			return
		}

		// 如果用戶未登入，則顯示註冊頁面
		c.HTML(http.StatusOK, "signup.tmpl", gin.H{
			"content": "",
		})
	}
}

// 根據LoginPost去重構Signup，並使用繁體中文說明

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

		middleware.SaveIDSession(c, newUser.ID)

		// Redirect to login page
		c.Redirect(http.StatusSeeOther, "/auth/getlogin")
	}
}

func (sc *SessionController) LoginGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		// 檢查Session中是否已有用戶ID，來判斷用戶是否已登入
		userID := session.Get("ID")
		if userID != nil {
			// 如果已經登入，則重定向到菜單頁面
			c.Redirect(http.StatusSeeOther, "/menu")
			return
		}

		// 如果用戶未登入，則顯示登入頁面
		c.HTML(http.StatusOK, "login.tmpl", gin.H{"content": ""})

	}
}

func (sc *SessionController) LoginPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 獲取Session
		session := sessions.Default(c)

		// 先清除可能存在的Session
		session.Clear()

		// 獲取表單數據
		email := c.PostForm("email")
		password := c.PostForm("password")

		// 驗證郵箱和密碼是否為空
		if email == "" || password == "" {
			c.HTML(http.StatusBadRequest, "login.tmpl", gin.H{"content": "Email and password cannot be empty"})
			return
		}

		// 檢查用戶是否存在於數據庫
		var user models.User
		if err := utils.DB.Where("email = ?", email).First(&user).Error; err != nil {
			c.HTML(http.StatusBadRequest, "login.tmpl", gin.H{"content": "Invalid email or password"})
			return
		}

		// 比較密碼
		if !middleware.Compare(password, user.Password) { // 假設utils.Compare是比較密碼的函數
			c.HTML(http.StatusBadRequest, "login.tmpl", gin.H{"err": "Incorrect password"})
			return
		}

		// 認證成功，將用戶ID和名稱保存到Session
		session.Set(ID, user.ID.String())
		session.Set(userName, user.Name)
		if err := session.Save(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
			return
		}

		// 重定向到菜單頁面
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
		c.Redirect(http.StatusSeeOther, "/auth/getlogin")
	}
}
