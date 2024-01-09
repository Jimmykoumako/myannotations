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

func setupUserController() *gin.Engine {
	router := gin.Default()
	userController := controllers.UserController{}
	router.POST("/register", userController.RegisterUser)
	router.POST("/login", userController.LoginUser) // Add login route
	return router
}

func TestRegisterUser(t *testing.T) {
	// Set up the router
	router := setupUserController()

	// Create a test request with valid JSON payload for user registration
	validPayload := []byte(`{"email": "test@example.com", "password": "password"}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(validPayload))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code for successful registration
	assert.Equal(t, http.StatusCreated, w.Code)

	// Create a test request with invalid JSON payload for user registration
	invalidPayload := []byte(`{"invalid_field": "value"}`)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/register", bytes.NewBuffer(invalidPayload))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code for invalid registration
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestLoginUser(t *testing.T) {
	// Set up the router
	router := setupUserController()

	// Create a user for testing
	testUser := models.User{
		Email:    "test@example.com",
		Password: "password",
	}
	config.DB.Create(&testUser)

	// Create a test request with valid JSON payload for user login
	validPayload := []byte(`{"email": "test@example.com", "password": "password"}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(validPayload))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code for successful login
	assert.Equal(t, http.StatusOK, w.Code)

	// Create a test request with invalid JSON payload for user login
	invalidPayload := []byte(`{"invalid_field": "value"}`)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(invalidPayload))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code for invalid login
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestViewUserDetails(t *testing.T) {
	// Set up the router
	router := setupUserController()

	// Create a user for testing
	testUser := models.User{
		Email:    "test@example.com",
		Password: "password",
	}
	config.DB.Create(&testUser)

	// Create a test request for viewing user details
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/"+string(testUser.ID), nil)

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code for successful user details retrieval
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateUserDetails(t *testing.T) {
	// Set up the router
	router := setupUserController()

	// Create a user for testing
	testUser := models.User{
		Email:    "test@example.com",
		Password: "password",
	}
	config.DB.Create(&testUser)

	// Create a test request with valid JSON payload for updating user details
	validPayload := []byte(`{"email": "updated@example.com"}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/user/"+string(testUser.ID), bytes.NewBuffer(validPayload))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code for successful user details update
	assert.Equal(t, http.StatusOK, w.Code)

	// Create a test request with invalid JSON payload for updating user details
	invalidPayload := []byte(`{"invalid_field": "value"}`)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("PUT", "/user/"+string(testUser.ID), bytes.NewBuffer(invalidPayload))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code for invalid user details update
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestDeleteUser(t *testing.T) {
	// Set up the router
	router := setupUserController()

	// Create a user for testing
	testUser := models.User{
		Email:    "test@example.com",
		Password: "password",
	}
	config.DB.Create(&testUser)

	// Create a test request for deleting a user
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/user/"+string(testUser.ID), nil)

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code for successful user deletion
	assert.Equal(t, http.StatusNoContent, w.Code)
}
