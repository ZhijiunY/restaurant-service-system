package middleware

import (
	"fmt"
	"net/http"

	"github.com/ZhijiunY/restaurant-service-system/models"
	"github.com/ZhijiunY/restaurant-service-system/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/google/uuid"
)

var Secret = []byte("secret")

const (
	userkey     = "user"
	emailkey    = "email"
	sessionName = "my-session"
)

// Save session using cookies
func EnableCookieSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte(userkey))
	return sessions.Sessions("mysession", store)

}

// // UserAuthSessionMiddle
// // 中間鍵 驗證是否已登入
// func AuthSession() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		session := sessions.Default(c)
// 		sessionID := session.Get(userkey)
// 		if sessionID == nil {
// 			c.Redirect(http.StatusMovedPermanently, "/user/login")
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 				"message:": "need to login!",
// 			})
// 			return
// 		}
// 		c.Next()
// 	}
// }

func SaveAuthSession(c *gin.Context, userID uuid.UUID) {
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
	sessionValue := session.Get("userkey")
	if sessionValue == nil {
		return uuid.Nil
	}
	return sessionValue.(uuid.UUID)
}

// 函數從當前請求中獲取使用者的 session 資訊
// 包括是否有有效的 session、使用者名稱等
// 並將這些資訊封裝到一個 map 中返回。
// 在這個函數中，如果當前請求包含有效的 session
// 還會從資料庫中查詢出相應的使用者資訊。
func GetUserSession(c *gin.Context) map[string]interface{} {
	hasSession := HasSession(c)
	userID := uuid.Nil
	userName := ""
	if hasSession {
		userId := GetSessionUserId(c)

		var user models.User
		if err := utils.DB.Where("id = ?", userId).First(&user).Error; err == nil {
			userID = user.ID
		}

	}
	// 從HTTP請求參數中獲取名稱
	data := make(map[string]interface{})
	data["hasSession"] = hasSession
	data["userId"] = userID
	data["userName"] = userName
	return data

}

// CheckSession for User
func CheckSession(c *gin.Context) bool {
	session := sessions.Default(c)
	sessionID := session.Get(userkey)
	return sessionID != nil
}
