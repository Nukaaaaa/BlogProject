package handlers

import (
	"BackendProject/post-service/config"
	"BackendProject/post-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Получение всех пользователей
func GetUsers(c *gin.Context) {
	var users []models.User
	result := config.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении пользователей"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// Получение пользователя по ID
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	result := config.DB.First(&user, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении пользователя"})
		}
		return
	}

	c.JSON(http.StatusOK, user)
}

// Обновление пользователя
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	// Проверяем существование пользователя
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}

	// Привязываем новые данные
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	// Сохраняем изменения
	result := config.DB.Save(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении пользователя"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Удаление пользователя
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	// Проверяем существование пользователя
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}

	// Удаляем пользователя
	result := config.DB.Delete(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении пользователя"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Пользователь успешно удален"})
}

// Получение постов пользователя
func GetUserPosts(c *gin.Context) {
	userID := c.Param("id")
	var posts []models.Post

	result := config.DB.Where("user_id = ?", userID).Find(&posts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении постов пользователя"})
		return
	}

	c.JSON(http.StatusOK, posts)
}
