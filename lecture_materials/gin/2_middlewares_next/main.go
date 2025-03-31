package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func MW1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("MW1 Start")
		c.Next()
		fmt.Println("MW1 Finish")
	}
}

func MW2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("MW2 Start")
		c.Next()
		fmt.Println("MW2 Finish")
	}
}

func MW3() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("MW3 Start")
		c.Next()
		fmt.Println("MW3 Finish")
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Применить middleware
	r.Use(MW1())
	r.Use(MW2())
	r.Use(MW3())

	// Определить обработчик для корневого пути
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello world")
	})

	// Запустить сервер
	r.Run(":8080")
}
