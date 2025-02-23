package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `gorm:"not null" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Token struct {
	ID        uint   `gorm:"primarykey"`
	UserID    uint   `gorm:"not null"`
	Token     string `gorm:"not null"`
	Type      string `gorm:"not null"` // access or refresh
	Revoked   bool   `gorm:"default:false"`
	ExpiresAt time.Time
	CreatedAt time.Time
}
