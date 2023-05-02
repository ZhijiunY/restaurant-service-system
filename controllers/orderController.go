package controllers

import (
	"fmt"
	"net/http"

	"github.com/ZhijiunY/restaurant-service-system/middleware"
	"github.com/ZhijiunY/restaurant-service-system/models"

	"github.com/gin-gonic/gin"
)

func OrderAction(c *gin.Context) {
	// TotalPrice := models.Price * models.count
	menus := []models.Menu{
		{Name: "Hambuger", Description: "good", Price: 10},
		{Name: "Apple", Description: "good", Price: 20},
		{Name: "Banana", Description: "good", Price: 30},
		{Name: "Pizza", Description: "good", Price: 40},
		{Name: "Salads", Description: "good", Price: 50},
		{Name: "Cake", Description: "good", Price: 60},
	}

	c.HTML(http.StatusOK, "order.tmpl", gin.H{
		"Menus": menus,
	})
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
