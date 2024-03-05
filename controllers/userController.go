package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

// var user models.User

// get user
func GetUsers(c *gin.Context) {

	// var user models.User
	// result := utils.DB.Take(&user).Error

	// c.JSON(http.StatusOK, gin.H{
	// 	"code": 1,
	// 	"data": result,
	// })
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
