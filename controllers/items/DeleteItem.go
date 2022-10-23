package items

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mic3b/hack-backend/db"
	"github.com/mic3b/hack-backend/models"
)

// NOTE: items column struct in db are not changed

func DeleteItem(c *gin.Context) {
	DB := db.Init()
	var deletingItem models.Item

	err := c.ShouldBindJSON(&deletingItem)

	fmt.Println(deletingItem)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	DB.Table("items").Where("id = ?", deletingItem.Id).Delete(&deletingItem)

	var uItems []models.UniqItem

	DB.Table("uniqitems").Where("place = ? AND name = ?", deletingItem.Place, deletingItem.Name).Find(&uItems)

	if len(uItems) >= 1 {
		DB.Table("uniqitems").Where("place = ? AND name = ?", deletingItem.Place, deletingItem.Name).Delete(&deletingItem)
	}

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
