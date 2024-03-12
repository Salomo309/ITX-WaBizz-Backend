package handlers

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"net/http"
	"strings"

	"itx-wabizz/configs"
	"itx-wabizz/models"
	"itx-wabizz/repositories"
)

// Handler to start Google OAuth flow, redirect to Google Login page
func HandleGoogleLogin(c *gin.Context) {
	url := configs.GoogleOauthConfig.AuthCodeURL(configs.OauthStateString, oauth2.AccessTypeOffline)
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
	token, err := configs.GoogleOauthConfig.Exchange(context.Background(), authCode)
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

	if token.RefreshToken != "" {
		refreshToken := models.RefreshToken{
			Google_ID:     googleUserInfo.ID,
			Refresh_Token: token.RefreshToken,
		}
		err = repositories.RefreshTokenRepo.Insert(&refreshToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to save user's refresh token"})
			return
		}
	}

	// Send back user information to the front-end
	userResponseToken := models.UserResponseToken{
		Google_ID: googleUserInfo.ID,
		Token:     token.AccessToken,
		Email:     googleUserInfo.Email,
		Name:      googleUserInfo.Name,
		Picture:   googleUserInfo.Picture,
		Admin:     isAdmin,
	}
	c.JSON(http.StatusOK, userResponseToken)
}

// Handler to logout from application
func HandleLogout(c *gin.Context) {
	token := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")

	client := &http.Client{}
    req, err := http.NewRequest("POST", "https://oauth2.googleapis.com/revoke", nil)
    if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to revoke token"})
        return 
    }
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    
    // Include the access token in the request body
    q := req.URL.Query()
    q.Add("token", token)
    req.URL.RawQuery = q.Encode()

    _, err = client.Do(req)
    if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to revoke token"})
        return 
    }

	// Decode logout request
	var logoutRequestToken models.LogoutRequestToken
	err = json.NewDecoder(c.Request.Body).Decode(&logoutRequestToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Failed to decode request"})
		return
	}

	// Invalidate refresh token in database
	refreshToken := models.RefreshToken{
		Google_ID: logoutRequestToken.Google_ID,
		Refresh_Token: "",
	}
	err = repositories.RefreshTokenRepo.Insert(&refreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to invalidate user's refresh token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Logout successful"})
}

// Handler to give new token to front-end
func HandleNewAccessToken(c *gin.Context) {
	// Decode request sent
	var accessTokenRequest models.AccessTokenRequest
	err := json.NewDecoder(c.Request.Body).Decode(&accessTokenRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Failed to decode request"})
		return
	}

	// Get user refresh token from the database
	refreshToken, err := repositories.RefreshTokenRepo.GetRefreshToken(accessTokenRequest.Google_ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to retrieve user's refresh token"})
		return
	}

	// Check if refresh token is available
	if refreshToken == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Unauthorized user try to get access token"})
		return
	}

	// Create token source using the refresh token
	token := &oauth2.Token{
		RefreshToken: refreshToken.Refresh_Token,
	}
	tokenSource := configs.GoogleOauthConfig.TokenSource(context.Background(), token)

	// Refresh and get new access token
	newAccessToken, err := tokenSource.Token()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to refresh token"})
		return
	}

	// Send new access token back
	accessTokenResponse := models.AccessTokenResponse{
		Google_ID: accessTokenRequest.Google_ID,
		Token:     newAccessToken.AccessToken,
	}
	c.JSON(http.StatusOK, accessTokenResponse)
}
