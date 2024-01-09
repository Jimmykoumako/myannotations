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

func setupTextController() *gin.Engine {
	router := gin.Default()
	textController := controllers.TextController{}
	router.POST("/texts", textController.CreateText)
	router.GET("/texts/:textID", textController.ViewText)
	router.PUT("/texts/:textID", textController.UpdateText)
	router.DELETE("/texts/:textID", textController.DeleteText)
	return router
}

func TestCreateText(t *testing.T) {
	// Set up the router
	router := setupTextController()

	// Create a test request with valid JSON payload for creating a text
	validPayload := []byte(`{"content": "Sample text content"}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/texts", bytes.NewBuffer(validPayload))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code for successful text creation
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestCreateTextInvalidInput(t *testing.T) {
	// Set up the router
	router := setupTextController()

	// Create a test request with invalid JSON payload for creating a text
	invalidPayload := []byte(`{"invalid_field": "value"}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/texts", bytes.NewBuffer(invalidPayload))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code for invalid text creation
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestViewTextDetails(t *testing.T) {
	// Set up the router
	router := setupTextController()

	// Create a text for testing
	testText := models.Text{
		Content: "Sample text content",
	}
	config.DB.Create(&testText)

	// Create a test request for viewing text details
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/texts/"+string(testText.ID), nil)

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code for successful text details retrieval
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateTextDetails(t *testing.T) {
	// Set up the router
	router := setupTextController()

	// Create a text for testing
	testText := models.Text{
		Content: "Sample text content",
	}
	config.DB.Create(&testText)

	// Create a test request with valid JSON payload for updating text details
	validPayload := []byte(`{"content": "Updated text content"}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/texts/"+string(testText.ID), bytes.NewBuffer(validPayload))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code for successful text details update
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateTextDetailsInvalidInput(t *testing.T) {
	// Set up the router
	router := setupTextController()

	// Create a text for testing
	testText := models.Text{
		Content: "Sample text content",
	}
	config.DB.Create(&testText)

	// Create a test request with invalid JSON payload for updating text details
	invalidPayload := []byte(`{"invalid_field": "value"}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/texts/"+string(testText.ID), bytes.NewBuffer(invalidPayload))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code for invalid text details update
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestDeleteText(t *testing.T) {
	// Set up the router
	router := setupTextController()

	// Create a text for testing
	testText := models.Text{
		Content: "Sample text content",
	}
	config.DB.Create(&testText)

	// Create a test request for deleting a text
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/texts/"+string(testText.ID), nil)

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code for successful text deletion
	assert.Equal(t, http.StatusNoContent, w.Code)
}
