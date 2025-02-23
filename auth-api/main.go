package main

import (
	"auth-api/config"
	"auth-api/core"
	"auth-api/database"
	"auth-api/handlers"
	"auth-api/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize router
	router := gin.Default()
	db, err := database.Initialize()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	userRepo := database.NewUserRepository(db)
	tokenService := core.NewTokenService(config.JWTSecretKey)
	authService := core.NewAuthService(userRepo, tokenService)
	authHandler := handlers.NewAuthHandler(authService)
	authMiddleware := middleware.NewAuthMiddleware(tokenService)

	// Initialize routes
	initializeRoutes(router, authHandler, authMiddleware)

	// Start server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func initializeRoutes(router *gin.Engine, authHandler *handlers.AuthHandler, authMiddleware *middleware.AuthMiddleware) {
	// Group auth routes
	auth := router.Group("/api/auth")
	{
		auth.POST("/signup", authHandler.Signup)
		auth.POST("/signin", authHandler.Signin)
		auth.POST("/refresh", authHandler.RefreshToken)
		// auth.POST("/revoke", middleware.AuthRequired(), handlers.RevokeToken)
		auth.GET("/check", authMiddleware.AuthRequired(), authHandler.Ping)
	}
}
