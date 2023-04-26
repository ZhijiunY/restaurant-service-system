package controllers

import (
	"net/http"

	"github.com/ZhijiunY/restaurant-service-system/models"
	"github.com/gin-gonic/gin"
)

func OrderAction(c *gin.Context) {
	menus := []models.Menu{
		{Name: "Hambuger", Description: "good", Price: 10},
		{Name: "Apple", Description: "good", Price: 20},
		{Name: "Banana", Description: "good", Price: 30},
		{Name: "Pizza", Description: "good", Price: 40},
		{Name: "Salads", Description: "good", Price: 50},
		{Name: "Cake", Description: "good", Price: 60},
	}
	orders := []models.Order{
		{TotalPrice: 100},
	}

	data := struct {
		Menus  []models.Menu
		Orders []models.Order
	}{
		Menus:  menus,
		Orders: orders,
	}

	c.HTML(http.StatusOK, "order.tmpl", gin.H{
		"Menus":  menus,
		"Orders": orders,
		"data":   data,
	})
}
