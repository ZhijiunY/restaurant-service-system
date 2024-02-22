package middleware

import (
	"fmt"
	"net/http"

	"github.com/ZhijiunY/restaurant-service-system/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/google/uuid"
)

var Secret = []byte("secret")

// const (
//
//	userkey   = "user"
//	ID        = "ID"
//	userNamee = "Name"
//	emailkey  = "email"
//	mysession = "mysession"
//
// )
const (
	userkey   = "user"
	ID        = "userID"
	Name      = "userName"
	emailkey  = "email"
	mysession = "mysession"
)

// // Save session using cookies
// func EnableCookieSession() gin.HandlerFunc {
// 	store := cookie.NewStore([]byte(userkey))
// 	return sessions.Sessions(mysession, store)
// }

func SaveIDSession(c *gin.Context, userID uuid.UUID) {
	session := sessions.Default(c)
	session.Set(userkey, userID.String()) // 將UUID轉換成字串
	err := session.Save()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		fmt.Println("store session error 400")
		return
	}
}

// Check if the current request contains a valid user session and return a boolean value
func HasSession(c *gin.Context) bool {
	session := sessions.Default(c)
	if sessionValue := session.Get(userkey); sessionValue == nil {
		return false
	}
	return true
}

// ClearAuthSession for User
func ClearAuthSession(c *gin.Context) error {
	session := sessions.Default(c)
	session.Clear()
	err := session.Save()
	if err != nil {
		return errors.Wrap(err, "failed to delete session")
	}
	return nil
}

// 函數從 session 中獲取使用者的 ID
// 如果 session 不存在或 ID 為空，則返回一個空的 UUID
// Get Session for User
func GetSessionUserId(c *gin.Context) uuid.UUID {
	session := sessions.Default(c)
	sessionID := session.Get(ID)
	if sessionID == nil {
		return uuid.Nil
	}
	return sessionID.(uuid.UUID)
}

// CheckSession for User
func CheckSession(c *gin.Context) bool {
	var user models.User
	session := sessions.Default(c)
	sessionID := session.Get(user.ID)
	return sessionID != nil
}

func UserAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userName := session.Get("Name")
		if userName == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User Name not found"})
			c.Abort()
			return
		}
		// 將 userName 設置到 Gin 上下文中，以便後續的處理器可以使用
		c.Set("userName", userName)
		c.Next()
	}
}
