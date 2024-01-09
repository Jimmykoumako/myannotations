// middleware/logging_middleware.go
package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// LoggingMiddleware is a middleware for logging
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Now()
		latency := endTime.Sub(startTime)

		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path
		statusCode := c.Writer.Status()

		fmt.Printf("[%s] %s %s %d %v\n", clientIP, method, path, statusCode, latency)
	}
}
