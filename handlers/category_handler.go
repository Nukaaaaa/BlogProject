package handlers

import (
	"BackendProject/config"
	"BackendProject/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Получить все категории
func GetCategories(c *gin.Context) {
	var categories []models.Category
	result := config.DB.Find(&categories)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить категории"})
		return
	}

	c.JSON(http.StatusOK, categories)
}

// Добавить новую категорию
func AddCategory(c *gin.Context) {
	var newCategory models.Category
	if err := c.ShouldBindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	if newCategory.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Название категории обязательно"})
		return
	}

	result := config.DB.Create(&newCategory)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось добавить категорию"})
		return
	}

	c.JSON(http.StatusCreated, newCategory)
}
