package main

import (
	"encoding/json"
	"fmt"
	"gin-app/db"
	"gin-app/handlers"
	"gin-app/todo"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.GET("/todo", handlers.GORMGetAllTodoes)
	r.GET("/todo/:id", handlers.GORMGetTodo)
	r.POST("/todo", handlers.GORMCreateTodo)
	r.DELETE("/todo/:id", handlers.GORMDeleteTodo)

	return r
}

func TestGORMGetAllTodoes(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("GET", "/todo", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "{")
}

func TestGORMGetTodo(t *testing.T) {
	r := setupRouter()
	n := 1

	req, _ := http.NewRequest("GET", fmt.Sprintf("/todo/%v", n), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	body := w.Body.String()
	var td todo.ToDo
	err := json.Unmarshal([]byte(body), &td)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, td.Id, n)
}

func TestGORMCreateTodo(t *testing.T) {
	r := setupRouter()
	body := `{"id":2,"title":"Тест таска добавлена!","description":"описание тест таски","is_сompleted":false}`

	req, _ := http.NewRequest("POST", "/todo", strings.NewReader(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var td todo.ToDo
	err := json.Unmarshal([]byte(w.Body.String()), &td)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, td.Title, "Тест таска добавлена!")

	d_req, _ := http.NewRequest("DELETE", "/todo/2", nil)
	r.ServeHTTP(w, d_req)
}

func TestGORMDeleteTodo(t *testing.T) {
	r := setupRouter()

	td := todo.ToDo{Id: 10, Title: "Для удаления", Description: "эта таска будет удалена", IsCompleted: false}
	db.DB.Create(&td)

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/todo/%v", td.Id), nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var check todo.ToDo
	res := db.DB.First(&check, td.Id)
	assert.Error(t, res.Error)
}

func TestMain(m *testing.M) {
	db.Init()
	code := m.Run()
	os.Exit(code)
}
