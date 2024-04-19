package infobip

// import (
// 	"bytes"
// 	"fmt"
// 	"time"
// 	"net/http"
// 	"github.com/gin-gonic/gin"
	
// 	"itx-wabizz/models"
// 	"itx-wabizz/repositories"
// )

// type Message struct {
// 	From         string     `json:"from"`
// 	To           string     `json:"to"`
// 	MessageID    string     `json:"messageId"`
// 	Content      Content    `json:"content"`
// 	CallbackData string     `json:"callbackData"`
// 	NotifyURL    string     `json:"notifyUrl"`
// 	URLOptions   URLOptions `json:"urlOptions"`
// }

// type Content struct {
// 	Text string `json:"text"`
// }

// type URLOptions struct {
// 	ShortenURL    bool   `json:"shortenUrl"`
// 	TrackClicks   bool   `json:"trackClicks"`
// 	TrackingURL   string `json:"trackingUrl"`
// 	RemoveProtocol bool   `json:"removeProtocol"`
// 	CustomDomain  string `json:"customDomain"`
// }

// var receivedMessage string

// func HandleSendMessage(c *gin.Context) {
// 	var message Message
// 	if err := c.ShouldBindJSON(&message); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	currentTime := time.Now().Format(time.RFC3339)

// 	// Buat struktur Chat dari pesan yang diterima
// 	chat := models.Chat{
// 		ChatId:     message.MessageID,
// 		ChatroomId: message.To,
// 		Timendate:  currentTime,
// 		IsRead:     "",
// 		Content:    message.Content.Text,
// 	}

// 	// Simpan chat ke dalam database
// 	if err := repositories.ChatRepo.CreateChat(&chat); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save chat to database"})
// 		return
// 	}

// 	// Kirim pesan ke penerima
// 	resp, err := http.Post("http://localhost"+recvPort+"/receive-message", "application/json", bytes.NewBuffer([]byte(dummyJSON)))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	defer resp.Body.Close()

// 	c.JSON(http.StatusOK, gin.H{"message": "Message sent and saved to database"})
// }
