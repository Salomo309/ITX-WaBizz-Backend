package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"itx-wabizz/handlers"
	"itx-wabizz/middlewares"
	"itx-wabizz/repositories"
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
	router := gin.Default()

	// Enable CORS for all endpoints
	router.Use(middlewares.CorsMiddleware())

	// Define endpoints for back-end services
	// General handler
	router.GET("/api", handlers.HelloHandler)
	apis := router.Group("/api")

	// Authorization handlers
	apis.POST("/login", handlers.HandleGoogleLogin)
	apis.GET("/auth/google/callback", handlers.HandleGoogleCallback)
	apis.POST("/logout", handlers.HandleLogout)

	// Infobip handlers
	apis.POST("/send-message", handlers.HandleSendMessage)
	apis.POST("/receive-message", handlers.HandleReceiveMessage)

	// Get port from .env and start server
	router.Run(getEnvPortOr("8080"))
}

func getEnvPortOr(port string) string {
	// If `PORT` variable in environment exists, return it
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	// Otherwise, return the value of `port` variable from function argument
	return ":" + port
}
