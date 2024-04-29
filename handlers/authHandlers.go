package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"itx-wabizz/models"
	"itx-wabizz/repositories"
)

func HandleCheckUserLogin(c *gin.Context) {
	var loginToken models.LoginToken
	err := c.BindJSON(&loginToken) 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request body"})
		return
	}

	existingUser, err := repositories.UserRepo.GetUser(loginToken.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "User not found"})
		return
	}

	if !existingUser.IsActive {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Inactive user"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"Message": "Authorized"})
		return
	}
}