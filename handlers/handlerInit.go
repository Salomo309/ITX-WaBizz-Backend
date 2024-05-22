package handlers

import (
	"context"
	"firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"itx-wabizz/configs"
)

var (
	FirebaseClient *messaging.Client
)

/*
Function: InitMessagingClient

Init the firebase messaging client for push notification use
*/
func InitMessagingClient() {
	opt := option.WithCredentialsFile(configs.MessagingCredentialsPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatal("Error initiating messaging client: ", err)
	}

	FirebaseClient, err = app.Messaging(context.Background())
	if err != nil {
		log.Fatal("Error initiating messaging client: ", err)
	}
}

/*
Function: Welcome

For testing purpose and see if server can be reached or not
*/
func Welcome(c *gin.Context) {
	fmt.Fprintf(c.Writer, "Hello, Welcome to ITX WABizz Web Service!")
}
