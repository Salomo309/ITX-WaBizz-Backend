package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Message struct {
	From         string     `json:"from"`
	To           string     `json:"to"`
	MessageID    string     `json:"messageId"`
	Content      Content    `json:"content"`
	CallbackData string     `json:"callbackData"`
	NotifyURL    string     `json:"notifyUrl"`
	URLOptions   URLOptions `json:"urlOptions"`
}

type Content struct {
	Text string `json:"text"`
}

type URLOptions struct {
	ShortenURL    bool   `json:"shortenUrl"`
	TrackClicks   bool   `json:"trackClicks"`
	TrackingURL   string `json:"trackingUrl"`
	RemoveProtocol bool   `json:"removeProtocol"`
	CustomDomain  string `json:"customDomain"`
}

var receivedMessage string

func HandlerSendMessage(c *gin.Context) {
	var message Message
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Buat struktur Chat dari pesan yang diterima
	chat := models.Chat{
		ChatId:     message.MessageID,
		ChatroomId: message.To, // Anda mungkin perlu menyesuaikan ini sesuai dengan struktur model Chat Anda
		Timendate:  "",         // Atur tanggal dan waktu jika diperlukan
		IsRead:     "",         // Setel status baca ke nilai default
		Content:    message.Content.Text,
	}

	// Simpan chat ke dalam database
	if err := repositories.ChatRepo.CreateChat(&chat); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save chat to database"})
		return
	}

	// Kirim pesan ke penerima
	resp, err := http.Post("http://localhost"+recvPort+"/receive-message", "application/json", bytes.NewBuffer([]byte(dummyJSON)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	c.JSON(http.StatusOK, gin.H{"message": "Message sent and saved to database"})
}

// func SendHandler(c *gin.Context) {
// 	// Dummy JSON data
// 	dummyJSON := `{
// 		"from": "447860099299",
// 		"to": "6281360778689",
// 		"messageId": "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
// 		"content": {
// 			"text": "Some text"
// 			},
// 		"callbackData": "Callback data",
// 		"notifyUrl": "https://www.example.com/whatsapp",
// 		"urlOptions": {
// 			"shortenUrl": true,
// 			"trackClicks": true,
// 			"trackingUrl": "https://example.com/click-report",
// 			"removeProtocol": true,
// 			"customDomain": "example.com"
// 		}
// 	}`

// 	resp, err := http.Post("http://localhost"+recvPort+"/receive-message", "application/json", bytes.NewBuffer([]byte(dummyJSON)))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	defer resp.Body.Close()

// 	c.JSON(http.StatusOK, gin.H{"message": "Message sent"})
// }

func HandleReceiveMessage(c *gin.Context) {
	var message Message
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	receivedMessage = fmt.Sprintf("Received message: %+v", message)
	fmt.Println(receivedMessage)

	c.JSON(http.StatusOK, gin.H{"message": "Message received"})
}
