// tests/controllers_test.go
package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"mas/config"
	"mas/controllers"
	"mas/models"
)

func setupUserController() *gin.Engine {
	router := gin.Default()
	userController := controllers.UserController{}
	router.POST("/register", userController.RegisterUser)
	return router
}

func TestRegisterUser(t *testing.T) {
	// Set up the router
	router := setupUserController()

	// Create a test request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/register", nil)

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func setupBookController() *gin.Engine {
	router := gin.Default()
	bookController := controllers.BookController{}
	router.POST("/books", bookController.CreateBook)
	// Add other routes as needed
	return router
}

func TestCreateBook(t *testing.T) {
	// Set up the router
	router := setupBookController()

	// Create a test request with valid JSON payload
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/books", nil)

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code
	assert.Equal(t, http.StatusCreated, w.Code)
}