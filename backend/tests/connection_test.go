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

func setupConnectionController() *gin.Engine {
	router := gin.Default()
	connectionController := controllers.ConnectionController{}
	router.POST("/connections", connectionController.CreateConnection)
	router.GET("/connections/:connectionID", connectionController.ViewConnection)
	router.PUT("/connections/:connectionID", connectionController.UpdateConnection)
	router.DELETE("/connections/:connectionID", connectionController.DeleteConnection)
	router.POST("/connections/:connectionID/feedback", connectionController.ProvideConnectionFeedback)
	return router
}

func TestCreateConnection(t *testing.T) {
	// Set up the router
	router := setupConnectionController()

	// Create two annotations for testing
	sourceAnnotation := models.Annotation{Content: "Source Annotation"}
	targetAnnotation := models.Annotation{Content: "Target Annotation"}
	config.DB.Create(&sourceAnnotation)
	config.DB.Create(&targetAnnotation)

	// Create a test request with valid JSON payload for creating a connection
	validPayload := []byte(`{
		"relationship_type": "related",
		"source_annotation_id": ` + string(sourceAnnotation.ID) + `,
		"target_annotation_id": ` + string(targetAnnotation.ID) + `
	}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/connections", bytes.NewBuffer(validPayload))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code for successful connection creation
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestCreateConnectionInvalidInput(t *testing.T) {
	// Set up the router
	router := setupConnectionController()

	// Create a test request with invalid JSON payload for creating a connection
	invalidPayload := []byte(`{"invalid_field": "value"}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/connections", bytes.NewBuffer(invalidPayload))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code for invalid connection creation
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestViewConnectionDetails(t *testing.T) {
	// Set up the router
	router := setupConnectionController()

	// Create a connection for testing
	testConnection := models.Connection{
		RelationshipType:    "related",
		SourceAnnotationID:  1, // Assuming this ID exists
		TargetAnnotationID:  2, // Assuming this ID exists
	}
	config.DB.Create(&testConnection)

	// Create a test request for viewing connection details
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/connections/"+string(testConnection.ID), nil)

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code for successful connection details retrieval
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateConnectionDetails(t *testing.T) {
	// Set up the router
	router := setupConnectionController()

	// Create a connection for testing
	testConnection := models.Connection{
		RelationshipType:    "related",
		SourceAnnotationID:  1, // Assuming this ID exists
		TargetAnnotationID:  2, // Assuming this ID exists
	}
	config.DB.Create(&testConnection)

	// Create a test request with valid JSON payload for updating connection details
	validPayload := []byte(`{"relationship_type": "updated_related"}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/connections/"+string(testConnection.ID), bytes.NewBuffer(validPayload))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code for successful connection details update
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateConnectionDetailsInvalidInput(t *testing.T) {
	// Set up the router
	router := setupConnectionController()

	// Create a connection for testing
	testConnection := models.Connection{
		RelationshipType:    "related",
		SourceAnnotationID:  1, // Assuming this ID exists
		TargetAnnotationID:  2, // Assuming this ID exists
	}
	config.DB.Create(&testConnection)

	// Create a test request with invalid JSON payload for updating connection details
	invalidPayload := []byte(`{"invalid_field": "value"}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/connections/"+string(testConnection.ID), bytes.NewBuffer(invalidPayload))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code for invalid connection details update
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestDeleteConnection(t *testing.T) {
	// Set up the router
	router := setupConnectionController()

	// Create a connection for testing
	testConnection := models.Connection{
		RelationshipType:    "related",
		SourceAnnotationID:  1, // Assuming this ID exists
		TargetAnnotationID:  2, // Assuming this ID exists
	}
	config.DB.Create(&testConnection)

	// Create a test request for deleting a connection
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/connections/"+string(testConnection.ID), nil)

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code for successful connection deletion
	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestProvideConnectionFeedback(t *testing.T) {
	// Set up the router
	router := setupConnectionController()

	// Create a connection for testing
	testConnection := models.Connection{
		RelationshipType:    "related",
		SourceAnnotationID:  1, // Assuming this ID exists
		TargetAnnotationID:  2, // Assuming this ID exists
	}
	config.DB.Create(&testConnection)

	// Create a test request with valid JSON payload for providing feedback on a connection
	validPayload := []byte(`{"feedback": "This is a helpful connection."}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/connections/"+string(testConnection.ID)+"/feedback", bytes.NewBuffer(validPayload))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code for successful feedback provision
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestProvideConnectionFeedbackInvalidInput(t *testing.T) {
	// Set up the router
	router := setupConnectionController()

	// Create a connection for testing
	testConnection := models.Connection{
		RelationshipType:    "related",
		SourceAnnotationID:  1, // Assuming this ID exists
		TargetAnnotationID:  2, // Assuming this ID exists
	}
	config.DB.Create(&testConnection)

	// Create a test request with invalid JSON payload for providing feedback on a connection
	invalidPayload := []byte(`{"invalid_field": "value"}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/connections/"+string(testConnection.ID)+"/feedback", bytes.NewBuffer(invalidPayload))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code for invalid feedback provision
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
