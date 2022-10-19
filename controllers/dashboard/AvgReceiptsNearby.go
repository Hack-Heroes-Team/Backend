package dashboard

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mic3b/hack-backend/db"
	"github.com/mic3b/hack-backend/models"
)

func AvgReceiptsPriceNearby(c *gin.Context) {
	DB := db.Init()

	var userMail models.InputForm
	var receipts []models.Receipt

	if err := c.ShouldBindJSON(&userMail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DB.Table("receipts").Where("owner = ?", userMail.Owner).Find(&receipts)

	for i, v := range receipts {
		var items []models.Item
		DB.Table("items").Where("receiptid = ?", v.Id).Find(&items)
		receipts[i].Items = items
	}

	var items []models.Item

	DB.Table("items").Where("place = ?", userMail.Shop).Find(&items)

	c.JSON(http.StatusOK, gin.H{"Avg": addAllItems(items)})
}

//How Json Should look like:
/*

{
"shop": string,
"owner": string,
}
*/
