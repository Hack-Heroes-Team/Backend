package receipts

import (
	"github.com/gin-gonic/gin"
	"github.com/mic3b/hack-backend/db"
	"github.com/mic3b/hack-backend/models"
	"gorm.io/gorm/clause"
	"net/http"
)

func FindLastParagonsFromShop(c *gin.Context) {
	DB := db.Init()

	var shopAndOwnerNames models.InputForm
	var receipts []models.Receipt

	if err := c.ShouldBindJSON(&shopAndOwnerNames); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if shopAndOwnerNames.Owner == "" {
		DB.Table("receipts").Where("shop = ?", shopAndOwnerNames.Shop).Find(&receipts).Order(clause.OrderByColumn{Column: clause.Column{Name: "date"}, Desc: true}).Limit(5)
	} else {
		DB.Table("receipts").Where("shop AND owner = ?", shopAndOwnerNames.Shop, shopAndOwnerNames.Owner).Find(&receipts).Order(clause.OrderByColumn{Column: clause.Column{Name: "date"}, Desc: true}).Limit(5)
	}

	for i, v := range receipts {
		var items []models.Item
		DB.Table("items").Where("receiptid = ?", v.Id).Find(&items)
		receipts[i].Items = items
	}

	c.JSON(http.StatusOK, gin.H{"receipts": receipts})
}
