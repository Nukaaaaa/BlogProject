package handlers

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"myproject/mock"
	"myproject/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// ВАЖНО: Для тестов поменяем config.DB на мок
var mockDB *mock.MockDB

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	return router
}

func TestGetCategories(t *testing.T) {
	mockDB = new(mock.MockDB)
	categories := []models.Category{
		{ID: 1, Name: "Категория 1"},
		{ID: 2, Name: "Категория 2"},
	}
	mockDB.On("FindCategories").Return(categories, nil)

	router := setupRouter()
	router.GET("/categories", func(c *gin.Context) {
		cats, err := mockDB.FindCategories()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить категории"})
			return
		}
		c.JSON(http.StatusOK, cats)
	})

	req, _ := http.NewRequest(http.MethodGet, "/categories", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Категория 1")
	mockDB.AssertExpectations(t)
}

func TestAddCategory(t *testing.T) {
	mockDB = new(mock.MockDB)
	newCat := models.Category{Name: "Новая категория"}
	createdCat := newCat
	createdCat.ID = 1

	mockDB.On("CreateCategory", newCat).Return(&createdCat, nil)

	router := setupRouter()
	router.POST("/categories", func(c *gin.Context) {
		var cat models.Category
		if err := c.ShouldBindJSON(&cat); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		created, err := mockDB.CreateCategory(cat)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания"})
			return
		}
		c.JSON(http.StatusCreated, created)
	})

	body := `{"name":"Новая категория"}`
	req, _ := http.NewRequest(http.MethodPost, "/categories", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
	assert.Contains(t, resp.Body.String(), "Новая категория")
	mockDB.AssertExpectations(t)
}

func TestGetCategoryByID(t *testing.T) {
	mockDB = new(mock.MockDB)
	cat := models.Category{ID: 1, Name: "Категория 1"}

	mockDB.On("FindCategoryByID", uint(1)).Return(&cat, nil)

	router := setupRouter()
	router.GET("/categories/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
			return
		}
		category, err := mockDB.FindCategoryByID(uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Категория не найдена"})
			return
		}
		c.JSON(http.StatusOK, category)
	})

	req, _ := http.NewRequest(http.MethodGet, "/categories/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Категория 1")
	mockDB.AssertExpectations(t)
}

func TestUpdateCategory(t *testing.T) {
	mockDB = new(mock.MockDB)
	id := uint(1)
	updatedCat := models.Category{Name: "Обновленная категория"}

	mockDB.On("UpdateCategory", id, updatedCat).Return(&updatedCat, nil)

	router := setupRouter()
	router.PUT("/categories/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		idNum, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
			return
		}
		var cat models.Category
		if err := c.ShouldBindJSON(&cat); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
			return
		}
		updated, err := mockDB.UpdateCategory(uint(idNum), cat)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления"})
			return
		}
		c.JSON(http.StatusOK, updated)
	})

	body := `{"name":"Обновленная категория"}`
	req, _ := http.NewRequest(http.MethodPut, "/categories/1", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Обновленная категория")
	mockDB.AssertExpectations(t)
}

func TestDeleteCategory(t *testing.T) {
	mockDB = new(mock.MockDB)
	mockDB.On("DeleteCategory", uint(1)).Return(nil)

	router := setupRouter()
	router.DELETE("/categories/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
			return
		}
		err = mockDB.DeleteCategory(uint(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удаления"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Категория удалена"})
	})

	req, _ := http.NewRequest(http.MethodDelete, "/categories/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Категория удалена")
	mockDB.AssertExpectations(t)
}
