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

func setupAnnotationController() *gin.Engine {
	router := gin.Default()
	annotationController := controllers.AnnotationController{}
	router.POST("/annotations", annotationController.CreateAnnotation)
	router.GET("/annotations/:annotationID", annotationController.ViewAnnotation)
	router.PUT("/annotations/:annotationID", annotationController.UpdateAnnotation)
	router.DELETE("/annotations/:annotationID", annotationController.DeleteAnnotation)
	router.POST("/annotations/:annotationID1/connect/:annotationID2", annotationController.ConnectAnnotations)
	router.GET("/annotations/:annotationID/connections", annotationController.ViewConnections)
	return router
}

func TestCreateAnnotation(t *testing.T) {
	// Set up the router
	router := setupAnnotationController()

	// Create a test request with valid JSON payload for creating an annotation
	validPayload := []byte(`{"content": "This is a test annotation."}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/annotations", bytes.NewBuffer(validPayload))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code for successful annotation creation
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestCreateAnnotationInvalidInput(t *testing.T) {
	// Set up the router
	router := setupAnnotationController()

	// Create a test request with invalid JSON payload for creating an annotation
	invalidPayload := []byte(`{"invalid_field": "value"}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/annotations", bytes.NewBuffer(invalidPayload))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code for invalid annotation creation
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestViewAnnotationDetails(t *testing.T) {
	// Set up the router
	router := setupAnnotationController()

	// Create an annotation for testing
	testAnnotation := models.Annotation{Content: "Test Annotation"}
	config.DB.Create(&testAnnotation)

	// Create a test request for viewing annotation details
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/annotations/"+string(testAnnotation.ID), nil)

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code for successful annotation details retrieval
	assert.Equal(t, http.StatusOK, w.Code)
}

// Continue with other tests for updating, deleting, connecting annotations, and viewing connections
// ...


func TestUpdateAnnotationDetails(t *testing.T) {
	// Set up the router
	router := setupAnnotationController()

	// Create an annotation for testing
	testAnnotation := models.Annotation{Content: "Test Annotation"}
	config.DB.Create(&testAnnotation)

	// Create a test request with valid JSON payload for updating annotation details
	validPayload := []byte(`{"content": "Updated Annotation"}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/annotations/"+string(testAnnotation.ID), bytes.NewBuffer(validPayload))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code for successful annotation details update
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateAnnotationDetailsInvalidInput(t *testing.T) {
	// Set up the router
	router := setupAnnotationController()

	// Create an annotation for testing
	testAnnotation := models.Annotation{Content: "Test Annotation"}
	config.DB.Create(&testAnnotation)

	// Create a test request with invalid JSON payload for updating annotation details
	invalidPayload := []byte(`{"invalid_field": "value"}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/annotations/"+string(testAnnotation.ID), bytes.NewBuffer(invalidPayload))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code for invalid annotation details update
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestDeleteAnnotation(t *testing.T) {
	// Set up the router
	router := setupAnnotationController()

	// Create an annotation for testing
	testAnnotation := models.Annotation{Content: "Test Annotation"}
	config.DB.Create(&testAnnotation)

	// Create a test request for deleting an annotation
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/annotations/"+string(testAnnotation.ID), nil)

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code for successful annotation deletion
	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestConnectAnnotations(t *testing.T) {
	// Set up the router
	router := setupAnnotationController()

	// Create two annotations for testing
	sourceAnnotation := models.Annotation{Content: "Source Annotation"}
	targetAnnotation := models.Annotation{Content: "Target Annotation"}
	config.DB.Create(&sourceAnnotation)
	config.DB.Create(&targetAnnotation)

	// Create a test request for connecting two annotations
	validPayload := []byte(`{"relationship_type": "related"}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/annotations/"+string(sourceAnnotation.ID)+"/connect/"+string(targetAnnotation.ID), bytes.NewBuffer(validPayload))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code for successful annotation connection
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestViewAnnotationConnections(t *testing.T) {
	// Set up the router
	router := setupAnnotationController()

	// Create two connected annotations for testing
	sourceAnnotation := models.Annotation{Content: "Source Annotation"}
	targetAnnotation := models.Annotation{Content: "Target Annotation"}
	config.DB.Create(&sourceAnnotation)
	config.DB.Create(&targetAnnotation)

	// Connect the two annotations
	connection := models.Connection{
		RelationshipType:    "related",
		SourceAnnotationID:  sourceAnnotation.ID,
		TargetAnnotationID:  targetAnnotation.ID,
	}
	config.DB.Create(&connection)

	// Create a test request for viewing annotation connections
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/annotations/"+string(sourceAnnotation.ID)+"/connections", nil)

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response code for successful annotation connections retrieval
	assert.Equal(t, http.StatusOK, w.Code)
}

// Continue with other tests as needed
// ...
