package config

import (
	"BackendProject/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Подключение к базе данных PostgreSQL
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}

	// Автоматическая миграция моделей
	err = database.AutoMigrate(&models.User{}, &models.Post{}, &models.Category{})
	if err != nil {
		log.Fatal("Ошибка при миграции моделей:", err)
	}

	log.Println("Подключение к базе данных выполнено успешно.")
	DB = database
}
