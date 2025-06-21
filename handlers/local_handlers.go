package handlers

import (
	"gin-app/todo"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func LocalGetAllTodoes(c *gin.Context) {
	todoes, err := todo.LoadAllTodoes("db.json")

	if err != nil {
		c.JSON(500, gin.H{"error": "Ошибка чтения БД"})
		return
	}

	c.JSON(200, todoes)
}

func LocalGetTodo(c *gin.Context) {
	todoes, err := todo.LoadAllTodoes("db.json")
	if err != nil {
		c.JSON(500, gin.H{"error": "Ошибка чтения БД"})
		return
	}

	tdId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": "Ошибка преобразования числа"})
		return
	}

	for _, v := range todoes {
		if v.Id == tdId {
			c.JSON(200, v)
		}
	}
}

func LocalCreateTodo(c *gin.Context) {
	var td todo.ToDo

	err := c.BindJSON(&td)
	if err != nil {
		c.JSON(500, gin.H{"error": "Ошибка парсинга тела"})
		return
	}

	todoes, err := todo.AddOneTodoes("db.json", td)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, todoes)
}

func LocalDeleteTodo(c *gin.Context) {
	todoes, err := todo.LoadAllTodoes("db.json")
	if err != nil {
		c.JSON(500, gin.H{"error": "Ошибка чтения БД"})
		return
	}

	tdId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": "Ошибка преобразования числа"})
		return
	}

	for i, v := range todoes {
		if v.Id == tdId {
			arr := append(append(todoes[:i], todoes[i+1:]...))

			_, err := todo.UpdateAllTodoes("db.json", arr)
			if err != nil {
				c.JSON(500, gin.H{"error": err})
			}

			c.JSON(200, arr)
		}
	}
}
