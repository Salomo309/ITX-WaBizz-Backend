package handlers

import (
	"itx-wabizz/models"
	"itx-wabizz/repositories"
	"net/http"
	"strconv"
	// "encoding/json"
	// "bytes"

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

func HandleSendMessage(c *gin.Context) {
	// Bind request body to ChatMessage struct
	var chatMessage models.Chat
	if err := c.BindJSON(&chatMessage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Create Chat object
	chat := &models.Chat{
		ChatID:		0,
		Email:      chatMessage.Email,
		ChatroomID: chatMessage.ChatroomID,
		Timendate:  chatMessage.Timendate,
		IsRead:     chatMessage.IsRead,
		Content:    chatMessage.Content,
		MessageType: chatMessage.MessageType,
	}

	// Insert chat message into database
	if err := repositories.ChatRepo.CreateChat(chat); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message"})
		return
	}

	// Send the message to http://localhost:8081/receive
	// messageBody, err := json.Marshal(chat)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal message body"})
	// 	return
	// }

	// resp, err := http.Post("http://localhost:8081/receive", "application/json", bytes.NewBuffer(messageBody))
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message to receive endpoint"})
	// 	return
	// }
	// defer resp.Body.Close()

	// if resp.StatusCode != http.StatusOK {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected response from receive endpoint"})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
}


func HandleReceiveMessage(c *gin.Context) {
	// Bind request body to ChatMessage struct
	var chatMessage models.Chat
	if err := c.BindJSON(&chatMessage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Set ChatID to 0
	chatMessage.ChatID = 0

	// Insert received chat message into database
	if err := repositories.ChatRepo.CreateChat(&chatMessage); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save received message"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message received and saved successfully"})
}

