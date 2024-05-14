package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"

	"itx-wabizz/models"
)

func HandleInfobipSend(c *gin.Context) {
	infobipMessage := models.ReceivedMessage{
		Results: []models.Result{
			{
				From:            "0888-8888-8888",
				To:              "0999-9999-9999",
				IntegrationType: "API",
				ReceivedAt:      time.Now(),
				MessageID:       "AB09JC",
				PairedMessageID: "XYZ123",
				CallbackData:    "some-callback-data",
				Message: models.MessageContent{
					Type: "TEXT",
					Text: "Halo Selamat Pagi Dunia",
				},
				Price: models.Price{
					PricePerMessage: 0,
					Currency:        "USD",
				},
			},
		},
		MessageCount:        1,
		PendingMessageCount: 0,
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
