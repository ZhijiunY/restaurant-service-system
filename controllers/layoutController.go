package controllers

import (
	"net/http"

	"github.com/ZhijiunY/restaurant-service-system/models"
	"github.com/gin-gonic/gin"
)

// Index page
func GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Home website",
	})
}

// Menu page
func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		// tmpl := template.Must(template.ParseFiles("layout.html"))
		// userName := middleware.GetUserSession(c)
		c.HTML(http.StatusOK, "menu.tmpl", gin.H{
			"title": "Menu website",
			// "Name":  userName,
		})
	}
}

// Manager page
// func GetOrder(c *gin.Context) {
func GetOrder() gin.HandlerFunc {
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
			"title":          "Order website",
			"menuItems":      menusItems,
			"CalculateTotal": totalPrice,
		})
	}
}
