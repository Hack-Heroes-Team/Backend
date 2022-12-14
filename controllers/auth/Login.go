package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mic3b/hack-backend/db"
	"github.com/mic3b/hack-backend/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var formFromInput models.User
	var Users models.User

	// Here i'am binding data from JSON request with models.User
	if err := c.ShouldBindJSON(&formFromInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Here i'am passing data from FormFromInput to variable User for better clarity

	var User models.User

	User.Name = ""
	User.Surname = " "
	User.Mail = formFromInput.Mail

	User.Password = formFromInput.Password

	// Authorization part

	DB := db.Init()

	err := DB.Table("users").Where("mail = ?", User.Mail).First(&Users).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"authorized": false, "error": "User Didn't find"})
		return
	}

	DB.Table("users").Where("mail = ?", User.Mail).First(&Users)

	checkPassword(c, Users.City, Users.Password, User.Password)

}

// Deserialize Hashed Password and Password from JSON request
func checkPassword(c *gin.Context, mail, userPassword, formPassword string) {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(formPassword))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"authorized": false, "error": "Bad Password, Try again"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"city": mail})
		return
	}
}
