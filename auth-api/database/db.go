package database

import (
	"auth-api/config"
	"auth-api/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Initialize() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto migrate the schemas
	err = db.AutoMigrate(&models.User{}, &models.Token{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
