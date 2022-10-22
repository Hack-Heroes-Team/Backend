package dashboard

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mic3b/hack-backend/db"
	"github.com/mic3b/hack-backend/models"
)

func PriceForEveryItemInShop(c *gin.Context) {
	DB := db.Init()

	var userMail models.InputForm
	var items []models.UniqItem
	var itemsSec []models.Item

	if err := c.ShouldBindJSON(&userMail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	DB.Table("items").Where("shop = ?", userMail.Shop).Find(&itemsSec)
	DB.Table("uniqitems").Where("shop = ?", userMail.Shop).Find(&items)

	for i, v := range items {
		items[i].AvgPrice = FindForItems(v.Name, itemsSec)
		items[i].AvgPrice = items[i].AvgPrice / FindSecForItems(v.Name, itemsSec)
	}

	c.JSON(http.StatusOK, gin.H{"items": items})
}

func FindForItems(elem string, elems []models.Item) float64 {
	var AvgPrice float64
	for _, v := range elems {
		if elem == v.Name {
			AvgPrice = AvgPrice + v.Price
		}
	}

	return AvgPrice
}

func FindSecForItems(elem string, elems []models.Item) float64 {
	var Counter float64
	for _, v := range elems {
		if elem == v.Name {
			Counter = Counter + 1
		}
	}

	return Counter
}
