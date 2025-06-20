package handlers

import (
	"gin-app/todo"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllTodoes(c *gin.Context) {
	todoes, err := todo.LoadAllTodoes("db.json")

	if err != nil {
		c.JSON(500, gin.H{"error": "Ошибка чтения БД"})
		return
	}

	c.JSON(200, todoes)
}

func GetTodo(c *gin.Context) {
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

func CreateTodo(c *gin.Context) {
	var td todo.ToDo

	err := c.BindJSON(&td)
	if err != nil {
		c.JSON(500, gin.H{"error": "Ошибка парсинга тела"})
		return
	}

	todoes, err := todo.UpdateAllTodoes("db.json", td)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, todoes)
}
