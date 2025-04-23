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

func AddCategory(c *gin.Context) {
	var newCategory models.Category
	if err := c.ShouldBindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Проверка, что база данных инициализирована
	if config.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is nil"})
		return
	}

	// Попытка создать категорию в базе данных
	result := config.DB.Create(&newCategory)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, newCategory)
}
func GetCategoryByID(c *gin.Context) {
	id := c.Param("id")
	var category models.Category
	result := config.DB.First(&category, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Категория не найдена"})
		return
	}
	c.JSON(http.StatusOK, category)
}

func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category
	if err := config.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Категория не найдена"})
		return
	}

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	config.DB.Save(&category)
	c.JSON(http.StatusOK, category)
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	result := config.DB.Delete(&models.Category{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить категорию"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Категория удалена"})
}
