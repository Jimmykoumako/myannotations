// middleware/error_handler.go
package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// ErrorHandlerMiddleware is a middleware for centralized error handling
func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic: %v", err)
				// Log the error or perform additional actions if needed

				// Respond to the client with a 500 Internal Server Error
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			}
		}()

		c.Next()

		// Check if there's an error in the response
		if len(c.Errors) > 0 {
			// Log the first error and respond with a 400 Bad Request
			err := c.Errors[0]
			log.Printf("Error: %v", err.Err)

			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		}
	}
}
