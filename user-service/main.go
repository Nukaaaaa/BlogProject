package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"user-service/config"
	"user-service/handlers"
	"user-service/middleware"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Ошибка загрузки .env файла")
	}
	// Подключение к базе данных (можно вынести в отдельную конфигурацию)
	config.ConnectDatabase()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Разрешаем запросы с React фронта
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	r.Use(middleware.Logger())
	// Маршруты для работы с пользователями
	r.GET("/users", handlers.GetUsers)
	r.GET("/users/:id", handlers.GetUserByID)
	r.GET("/users/:id/posts", handlers.GetUserPosts)
	r.PUT("/users/:id", middleware.AuthMiddleware(), handlers.UpdateUser)
	r.DELETE("/users/:id", middleware.AuthMiddleware(), handlers.DeleteUser)

	// Аутентификация
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	// Запуск сервера на порту 8081 (чтобы не конфликтовать с PostService)
	r.Run(":8081")
}
