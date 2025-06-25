package main

import (
	"context"
	"fmt"
	"gin-app/db"
	"gin-app/handlers"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

func main() {
	r := gin.Default()
	db.Init()

	//	r.Use(middleware.Logger())

	r.GET("/todo", handlers.GORMGetAllTodoes)
	r.GET("/todo/:id", handlers.GORMGetTodo)
	r.POST("/todo", handlers.GORMCreateTodo)
	r.DELETE("/todo/:id", handlers.GORMDeleteTodo)

	// Prometheus
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// RabbitMQ
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
		"todo_queue",
		false,
		false,
		false,
		false,
		nil,
	)
	if q_err != nil {
		fmt.Println("error", q_err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "Hello from Rabbit епта"
	pub_err := ch.PublishWithContext(ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if pub_err != nil {
		fmt.Println("error", pub_err.Error())
	}

	log.Printf(" [x] Sent %s", body)

	fmt.Println("Сервер Gin запущен на http://localhost:8080")
	err := r.Run(":8080")
	if err != nil {
		fmt.Printf("ошибка запуска сервера: %s\n", err)
	}
}
