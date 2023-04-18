package middleware

import (
	"fmt"
	"net/http"

	"github.com/ZhijiunY/restaurant-service-system/models"
	"github.com/ZhijiunY/restaurant-service-system/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	//"github.com/gofrs/uuid"
	"github.com/google/uuid"
)

var Secret = []byte("secret")

// const Userkey = "user"
// const User = "user_id"
const (
	userkey  = "user"
	emailkey = "email"
)

// Save session using cookies
func EnableCookieSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte(userkey))
	return sessions.Sessions("mysession", store)

}

// func EnableCookieSession() gin.HandlerFunc {
// 	sessionName := "mysession"
// 	store := cookie.NewStore([]byte(User))
// 	sessionMiddleware := sessions.Sessions(sessionName, store)

// 	sessionController := controllers.NewSessionController(store)
// 	sessionController.LoadAndSave()

// 	return func(c *gin.Context) {
// 		sessionMiddleware(c)
// 		c.Set("sessionController", sessionController)
// 	}
// }

// UserAuthSessionMiddle
// 中間鍵 驗證是否已登入
func AuthSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionID := session.Get(userkey)
		if sessionID == nil {
			c.Redirect(http.StatusMovedPermanently, "/user/login")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message:": "need to login!",
			})
			return
		}
		c.Next()
	}
}

func SaveAuthSession(c *gin.Context, userID uuid.UUID) {
	session := sessions.Default(c)
	session.Set(userkey, userID.String()) // 將UUID轉換成字串
	err := session.Save()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		fmt.Println("store session error 400")
		return
	}
	session.Save()
}

// ClearAuthSession for User
func ClearAuthSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

// 檢查當前請求是否包含有效的使用者 session，返回值為布林值；
func HasSession(c *gin.Context) bool {
	session := sessions.Default(c)
	if sessionValue := session.Get("userId"); sessionValue == nil {
		return false
	}
	return true
}

// 函數從 session 中獲取使用者的 ID
// 如果 session 不存在或 ID 為空，則返回一個空的 UUID
func GetSessionUserId(c *gin.Context) uuid.UUID {
	session := sessions.Default(c)
	sessionValue := session.Get("userId")
	if sessionValue == nil {
		return uuid.UUID{}
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
	userName := ""
	if hasSession {
		userId := GetSessionUserId(c)
		var user models.User
		if err := utils.DB.Where("id = ?", userId).First(&user).Error; err == nil {
			userName = user.Name
		}
	}
	data := make(map[string]interface{})
	data["hasSession"] = hasSession
	data["userName"] = userName
	return data
}

// CheckSession for User
func CheckSession(c *gin.Context) bool {
	session := sessions.Default(c)
	sessionID := session.Get(userkey)
	return sessionID != nil
}
