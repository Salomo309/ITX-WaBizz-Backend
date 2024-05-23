package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"itx-wabizz/models"
	"itx-wabizz/repositories"
)

/*
Function: CheckUserLogin

Check if user that want to login is in the database and is an active user (Activeness determined by admin).
If true, then store the user device registration token (for FCM) and accept the user.
If false, then deny the user.
*/
func CheckUserLogin(c *gin.Context) {
	// Bind the request body sent, it should comply with the LoginRequest structure.
	var loginRequest models.LoginRequest
	err := c.BindJSON(&loginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request body"})
		return
	}

	// Find the requested user in the database. If user is not found, deny the user.
	existingUser, err := repositories.UserRepo.GetUser(loginRequest.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "User not found"})
		return
	}

	// Check user activeness.
	if !existingUser.IsActive {
		// If not active, deny the user.

		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Inactive user"})
		return
	} else {
		// If active, store user device registration token and accept the user.

		err = repositories.UserRepo.UpdateDeviceToken(loginRequest.Email, loginRequest.DeviceToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to check user login"})
			return
		}

		c.JSON(http.StatusOK, models.LoginResponse{
			Message: "Authorized",
			IsAdmin: existingUser.IsAdmin})
		return
	}
}

/*
Function: Logout

Delete user device token from database to avoid sending FCM notifications
*/
func Logout(c *gin.Context) {
	// Bind the request body sent, it should comply with the LogoutRequest structure.
	var logoutRequest models.LogoutRequest
	err := c.BindJSON(&logoutRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request body"})
		return
	}

	// Make user device token an empty string again
	err = repositories.UserRepo.UpdateDeviceToken(logoutRequest.Email, "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to logout"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Success"})
}
