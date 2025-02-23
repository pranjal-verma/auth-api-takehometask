package handlers

import (
	"auth-api/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService core.AuthService
}

type CreateUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
type UserResponse struct {
	Email string `json:"email"`
	ID    uint   `json:"id"`
}

type SigninInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (ah *AuthHandler) Signup(c *gin.Context) {

	var input CreateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ah.authService.CreateUser(core.User{
		Email:    input.Email,
		Password: input.Password,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully ", "email": user.Email})
}

func (ah *AuthHandler) Signin(c *gin.Context) {
	var input SigninInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, err := ah.authService.Authenticate(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})

}

func NewAuthHandler(authService core.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}
