package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"itx-wabizz/models"
	"itx-wabizz/repositories"
)

/*
Function: GetChatroomList

Find all chatroom list to be viewed by the user. This include the last chat in that chatroom.
*/
func GetChatroomList(c *gin.Context) {
	// Retrieve all chatroom list information from database
	var chatlists []models.ChatList
	chatlists, err := repositories.ChatlistRepo.GetChatroomList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to retrieve chatlist information"})
		return
	}

	// Send the information back as response
	c.JSON(http.StatusOK, gin.H{"ChatList": chatlists})
}

/*
Function: SearchChatroomByContact

Retrieve all chatroom with customer name that has matching substring
*/
func SearchChatroomByContact(c *gin.Context) {
	// Find search query or keyword in query parameter.
	queryParam := c.Query("keyword")

	// If query parameter is not present, deny the request.
	if queryParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Query parameter 'keyword' is missing"})
		return
	}

	// Retrieve all chatroom with matching condition
	var chatlists []models.ChatList
	chatlists, err := repositories.ChatlistRepo.SearchChatroomByContact(queryParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to retrieve chatroom information"})
		return
	}

	// Send information back as response
	c.JSON(http.StatusOK, gin.H{"ChatList": chatlists})
}

/*
Function: SearchChatroomByContact

Retrieve all chatroom with chats content that has matching substring
*/
func SearchChatroomByMessage(c *gin.Context) {
	// Find search query or keyword in query parameter.
	queryParam := c.Query("keyword")

	// If query parameter is not present, deny the request.
	if queryParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Query parameter 'keyword' is missing"})
		return
	}

	// Retrieve all chatroom with matching condition
	var chatlists []models.MessageSearchResult
	chatlists, err := repositories.ChatlistRepo.SearchChatroomByMessage(queryParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to retrieve chatroom information"})
		return
	}

	// Send information back as response
	c.JSON(http.StatusOK, gin.H{"ChatList": chatlists})
}
