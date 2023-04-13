package controllers

import (
	"net/http"

	"github.com/ZhijiunY/restaurant-service-system/models"
	"github.com/ZhijiunY/restaurant-service-system/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

// var user models.User

// get user
func GetUser(c *gin.Context) {
	// db := utils.DB

	// // 从URL参数中获取用户ID
	// userID := c.Param("id")

	// // 在数据库中查找指定ID的用户
	// if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	// 	return
	// }

	// c.JSON(http.StatusOK, user)

	var user models.User
	result := utils.DB.Take(&user).Error //开头就写最简单的，获取一条记录，不指定排序

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": result,
	})
}

func CreateUser(c *gin.Context) {

	// users := models.User{
	// 	Model:      gorm.Model{},
	// 	ID:         user.ID,
	// 	Name:       "Simba",
	// 	Password:   123456,
	// 	Email:      "simba@gmail.com",
	// 	Created_at: time.Time{},
	// 	Updated_at: time.Time{},
	// }

	// result := utils.DB.Create(&user)
	// if result.Error != nil {
	// 	// 如果出現錯誤，進行相應處理
	// 	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
	// 	return
	// }

	// if err := c.BindJSON(&users); err != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// c.JSON(http.StatusOK, users)

}

func DeleteUser(c *gin.Context) {
	// var user models.User
	// utils.DB.Where("id = ?", c.Param("id")).Delete(&user)
	// c.JSON(200, &user)
}

func UpdateUser(c *gin.Context) {
	// SON(200, &user)

}

// func ChackUserPassword(c *gin.Context) {

// }
