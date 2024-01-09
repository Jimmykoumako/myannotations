package controllers

import (
	"github.com/gin-gonic/gin"
	"mas/models"
	"mas/config"
	"strconv"
)

// UserController handles user-related operations
type UserController struct{}

// RegisterUser creates a new user
func (uc *UserController) RegisterUser(c *gin.Context) {
	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	config.DB.Create(&newUser)

	c.JSON(201, gin.H{"data": newUser})
}

// AddBookToUser adds a new book to the user's collection
func (uc *UserController) AddBookToUser(c *gin.Context) {
	userID := c.Param("userID")
	var newBook models.Book

	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	newBook.UserID = getUserIDFromParam(c, userID)

	config.DB.Create(&newBook)

	c.JSON(201, gin.H{"data": newBook})
}

// ViewUserBooks retrieves and displays the user's book collection
func (uc *UserController) ViewUserBooks(c *gin.Context) {
	userID := c.Param("userID")
	books := []models.Book{}

	config.DB.Where("user_id = ?", userID).Find(&books)

	c.JSON(200, gin.H{"data": books})
}

// Utility function to get user ID from URL parameter
func getUserIDFromParam(c *gin.Context, param string) uint {
	userID, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		c.Abort()
		return 0
	}
	return uint(userID)
}

// CreateUser creates a new user
func (uc *UserController) CreateUser(c *gin.Context) {
	// Parse request and create a new user
	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// Add the new user to the database
	config.DB.Create(&newUser)

	c.JSON(201, gin.H{"data": newUser})
}

// GetUserByID retrieves a user by ID
func (uc *UserController) GetUserByID(c *gin.Context) {
	// Get user ID from the URL parameter
	userID := c.Param("id")

	// Query the database for the user
	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, gin.H{"data": user})
}

// UpdateUser updates a user by ID
func (uc *UserController) UpdateUser(c *gin.Context) {
	// Get user ID from the URL parameter
	userID := c.Param("id")

	// Query the database for the user
	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	// Parse request and update the user
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// Save the updated user to the database
	config.DB.Save(&user)

	c.JSON(200, gin.H{"data": user})
}

// DeleteUser deletes a user by ID
func (uc *UserController) DeleteUser(c *gin.Context) {
	// Get user ID from the URL parameter
	userID := c.Param("id")

	// Delete the user from the database
	config.DB.Delete(&models.User{}, userID)

	c.JSON(204, nil)
}
