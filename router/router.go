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
	// General endpoints
	router.GET("/api", handlers.Welcome)
	apis := router.Group("/api")

	// Auth endpoints
	apis.POST("/login", handlers.CheckUserLogin)

	// Chatlist endpoints
	apis.GET("/chatlist", handlers.GetChatroomList)
	apis.GET("/chatlist/search/contact", handlers.SearchChatroomByContact)
	apis.GET("/chatlist/search/message", handlers.SearchChatroomByMessage)

	// Chat Endpoints
	apis.GET("/chatroom", handlers.GetChatroom)
	apis.GET("/chatroom/websocket", handlers.HandleNewWebsocket)
	apis.POST("/chatroom/send", handlers.HandleSendMessage)
	apis.POST("/chatroom/receive", handlers.HandleReceiveMessage)

	// User endpoints
	apis.POST("/user/insert", handlers.InsertUser)
	apis.GET("/user/info", handlers.GetUserInfo)
	apis.GET("/user/all", handlers.GetAllUserInfo)
	apis.GET("/user/active", handlers.MakeUserActive)
	apis.GET("/user/inactive", handlers.MakeUserInactive)
}

func applyCorsMiddleware(router *gin.Engine) {
	router.Use(middlewares.CorsMiddleware())
}
