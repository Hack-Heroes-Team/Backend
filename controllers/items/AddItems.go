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

	var uniqItems []models.UniqItem
	var uniqItem models.UniqItem
	DB.Table("uniqitems").Where("shop = ? ", newItem.Shop).Find(&uniqItems)

	for _, v := range uniqItems {
		if v.Name == newItem.Shop {
			fmt.Println("Matching")
		} else {
			uniqItem = models.UniqItem{Name: newItem.Shop}
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
