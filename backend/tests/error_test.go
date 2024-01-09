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

func TestInvalidUserIDParam(t *testing.T) {
	// Set up the router
	router := setupUserController()

	// Create a test request with an invalid user ID parameter
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/invalid_user_id/books", nil)

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code for an invalid user ID
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestInvalidBookPayload(t *testing.T) {
	// Set up the router
	router := setupBookController()

	// Create a test request with an invalid JSON payload for creating a book
	invalidPayload := []byte(`{"invalid_field": "value"}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(invalidPayload))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code for an invalid book creation payload
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// Continue with other error handling and edge case tests
// ...
