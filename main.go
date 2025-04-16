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
	r.GET("/posts", handlers.GetPosts)                          // Получить все посты
	r.POST("/posts", handlers.AddPost)                          // Добавить новый пост
	r.GET("/posts/:id", handlers.GetPostByID)                   // Получить пост по ID
	r.PUT("/posts/:id", handlers.UpdatePost)                    // Обновить пост по ID
	r.DELETE("/posts/:id", handlers.DeletePost)                 // Удалить пост по ID
	r.GET("/posts/user/:id", handlers.GetPostsByUserID)         // Получить все посты по ID клиенты
	r.GET("/posts/category/:id", handlers.GetPostsByCategoryID) // Получить все посты по ID категории

	// Маршруты для работы с категориями (Category)
	r.GET("/categories", handlers.GetCategories) // Получить все категории
	r.POST("/categories", handlers.AddCategory)  // Добавить новую категорию
	auth := r.Group("/")

	auth.Use(middleware.AuthMiddleware()) // защищённые маршруты
	{
		auth.GET("/users", handlers.GetUsers)
		auth.POST("/users", handlers.AddUser)
	}
	// Маршрут регистрации и входа
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)
	// Запуск сервера на порту 8080
	r.Run(":8080")
}
