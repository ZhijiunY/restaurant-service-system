package controllers

import (
	"fmt"
	"net/http"

	"github.com/ZhijiunY/restaurant-service-system/middleware"
	"github.com/ZhijiunY/restaurant-service-system/models"

	"github.com/gin-gonic/gin"
)

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
