package main

import (
	"BackendProject/post-service/config"
	"BackendProject/post-service/middleware"
	"BackendProject/user-service/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Ошибка загрузки .env файла")
	}
	// Подключение к базе данных (можно вынести в отдельную конфигурацию)
	config.ConnectDatabase()

	r := gin.Default()

	r.Use(middleware.Logger())
	// Маршруты для работы с пользователями
	r.GET("/users", handlers.GetUsers)
	r.GET("/users/:id", handlers.GetUserByID)
	r.PUT("/users/:id", middleware.AuthMiddleware(), handlers.UpdateUser)
	r.DELETE("/users/:id", middleware.AuthMiddleware(), handlers.DeleteUser)

	// Аутентификация
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	// Запуск сервера на порту 8081 (чтобы не конфликтовать с PostService)
	r.Run(":8081")
}
