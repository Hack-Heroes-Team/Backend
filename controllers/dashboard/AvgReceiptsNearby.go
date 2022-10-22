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

	if err := c.ShouldBindJSON(&userMail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
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