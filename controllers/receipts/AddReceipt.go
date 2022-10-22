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
