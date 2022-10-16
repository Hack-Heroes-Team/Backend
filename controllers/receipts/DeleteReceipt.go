package receipts

import (
	"github.com/gin-gonic/gin"
	"github.com/mic3b/hack-backend/db"
	"github.com/mic3b/hack-backend/models"
	"net/http"
)

func DeleteReceipt(c *gin.Context) {
	DB := db.Init()
	var deletingReceipt models.Receipt

	err := c.ShouldBindJSON(deletingReceipt)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	DB.Table("receipts").Where("id = ?", deletingReceipt.Id).Delete(&deletingReceipt)
}
