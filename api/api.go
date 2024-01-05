package api

import (
	"DigBGM/database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	rdb *redis.Client
)

// 处理 GET 请求
func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, Gin!",
	})
}

// 注册 POST 请求
func RegisterHandler(c *gin.Context) {
	var newUser database.User
	// 解析请求体中的JSON数据
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 检查数据库中是否已存在相同的用户名
	var existingUser database.User
	result := db.Where("username = ?", newUser.Username).First(&existingUser)
	if result.Error == nil {
		// 用户名已存在，返回错误响应
		c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		return
	}
	// 在数据库中插入新用户
	result = db.Create(&newUser)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("User %s registered successfully", newUser.Username)})
}

func RedisSet(c *gin.Context) {
	// 设置Redis键值对
	err := rdb.Set(c, "example_key", "example_value", 0).Err()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to set Redis key",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Redis key set successfully",
	})
}

func RedisHandel(c *gin.Context) {
	// 获取Redis键值对
	val, err := rdb.Get(c, "example_key").Result()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to get Redis key",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Redis key value: " + val,
	})
}
