package controllers

import (
	"github.com/ZhijiunY/restaurant-service-system/database"
	"github.com/ZhijiunY/restaurant-service-system/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {

	// find all users
	users := []models.User{}
	database.DB.Find(&users)

	c.JSON(200, &users)
}

func CreateUsers(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	database.DB.Create(&user)
	c.JSON(200, &user)
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
