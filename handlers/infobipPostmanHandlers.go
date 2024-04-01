package handlers
// package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Message struct {
	From         string       `json:"from"`
	To           string       `json:"to"`
	MessageID    string       `json:"messageId"`
	Content      Content      `json:"content"`
	CallbackData string       `json:"callbackData"`
	NotifyURL    string       `json:"notifyUrl"`
	URLOptions   URLOptions   `json:"urlOptions"`
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

const (
	sendPort = ":8080"
	recvPort = ":8081"
)

var receivedMessage string

func sendHandler(w http.ResponseWriter, r *http.Request) {
	// Dummy JSON data
	dummyJSON := `{
		"from": "447860099299",
		"to": "6281360778689",
		"messageId": "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
		"content": {
			"text": "Some text"
			},
		"callbackData": "Callback data",
		"notifyUrl": "https://www.example.com/whatsapp",
		"urlOptions": {
			"shortenUrl": true,
			"trackClicks": true,
			"trackingUrl": "https://example.com/click-report",
			"removeProtocol": true,
			"customDomain": "example.com"
		}
	}`

	resp, err := http.Post("http://localhost"+recvPort+"/receive-message", "application/json", bytes.NewBuffer([]byte(dummyJSON)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.WriteHeader(http.StatusOK)
}

func receiveHandler(w http.ResponseWriter, r *http.Request) {
	var message Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

	receivedMessage = fmt.Sprintf("Received message: %+v", message)
	fmt.Println(receivedMessage)
}

// func main() {
// 	http.HandleFunc("/send-message", sendHandler)
// 	http.HandleFunc("/receive-message", receiveHandler)

// 	go func() {
// 		fmt.Println("Send server is running on port", sendPort)
// 		log.Fatal(http.ListenAndServe(sendPort, nil))
// 	}()

// 	fmt.Println("Receive server is running on port", recvPort)
// 	log.Fatal(http.ListenAndServe(recvPort, nil))
// }
