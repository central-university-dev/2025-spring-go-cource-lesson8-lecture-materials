package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func MiddlewareOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Устанавливаем значение в контексте
		c.Set("key", "value from middleware one")
		// Продолжаем выполнение следующего middleware или обработчика
		c.Next()
	}
}

func MiddlewareTwo() gin.HandlerFunc {
	return func(c *gin.Context) {
		if value, exists := c.Get("key"); exists {
			fmt.Println("Middleware Two: Retrieved value:", value)
		}
		c.Next()
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Применяем middleware
	r.Use(MiddlewareOne())
	r.Use(MiddlewareTwo())

	// Определяем обработчик для корневого пути
	r.GET("/", func(c *gin.Context) {
		// Получаем значение из контекста
		if value, exists := c.Get("key"); exists {
			c.String(200, "Value: %s", value)
		} else {
			c.String(200, "No value found")
		}
	})

	// Запускаем сервер
	r.Run(":8080")
}
