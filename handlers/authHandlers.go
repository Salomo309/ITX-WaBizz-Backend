package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"itx-wabizz/configs"
	"itx-wabizz/models"
	"itx-wabizz/repositories"
)

// Handler to start Google OAuth flow, redirect to Google Login page
func HandleGoogleLogin(c *gin.Context) {
	url := configs.GoogleOauthConfig.AuthCodeURL(configs.OauthStateString)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// Handler for Google callback after user log in
func HandleGoogleCallback(c *gin.Context) {
	// Check state string to avoid CRSF attack
	state := c.Query("state")
	if state != configs.OauthStateString {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid OAuth State"})
		return
	}

	// Exchange code for Google token
	authCode := c.Query("code")
	token, err := configs.GoogleOauthConfig.Exchange(c, authCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to exchange code"})
		return
	}

	// Use token to get response about user information
	res, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to get user information"})
		return
	}
	defer res.Body.Close()

	// Decode response from Google to get user information
	var googleUserInfo models.GoogleUserInfo
	err = json.NewDecoder(res.Body).Decode(&googleUserInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to decode user information"})
		return
	}

	// Check if existingUser information already store in the database
	var existingUser *models.User
	var isAdmin bool
	existingUser, err = repositories.UserRepo.GetUserByGoogleID(googleUserInfo.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to retrieve user information"})
		return
	}

	// If user information does not exist, save it to database for future use and reference
	if existingUser == nil {
		isAdmin = false
		user := models.User{
			User_ID:   0,
			Google_ID: googleUserInfo.ID,
			Email:     googleUserInfo.Email,
			Name:      googleUserInfo.Name,
			Picture:   googleUserInfo.Picture,
			Admin:     isAdmin,
		}
		err = repositories.UserRepo.Insert(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to save user information"})
			return
		}
	} else {
		isAdmin = existingUser.Admin
	}

	// Send back user information to the front-end
	userResponseToken := models.UserResponseToken{
		Token:   token.AccessToken,
		Email:   googleUserInfo.Email,
		Name:    googleUserInfo.Name,
		Picture: googleUserInfo.Picture,
		Admin:   isAdmin,
	}
	c.JSON(http.StatusOK, userResponseToken)
}
