package core

import (
	"auth-api/config"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenService interface {
	GenerateTokenPair(userID uint) (accessToken, refreshToken string, err error)
	ValidateToken(tokenString string) (*Claims, error)
}

type tokenService struct {
	secretKey string
}

type Claims struct {
	UserID uint
	Type   string
	jwt.RegisteredClaims
}

func NewTokenService(secretKey string) TokenService {
	return &tokenService{secretKey: secretKey}
}

func (s *tokenService) GenerateTokenPair(userID uint) (string, string, error) {
	// Generate access token
	accessToken, err := s.generateToken(userID, "access", config.AccessTokenDuration)
	if err != nil {
		return "", "", err
	}

	// Generate refresh token
	refreshToken, err := s.generateToken(userID, "refresh", config.RefreshTokenDuration)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *tokenService) generateToken(userID uint, tokenType string, duration time.Duration) (string, error) {
	claims := Claims{
		UserID: userID,
		Type:   tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secretKey))
}

func (s *tokenService) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
