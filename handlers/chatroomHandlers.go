package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"github.com/gin-gonic/gin"

	// "os"

	"itx-wabizz/models"
	"itx-wabizz/repositories"
)

/*
Function: GetChatroom

Retrieve all chats of a spesific chatroom
*/
func GetChatroom(c *gin.Context) {
	// Find chatroom ID in query parameter.
	chatroomID := c.Query("chatroomID")

	// If query parameter is not present, deny the request.
	if chatroomID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Query parameter 'chatroomID' is missing"})
		return
	}

	// Parse query parameter into integer
	intChatroomID, err := strconv.Atoi(chatroomID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Query parameter 'chatroomID' is not a number"})
		return
	}

	// Mark all chats in the chatroom as read
	if err := repositories.ChatRepo.MarkAllChatsAsRead(intChatroomID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to mark all chats as read"})
		return
	}

	// Retrieve the chats of that chatroom
	var chats []models.Chat
	chats, err = repositories.ChatRepo.GetChats(intChatroomID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to retrieve chats information"})
		return
	}

	// Send information back as response
	c.JSON(http.StatusOK, gin.H{"Chats": chats})
}

func HandleSendMessage(c *gin.Context) {
	chatJSON := c.PostForm("chatJSON")

	// Bind request body to ChatMessage struct
	var chat models.Chat
	if err := json.Unmarshal([]byte(chatJSON), &chat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid JSON in form-data"})
		return
	}

	// Check message type and content validity
	if chat.MessageType == "text"{

		if err := repositories.ChatRepo.CreateChat(&chat); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to send message"})
			return
		}

	} else {
		if chat.MessageType == "photo" ||  chat.MessageType == "video" || chat.MessageType == "file" {

			fileHeader, err := c.FormFile("file")
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"Error": "File is required"})
				return
			}

			// Save file and get the file path
			filePath, err := SaveFile(fileHeader)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to save file"})
				return
			}

			chat.Content = filePath

			// Insert chat (text type) message into database
			if err := repositories.ChatRepo.CreateChat(&chat); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to send message"})
				return
			}
		}
	}

	existingChatroom, err := repositories.ChatlistRepo.GetChatroomByID(chat.ChatroomID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to retrieve chatroom information"})
		return
	}

	infobipMessage := models.Message{
		From:         "0999-9999-9999",
		To:           existingChatroom.CustomerPhone,
		MessageID:    "AB09JC",
		Content:      models.Content{Text: chat.Content},
		CallbackData: "",
		NotifyURL:    "",
		URLOptions: models.URLOptions{
			ShortenURL:     true,
			TrackClicks:    true,
			TrackingURL:    "",
			RemoveProtocol: true,
			CustomDomain:   "",
		},
	}

	requestBody, _ := json.Marshal(infobipMessage)
	http.Post("http://host.docker.internal:8081/receive", "application/json", bytes.NewBuffer(requestBody))

	c.JSON(http.StatusOK, gin.H{"Message": chat})
}

func HandleReceiveMessage(c *gin.Context) {
	// var infobipMessage models.Message
	// if err := c.BindJSON(&infobipMessage); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request body"})
	// 	return
	// }

	// customerPhone := infobipMessage.From
	// existingChatroom, err := repositories.ChatlistRepo.GetChatroomByPhone(customerPhone)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to retrieve chatroom information"})
	// 	return
	// }

	var infobipMessage models.ReceivedMessage
	if err := c.BindJSON(&infobipMessage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request body"})
		return
	}

	customerPhone := infobipMessage.Results[0].From
	existingChatroom, err := repositories.ChatlistRepo.GetChatroomByPhone(customerPhone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to retrieve chatroom information"})
		return
	}

	var chatroomID int
	if existingChatroom != nil {
		chatroomID = existingChatroom.ChatroomID
	} else {
		newChatroom := models.Chatroom{
			ChatroomID:    0,
			CustomerPhone: customerPhone,
			CustomerName:  "",
		}
		err := repositories.ChatlistRepo.Insert(&newChatroom)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to save chatroom information"})
		}

		existingChatroom, err = repositories.ChatlistRepo.GetChatroomByPhone(customerPhone)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to retrieve chatroom information"})
			return
		}

		chatroomID = existingChatroom.ChatroomID
	}

	isRead := "0"
	var newChat models.Chat
	// newChat := models.Chat{
	// 	ChatID:      0,
	// 	Email:       nil,
	// 	ChatroomID:  chatroomID,
	// 	Timendate:   time.Now().Format("2006-01-02 15:04:05"),
	// 	IsRead:      &isRead,
	// 	StatusRead:  nil,
	// 	Content:     infobipMessage.Content.Text,
	// 	MessageType: "text",
	// }

	switch infobipMessage.Results[0].Message.Type {
	case "TEXT":
		newChat = models.Chat{
			ChatID:      0,
			Email:       nil,
			ChatroomID:  chatroomID,
			Timendate:   time.Now().Format("2006-01-02 15:04:05"),
			IsRead:      &isRead,
			StatusRead:  nil,
			Content:     infobipMessage.Results[0].Message.Text,
			MessageType: "text",
		}
	case "IMAGE":
		newChat = models.Chat{
			ChatID:      0,
			Email:       nil,
			ChatroomID:  chatroomID,
			Timendate:   time.Now().Format("2006-01-02 15:04:05"),
			IsRead:      &isRead,
			StatusRead:  nil,
			Content:     infobipMessage.Results[0].Message.URL,
			MessageType: "photo",
		}
	case "VIDEO":
		newChat = models.Chat{
			ChatID:      0,
			Email:       nil,
			ChatroomID:  chatroomID,
			Timendate:   time.Now().Format("2006-01-02 15:04:05"),
			IsRead:      &isRead,
			StatusRead:  nil,
			Content:     infobipMessage.Results[0].Message.URL,
			MessageType: "video",
		}
	case "DOCUMENT":
		newChat = models.Chat{
			ChatID:      0,
			Email:       nil,
			ChatroomID:  chatroomID,
			Timendate:   time.Now().Format("2006-01-02 15:04:05"),
			IsRead:      &isRead,
			StatusRead:  nil,
			Content:     infobipMessage.Results[0].Message.URL,
			MessageType: "file",
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Unsupported message type"})
		return
	}

	err = repositories.ChatRepo.CreateChat(&newChat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to save chat information"})
		return
	}

	chatJSON, err := json.Marshal(newChat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to send chat"})
	}
	SendMessageToAll(c, chatJSON)
}
