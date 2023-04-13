package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const ID = "session_id"

// Save session using cookies
func EnableCookieSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte(ID))
	return sessions.Sessions("mysession", store)
}

// User Auth Session Middle
// 中間鍵
func AuthSessionMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionID := session.Get(ID)
		if sessionID == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message:": "need to login!",
			})
			return
		}
		c.Next()
	}
}

// Save Session for User
func SaveAuthSession(c *gin.Context, userID int) {
	session := sessions.Default(c)
	session.Set(ID, userID)
	session.Save()
}

// Clear Session for User
func ClearAuthSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

// Get Session for User
func GetSession(c *gin.Context) int {
	session := sessions.Default(c)
	sessionID := session.Get(ID)
	if sessionID == nil {
		return 0
	}
	return sessionID.(int)
}

// Check Session for User
func CheckSession(c *gin.Context) bool {
	session := sessions.Default(c)
	sessionID := session.Get(ID)
	return sessionID != nil
}

func HasSession(c *gin.Context) bool {
	session := sessions.Default(c)
	if sessionValue := session.Get("ID"); sessionValue == nil {
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
