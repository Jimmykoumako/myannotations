package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"mas/models"
	"mas/config"
	"strconv"
)

// ConnectionController handles connection-related operations
type ConnectionController struct{}

// CreateConnection creates a new connection between two annotations
func (cc *ConnectionController) CreateConnection(c *gin.Context) {
	var newConnection models.Connection
	if err := c.BindJSON(&newConnection); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	config.DB.Create(&newConnection)

	c.JSON(201, gin.H{"data": newConnection})
}

// ViewConnection retrieves and displays connection details
func (cc *ConnectionController) ViewConnection(c *gin.Context) {
	connectionID := c.Param("connectionID")

	var connection models.Connection
	if err := config.DB.First(&connection, connectionID).Error; err != nil {
		c.JSON(404, gin.H{"error": "Connection not found"})
		return
	}

	c.JSON(200, gin.H{"data": connection})
}

// UpdateConnection updates connection details
func (cc *ConnectionController) UpdateConnection(c *gin.Context) {
	connectionID := c.Param("connectionID")

	var connection models.Connection
	if err := config.DB.First(&connection, connectionID).Error; err != nil {
		c.JSON(404, gin.H{"error": "Connection not found"})
		return
	}

	if err := c.BindJSON(&connection); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	config.DB.Save(&connection)

	c.JSON(200, gin.H{"data": connection})
}

// DeleteConnection deletes a connection from the system
func (cc *ConnectionController) DeleteConnection(c *gin.Context) {
	connectionID := c.Param("connectionID")

	config.DB.Delete(&models.Connection{}, connectionID)

	c.JSON(204, nil)
}

// ProvideConnectionFeedback allows users to provide feedback on why two annotations are connected
func (cc *ConnectionController) ProvideConnectionFeedback(c *gin.Context) {
	connectionID := c.Param("connectionID")

	var feedback models.ConnectionFeedback
	if err := c.BindJSON(&feedback); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	feedback.ConnectionID = getConnectionIDFromParam(c, connectionID)

	config.DB.Create(&feedback)

	c.JSON(201, gin.H{"data": feedback})
}

// Utility function to get connection ID from URL parameter
func getConnectionIDFromParam(c *gin.Context, param string) uint {
	connectionID, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid connection ID"})
		c.Abort()
		return 0
	}
	return uint(connectionID)
}

// GetConnectionFeedbackForConnection retrieves all feedback for a specific connection
func (cc *ConnectionController) GetConnectionFeedbackForConnection(c *gin.Context) {
	connectionID := c.Param("connectionID")

	// Query the database for feedback related to the connection
	var feedback []models.ConnectionFeedback
	config.DB.Where("connection_id = ?", connectionID).Find(&feedback)

	c.JSON(http.StatusOK, gin.H{"data": feedback})
}

// GetConnectionFeedbackByUser retrieves all feedback provided by a specific user
func (cc *ConnectionController) GetConnectionFeedbackByUser(c *gin.Context) {
	userID := c.Param("userID")

	// Query the database for feedback provided by the user
	var feedback []models.ConnectionFeedback
	config.DB.Where("user_id = ?", userID).Find(&feedback)

	c.JSON(http.StatusOK, gin.H{"data": feedback})
}

func (cc *ConnectionController) UpdateConnectionFeedback(c *gin.Context) {
	feedbackID := c.Param("feedbackID")

	var feedback models.ConnectionFeedback
	if err := config.DB.First(&feedback, feedbackID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Feedback not found"})
		return
	}

	// Update the feedback based on the request body
	if err := c.ShouldBindJSON(&feedback); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&feedback)

	c.JSON(http.StatusOK, gin.H{"data": feedback})
}

// DeleteConnectionFeedback deletes connection feedback by ID
func (cc *ConnectionController) DeleteConnectionFeedback(c *gin.Context) {
	feedbackID := c.Param("feedbackID")

	config.DB.Delete(&models.ConnectionFeedback{}, feedbackID)

	c.JSON(http.StatusNoContent, nil)
}
