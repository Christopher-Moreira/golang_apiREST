package config

import (
	"fmt"
	"os"

	"github.com/Christopher-Moreira/golang_apiREST/schemas"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgres() (*gorm.DB, error) {
	//Load Logger for error conditioning
	logger := GetLogger("postgres")

	//Check for .env
	err := godotenv.Load()
	if err != nil {
		logger.Errorf(".Env Loading error: %v", err)
		return nil, err
	}

	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_SSLMODE := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		DB_HOST,
		DB_USER,
		DB_PASSWORD,
		DB_NAME,
		DB_PORT,
		DB_SSLMODE,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("postgres opening error: %v", err)
		return nil, err
	}

	if err := db.AutoMigrate(&schemas.Opening{}); err != nil {
		logger.Errorf("postgres automigrate error: %v", err)
		return nil, err
	}

	return db, nil
}
