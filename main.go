package main

import (
	"BackendProject/post-service/config"
	handlers2 "BackendProject/post-service/handlers"
	"BackendProject/post-service/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Ошибка загрузки .env файла")
	}
	// Подключение к базе данных (можно использовать ту же БД или отдельную)
	config.ConnectDatabase()

	r := gin.Default()

	r.Use(middleware.Logger())
	r.Use(middleware.RequestIdMiddleware())
	r.Use(middleware.RequestLogger())
	// Маршруты для работы с постами
	r.GET("/posts", handlers2.GetPosts)
	r.POST("/posts", handlers2.AddPost)
	r.GET("/posts/:id", middleware.AuthMiddleware(), handlers2.GetPostByID)
	r.PUT("/posts/:id", middleware.AuthMiddleware(), handlers2.UpdatePost)
	r.DELETE("/posts/:id", middleware.AuthMiddleware(), handlers2.DeletePost)
	r.GET("/posts/user/:id", middleware.AuthMiddleware(), handlers2.GetPostsByUserID)
	r.GET("/posts/category/:id", middleware.AuthMiddleware(), handlers2.GetPostsByCategoryID)
	r.GET("/posts/search", middleware.AuthMiddleware(), handlers2.SearchPostsByTitle)

	// Маршруты для работы с категориями
	r.GET("/categories", handlers2.GetCategories)
	r.POST("/categories", handlers2.AddCategory)
	r.GET("/categories/:id", handlers2.GetCategoryByID)
	r.PUT("/categories/:id", middleware.AuthMiddleware(), handlers2.UpdateCategory)
	r.DELETE("/categories/:id", middleware.AuthMiddleware(), handlers2.DeleteCategory)

	r.GET("/users/:id", middleware.AuthMiddleware(), handlers2.GetUserFromUserService)
	// Запуск сервера на порту 8080
	r.Run(":8080")
}
