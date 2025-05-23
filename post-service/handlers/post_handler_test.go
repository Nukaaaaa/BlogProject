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

func TestGetPostsByUserID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockPosts := []models.Post{
		{ID: 1, Title: "Тестовый пост", Content: "Содержимое", UserID: 42, CategoryID: 1},
	}

	mockDB := new(mock.MockDB)
	mockDB.On("FindPosts", uint(42)).Return(mockPosts, nil)

	router.GET("/users/:id/posts", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
			return
		}
		posts, err := mockDB.FindPosts(uint(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения постов"})
			return
		}
		c.JSON(http.StatusOK, posts)
	})

	req, _ := http.NewRequest(http.MethodGet, "/users/42/posts", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Тестовый пост")
	mockDB.AssertExpectations(t)
}

func TestAddPost(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockDB := new(mock.MockDB)
	post := models.Post{Title: "Новый пост", Content: "Текст", UserID: 1, CategoryID: 1}

	mockDB.On("CreatePost", post).Return(&post, nil)

	router.POST("/posts", func(c *gin.Context) {
		var p models.Post
		if err := c.ShouldBindJSON(&p); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Плохой JSON"})
			return
		}
		createdPost, err := mockDB.CreatePost(p)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания поста"})
			return
		}
		c.JSON(http.StatusCreated, createdPost)
	})

	body := `{"title":"Новый пост","content":"Текст","user_id":1,"category_id":1}`
	req, _ := http.NewRequest(http.MethodPost, "/posts", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
	assert.Contains(t, resp.Body.String(), "Новый пост")
	mockDB.AssertExpectations(t)
}

func TestGetPostByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockPost := models.Post{ID: 1, Title: "Пост по ID"}

	mockDB := new(mock.MockDB)
	mockDB.On("FindPostByID", uint(1)).Return(&mockPost, nil)

	router.GET("/posts/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
			return
		}
		post, err := mockDB.FindPostByID(uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Пост не найден"})
			return
		}
		c.JSON(http.StatusOK, post)
	})

	req, _ := http.NewRequest(http.MethodGet, "/posts/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Пост по ID")
	mockDB.AssertExpectations(t)
}

func TestUpdatePost(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockDB := new(mock.MockDB)

	// Полный объект, соответствующий телу запроса
	inputPost := models.Post{Title: "Обновленный пост"}
	returnPost := models.Post{ID: 1, Title: "Обновленный пост"}

	// Ожидаем ровно этот пост (важно, чтобы поля совпадали!)
	mockDB.On("UpdatePost", uint(1), inputPost).Return(&returnPost, nil)

	router.PUT("/posts/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
			return
		}
		var post models.Post
		if err := c.ShouldBindJSON(&post); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
			return
		}
		updatedPost, err := mockDB.UpdatePost(uint(id), post)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления"})
			return
		}
		c.JSON(http.StatusOK, updatedPost)
	})

	body := `{"title":"Обновленный пост"}`
	req, _ := http.NewRequest(http.MethodPut, "/posts/1", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Обновленный пост")
	mockDB.AssertExpectations(t)
}

func TestDeletePost(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockDB := new(mock.MockDB)
	mockDB.On("DeletePost", uint(1)).Return(nil)

	router.DELETE("/posts/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
			return
		}
		err = mockDB.DeletePost(uint(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удаления"})
			return
		}
		c.Status(http.StatusNoContent)
	})

	req, _ := http.NewRequest(http.MethodDelete, "/posts/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNoContent, resp.Code)
	mockDB.AssertExpectations(t)
}

func TestGetPostsByCategoryID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockPosts := []models.Post{
		{ID: 2, Title: "Пост категории", Content: "Контент", UserID: 2, CategoryID: 5},
	}

	mockDB := new(mock.MockDB)
	mockDB.On("FindPostsByCategory", uint(5)).Return(mockPosts, nil)

	router.GET("/categories/:id/posts", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
			return
		}
		posts, err := mockDB.FindPostsByCategory(uint(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения постов"})
			return
		}
		c.JSON(http.StatusOK, posts)
	})

	req, _ := http.NewRequest(http.MethodGet, "/categories/5/posts", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Пост категории")
	mockDB.AssertExpectations(t)
}
