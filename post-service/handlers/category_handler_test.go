package handlers_test

import (
	"BackendProject/post-service/config"
	"BackendProject/post-service/handlers"
	"BackendProject/post-service/mock"
	"BackendProject/post-service/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAddCategory(t *testing.T) {
	// Инициализация моковой базы и замена основной
	config.DB = mock.InitMockDB()

	router := gin.Default()
	router.POST("/categories", handlers.AddCategory)

	category := models.Category{Name: "Test Category"}
	jsonValue, _ := json.Marshal(category)
	req, _ := http.NewRequest("POST", "/categories", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Ожидался статус 201 Created, получен: %d", w.Code)
	}
}
