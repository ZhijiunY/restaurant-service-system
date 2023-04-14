package routes

import (
	"github.com/ZhijiunY/restaurant-service-system/controllers"
	"github.com/ZhijiunY/restaurant-service-system/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var secret = []byte("secret")

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(middleware.Logger())
	router.Use(gin.Recovery())
	router.Use(sessions.Sessions("mysession", cookie.NewStore(secret)))

	router.Use(gin.BasicAuth(gin.Accounts{"Simba": "1234"}))

	// connect to template
	// Static file
	router.LoadHTMLGlob("./templates/**/*")
	router.Static("/static", "./static")

	// Grouping routes
	MainRoutes := router.Group("/")
	{
		MainRoutes.GET("/", controllers.GetIndex)
		MainRoutes.GET("/menu", controllers.GetMenu)
		MainRoutes.GET("/order", controllers.GetManager)
	}

	AuthRoutes := router.Group("/user")
	{
		AuthRoutes.GET("/login", controllers.LoginPage)
		AuthRoutes.GET("/signup", controllers.SignupPage)
		AuthRoutes.POST("/login", controllers.Login)
		AuthRoutes.POST("/logout", controllers.Logout)
		AuthRoutes.POST("/signup", controllers.Signup)

		AuthRoutes.Static("/static", "./static")
	}

	UserRoutes := router.Group("/user")
	{
		UserRoutes.GET("/", controllers.GetUser)
		UserRoutes.POST("/", controllers.CreateUser)
		UserRoutes.PUT("/:id", controllers.UpdateUser)
		UserRoutes.DELETE("/:id", controllers.DeleteUser)
	}

	return router
}

// 上面的代碼中，我們創建了一個名為api的路由組，然後為這個組設置了幾個API路由。最後，我們使用Static()方法為api路由組設置了一個靜態文件路由，將所有以/api/static開頭的請求路徑映射到./static目錄下。

// 當然，你也可以在分組路由的結構中嵌套子分組路由，然後為子分組路由設置靜態文件路由，例如：

// go
// Copy code
// func main() {
//     router := gin.Default()

//     // Grouping routes
//     api := router.Group("/api")
//     {
//         api.GET("/users", getUsers)
//         api.GET("/users/:id", getUser)
//         api.POST("/users", createUser)
//         api.PUT("/users/:id", updateUser)
//         api.DELETE("/users/:id", deleteUser)

//         // Nested grouping
//         v1 := api.Group("/v1")
//         {
//             v1.GET("/users", getUsersV1)
//             v1.GET("/users/:id", getUserV1)
//             v1.POST("/users", createUserV1)
//             v1.PUT("/users/:id", updateUserV1)
//             v1.DELETE("/users/:id", deleteUserV1)

//             // Serve static files under "/api/v1/static"
//             v1.Static("/static", "./static")
//         }

//         v2 := api.Group("/v2")
//         {
//             v2.GET("/users", getUsersV2)
//             v2.GET("/users/:id", getUserV2)
//             v2.POST("/users", createUserV2)
//             v2.PUT("/users/:id", updateUserV2)
//             v2.DELETE("/users/:id", deleteUserV2)

//             // Serve static files under "/api/v2/static"
//             v2.Static("/static", "./static")
//         }
//     }

//     router.Run(":8080")
// }
