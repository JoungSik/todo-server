package configs

import (
	"todo-server/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// migrate database
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Event{})

	return db, nil
}
