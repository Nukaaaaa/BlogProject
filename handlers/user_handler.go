package handlers

import (
	"BackendProject/config"
	"BackendProject/models"
	"net/http"

	"github.com/gin-gonic/gin"
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
