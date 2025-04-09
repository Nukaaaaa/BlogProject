package handlers

import (
	"BackendProject/config"
	"BackendProject/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Получить все посты с пагинацией и фильтрацией
func GetPosts(c *gin.Context) {
	var posts []models.Post

	// Получаем параметры для пагинации
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "5"))
	if err != nil {
		limit = 5
	}

	// Получаем параметры фильтрации
	categoryID := c.DefaultQuery("category_id", "")
	userID := c.DefaultQuery("user_id", "")

	// Формируем запрос с фильтрацией
	query := config.DB.Model(&models.Post{}).Preload("Category").Preload("User")

	// Фильтрация по категории
	if categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}

	// Фильтрация по пользователю
	if userID != "" {
		query = query.Where("user_id = ?", userID)
	}

	// Пагинация
	offset := (page - 1) * limit
	query = query.Offset(offset).Limit(limit)

	// Выполнение запроса
	result := query.Find(&posts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить посты"})
		return
	}

	// Ответ с постами
	c.JSON(http.StatusOK, posts)
}

// Добавить новый пост
func AddPost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	if post.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Название поста обязательно"})
		return
	}

	result := config.DB.Create(&post)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось добавить пост"})
		return
	}

	c.JSON(http.StatusCreated, post)
}

// Получить пост по ID
func GetPostByID(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	result := config.DB.First(&post, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пост не найден"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// Обновить пост по ID
func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	// Проверим, существует ли пост
	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пост не найден"})
		return
	}

	// Привязываем новые данные
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	// Обновляем пост
	result := config.DB.Save(&post)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить пост"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// Удалить пост по ID
func DeletePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	// Проверим, существует ли пост
	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пост не найден"})
		return
	}

	// Удаляем пост
	result := config.DB.Delete(&post)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить пост"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Пост удален"})
}
