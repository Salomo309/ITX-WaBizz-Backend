package handlers

import (
    "fmt"
    "net/http"
    "io/ioutil"

	"context"

	"github.com/gin-gonic/gin"
	"github.com/infobip-community/infobip-api-go-sdk/v3/pkg/infobip"
	"github.com/infobip-community/infobip-api-go-sdk/v3/pkg/infobip/models"
)

const (
	apiKey  = "7dc75baa5ee04eecd96b9de85644da23-7215f447-3122-4036-9983-21dd3b6ff449"
	baseURL = "n89mny.api.infobip.com"
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

func HandleSendTextMessage(c *gin.Context) {
	fmt.Println("API Key:", apiKey)
	fmt.Println("Base URL:", baseURL)
	client, err := infobip.NewClient(baseURL, apiKey)

	message := models.WATextMsg{
		MsgCommon: models.MsgCommon{
			From: "447860099299",
			To:   "6281360778689",
		},
		Content: models.TextContent{
			Text: "Hai sallll",
		},
	}
	msgResp, respDetails, err := client.WhatsApp.SendText(context.Background(), message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("%+v\n", msgResp)
	fmt.Printf("%+v\n", respDetails)
}

func HandleSendDocumentMessage(c *gin.Context) {
	fmt.Println("API Key:", apiKey)
	fmt.Println("Base URL:", baseURL)
	client, err := infobip.NewClient(baseURL, apiKey)

	message := models.WATextMsg{
		MsgCommon: models.MsgCommon{
			From: "447860099299",
			To:   "6281360778689",
		},
		Content: models.DocumentContent{MediaURL: "https://myurl.com/doc1.doc"},
	}
	msgResp, respDetails, err := client.WhatsApp.SendDocument(context.Background(), message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("%+v\n", msgResp)
	fmt.Printf("%+v\n", respDetails)
}

func HandleSendDocumentMessage(c *gin.Context) {
	fmt.Println("API Key:", apiKey)
	fmt.Println("Base URL:", baseURL)
	client, err := infobip.NewClient(baseURL, apiKey)

	message := models.WATextMsg{
		MsgCommon: models.MsgCommon{
			From: "447860099299",
			To:   "6281360778689",
		},
		Content: models.ImageContent{
			MediaURL: "https://thumbs.dreamstime.com/z/example-red-tag-example-red-square-price-tag-117502755.jpg",
		},
	}
	msgResp, respDetails, err := client.WhatsApp.SendImage(context.Background(), message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("%+v\n", msgResp)
	fmt.Printf("%+v\n", respDetails)
}

func HandleSendAudioMessage(c *gin.Context) {
	fmt.Println("API Key:", apiKey)
	fmt.Println("Base URL:", baseURL)
	client, err := infobip.NewClient(baseURL, apiKey)

	message := models.WATextMsg{
		MsgCommon: models.MsgCommon{
			From: "447860099299",
			To:   "6281360778689",
		},
		Content: models.AudioContent{MediaURL: "https://dl.espressif.com/dl/audio/ff-16b-2c-44100hz.aac"},
	}
	msgResp, respDetails, err := client.WhatsApp.SendAudio(context.Background(), message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("%+v\n", msgResp)
	fmt.Printf("%+v\n", respDetails)
}

func HandleSendAudioMessage(c *gin.Context) {
	fmt.Println("API Key:", apiKey)
	fmt.Println("Base URL:", baseURL)
	client, err := infobip.NewClient(baseURL, apiKey)

	message := models.WATextMsg{
		MsgCommon: models.MsgCommon{
			From: "447860099299",
			To:   "6281360778689",
		},
		Content: models.VideoContent{MediaURL: "https://download.samplelib.com/mp4/sample-5s.mp4"},
	}
	msgResp, respDetails, err := client.WhatsApp.SendVideo(context.Background(), message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("%+v\n", msgResp)
	fmt.Printf("%+v\n", respDetails)
}
