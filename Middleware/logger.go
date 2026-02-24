package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now()
		path := c.Request.URL.Path

		c.Next()

		duration := time.Since(start)

		log.Printf(
			"%s | %d | %v | %s",
			c.Request.Method,
			c.Writer.Status(),
			duration,
			path,
		)
	}
}
