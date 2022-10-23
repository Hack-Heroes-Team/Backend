package dashboard

import (
	"fmt"
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
		DB.Table("items").Where("receipt_id = ?", v.Id).Find(&items)
		receipts[i].Items = items
		total := 0.00
		for _, v := range items {
			total = total + v.Price
		}
		receipts[i].Price = total
	}

	for i, v := range shops {
		shops[i].AvgPrice, _ = Find(v.Place, receipts)
		_, d := Find(v.Place, receipts)
		shops[i].AvgPrice = shops[i].AvgPrice / float64(len(d))
		if shops[i].AvgPrice == 0 {
			shops[i].AvgPrice = 0.00
		}
	}

	fmt.Println(shops)

	c.JSON(http.StatusOK, gin.H{"stats": shops})

}

func Find(elem string, elems []models.Receipt) (float64, []string) {
	var AvgPrice float64
	var sliceForLen []string
	for _, v := range elems {
		if elem == v.Place {
			sliceForLen = append(sliceForLen, v.Place)
			AvgPrice = AvgPrice + v.Price
		}
	}

	return AvgPrice, sliceForLen
}
