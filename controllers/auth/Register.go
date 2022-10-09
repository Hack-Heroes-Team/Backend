package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mic3b/hack-backend/models"
)

// DURING WORK
func Register(c *gin.Context) {
	var formFromInput models.User

	// Here i'am binding data from JSON request with models.User
	if err := c.ShouldBindJSON(&formFromInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Here i'am passing data from FormFromInput to variable User for better clarity

	var User models.User

	User.Name = formFromInput.Name
	User.Surname = formFromInput.Surname
	User.Mail = formFromInput.Mail
	User.Password = formFromInput.Password
}
