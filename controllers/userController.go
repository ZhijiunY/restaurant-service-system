package controllers

import (
	"github.com/ZhijiunY/restaurant-service-system/database"
	"github.com/ZhijiunY/restaurant-service-system/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	c.HTML(200, "index.tmpl")
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
