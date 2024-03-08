package handlers

import (
    "fmt"
    "strings"
    "net/http"
    "io/ioutil"

	"github.com/gin-gonic/gin"

	// "itx-wabizz/models"
)

func HandleReceiveMessage(c *gin.Context) {
	// Read the body of the request
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to read request body: %v", err))
		return
	}

	// Print the received message
	fmt.Println("Received message from Infobip WhatsApp:")
	fmt.Println(string(body))

	// Respond with success
	c.JSON(http.StatusOK, gin.H{"status": string(body)})
}

func HandleSendMessage(c *gin.Context) {
	fmt.Fprintf(c.Writer, "HandleSendMessage is called")
	url := "https://n89mny.api.infobip.com/whatsapp/1/message/template"
	method := "POST"

	payload := strings.NewReader(`{"messages":[{"from":"447860099299","to":"6281360778689","messageId":"90db5876-18e5-4aae-af3c-2567185b4af7","content":{"templateName":"message_test","templateData":{"body":{"placeholders":["Salomo"]}},"language":"en"}}]}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to create request: %v", err))
		return
	}
	req.Header.Add("Authorization", "App 7dc75baa5ee04eecd96b9de85644da23-7215f447-3122-4036-9983-21dd3b6ff449")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to send message: %v", err))
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to send message: status code %d", res.StatusCode))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Message sent successfully",
	})
}

