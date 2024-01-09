// controllers/user_controller.go
package controllers

import (
	"github.com/gin-gonic/gin"
	"mas/models"
	"mas/config"
	// "golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"time"
	// "log"
)



// LoginUser handles user login
func (uc *UserController) LoginUser(c *gin.Context) {
	var credentials models.User
	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	var user models.User
	if err := config.DB.Where("username = ?", credentials.Username).First(&user).Error; err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	// err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	// if err != nil {
		// log.Printf(credentials.Password)
		// log.Printf(user.Password)
	// 	c.JSON(401, gin.H{"error": "Invalid credentials"})
	// 	return
	// }

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})

	tokenString, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	user.Token = tokenString

	c.JSON(200, gin.H{"token": tokenString})
}
