package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"itx-wabizz/models"
	"itx-wabizz/repositories"
)

/*
Function: InsertUser

Insert a requested user to the database. User should not be admin and their device token set to empty string.
*/
func InsertUser(c *gin.Context) {
	// Bind the request body sent, it should comply with the UserInsertRequest structure.
	var userInsertRequest models.UserInsertRequest
	if err := c.BindJSON(&userInsertRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request body"})
		return
	}

	// Check if user email is empty or not. It must not be empty.
	if userInsertRequest.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request body"})
		return
	}

	// Create user structure.
	user := models.User{
		Email:       userInsertRequest.Email,
		IsActive:    userInsertRequest.IsActive,
		IsAdmin:     false,
		DeviceToken: "",
	}

	// Insert new user to the database.
	retrievedUser, err := repositories.UserRepo.Insert(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to insert user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"User": retrievedUser})
}

/*
Function: GetUserInfo

Find user in database and send the user information.
*/
func GetUserInfo(c *gin.Context) {
	// Find user email in query parameter.
	queryParam := c.Query("email")

	// If query parameter is not present, deny the request.
	if queryParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Query parameter 'email' is missing"})
		return
	}

	// Get the user from the database.
	var user *models.User
	user, err := repositories.UserRepo.GetUser(queryParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to retrieve user information"})
		return
	}

	// Send information of the user.
	c.JSON(http.StatusOK, gin.H{"User": user})
}

/*
Function: GetAllUserInfo

Find the information of all user in database.
*/
func GetAllUserInfo(c *gin.Context) {
	// Retrieve all user from database.
	var users []models.User
	users, err := repositories.UserRepo.GetAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to retrieve user information"})
		return
	}

	// Send information of all user.
	c.JSON(http.StatusOK, gin.H{"Users": users})
}

/*
Function MakeUserActive

Set status of a spesific user to active.
*/
func MakeUserActive(c *gin.Context) {
	// Find user email in query parameter.
	queryParam := c.Query("email")

	// If query parameter is not present, deny the request.
	if queryParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Query parameter 'email' is missing"})
		return
	}

	// Update user status with that email to active.
	err := repositories.UserRepo.MakeActive(queryParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to update user"})
		return
	}

	// Get user information again.
	retrievedUser, err := repositories.UserRepo.GetUser(queryParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to retrieve user information"})
		return
	}

	// Send information of the user.
	c.JSON(http.StatusOK, gin.H{"User": retrievedUser})
}

/*
Function MakeUserInactive

Set status of a spesific user to inactive.
*/
func MakeUserInactive(c *gin.Context) {
	// Find user email in query parameter
	queryParam := c.Query("email")

	// If query parameter is not present, deny the request.
	if queryParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Query parameter 'email' is missing"})
		return
	}

	// Update user status with that email to inactive.
	err := repositories.UserRepo.MakeInactive(queryParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to update user"})
		return
	}

	// Get user information again.
	retrievedUser, err := repositories.UserRepo.GetUser(queryParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user information"})
		return
	}

	// Send information of the user.
	c.JSON(http.StatusOK, gin.H{"User": retrievedUser})
}
