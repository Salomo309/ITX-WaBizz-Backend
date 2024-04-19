package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"

	"itx-wabizz/models"
)

func HandleInfobipSend(c *gin.Context) {
	infobipMessage := models.Message{
		From: "0888-8888-8888",
		To:"0999-9999-9999",
		MessageID: "AB09JC",
		Content: models.Content{Text: "Halo Selamat Pagi Dunia"},
		CallbackData: "",
		NotifyURL: "",
		URLOptions: models.URLOptions{
			ShortenURL: true,
			TrackClicks: true,
			TrackingURL: "",
			RemoveProtocol: true,
			CustomDomain: "",
		},
	}

	requestBody, _ := json.Marshal(infobipMessage)
	response, _ := http.Post("http://localhost:8080/api/chatroom/receive", "application/json", bytes.NewBuffer(requestBody))
	defer response.Body.Close()
	
	var responseData map[string]interface{}
	json.NewDecoder(response.Body).Decode(&responseData)
	c.JSON(response.StatusCode, responseData)
}

func HandleInfobipReceive(c *gin.Context) {
	var infobipMessage models.Message
	c.BindJSON(&infobipMessage)
	fmt.Println("New Message: Received message from ITX WABizz Back-end Server")
	fmt.Println("Message Content: " + infobipMessage.Content.Text)
}