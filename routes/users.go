package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func userSignup(c *gin.Context) {
	var newUser models.User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	err := models.UserSignup(newUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "User not created"})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

func userLogin(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error binding"})
		return
	}

	err, comp, id := models.UserLogin(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if comp {

		token, err := utils.GenerateToken(user.Email, id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		}

		c.JSON(http.StatusFound, gin.H{"message": "Login Successful", "token": token})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Wrong Credentials"})
	}
}
