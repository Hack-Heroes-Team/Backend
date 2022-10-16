package receipts

import (
	"github.com/gin-gonic/gin"
	"github.com/mic3b/hack-backend/db"
	"github.com/mic3b/hack-backend/models"
	"net/http"
)

// NOTE: items column struct in db are not changed

func DeleteItem(c *gin.Context) {
	DB := db.Init()
	var deletingItem models.Item

	err := c.ShouldBindJSON(deletingItem)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	DB.Table("items").Where("id = ?", deletingItem.Id).Delete(&deletingItem)
}
