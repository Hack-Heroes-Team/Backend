package auth

import (
	"github.com/mic3b/hack-backend/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mic3b/hack-backend/models"
)

func Delete(c *gin.Context) {
	DB := db.Init()
	var formFromInput models.User

	// Here i'am binding data from JSON request with models.User
	if err := c.ShouldBindJSON(&formFromInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"registered": false, "error": err.Error()})
		return
	}

	DB.Table("users").Where("mail = ?", formFromInput.Mail).Delete(&models.User{})

	c.JSON(http.StatusAccepted, gin.H{"registered": true})
}
