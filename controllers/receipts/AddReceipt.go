package receipts

import (
	"github.com/gin-gonic/gin"
	"github.com/mic3b/hack-backend/db"
	"github.com/mic3b/hack-backend/models"
	"net/http"
	"time"
)

func AddReceipts(c *gin.Context) {
	DB := db.Init()
	var newReceipt models.Receipt

	err := c.ShouldBindJSON(newReceipt)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	newReceipt.Date = time.Now()

	DB.Table("receipts").Create(&newReceipt)
}
