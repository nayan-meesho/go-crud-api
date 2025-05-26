package middleware

import (
    "log"
    "time"

    "github.com/gin-gonic/gin"
)

func CustomLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		log.Printf("Custom log: %s %s (%v)", c.Request.Method, c.Request.URL.Path, duration)
	}
}