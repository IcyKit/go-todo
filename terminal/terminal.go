package terminal

import (
	"bufio"
	"fmt"
	"gin-app/todo"
	"os"
	"strings"
)

func AskTerminal() todo.ToDo {
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

	return todo.ToDo{
		Title:       title,
		Description: description,
		IsCompleted: false,
	}
}
