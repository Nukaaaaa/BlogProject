package main

import (
	"BackendProject/config"
	"BackendProject/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Подключение к базе данных
	config.ConnectDatabase()

	// Инициализация маршрутизатора Gin
	r := gin.Default()

	// Маршруты для работы с постами (Post)
	r.GET("/posts", handlers.GetPosts)          // Получить все посты
	r.POST("/posts", handlers.AddPost)          // Добавить новый пост
	r.GET("/posts/:id", handlers.GetPostByID)   // Получить пост по ID
	r.PUT("/posts/:id", handlers.UpdatePost)    // Обновить пост по ID
	r.DELETE("/posts/:id", handlers.DeletePost) // Удалить пост по ID

	// Маршруты для работы с категориями (Category)
	r.GET("/categories", handlers.GetCategories) // Получить все категории
	r.POST("/categories", handlers.AddCategory)  // Добавить новую категорию

	r.GET("/users", handlers.GetUsers) // Получить всех пользователей
	r.POST("/users", handlers.AddUser)
	// Запуск сервера на порту 8080
	r.Run(":8080")
}
