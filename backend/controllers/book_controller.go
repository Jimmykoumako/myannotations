// controllers/book_controller.go

package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"

    "mas/models"
)

// GetBooks handles the request to fetch all books.
func GetBooks(c *gin.Context) {
    db, err := config.GetDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
        return
    }
    defer db.Close()

    var books []models.Book
    if err := db.Find(&books).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
        return
    }

    c.JSON(http.StatusOK, books)
}
