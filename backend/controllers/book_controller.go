package controllers

import (
	"github.com/gin-gonic/gin"
	"mas/models"
	"mas/config"
	"net/http"
	"strconv"
	"fmt"
)

// BookController handles book-related operations
type BookController struct{}

// CreateBook creates a new book
func (bc *BookController) CreateBook(c *gin.Context) {
	var newBook models.Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Additional input validation if needed
	if len(newBook.Title) == 0 || len(newBook.Author) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title and author are required"})
		return
	}

	// Save the new book to the database
	config.DB.Create(&newBook)

	c.JSON(http.StatusCreated, gin.H{"data": newBook})
}

// ViewBook retrieves a book by ID
func (bc *BookController) ViewBook(c *gin.Context) {
	bookID := c.Param("bookID")

	var book models.Book
	if err := config.DB.First(&book, bookID).Error; err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		// Log other types of errors
		fmt.Printf("Error retrieving book: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// UpdateBook updates book details
func (bc *BookController) UpdateBook(c *gin.Context) {
	bookID := c.Param("bookID")

	var book models.Book
	if err := config.DB.First(&book, bookID).Error; err != nil {
		c.JSON(404, gin.H{"error": "Book not found"})
		return
	}

	if err := c.BindJSON(&book); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	config.DB.Save(&book)

	c.JSON(200, gin.H{"data": book})
}

// DeleteBook deletes a book from the system
func (bc *BookController) DeleteBook(c *gin.Context) {
	bookID := c.Param("bookID")

	config.DB.Delete(&models.Book{}, bookID)

	c.JSON(204, nil)
}

// CreateTextInBook creates a new text (chapter, paragraph) in a book
func (bc *BookController) CreateTextInBook(c *gin.Context) {
	bookID := c.Param("bookID")
	var newText models.Text

	if err := c.BindJSON(&newText); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	newText.BookID = getBookIDFromParam(c, bookID)

	config.DB.Create(&newText)

	c.JSON(201, gin.H{"data": newText})
}

// Utility function to get book ID from URL parameter
func getBookIDFromParam(c *gin.Context, param string) uint {
	bookID, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid book ID"})
		c.Abort()
		return 0
	}
	return uint(bookID)
}
