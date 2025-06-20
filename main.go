package main

import (
	"fmt"
	"gin-app/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/todo", handlers.GetAllTodoes)
	r.GET("/todo/:id", handlers.GetTodo)
	r.POST("/todo", handlers.CreateTodo)

	fmt.Println("Сервер Gin запущен на http://localhost:8080")
	err := r.Run(":8080")
	if err != nil {
		fmt.Printf("ошибка запуска сервера: %s\n", err)
	}
}
