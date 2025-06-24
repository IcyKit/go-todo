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
	n := 10

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

func TestMain(m *testing.M) {
	db.Init()
	code := m.Run()
	os.Exit(code)
}
