package core

import (
	"auth-api/config"
	"auth-api/utils"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID       uint
	Email    string
	Password string
}

type AuthService interface {
	CreateUser(user User) (User, error)
	GetUserByEmail(email string) (User, error)
	Authenticate(email, password string) (string, string, error)
	// ValidateToken(tokenString string) (*Claims, error)
	RefreshToken(refreshToken string) (string, error)
}

// Deals with user CRUD
type UserRepository interface {
	CreateUser(user User) (User, error)
	GetUserByEmail(email string) (User, error)
}

type authServiceImpl struct {
	userRepo     UserRepository
	tokenService TokenService
}

func (as *authServiceImpl) CreateUser(user User) (User, error) {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return User{}, err
	}
	user.Password = hashedPassword
	fmt.Println("USER", user)
	return as.userRepo.CreateUser(user)
}

func (as *authServiceImpl) Authenticate(email, password string) (string, string, error) {
	user, err := as.userRepo.GetUserByEmail(email)
	if err != nil {
		return "", "", err
	}

	if !utils.CheckPassword(password, user.Password) {
		return "", "", fmt.Errorf("invalid password")
	}

	accessToken, refreshToken, err := as.tokenService.GenerateTokenPair(user.ID)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func (as *authServiceImpl) GetUserByEmail(email string) (User, error) {
	return as.userRepo.GetUserByEmail(email)
}
func (as *authServiceImpl) RefreshToken(refreshToken string) (string, error) {
	claims, err := ValidateToken(refreshToken)
	if err != nil {
		return "", err
	}
	// check if refresh token is expired
	if claims.ExpiresAt.Before(time.Now()) {
		return "", fmt.Errorf("refresh token expired")
	}
	return GenerateToken(claims.UserID, "access")
}

func GenerateToken(userID uint, tokenType string) (string, error) {
	duration := config.AccessTokenDuration
	if tokenType == "refresh" {
		duration = config.RefreshTokenDuration
	}

	claims := Claims{
		UserID: userID,
		Type:   tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWTSecretKey))
}

func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWTSecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func NewAuthService(userRepo UserRepository, tokenService TokenService) AuthService {
	return &authServiceImpl{userRepo: userRepo, tokenService: tokenService}
}
