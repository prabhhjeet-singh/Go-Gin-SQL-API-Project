package routes

import (
	"net/http"

	"example.com/rest-api/models"
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

	err, comp := models.UserLogin(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if comp {
		c.JSON(http.StatusFound, gin.H{"message": "Login Successful"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Wrong Credentials"})
	}
}
