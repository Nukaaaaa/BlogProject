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
	r.GET("/posts", middleware.AuthMiddleware(), handlers.GetPosts)
	r.POST("/posts", middleware.AuthMiddleware(), handlers.AddPost)
	r.GET("/posts/:id", middleware.AuthMiddleware(), handlers.GetPostByID)
	r.PUT("/posts/:id", middleware.AuthMiddleware(), handlers.UpdatePost)
	r.DELETE("/posts/:id", middleware.AuthMiddleware(), handlers.DeletePost)
	r.GET("/posts/user/:id", middleware.AuthMiddleware(), handlers.GetPostsByUserID)
	r.GET("/posts/category/:id", middleware.AuthMiddleware(), handlers.GetPostsByCategoryID)
	r.GET("/posts/search", middleware.AuthMiddleware(), handlers.SearchPostsByTitle)

	// Маршруты для работы с категориями (Category)
	r.GET("/categories", handlers.GetCategories)
	r.POST("/categories", handlers.AddCategory)
	r.GET("/categories/:id", handlers.GetCategoryByID)
	r.PUT("/categories/:id", middleware.AuthMiddleware(), handlers.UpdateCategory)
	r.DELETE("/categories/:id", middleware.AuthMiddleware(), handlers.DeleteCategory)

	// Маршруты для работы с пользователями (User)
	r.GET("/users", handlers.GetUsers)
	r.GET("/users/:id/posts", middleware.AuthMiddleware(), handlers.GetUserPosts)
	r.GET("/users/:id", handlers.GetUserByID)
	r.PUT("/users/:id", middleware.AuthMiddleware(), handlers.UpdateUser)
	r.DELETE("/users/:id", middleware.AuthMiddleware(), handlers.DeleteUser)

	// Маршрут регистрации и входа
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	// Получить данные текущего пользователя
	r.GET("/me", middleware.AuthMiddleware(), handlers.GetMe)

	// Запуск сервера на порту 8080
	r.Run(":8080")
}
