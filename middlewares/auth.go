package middlewares

import (
	"net/http"

	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "No Authentication Token"})
		return
	}

	err, userId := utils.VerifyToken(token)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized", "error": err})
		return
	}

	c.Set("userId", userId)
	c.Next()
}
