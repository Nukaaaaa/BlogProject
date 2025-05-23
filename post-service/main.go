package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"myproject/config"
	"myproject/handlers"
	"myproject/middleware"
)

func main() {
	// Загрузка .env файла
	err := godotenv.Load()
	if err != nil {
		log.Println("Ошибка загрузки .env файла")
	}

	// Подключение к базе данных
	config.ConnectDatabase()

	// Создаем роутер Gin
	r := gin.Default()

	// Настройка CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Разрешаем запросы с React фронта
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Подключаем другие middleware
	r.Use(middleware.Logger())
	r.Use(middleware.RequestIdMiddleware())
	r.Use(middleware.RequestLogger())

	// Маршруты для работы с постами
	r.GET("/posts", handlers.GetPosts)
	r.POST("/posts", handlers.AddPost)
	r.GET("/posts/:id", middleware.AuthMiddleware(), handlers.GetPostByID)
	r.PUT("/posts/:id", middleware.AuthMiddleware(), handlers.UpdatePost)
	r.DELETE("/posts/:id", middleware.AuthMiddleware(), handlers.DeletePost)
	r.GET("/posts/user/:id", middleware.AuthMiddleware(), handlers.GetPostsByUserID)
	r.GET("/posts/category/:id", middleware.AuthMiddleware(), handlers.GetPostsByCategoryID)
	r.GET("/posts/search", middleware.AuthMiddleware(), handlers.SearchPostsByTitle)

	// Маршруты для работы с категориями
	r.GET("/categories", handlers.GetCategories)
	r.POST("/categories", handlers.AddCategory)
	r.GET("/categories/:id", handlers.GetCategoryByID)
	r.PUT("/categories/:id", middleware.AuthMiddleware(), handlers.UpdateCategory)
	r.DELETE("/categories/:id", middleware.AuthMiddleware(), handlers.DeleteCategory)

	// Маршрут для пользователей
	r.GET("/users/:id", middleware.AuthMiddleware(), handlers.GetUserFromUserService)

	// Запуск сервера на порту 8080
	r.Run(":8080")
}
