package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mic3b/hack-backend/controllers/auth"
)

func main() {
	router := gin.Default()

	// Auth part
	router.POST("/login", auth.Login)

	router.Run(":" + "8080")
}
