package receipts

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mic3b/hack-backend/db"
	"github.com/mic3b/hack-backend/models"
)

func UpdateReceipt(c *gin.Context) {
	DB := db.Init()
	var updatingReceipt models.Receipt

	err := c.ShouldBindJSON(updatingReceipt)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	DB.Table("receipts").Where("id = ?", updatingReceipt.Id).Updates(&updatingReceipt)
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
