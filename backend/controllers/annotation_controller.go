package controllers

import (
	"github.com/gin-gonic/gin"
	"mas/models"
	"mas/config"
	"strconv"
)

// AnnotationController handles annotation-related operations
type AnnotationController struct{}

// CreateAnnotation creates a new annotation (note, highlight, underline)
func (ac *AnnotationController) CreateAnnotation(c *gin.Context) {
	var newAnnotation models.Annotation
	if err := c.BindJSON(&newAnnotation); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	config.DB.Create(&newAnnotation)

	c.JSON(201, gin.H{"data": newAnnotation})
}

// ViewAnnotation retrieves and displays annotation details
func (ac *AnnotationController) ViewAnnotation(c *gin.Context) {
	annotationID := c.Param("annotationID")

	var annotation models.Annotation
	if err := config.DB.First(&annotation, annotationID).Error; err != nil {
		c.JSON(404, gin.H{"error": "Annotation not found"})
		return
	}

	c.JSON(200, gin.H{"data": annotation})
}

// UpdateAnnotation updates annotation details
func (ac *AnnotationController) UpdateAnnotation(c *gin.Context) {
	annotationID := c.Param("annotationID")

	var annotation models.Annotation
	if err := config.DB.First(&annotation, annotationID).Error; err != nil {
		c.JSON(404, gin.H{"error": "Annotation not found"})
		return
	}

	if err := c.BindJSON(&annotation); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	config.DB.Save(&annotation)

	c.JSON(200, gin.H{"data": annotation})
}

// DeleteAnnotation deletes an annotation from the system
func (ac *AnnotationController) DeleteAnnotation(c *gin.Context) {
	annotationID := c.Param("annotationID")

	config.DB.Delete(&models.Annotation{}, annotationID)

	c.JSON(204, nil)
}

// ConnectAnnotations establishes a connection between two annotations
func (ac *AnnotationController) ConnectAnnotations(c *gin.Context) {
	annotationID1 := c.Param("annotationID1")
	annotationID2 := c.Param("annotationID2")

	var connection models.Connection
	if err := c.BindJSON(&connection); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	connection.SourceAnnotationID = getAnnotationIDFromParam(c, annotationID1)
	connection.TargetAnnotationID = getAnnotationIDFromParam(c, annotationID2)

	config.DB.Create(&connection)

	c.JSON(201, gin.H{"data": connection})
}

// ViewConnections retrieves and displays connections between annotations
func (ac *AnnotationController) ViewConnections(c *gin.Context) {
	annotationID := c.Param("annotationID")

	var connections []models.Connection
	if err := config.DB.Where("annotation_id1 = ? OR annotation_id2 = ?", annotationID, annotationID).Find(&connections).Error; err != nil {
		c.JSON(404, gin.H{"error": "Connections not found"})
		return
	}

	c.JSON(200, gin.H{"data": connections})
}

// Utility function to get annotation ID from URL parameter
func getAnnotationIDFromParam(c *gin.Context, param string) uint {
	annotationID, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid annotation ID"})
		c.Abort()
		return 0
	}
	return uint(annotationID)
}