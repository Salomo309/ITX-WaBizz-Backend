package main

import (
	"github.com/gin-gonic/gin"

	"itx-wabizz/handlers"
)


func main() {
	r := gin.Default()

	r.POST("/send", handlers.HandleInfobipSend)
	r.POST("/receive", handlers.HandleInfobipReceive)

	r.Run(":8081")
}