package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/ZhijiunY/restaurant-service-system/models"
	"github.com/ZhijiunY/restaurant-service-system/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/skip2/go-qrcode"
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
		session := sessions.Default(c)
		userName := session.Get("Name")
		if userName == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User Name not found"})
			return
		}
		// userName, _ := c.Get("userName")

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

		for _, menuItem := range menusItems {
			result := utils.DB.Create(&menuItem)
			if result.Error != nil {
				log.Printf("插入菜單錯誤: %v\n", result.Error)
			}
		}

		// 按 FoodType 分類的菜單項
		categorizedMenu := make(map[string][]models.Menu)
		for _, item := range menusItems {
			categorizedMenu[item.FoodType] = append(categorizedMenu[item.FoodType], item)
		}

		// 將分類後的菜單和總價傳遞給模板
		c.HTML(http.StatusOK, "order.tmpl", gin.H{
			"title":           "Order Website",
			"categorizedMenu": categorizedMenu,
			"userName":        userName,
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
	}
}

func (oc *OrderController) ShowOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userName := session.Get("Name")
		if userName == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User Name not found"})
			return
		}
		// 從Redis 獲取數據
		jsonData, err := oc.RedisClient.Get(oc.Ctx, "orderItems").Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve order items from Redis"})
			return
		}

		var OrderItems []models.OrderItem
		// 解析 JSON 數據到 orderItems
		err = json.Unmarshal([]byte(jsonData), &OrderItems)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal order items"})
			return
		}

		var totalAmount float64
		for _, item := range OrderItems {
			totalAmount += item.Price * float64(item.Quantity)
		}

		c.HTML(http.StatusOK, "show-orders.tmpl", gin.H{
			"OrderItems":  OrderItems,
			"TotalAmount": totalAmount,
			"userName":    userName,
		})
	}
}

func (oc *OrderController) GenerateOrderQRCode() gin.HandlerFunc {
	return func(c *gin.Context) {

		session := sessions.Default(c)
		userID := session.Get("ID")
		userName := session.Get("Name")

		// 檢查用戶 ID 和 Name 是否存在
		if userID == nil || userName == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未找到用戶信息或用戶未登入"})
			return
		}

		// 從 Redis 獲取 orderItems 數據，這部分代碼保持不變
		jsonData, err := oc.RedisClient.Get(oc.Ctx, "orderItems").Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to retrieve order items from Redis",
			})
			return
		}

		var orderItems []models.OrderItem
		err = json.Unmarshal([]byte(jsonData), &orderItems)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to unmarshal order items",
			})
			return
		}

		// 使用查詢到的用戶資訊和訂單項目生成 QR 碼數據
		simplifiedData := make([]map[string]interface{}, len(orderItems))
		for i, item := range orderItems {
			simplifiedData[i] = map[string]interface{}{
				"UserID":   userID,
				"UserName": userName,
				"品名":       item.Name,
				"數量":       item.Quantity,
				"價格":       item.Price,
			}
		}
		simplifiedJSON, _ := json.Marshal(simplifiedData)

		// 生成 QR 碼
		qrCode, err := qrcode.Encode(string(simplifiedJSON), qrcode.Medium, 256)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to generate QR code",
			})
			return
		}

		// 直接返回 QR 碼圖片
		c.Writer.Header().Set("Content-Type", "image/png")
		c.Writer.Write(qrCode)
	}
}
