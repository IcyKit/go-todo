package handlers

import (
	"gin-app/db"
	"gin-app/todo"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GORMGetAllTodoes(c *gin.Context) {
	var td []todo.ToDo

	// result := db.DB.Find(&td)

	// if result.Error != nil {
	// 	c.JSON(500, gin.H{"error": result.Error.Error()})
	// 	return
	// }

	c.JSON(200, td)
}

func GORMGetTodo(c *gin.Context) {
	var td []todo.ToDo

	tdId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": "Ошибка преобразования числа"})
		return
	}

	result := db.DB.First(&td, tdId)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, td[0])
}

func GORMCreateTodo(c *gin.Context) {
	var td todo.ToDo

	err := c.BindJSON(&td)
	if err != nil {
		c.JSON(500, gin.H{"error": "Ошибка парсинга тела"})
		return
	}

	result := db.DB.Create(&td)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, td)
}

func GORMDeleteTodo(c *gin.Context) {
	var td todo.ToDo

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": "Ошибка преобразования числа"})
		return
	}

	findResult := db.DB.First(&td, id)
	if findResult.Error != nil {
		c.JSON(404, gin.H{"error": findResult.Error.Error()})
		return
	}

	delResult := db.DB.Delete(&td, id)
	if delResult.Error != nil {
		c.JSON(500, gin.H{"error": delResult.Error.Error()})
		return
	}

	c.JSON(200, td)
}
