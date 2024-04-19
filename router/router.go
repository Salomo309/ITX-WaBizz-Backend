package router

import (
	"github.com/gin-gonic/gin"

	"itx-wabizz/handlers"
	"itx-wabizz/middlewares"
)

func ConfigureRouter(router *gin.Engine) {
	// Enable CORS middleware for all endpoints
	applyCorsMiddleware(router)

	// Define endpoints for back-end services
	// General handler
	router.GET("/api", handlers.HelloHandler)
	apis := router.Group("/api")

	// Chatlist endpoints
	apis.GET("/chatlist", handlers.HandleChatlist)
	apis.GET("/chatlist-search-by-contact", handlers.HandleChatlistSearchByContact)
	apis.GET("/chatlist-search-by-message", handlers.HandleChatlistSearchByMessage)

	// Chat Endpoints
	apis.POST("/send-msg", handlers.HandleSendMessage)
	apis.POST("rcv-msg", handlers.HandleReceiveMessage)

}

func applyCorsMiddleware(router *gin.Engine) {
	router.Use(middlewares.CorsMiddleware())
}
