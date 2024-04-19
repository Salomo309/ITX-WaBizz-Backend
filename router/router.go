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

	// Authorization endpoints
	apis.POST("/auth/login", handlers.HandleGoogleLogin)
	apis.GET("/auth/google/callback", handlers.HandleGoogleCallback)
	apis.POST("/auth/refresh-token", handlers.HandleNewAccessToken)
	apis.POST("/auth/logout", middlewares.VerifyTokenMiddleware(), handlers.HandleLogout)

	// Chatlist endpoints
	apis.GET("/chatlist", handlers.HandleChatlist)
	apis.GET("/chatlist-search-by-contact", handlers.HandleChatlistSearchByContact)
	apis.GET("/chatlist-search-by-message", handlers.HandleChatlistSearchByMessage)

}

func applyCorsMiddleware(router *gin.Engine) {
	router.Use(middlewares.CorsMiddleware())
}
