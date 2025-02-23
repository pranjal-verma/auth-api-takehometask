package database

import (
	"auth-api/core"
	"errors"
	"time"

	"gorm.io/gorm"
)

// consider index on email
type User struct {
	ID        uint      `gorm:"primarykey;column:id"`
	Email     string    `gorm:"unique;not null;column:email"`
	Password  string    `gorm:"not null;column:password"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type userRepository struct {
	db *gorm.DB
}

// TODO implement
func (ur *userRepository) CreateUser(user core.User) (core.User, error) {
	dbUser := toUser(user)
	// find one user with the same email
	var existingUser User
	err := ur.db.Where("email = ?", dbUser.Email).First(&existingUser).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return core.User{}, err
		}
	}
	if existingUser.ID != 0 {
		return core.User{}, errors.New("email already exists")
	}
	err = ur.db.Create(&dbUser).Error
	if err != nil {
		return core.User{}, err
	}

	return toCoreUser(dbUser), nil
}

// TODO implement
func (ur *userRepository) GetUserByEmail(email string) (core.User, error) {
	var user User
	err := ur.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return core.User{}, err
	}
	return toCoreUser(user), nil
}

func NewUserRepository(db *gorm.DB) core.UserRepository {
	return &userRepository{db: db}
}

func toUser(user core.User) User {
	return User{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}
}

func toCoreUser(user User) core.User {
	return core.User{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}
}
