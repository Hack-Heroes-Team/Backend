package receipts

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mic3b/hack-backend/db"
	"github.com/mic3b/hack-backend/models"
)

func FindReceipt(c *gin.Context) {
	DB := db.Init()

	var userMail models.Receipt
	var receipts []models.Receipt

	if err := c.ShouldBindJSON(&userMail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(userMail)

	DB.Table("receipts").Where("id = ?", userMail.Id).Find(&receipts)

	for i, v := range receipts {
		var items []models.Item
		DB.Table("items").Where("receipt_id = ?", v.Id).Find(&items)
		receipts[i].Items = items
		var sum float64
		for _, v := range items {
			sum = sum + v.Price
		}
		receipts[i].Price = sum
	}

	c.JSON(http.StatusOK, gin.H{"receipt": receipts[1]})
}
