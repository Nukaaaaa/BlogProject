package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Генерация RequestID
		requestID := uuid.New().String()

		// Логируем запрос
		log.Printf("[%s] [%s] %s %s - Start", start.Format(time.RFC3339), requestID, c.Request.Method, c.Request.URL.Path)

		// Пропускаем запрос дальше
		c.Next()

		// Время выполнения запроса
		duration := time.Since(start).Milliseconds()

		// Логируем ответ
		log.Printf("[%s] [%s] %s %s - %d - Duration: %dms", start.Format(time.RFC3339), requestID, c.Request.Method, c.Request.URL.Path, c.Writer.Status(), duration)
	}
}

func RequestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Генерация уникального RequestId
		requestId := uuid.New().String()

		// Добавляем RequestId в заголовки ответа
		c.Header("X-Request-Id", requestId)

		// Логируем RequestId и информацию о запросе
		log.Printf("RequestId: %s | Method: %s | Path: %s | Time: %s", requestId, c.Request.Method, c.Request.URL.Path, time.Now().Format(time.RFC3339))

		// Добавляем RequestId в контекст для использования в других частях приложения
		c.Set("RequestId", requestId)

		// Пропускаем запрос дальше
		c.Next()
	}
}
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Генерация уникального RequestId для каждого запроса
		requestID := uuid.New().String()
		c.Set("RequestID", requestID)
		start := time.Now() // Время начала запроса

		// Логируем начало запроса
		log.Printf("[%s] %s %s - Start", requestID, c.Request.Method, c.Request.URL.Path)

		// Обрабатываем запрос
		c.Next()

		// Вычисляем продолжительность запроса
		duration := time.Since(start).Milliseconds()

		// Логируем завершение запроса с информацией о статусе, времени и продолжительности
		log.Printf("[%s] %s %s - %d - Duration: %dms", requestID, c.Request.Method, c.Request.URL.Path, c.Writer.Status(), duration)
	}
}
