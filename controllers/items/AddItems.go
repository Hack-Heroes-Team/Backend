package items

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mic3b/hack-backend/db"
	"github.com/mic3b/hack-backend/models"
)

func AddItems(c *gin.Context) {
	DB := db.Init()
	var newItem []models.Item

	err := c.ShouldBindJSON(&newItem)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	fmt.Println(newItem)

	var uniqItems []models.UniqItem

	for _, v := range newItem {
		DB.Table("uniqitems").Where("place = ? ", v.Place).Find(&uniqItems)
	}

	unique := difference(newItem, uniqItems)
	fmt.Println(unique)

	for _, v := range unique {
		var uItem models.UniqItem
		uItem.Shop = v.Shop
		uItem.Name = v.Name
		uItem.City = v.City
		uItem.Place = v.Place
		uItem.AvgPrice = v.Price
		DB.Table("uniqitems").Create(&uItem)
	}

	for _, v := range newItem {
		DB.Table("items").Updates(&v)
	}

	c.JSON(http.StatusAccepted, gin.H{"result": "Added"})
}

func difference(a []models.Item, b []models.UniqItem) []models.Item {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x.Name] = struct{}{}
	}
	var diff []models.Item
	for _, x := range a {
		if _, found := mb[x.Name]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

//How Json Should look like:
/*

{
"id": int,
"receiptid": int,
"owner": string,
"name": string,
"place": string,
"price": float64
}
*/
