package routes

import (
	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func login(c *gin.Context) {
	var user models.User
	var err error

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request data."})
		return
	}

	err = user.ValidateUser()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not authenticate user."})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token})
}

func signup(c *gin.Context) {
	var user models.User
	var err error
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request data."})
		return
	}

	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save user."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created!", "user": user})
}
