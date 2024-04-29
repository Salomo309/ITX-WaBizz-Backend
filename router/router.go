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

	// Auth endpoints
	apis.POST("/login", handlers.HandleCheckUserLogin)
	
	// Chatlist endpoints
	apis.GET("/chatlist", handlers.HandleChatlist)
	apis.GET("/chatlist/search/contact", handlers.HandleChatlistSearchByContact)
	apis.GET("/chatlist/search/message", handlers.HandleChatlistSearchByMessage)

	// Chat Endpoints
	apis.GET("/chatroom", handlers.HandleGetChatroom)
	apis.GET("/chatroom/websocket", handlers.HandleNewWebsocket)
	apis.POST("/chatroom/send", handlers.HandleSendMessage)
	apis.POST("/chatroom/receive", handlers.HandleReceiveMessage)

}

func applyCorsMiddleware(router *gin.Engine) {
	router.Use(middlewares.CorsMiddleware())
}
