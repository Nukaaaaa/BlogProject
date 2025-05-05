package mock

import (
	models "BackendProject/post-service/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var TestDB *gorm.DB

// InitMockDB инициализирует in-memory базу с modernc.org/sqlite
func InitMockDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		log.Fatalf("не удалось создать мок-базу: %v", err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Post{}, &models.Category{})
	if err != nil {
		log.Fatalf("не удалось провести миграцию: %v", err)
	}

	TestDB = db
	return db
}
