package models

import (
	"gorm.io/driver/sqlite" 
	"gorm.io/gorm"
)

// Database exported GORM database.
var Database *gorm.DB 

func InitialiseDatabase(path string) error {
	var err error
	Database, err = gorm.Open(sqlite.Open(path), &gorm.Config{})

	if err != nil {
		return err 
	}

	Database.AutoMigrate(&Transaction{}, &Statement{})

	return nil
}

