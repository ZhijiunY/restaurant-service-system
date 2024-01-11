package controllers

import (
	"fmt"
	"net/http"

	"github.com/ZhijiunY/restaurant-service-system/middleware"
	"github.com/ZhijiunY/restaurant-service-system/models"

	"github.com/gin-gonic/gin"
)

// Order page
// func GetOrder(c *gin.Context) {
func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		menusItems := []models.Menu{
			{FoodType: "主食", Name: "北京烤鴨", Description: "**", Price: 380},
			{FoodType: "主食", Name: "意式千層麵", Description: "*", Price: 240},
			{FoodType: "主食", Name: "日式壽司", Description: "**", Price: 185},
			{FoodType: "點心", Name: "提拉米蘇", Description: "**", Price: 80},
			{FoodType: "點心", Name: "馬卡龍", Description: "*", Price: 90},
			{FoodType: "點心", Name: "芝士蛋糕", Description: "**", Price: 100},
			{FoodType: "飲料", Name: "珍珠奶茶", Description: "****", Price: 50},
			{FoodType: "飲料", Name: "抹茶拿鐵", Description: "*****", Price: 70},
			{FoodType: "飲料", Name: "鮮榨果汁", Description: "*****", Price: 60},
		}

		// 按 FoodType 分类的菜单项
		categorizedMenu := make(map[string][]models.Menu)
		for _, item := range menusItems {
			categorizedMenu[item.FoodType] = append(categorizedMenu[item.FoodType], item)
		}

		// 计算总价格
		var totalPrice float64
		for _, item := range menusItems {
			totalPrice += item.Price
		}

		// 将分类后的菜单和总价传递给模板
		c.HTML(http.StatusOK, "order.tmpl", gin.H{
			"title":           "Order Website",
			"categorizedMenu": categorizedMenu,
			"CalculateTotal":  totalPrice,
		})
	}
}

// get order page and items
func OrderAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		menusItems := []models.Menu{
			{Name: "Burger", Description: "good", Price: 10.0},
			{Name: "Apple", Description: "good", Price: 20.0},
			{Name: "Banana", Description: "good", Price: 30.0},
			{Name: "Pizza", Description: "good", Price: 40.0},
			{Name: "Salads", Description: "good", Price: 50.0},
			{Name: "Cake", Description: "good", Price: 60.0},
		}

		// 计算总价格
		var totalPrice float64
		for _, menu := range menusItems {
			totalPrice += menu.Price
		}

		c.HTML(http.StatusOK, "order.tmpl", gin.H{
			"menuItems":      menusItems,
			"CalculateTotal": totalPrice,
		})
	}

}

// confrom prices
func ConfirmPrice(c *gin.Context) {
	var order models.Order

	hasSession := middleware.HasSession(c)
	if hasSession {
		c.Redirect(http.StatusSeeOther, "/menu")
		fmt.Println("Authentication login")
	} else {
		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var totalPrice float64
		for _, Menu := range order.OrderItems {
			totalPrice += Menu.Menu.Price * float64(Menu.Quantity)
		}

		c.JSON(http.StatusOK, gin.H{"TotalPrice": totalPrice})
	}
}
