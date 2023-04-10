package controllers

import (
	"net/http"

	"github.com/ZhijiunY/restaurant-service-system/database"
	"github.com/ZhijiunY/restaurant-service-system/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

var user models.User

// get user
func GetUser(c *gin.Context) {
	id := c.Param("id")
	// 使用GORM查詢用户
	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUsers(c *gin.Context) {

	user := models.User{
		Model:    gorm.Model{},
		ID:       "1",
		UserName: Simba,
		Password: 123456,
		Email:    "simba@gmail.com",
	}

	result := database.DB.Create(&user)

	users := models.User{}
	err := c.BindJSON(&users)
	if err != nil {
		c.String(400, "Error:%s", err.Error())
		return
	}
	c.JSON(http.StatusOK, users)

}

func DeleteUsers(c *gin.Context) {
	var user models.User
	database.DB.Where("id = ?", c.Param("id")).Delete(&user)
	c.JSON(200, &user)
}

func UpdateUsers(c *gin.Context) {
	var user models.User
	database.DB.Where("id = ?", c.Param("id")).First(&user)
	c.BindJSON(&user)
	database.DB.Save(&user)
	c.JSON(200, &user)

}

// func ChackUserPassword(c *gin.Context) {

// }
