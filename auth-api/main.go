package main

import (
	"auth-api/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize router
	router := gin.Default()

	// Initialize routes
	initializeRoutes(router)

	// Start server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func initializeRoutes(router *gin.Engine) {
	// Group auth routes
	auth := router.Group("/api/auth")
	{
		auth.POST("/signup", handlers.Signup)
		auth.POST("/signin", handlers.Signin)
		// auth.POST("/refresh", middleware.AuthRequired(), handlers.RefreshToken)
		// auth.POST("/revoke", middleware.AuthRequired(), handlers.RevokeToken)
	}
}
