package models

import "time"

/*
Struct: Message ->
Represent the data structure that is accepted by Infobip API
*/
type Message struct {
	From         string     `json:"from"`
	To           string     `json:"to"`
	MessageID    string     `json:"messageId"`
	Content      Content    `json:"content"`
	CallbackData string     `json:"callbackData"`
	NotifyURL    string     `json:"notifyUrl"`
	URLOptions   URLOptions `json:"urlOptions"`
}

/*
Struct: Content ->
Represent the content inside of Message accepted or sent by Infobip API
*/
type Content struct {
	Text string `json:"text"`
}

/*
Struct: URLOptions ->
Represent the URL Options that can be added for Message accepted or sent by Infobip API
*/
type URLOptions struct {
	ShortenURL    	bool   `json:"shortenUrl"`
	TrackClicks   	bool   `json:"trackClicks"`
	TrackingURL   	string `json:"trackingUrl"`
	RemoveProtocol 	bool   `json:"removeProtocol"`
	CustomDomain  	string `json:"customDomain"`
}

type ReceivedMessage struct {
	Results             []Result `json:"results"`
	MessageCount        int      `json:"messageCount"`
	PendingMessageCount int      `json:"pendingMessageCount"`
}

type Result struct {
	From            string        `json:"from"`
	To              string        `json:"to"`
	IntegrationType string        `json:"integrationType"`
	ReceivedAt      time.Time     `json:"receivedAt"`
	MessageID       string        `json:"messageId"`
	PairedMessageID string        `json:"pairedMessageId"`
	CallbackData    string        `json:"callbackData"`
	Message         MessageContent `json:"message"`
	Price           Price         `json:"price"`
}

type MessageContent struct {
	Type string `json:"type"`
	Text string `json:"text,omitempty"`
	URL  string `json:"url,omitempty"`
}

type Price struct {
	PricePerMessage int    `json:"pricePerMessage"`
	Currency        string `json:"currency"`
}
