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
	apis.POST("/login", handlers.HandleGoogleLogin)
	apis.GET("/auth/google/callback", handlers.HandleGoogleCallback)
	apis.POST("/refresh-token", handlers.HandleNewAccessToken)
	apis.POST("/logout", middlewares.VerifyTokenMiddleware(), handlers.HandleLogout)

}

func applyCorsMiddleware(router *gin.Engine) {
	router.Use(middlewares.CorsMiddleware())
}
