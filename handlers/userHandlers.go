package handlers

import (
	"net/http"

	"itx-wabizz/models"
	"itx-wabizz/repositories"

	"github.com/gin-gonic/gin"
)

func HandleInsertUser(c *gin.Context){
	var user models.User;
	if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
        return
    }

    if user.Email == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Email field is required"})
        return
    }

    retrievedUser, err := repositories.UserRepo.Insert(&user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user information"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"User": retrievedUser})
}

func HandleUser(c *gin.Context) {
	var user *models.User
	queryParam := c.Query("email")
	if queryParam == "" {
		// If query parameter is not present
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Query parameter 'email' is missing",
		})
		return
	}
	user, err := repositories.UserRepo.GetUser(queryParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to retrieve user information"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User": user})
}

func HandleGetAllUser(c *gin.Context) {
	var users []models.User
	users, err := repositories.UserRepo.GetAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to retrieve users information"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Users": users})
}

func HandleMakeActive(c *gin.Context) {
	queryParam := c.Query("email")
	if queryParam == "" {
		// If query parameter is not present
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Query parameter 'email' is missing",
		})
		return
	}
	err := repositories.UserRepo.MakeActive(queryParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to make user active"})
		return
	}
	retrievedUser, err := repositories.UserRepo.GetUser(queryParam)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user information"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"User": retrievedUser})
}

func HandleMakeInactive(c *gin.Context) {
	queryParam := c.Query("email")
	if queryParam == "" {
		// If query parameter is not present
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Query parameter 'email' is missing",
		})
		return
	}
	err := repositories.UserRepo.MakeInactive(queryParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to make user inactive"})
		return
	}
	retrievedUser, err := repositories.UserRepo.GetUser(queryParam)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user information"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"User": retrievedUser})
}