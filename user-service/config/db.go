package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	// Получаем параметры подключения из переменных окружения
	host := os.Getenv("DB_HOST")
	portStr := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Неверный порт: %v", err)
	}

	// Формируем строку подключения
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		host, user, password, dbname, port,
	)

	log.Printf("DSN: %s", dsn)

	// Создаём подключение к sql.DB
	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}

	// Создаём драйвер миграций
	driver, err := migratepg.WithInstance(sqlDB, &migratepg.Config{})
	if err != nil {
		log.Fatal("Ошибка создания драйвера миграций:", err)
	}

	// Создаём мигратор
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations", // Убедитесь, что путь правильный!
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal("Ошибка инициализации мигратора:", err)
	}

	// Обработка грязной миграции
	if version, dirty, err := m.Version(); err == nil && dirty {
		log.Printf("Грязная миграция на версии %d — принудительно исправляем...", version)
		if err := m.Force(int(version)); err != nil {
			log.Fatal("Ошибка force-модификации миграции:", err)
		}
	}

	// Применяем миграции
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Ошибка применения миграций:", err)
	}

	// Подключение GORM
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения GORM:", err)
	}

	DB = db
	log.Println("База данных подключена и миграции применены!")
}
