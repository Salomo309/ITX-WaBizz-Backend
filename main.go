package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"database/sql"

	"itx-wabizz/handlers"
	"itx-wabizz/middlewares"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize Gin router
	router := gin.Default()

	// Enable CORS for all endpoints
	router.Use(middlewares.CorsMiddleware())

	// Define endpoints for back-end services
	// General handler
	router.GET("/api", handlers.HelloHandler)

	// Get port from .env and start server
	fmt.Println()
	router.Run(GetEnvPortOr("8080"))

	db, err := sql.Open("mysql", "root:admin123@tcp(mysql:3306)/itxwabizzdb")

    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    // Coba lakukan ping ke database
    err = db.Ping()
    if err != nil {
        panic(err.Error())
    }

    fmt.Println("Koneksi ke database MySQL berhasil!")
}

func GetEnvPortOr(port string) string {
	// If `PORT` variable in environment exists, return it
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	// Otherwise, return the value of `port` variable from function argument
	return ":" + port
}
