package dashboard

import (
	"github.com/gin-gonic/gin"
	"github.com/mic3b/hack-backend/db"
	"github.com/mic3b/hack-backend/models"
	"net/http"
)

func ShopStats(c *gin.Context) {
	DB := db.Init()

	var receipts []models.Receipt
	var shops []models.Shop
	DB.Table("shops").Find(&shops)

	DB.Table("receipts").Find(&receipts)

	for i, v := range receipts {
		var items []models.Item
		DB.Table("items").Where("receiptid = ?", v.Id).Find(&items)
		receipts[i].Items = items
		total := 0.00
		for _, v := range items {
			total = total + v.Price
		}
		receipts[i].Price = total
	}

	for i, v := range shops {
		shops[i].AvgPrice = Find(v.Name, receipts)
		shops[i].AvgPrice = shops[i].AvgPrice / FindSec(v.Name, receipts)
	}

	c.JSON(http.StatusOK, gin.H{"stats": shops})

}

func Find(elem string, elems []models.Receipt) float64 {
	var AvgPrice float64
	for _, v := range elems {
		if elem == v.Shop {
			AvgPrice = AvgPrice + v.Price
		}
	}

	return AvgPrice
}

func FindSec(elem string, elems []models.Receipt) float64 {
	var Counter float64
	for _, v := range elems {
		if elem == v.Shop {
			Counter = Counter + 1
		}
	}

	return Counter
}
