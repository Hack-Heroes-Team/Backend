package auth

import (
	"github.com/mic3b/hack-backend/db"
	"golang.org/x/crypto/bcrypt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mic3b/hack-backend/models"
)

type MailError struct{}

func (m *MailError) Error() string {
	return "Mail Already Existing"
}

func Register(c *gin.Context) {
	DB := db.Init()
	var formFromInput models.User

	// Here i'am binding data from JSON request with models.User
	if err := c.ShouldBindJSON(&formFromInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"registered": false, "error": err.Error()})
		return
	}

	// Here i'am passing data from FormFromInput to variable User for better clarity

	var User models.User

	hashedPassword, err := HashPassword(c, formFromInput.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"registered": false, "error": err.Error()})
		return
	}

	User.Name = formFromInput.Name
	User.Surname = formFromInput.Surname
	User.Mail = formFromInput.Mail
	User.City = formFromInput.City
	User.Password = hashedPassword

	err = CheckIfMailAlreadyExisting(c, User.Mail)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DB.Table("users").Create(&User)

	c.JSON(http.StatusAccepted, gin.H{"registered": User})
}

func HashPassword(c *gin.Context, password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 4)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return "", err
	}

	return string(hashedPassword), nil
}

func CheckIfMailAlreadyExisting(c *gin.Context, mail string) error {
	DB := db.Init()

	var u models.User

	DB.Where("mail = ?", mail).First(&u)

	if u.Mail == mail {
		return &MailError{}
	}
	return nil
}
