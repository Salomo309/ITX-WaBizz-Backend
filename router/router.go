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

	protected := apis.Group("")
	protected.Use(middlewares.VerifyEmailMiddleware())

	// Auth endpoints
	apis.POST("/login", handlers.CheckUserLogin)
	protected.POST("/logout", handlers.Logout)

	// Chatlist endpoints
	protected.GET("/chatlist", handlers.GetChatroomList)
	protected.GET("/chatlist/search/contact", handlers.SearchChatroomByContact)
	protected.GET("/chatlist/search/message", handlers.SearchChatroomByMessage)

	// Chat Endpoints
	protected.GET("/chatroom", handlers.GetChatroom)
	protected.POST("/chatroom/send", handlers.HandleSendMessage)
	protected.POST("/chatroom/receive", handlers.HandleReceiveMessage)

	// User endpoints
	protected.POST("/user/insert", handlers.InsertUser)
	protected.GET("/user/info", handlers.GetUserInfo)
	protected.GET("/user/all", handlers.GetAllUserInfo)
	protected.GET("/user/active", handlers.MakeUserActive)
	protected.GET("/user/inactive", handlers.MakeUserInactive)
}

func applyCorsMiddleware(router *gin.Engine) {
	router.Use(middlewares.CorsMiddleware())
}
