package handlers

import (
	"itx-wabizz/models"
	"itx-wabizz/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleGetChatroom(c *gin.Context) {
	chatroomID := c.Query("chatroomID")
	if chatroomID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Query parameter 'chatroomID' is missing"})
		return
	}
	
	intChatroomID, err := strconv.Atoi(chatroomID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Query parameter 'chatroomID' is not a number"})
		return
	} 

	var chats []models.Chat
	chats, err = repositories.ChatRepo.GetChats(intChatroomID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to retrieve chats information"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Chats": chats})
}

func HandleReceiveMessage(c *gin.Context) {
	
}

func HandleSendMessage(c *gin.Context) {
	
}