// middleware/auth.go
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"mas/config"
	"mas/models"
)

// AuthMiddleware is a middleware for user authentication
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "Missing or empty Authorization header"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("your_secret_key"), nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(401, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		userID, ok := claims["id"].(float64)
		if !ok {
			c.JSON(401, gin.H{"error": "Invalid user ID in token"})
			c.Abort()
			return
		}

		var user models.User
		if err := config.DB.First(&user, uint(userID)).Error; err != nil {
			c.JSON(401, gin.H{"error": "User not found or database error"})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

