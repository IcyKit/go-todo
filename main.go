package main

import (
	"fmt"
	"gin-app/db"
	"gin-app/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db.Init()

	r.GET("/todo", handlers.GORMGetAllTodoes)
	r.GET("/todo/:id", handlers.GORMGetTodo)
	r.POST("/todo", handlers.GORMCreateTodo)
	r.DELETE("/todo/:id", handlers.GORMDeleteTodo)

	fmt.Println("Сервер Gin запущен на http://localhost:8080")
	err := r.Run(":8080")
	if err != nil {
		fmt.Printf("ошибка запуска сервера: %s\n", err)
	}
}

// 1. Изучить что такое RPC и попробовать gRPC
// 2. Добавить тесты
// 3. Реализовать логирование как middleware
// 4. Поднять Кафку или RabbitMQ, предварительно разбив на микросервисы
// 5. Аутентефикация/Авторизация через JWT
// 6. Grafana + Prometeus
// 7. Обернуть все в Docker контейнер и потом поднять в Kubernetes
// 8. Добавить документацию через Swagger
