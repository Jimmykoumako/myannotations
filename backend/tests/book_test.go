// tests/controllers_test.go
package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"mas/config"
	"mas/controllers"
	"mas/models"
)

func setupBookController() *gin.Engine {
	router := gin.Default()
	bookController := controllers.BookController{}
	router.POST("/books", bookController.CreateBook)
	router.GET("/books/:bookID", bookController.ViewBook)
	router.PUT("/books/:bookID", bookController.UpdateBook)
	router.DELETE("/books/:bookID", bookController.DeleteBook)
	router.POST("/books/:bookID/texts", bookController.CreateTextInBook) // Add route for creating text within a book
	return router
}

func TestCreateBook(t *testing.T) {
	// Set up the router
	router := setupBookController()

	// Create a test request with valid JSON payload for creating a book
	validPayload := []byte(`{"title": "Test Book", "author": "Test Author", "isbn": "1234567890"}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(validPayload))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code for successful book creation
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestCreateBookInvalidInput(t *testing.T) {
	// Set up the router
	router := setupBookController()

	// Create a test request with invalid JSON payload for creating a book
	invalidPayload := []byte(`{"invalid_field": "value"}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(invalidPayload))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code for invalid book creation
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestViewBookDetails(t *testing.T) {
	// Set up the router
	router := setupBookController()

	// Create a book for testing
	testBook := models.Book{
		Title:  "Test Book",
		Author: "Test Author",
		ISBN:   "1234567890",
	}
	config.DB.Create(&testBook)

	// Create a test request for viewing book details
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books/"+string(testBook.ID), nil)

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code for successful book details retrieval
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateBookDetails(t *testing.T) {
	// Set up the router
	router := setupBookController()

	// Create a book for testing
	testBook := models.Book{
		Title:  "Test Book",
		Author: "Test Author",
		ISBN:   "1234567890",
	}
	config.DB.Create(&testBook)

	// Create a test request with valid JSON payload for updating book details
	validPayload := []byte(`{"title": "Updated Book Title"}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/books/"+string(testBook.ID), bytes.NewBuffer(validPayload))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code for successful book details update
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateBookDetailsInvalidInput(t *testing.T) {
	// Set up the router
	router := setupBookController()

	// Create a book for testing
	testBook := models.Book{
		Title:  "Test Book",
		Author: "Test Author",
		ISBN:   "1234567890",
	}
	config.DB.Create(&testBook)

	// Create a test request with invalid JSON payload for updating book details
	invalidPayload := []byte(`{"invalid_field": "value"}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/books/"+string(testBook.ID), bytes.NewBuffer(invalidPayload))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code for invalid book details update
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestDeleteBook(t *testing.T) {
	// Set up the router
	router := setupBookController()

	// Create a book for testing
	testBook := models.Book{
		Title:  "Test Book",
		Author: "Test Author",
		ISBN:   "1234567890",
	}
	config.DB.Create(&testBook)

	// Create a test request for deleting a book
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/books/"+string(testBook.ID), nil)

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code for successful book deletion
	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestCreateTextInBook(t *testing.T) {
	// Set up the router
	router := setupBookController()

	// Create a book for testing
	testBook := models.Book{
		Title:  "Test Book",
		Author: "Test Author",
		ISBN:   "1234567890",
	}
	config.DB.Create(&testBook)

	// Create a test request with valid JSON payload for creating a text within a book
	validPayload := []byte(`{"content": "Sample text content"}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/books/"+string(testBook.ID)+"/texts", bytes.NewBuffer(validPayload))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code for successful text creation within a book
	assert.Equal(t, http.StatusCreated, w.Code)
}
