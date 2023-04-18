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

		if pwd, err := middleware.Encrypt(c.PostForm("password")); err == nil {
			password = pwd
		}

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
			"auth":    user,
		})
		fmt.Println("login user 200")
	}
}

func (sc *SessionController) LoginPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		sc.AuthRequired()
		session := sessions.Default(c)
		userkey := session.Get(userkey)

		// Get form values
		email := c.PostForm("email")
		password := c.PostForm("password")

		// check if login was successful
		if hasSession := middleware.HasSession(c); hasSession {
			c.String(200, "already logged")
		}

		// Validate email and password
		if email == "" || password == "" {
			c.HTML(http.StatusBadRequest, "login.tmpl", gin.H{
				"content": "Email and password cannot be empty",
				"user":    nil,
			})
			fmt.Println("empty error")
			return
		}

		// // Check password
		// middleware.Compare(user.Password, password)
		// check if password is correct
		if err := middleware.Compare(password, user.Password); err != nil {
			c.String(200, "incorrect password")
			return
		}

		// Check if user exists in database
		// Verify user credentials
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
			"auth":    "login",
		})
		fmt.Println("Status OK")

		// Create a new session for the user
		session.Set(userkey, user.ID)
		session.Set(emailkey, email)
		err = session.Save()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)

			return
		}

		// Save user ID to session
		middleware.SaveAuthSession(c, user.ID)

		// 重定向到主頁
		c.Redirect(http.StatusSeeOther, "/menu")
	}

	// 下面這些跑得動，但還是無法解密
	// 	// get form value
	// 	email := c.PostForm("email")
	// 	password := c.PostForm("password")

	// 	// check if user exists
	// 	user := models.UserDetailByEmail(email)
	// 	if user.ID == uuid.Nil {
	// 		c.String(200, "user does not exist")
	// 		return
	// 	}

	// 	// // Compare entered password with stored hash password
	// 	// if err := middleware.Compare(user.Password, password); err != nil {
	// 	// 	// Invalid password
	// 	// 	c.HTML(http.StatusBadRequest, "login.tmpl", gin.H{
	// 	// 		"content": "Invalid Email or Password",
	// 	// 		"user":    nil,
	// 	// 	})
	// 	// 	return
	// 	// }

	// 	// check if password is correct
	// 	if err := middleware.Compare(password, user.Password); err != nil {
	// 		c.String(200, "incorrect password")
	// 		return
	// 	}

	// 	// save user info in session
	// 	middleware.SaveAuthSession(c, user.ID)

	// 	// Redirect to home page
	// 	c.Redirect(http.StatusSeeOther, "/")
	// }
}

func (sc *SessionController) LogoutPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 獲取現有會話並刪除已經存儲的用戶ID和郵件
		// Get the existing session and remove the stored user ID and email
		session := sessions.Default(c)
		session.Delete(userkey)
		session.Delete(emailkey)
		err := session.Save()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			fmt.Println("Failed to delete session")
			return
		}

		// 從用戶中注銷，並將用戶重定向回主頁
		// Logout from the user and redirect the user back to the homepage.
		middleware.ClearAuthSession(c)
		c.Redirect(http.StatusSeeOther, "/")
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
