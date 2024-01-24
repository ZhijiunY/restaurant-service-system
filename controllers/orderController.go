package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ZhijiunY/restaurant-service-system/models"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
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

		// 按 FoodType 分類的菜單項
		categorizedMenu := make(map[string][]models.Menu)
		for _, item := range menusItems {
			categorizedMenu[item.FoodType] = append(categorizedMenu[item.FoodType], item)
		}

		// 計算總價格
		var totalPrice float64
		for _, item := range menusItems {
			totalPrice += item.Price
		}

		// 將分類後的菜單和總價傳遞給模板
		c.HTML(http.StatusOK, "order.tmpl", gin.H{
			"title":           "Order Website",
			"categorizedMenu": categorizedMenu,
		})
	}
}

func (oc *OrderController) SubmitOrder() gin.HandlerFunc {
	return func(c *gin.Context) {

		var orderItems []models.OrderItem
		if err := c.BindJSON(&orderItems); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 假設 rdb 是一個已經配置的 Redis 客戶端實例
		jsonData, err := json.Marshal(orderItems)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal order items"})
			return
		}

		err = oc.RedisClient.Set(oc.Ctx, "orderItems", jsonData, 0).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save order items to Redis"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":    "Order submitted successfully",
			"orderItems": orderItems,
		})

	}
}
