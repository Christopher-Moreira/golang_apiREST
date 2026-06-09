package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() error {

	//Initialize DB
	database, err := InitializePostgres()
	if err != nil {
		return fmt.Errorf("error: Initializando postgres: %v", err)
	}

	db = database
	return nil
}

func GetPostgres() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	return NewLogger(p)

}
