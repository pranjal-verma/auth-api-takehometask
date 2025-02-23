package config

import (
	"time"
)

// TODO: move to env variables
const (
	// JWT settings
	JWTSecretKey         = "your-secret-key" // Change this in production
	AccessTokenDuration  = time.Minute * 1
	RefreshTokenDuration = time.Hour * 24 * 7

	// Database settings
	DBHost = "db"
	// DBHost     = "localhost"
	DBUser     = "root"
	DBPassword = "password"
	DBName     = "auth_api"
	DBPort     = "3306"
)
