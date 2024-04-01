package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"itx-wabizz/repositories"
	"itx-wabizz/router"
)

func init() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// Initialize connection to database
	repositories.InitDatabaseConnection()
	defer repositories.Db.Close()

	// Initialize all repositories needed
	repositories.InitRepositories()

	// Initialize Gin router
	r := gin.Default()

	// Configure router to include routes and middlewares
	router.ConfigureRouter(r)

	// Infobip handlers
	apis.POST("/send-message/text", handlers.HandleSendTextMessage)
	apis.POST("/send-message/doc", handlers.HandleSendDocumentMessage)
	apis.POST("/send-message/img", handlers.HandleSendImageMessage)
	apis.POST("/send-message/audio", handlers.HandleSendAudioMessage)
	apis.POST("/send-message/video", handlers.HandleSendVideoMessage)
	apis.POST("/receive-message", handlers.HandleReceiveMessage)

	// Get port from .env and start server
	r.Run(getEnvPortOr("8080"))
}

func getEnvPortOr(port string) string {
	// If `PORT` variable in environment exists, return it
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	// Otherwise, return the value of `port` variable from function argument
	return ":" + port
}
