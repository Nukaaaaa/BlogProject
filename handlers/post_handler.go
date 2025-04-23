package handlers

import (
	"BackendProject/config"
	"BackendProject/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	var posts []models.Post

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

	c.JSON(http.StatusOK, posts)
}

func AddPost(c *gin.Context) {
	var post models.Post

	// Биндинг JSON в структуру post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	// Проверка обязательных полей
	if post.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Название поста обязательно"})
		return
	}

	// Проверка инициализации базы данных
	if config.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "База данных не инициализирована"})
		return
	}

	// Создание поста в базе данных
	result := config.DB.Create(&post)
	if result.Error != nil {
		// Логирование ошибки и отправка ответа
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Не удалось добавить пост",
			"details": result.Error.Error(), // Подробности об ошибке
		})
		return
	}

	// Успешный ответ с созданным постом
	c.JSON(http.StatusCreated, post)
}

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

func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пост не найден"})
		return
	}

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	result := config.DB.Save(&post)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить пост"})
		return
	}

	c.JSON(http.StatusOK, post)
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	// Проверим, существует ли пост
	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пост не найден"})
		return
	}

	result := config.DB.Delete(&post)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить пост"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Пост удален"})
}

func GetPostsByUserID(c *gin.Context) {
	userID := c.Param("id") // Получаем ID пользователя из параметра маршрута

	var posts []models.Post

	// Получаем посты конкретного пользователя
	result := config.DB.Where("user_id = ?", userID).Preload("User").Preload("Category").Find(&posts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить посты пользователя"})
		return
	}

	if len(posts) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Посты пользователя не найдены"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func GetPostsByCategoryID(c *gin.Context) {
	categoryID := c.Param("id") // Получаем ID категории из параметра маршрута

	var posts []models.Post

	// Получаем посты конкретной категории с подгрузкой связанных данных (User и Category)
	result := config.DB.Where("category_id = ?", categoryID).Preload("User").Preload("Category").Find(&posts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить посты по категории"})
		return
	}

	// Если посты не найдены
	if len(posts) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Посты категории не найдены"})
		return
	}

	// Возвращаем найденные посты
	c.JSON(http.StatusOK, posts)
}

func SearchPostsByTitle(c *gin.Context) {
	query := c.Query("q")
	var posts []models.Post

	result := config.DB.Where("title ILIKE ?", "%"+query+"%").Find(&posts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при поиске постов"})
		return
	}

	c.JSON(http.StatusOK, posts)
}
