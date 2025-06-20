package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type ToDo struct {
	title       string
	description string
	isCompleted bool
}

func ask() ToDo {
	reader := bufio.NewReader(os.Stdin)

	// Title
	fmt.Println("Привет! Введи название задачи: ")
	title, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка чтения ввода", err)
	}
	title = strings.TrimSpace(title)

	// Description
	fmt.Println("Супер! А теперь введи описание задачи: ")
	description, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Ошибка чтения ввода", err)
	}
	description = strings.TrimSpace(description)

	// Result
	str := fmt.Sprintf("Твоя задача сформирована, заголовок: %v, описание: %v", title, description)
	fmt.Println(str)

	return ToDo{
		title:       title,
		description: description,
		isCompleted: false,
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	tds := []ToDo{
		{title: "123", description: "321", isCompleted: false},
	}
	fmt.Printf("получен GET-запрос на /\n")
	str := fmt.Sprintf("title - %v, description - %v, isCompleted - %v", tds[0].title, tds[0].description, tds[0].isCompleted)
	io.WriteString(w, str)
}

func main() {
	tds := []ToDo{}

	http.HandleFunc("GET /", getRoot)
	fmt.Println("Сервер запущен на http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("ошибка запуска сервера: %s\n", err)
	}

	fmt.Println("Задачи:", tds)
}
