package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestUserServiceIntegration(t *testing.T) {
	client := resty.New()

	// Отправка запроса на user-service
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"id": "2",
		}).
		Get("http://localhost:8081/users/2")

	// Проверка, что запрос прошел успешно
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())

	// Проверка, что данные пришли правильные
	assert.Contains(t, resp.String(), "Nurt")
	assert.Contains(t, resp.String(), "nurt@narxoz.kz")
}
func TestGetPostsByUserID_Found(t *testing.T) {
	resp, err := resty.New().R().
		Get("http://localhost:8080/posts/user/1")

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode())
	assert.Contains(t, resp.String(), "user_id")
}

func TestGetPostByID_NotFound(t *testing.T) {
	resp, err := resty.New().R().
		Get("http://localhost:8080/posts/99999")

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode())
	assert.Contains(t, resp.String(), "Пост не найден")
}
func TestDeletePost_Success(t *testing.T) {
	// Сначала создаём пост
	post := map[string]interface{}{
		"title":       "Delete Test",
		"content":     "to delete",
		"user_id":     1,
		"category_id": 1,
	}
	createResp, _ := resty.New().R().
		SetHeader("Content-Type", "application/json").
		SetBody(post).
		Post("http://localhost:8080/posts")

	var created map[string]interface{}
	json.Unmarshal(createResp.Body(), &created)
	id := int(created["id"].(float64))

	// Теперь удаляем
	resp, err := resty.New().R().
		Delete(fmt.Sprintf("http://localhost:8080/posts/%d", id))

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode())
	assert.Contains(t, resp.String(), "Пост удален")
}
