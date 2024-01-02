package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 处理 GET 请求
func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, Gin!",
	})
}

// 处理 POST 请求
func createUserHandler(c *gin.Context) {
	// 在实际应用中，可以从请求中获取数据，进行处理，返回响应等
	// 这里只是一个简单的示例
	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
	})
}
