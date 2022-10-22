package receipts

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mic3b/hack-backend/db"
	"github.com/mic3b/hack-backend/models"
)

func AddReceipts(c *gin.Context) {
	DB := db.Init()
	var newReceipt models.Receipt

	err := c.ShouldBindJSON(&newReceipt)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	fmt.Println(newReceipt)

	newReceipt.Date = time.Now()

	var shops []models.Shop
	var shop models.Shop
	DB.Table("shops").Find(&shops)
	fmt.Println(shops)

	for _, v := range shops {
		if v.Place == newReceipt.Place {
			fmt.Println("Matching")
		}
		shop = models.Shop{Name: newReceipt.Shop}
		DB.Table("shops").Create(&shop)

	}

	DB.Table("receipts").Create(&newReceipt)
	c.JSON(http.StatusAccepted, gin.H{"result": "Added"})

}

//How Json Should look like:
/*

{
"id": int,
"name": string,
"shop": string,
"owner": string,
"date": timestamp
}
*/
