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

// Добавление нового пользователя
func AddUser(c *gin.Context) {
	var newUser models.User

	// Привязка JSON к структуре
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	// Проверка на пустые поля
	if newUser.Name == "" || newUser.Email == "" || newUser.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Все поля обязательны"})
		return
	}

	// Сохранение пользователя в базе данных
	result := config.DB.Create(&newUser)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при добавлении пользователя"})
		return
	}

	// Ответ с добавленным пользователем
	c.JSON(http.StatusCreated, newUser)
}
