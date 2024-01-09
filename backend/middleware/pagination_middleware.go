// middleware/pagination_middleware.go
package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// PaginationMiddleware is a middleware for pagination
func PaginationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
		if err != nil || page < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
			c.Abort()
			return
		}

		pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
		if err != nil || pageSize < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pageSize parameter"})
			c.Abort()
			return
		}

		c.Set("page", page)
		c.Set("pageSize", pageSize)

		c.Next()
	}
}
