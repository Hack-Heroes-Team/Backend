package receipts

import (
	"github.com/gin-gonic/gin"
	"github.com/mic3b/hack-backend/db"
	"github.com/mic3b/hack-backend/models"
	"net/http"
)

// NOTE: items column struct in db are not changed

func UpdateItem(c *gin.Context) {
	DB := db.Init()
	var updatingItem models.Item

	err := c.ShouldBindJSON(updatingItem)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	DB.Table("items").Where("id = ?", updatingItem.Id).Updates(&updatingItem)
}
