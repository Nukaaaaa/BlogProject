package handlers

import (
	"BackendProject/config"
	"BackendProject/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	// Мигрируем схемы
	err = db.AutoMigrate(&models.User{}, &models.Category{}, &models.Post{})
	if err != nil {
		t.Fatalf("failed to migrate models: %v", err)
	}

	// Создаем тестового пользователя и категорию
	testUser := models.User{Name: "Test User", Email: "test@example.com"}
	if err := db.Create(&testUser).Error; err != nil {
		t.Fatalf("failed to create test user: %v", err)
	}

	testCategory := models.Category{Name: "Test Category"}
	if err := db.Create(&testCategory).Error; err != nil {
		t.Fatalf("failed to create test category: %v", err)
	}

	return db
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/posts", AddPost)

	return router
}

func TestAddPost(t *testing.T) {
	// Настраиваем тестовую БД
	db := setupTestDB(t)
	config.DB = db // Подменяем глобальную переменную config.DB

	router := setupRouter()

	tests := []struct {
		name         string
		postData     map[string]interface{}
		wantStatus   int
		wantResponse gin.H
	}{
		{
			name: "Успешное добавление поста",
			postData: map[string]interface{}{
				"title":       "Test Post",
				"content":     "Test Content",
				"user_id":     1,
				"category_id": 1,
			},
			wantStatus: http.StatusCreated,
			wantResponse: gin.H{
				"title":       "Test Post",
				"content":     "Test Content",
				"user_id":     float64(1), // JSON numbers decode as float64
				"category_id": float64(1),
			},
		},
		{
			name: "Отсутствует заголовок",
			postData: map[string]interface{}{
				"content":     "Test Content",
				"user_id":     1,
				"category_id": 1,
			},
			wantStatus:   http.StatusBadRequest,
			wantResponse: gin.H{"error": "Название поста обязательно"},
		},
		{
			name: "Неверный user_id",
			postData: map[string]interface{}{
				"title":       "Test Post",
				"content":     "Test Content",
				"user_id":     999, // Несуществующий пользователь
				"category_id": 1,
			},
			wantStatus:   http.StatusInternalServerError,
			wantResponse: gin.H{"error": "Не удалось добавить пост"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Подготовка запроса
			jsonData, _ := json.Marshal(tt.postData)
			req, _ := http.NewRequest("POST", "/posts", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")

			// Выполнение запроса
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			// Проверка статуса
			assert.Equal(t, tt.wantStatus, resp.Code)

			// Проверка ответа
			var response gin.H
			err := json.Unmarshal(resp.Body.Bytes(), &response)
			assert.NoError(t, err)

			if tt.wantStatus == http.StatusCreated {
				// Проверяем, что пост действительно создан в БД
				var post models.Post
				err := db.First(&post, response["id"]).Error
				assert.NoError(t, err)
				assert.Equal(t, tt.postData["title"], post.Title)
				assert.Equal(t, tt.postData["content"], post.Content)
			}

			// Проверяем ожидаемые поля в ответе
			for key, expected := range tt.wantResponse {
				assert.Equal(t, expected, response[key])
			}
		})
	}
}
