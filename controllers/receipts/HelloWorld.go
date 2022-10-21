package receipts

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {
	

	c.JSON(http.StatusOK, gin.H{"Test": "HelloWorld"})

}