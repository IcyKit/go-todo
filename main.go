package main

import (
	"context"
	"fmt"
	"gin-app/handlers"
	"gin-app/middlewares"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	r := gin.Default()
	// db.Init()

	r.Use(middlewares.Logger())

	r.GET("/todo", handlers.GORMGetAllTodoes)
	r.GET("/todo/:id", handlers.GORMGetTodo)
	r.POST("/todo", handlers.GORMCreateTodo)
	r.DELETE("/todo/:id", handlers.GORMDeleteTodo)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	//RabbitMQ
	conn, cerr := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if cerr != nil {
		fmt.Println("error", cerr.Error())
	}
	defer conn.Close()

	ch, cherr := conn.Channel()
	if cherr != nil {
		fmt.Println("error", cherr.Error())
	}
	defer ch.Close()
	q, q_err := ch.QueueDeclare(
		"todo_queue", // имя очереди
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if q_err != nil {
		fmt.Println("error", q_err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "Hello from ToDo!"
	pub_err := ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
	})
	if pub_err != nil {
		fmt.Println("error", pub_err.Error())
	}

	log.Printf(" [x] Sent %s\n", body)
	//

	fmt.Println("Сервер Gin запущен на http://localhost:8080")
	err := r.Run(":8080")
	if err != nil {
		fmt.Printf("ошибка запуска сервера: %s\n", err)
	}
}

// 1. Изучить что такое RPC и попробовать gRPC
// 2. Добавить тесты ✅
// 3. Реализовать логирование как middleware ✅
// 4. Поднять Кафку или RabbitMQ, предварительно разбив на микросервисы
// 5. Аутентефикация/Авторизация через JWT
// 6. Grafana + Prometeus ✅
// 7. Обернуть все в Docker контейнер и потом поднять в Kubernetes
// 8. Добавить документацию через Swagger
// 9. Добавить elasticsearch
