package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ZhijiunY/restaurant-service-system/models"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var (
	ctx = context.Background()
)

type OrderController struct {
	RedisClient *redis.Client
	Ctx         context.Context
}

func NewOrderController(redisClient *redis.Client, ctx context.Context) *OrderController {
	return &OrderController{
		RedisClient: redisClient,
		Ctx:         ctx,
	}
}

type OrderItem struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

// Order page
// func GetOrder(c *gin.Context) {
func (oc *OrderController) GetOrder() gin.HandlerFunc {
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

		// 按 FoodType 分類的菜单项
		categorizedMenu := make(map[string][]models.Menu)
		for _, item := range menusItems {
			categorizedMenu[item.FoodType] = append(categorizedMenu[item.FoodType], item)
		}

		// 計算總價格
		var totalPrice float64
		for _, item := range menusItems {
			totalPrice += item.Price
		}

		// 将分类后的菜单和总价传递给模板
		c.HTML(http.StatusOK, "order.tmpl", gin.H{
			"title":           "Order Website",
			"categorizedMenu": categorizedMenu,
			// "CalculateTotal":  totalPrice,
		})
	}
}

func (oc *OrderController) SubmitOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var orderItems []struct {
			Name     string  `form:"name"`
			Quantity int     `form:"quantity"`
			Price    float64 `form:"price"`
		}

		if err := c.ShouldBind(&orderItems); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		jsonData, err := json.Marshal(orderItems)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to serialize data"})
			return
		}

		err = oc.RedisClient.Set(ctx, "orderData", jsonData, 0).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store data in Redis"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Order stored in Redis"})
	}
}
