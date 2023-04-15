package middleware

import (
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
const User = "user_id"

// Save session using cookies
func EnableCookieSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte(User))
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
		sessionID := session.Get(User)
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
	session.Set(User, userID.String()) // 將UUID轉換成字串
	session.Save()
}

// ClearAuthSession for User
func ClearAuthSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

func HasSession(c *gin.Context) bool {
	session := sessions.Default(c)
	if sessionValue := session.Get("userId"); sessionValue == nil {
		return false
	}
	return true
}

func GetSessionUserId(c *gin.Context) uint {
	session := sessions.Default(c)
	sessionValue := session.Get("userId")
	if sessionValue == nil {
		return 0
	}
	return sessionValue.(uint)
}

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

// // GetSession for User
// func GetSession(c *gin.Context) int {
// 	session := sessions.Default(c)
// 	sessionID := session.Get(User)
// 	if sessionID == nil {
// 		return 0
// 	}
// 	return sessionID.(int)
// }

// // CheckSession for User
// func CheckSession(c *gin.Context) bool {
// 	session := sessions.Default(c)
// 	sessionID := session.Get(User)
// 	return sessionID != nil
// }

// func HasSession(c *gin.Context) bool {
// 	session := sessions.Default(c)
// 	if sessionValue := session.Get("ID"); sessionValue == nil {
// 		return false
// 	}
// 	return true
// }

// func GetSessionUserId(c *gin.Context) uint {
// 	session := sessions.Default(c)
// 	sessionValue := session.Get("userId")
// 	if sessionValue == nil {
// 		return 0
// 	}
// 	return sessionValue.(uint)
// }

// func GetUserSession(c *gin.Context) map[string]interface{} {

// 	// hasSession := HasSession(c)
// 	// userName := ""
// 	// if hasSession {
// 	// 	ID := GetSessionUserId(c)
// 	// 	userName = models.User(ID).Name
// 	// }
// 	// data := make(map[string]interface{})
// 	// data["hasSession"] = hasSession
// 	// data["userName"] = userName
// 	return GetUserSession
// }
