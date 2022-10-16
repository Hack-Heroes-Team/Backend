package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mic3b/hack-backend/controllers/auth"
	"github.com/mic3b/hack-backend/controllers/items"
	"github.com/mic3b/hack-backend/controllers/receipts"
)

func main() {
	router := gin.Default()

	// Auth part:
	router.POST("/login", auth.Login)
	router.POST("/register", auth.Register)

	// Receipt part:
	router.POST("/receiptsForUser", receipts.FindAllReceiptsForUser)
	router.POST("/addReceipt", receipts.AddReceipts)
	router.POST("/deleteReceipt", receipts.DeleteReceipt)
	router.POST("/findLastReceiptsFromShop", receipts.FindLastReceiptsFromShop)
	router.POST("/updateReceipt", receipts.UpdateReceipt)

	// Items part:
	router.POST("/updateItem", items.UpdateItem)
	router.POST("/addItem", items.AddItems)
	router.POST("/deleteItem", items.DeleteItem)

	// Run Server
	err := router.Run(":" + "8080")

	if err != nil {
		return
	}
}
