package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(fmt.Sprintf("Запрос от %v", c.Request.RemoteAddr))
		c.Next()
	}
}