package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"

	"itx-wabizz/handlers"
	"itx-wabizz/middlewares"
	"itx-wabizz/configs"
)

var Db *sql.DB

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize connection to database
	initDatabaseConnection()
	defer Db.Close()

	// Initialize Gin router
	router := gin.Default()

	// Enable CORS for all endpoints
	router.Use(middlewares.CorsMiddleware())

	// Define endpoints for back-end services
	// General handler
	router.GET("/api", handlers.HelloHandler)

	// Get port from .env and start server
	router.Run(getEnvPortOr("8080"))
}

func initDatabaseConnection() {
	var err error
	Db, err = sql.Open("mysql", configs.MysqlUser + ":" + configs.MysqlPassword +"@tcp(" + configs.MysqlHost + ":3306)/" + configs.MysqlDatabase)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	// Ping database to check if the connection is valid
	if err := Db.Ping(); err != nil {
		log.Fatal("Error pinging database:", err)
	}
}

func getEnvPortOr(port string) string {
	// If `PORT` variable in environment exists, return it
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	// Otherwise, return the value of `port` variable from function argument
	return ":" + port
}
