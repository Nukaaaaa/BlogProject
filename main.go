package main

import (
	"BackendProject/config"
	"BackendProject/handlers"
	"BackendProject/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// Подключение к базе данных
	config.ConnectDatabase()

	// Инициализация маршрутизатора Gin
	r := gin.Default()

	// Маршруты для работы с постами (Post)
	r.GET("/posts", middleware.AuthMiddleware(), handlers.GetPosts)                          // Получить все посты
	r.POST("/posts", middleware.AuthMiddleware(), handlers.AddPost)                          // Добавить новый пост
	r.GET("/posts/:id", middleware.AuthMiddleware(), handlers.GetPostByID)                   // Получить пост по ID
	r.PUT("/posts/:id", middleware.AuthMiddleware(), handlers.UpdatePost)                    // Обновить пост по ID
	r.DELETE("/posts/:id", middleware.AuthMiddleware(), handlers.DeletePost)                 // Удалить пост по ID
	r.GET("/posts/user/:id", middleware.AuthMiddleware(), handlers.GetPostsByUserID)         // Получить все посты по ID клиенты
	r.GET("/posts/category/:id", middleware.AuthMiddleware(), handlers.GetPostsByCategoryID) // Получить все посты по ID категории

	// Маршруты для работы с категориями (Category)
	r.GET("/categories", middleware.AuthMiddleware(), handlers.GetCategories) // Получить все категории
	r.POST("/categories", middleware.AuthMiddleware(), handlers.AddCategory)  // Добавить новую категорию

	r.GET("/users", middleware.AuthMiddleware(), handlers.GetUsers) // Получить всех пользователей
	// Маршрут регистрации и входа
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)
	// Запуск сервера на порту 8080
	r.Run(":8080")
}
