package receipts

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mic3b/hack-backend/db"
	"github.com/mic3b/hack-backend/models"
)

func DeleteReceipt(c *gin.Context) {
	DB := db.Init()
	var deletingReceipt models.Receipt
	var deletingReceipts []models.Receipt

	err := c.ShouldBindJSON(&deletingReceipt)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Println(deletingReceipt)

	DB.Table("receipts").Where("id = ?", deletingReceipt.Id).Find(&deletingReceipts)
	DB.Table("receipts").Where("id = ?", deletingReceipt.Id).Delete(&deletingReceipt)
	c.JSON(http.StatusOK, gin.H{"deleted": deletingReceipts})
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
