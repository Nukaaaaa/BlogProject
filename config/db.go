package config

import (
	"log"

	"github.com/golang-migrate/migrate/v4"                     // Пакет для миграций
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // Псевдоним для пакета миграций
	_ "github.com/golang-migrate/migrate/v4/source/file"       // Пакет для источников миграций
	gormPostgres "gorm.io/driver/postgres"                     // Псевдоним для GORM PostgreSQL драйвера
	"gorm.io/gorm"                                             // Пакет GORM
)

var DB *gorm.DB

func ConnectDatabase() {
	// Строка подключения для GORM
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	var err error
	// Используем GORM PostgreSQL драйвер
	DB, err = gorm.Open(gormPostgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}

	// Строка подключения для миграций
	m, err := migrate.New(
		"file://migrations", // Путь к миграциям
		"postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable", // Строка подключения для миграций
	)
	if err != nil {
		log.Fatal("Ошибка при создании мигратора:", err)
	}

	// Применяем миграции
	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatal("Ошибка при применении миграций:", err)
	}

	log.Println("Подключение и миграция выполнены успешно.")
}
