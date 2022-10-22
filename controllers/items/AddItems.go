package items

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mic3b/hack-backend/db"
	"github.com/mic3b/hack-backend/models"
)

// NOTE: items column struct in db are not changed

func AddItems(c *gin.Context) {
	DB := db.Init()
	var newItem models.Item

	err := c.ShouldBindJSON(newItem)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	fmt.Println(newItem)

	var uniqItems []models.UniqItem
	var uniqItem models.UniqItem
	DB.Table("uniqitems").Where("place = ? ", newItem.Place).Find(&uniqItems)

	for _, v := range uniqItems {
		if v.Name == newItem.Name {
			fmt.Println("Matching")
		} else {
			uniqItem = models.UniqItem{Name: newItem.Name, Shop: newItem.Shop, Place: newItem.Place, City: newItem.City}
			DB.Table("uniqitems").Create(&uniqItem)
		}
	}

	DB.Table("items").Create(&newItem)
	c.JSON(http.StatusAccepted, gin.H{"result": "Added"})
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
