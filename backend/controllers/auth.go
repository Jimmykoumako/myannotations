package controllers

import (
   "mas/models"
   "golang.org/x/crypto/bcrypt"
   "github.com/gin-gonic/gin"
)

// RegisterUser registers a new user
func RegisterUser(c *gin.Context) {
   var user models.User
   if err := c.ShouldBindJSON(&user); err != nil {
      c.JSON(400, gin.H{"error": err.Error()})
      return
   }

   // Hash the password before storing it in the database
   hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
   if err != nil {
      c.JSON(500, gin.H{"error": "Failed to hash the password"})
      return
   }

   user.Password = string(hashedPassword)
   models.DB.Create(&user)
   c.JSON(200, user)
}

// LoginUser handles user login
func LoginUser(c *gin.Context) {
   var user models.User
   if err := c.ShouldBindJSON(&user); err != nil {
      c.JSON(400, gin.H{"error": err.Error()})
      return
   }

   // Retrieve the user from the database
   models.DB.Where("email = ?", user.Email).First(&user)

   // Compare the hashed password with the provided password
   if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(user.Password)); err != nil {
      c.JSON(401, gin.H{"error": "Invalid login credentials"})
      return
   }

   // TODO: Generate and return a JWT token for authenticated users
   // ...

   c.JSON(200, gin.H{"message": "Login successful"})
}
