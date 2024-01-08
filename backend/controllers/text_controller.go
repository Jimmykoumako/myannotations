package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"your_project/models"
	"your_project/config"
	"strconv"
)

// TextController handles text-related operations
type TextController struct{}

// CreateText adds a new text (chapter, paragraph) to the system
func (tc *TextController) CreateText(c *gin.Context) {
	var newText models.Text
	if err := c.BindJSON(&newText); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	config.DB.Create(&newText)

	c.JSON(201, gin.H{"data": newText})
}

// ViewText retrieves and displays text details
func (tc *TextController) ViewText(c *gin.Context) {
	textID := c.Param("textID")

	var text models.Text
	if err := config.DB.Preload("Annotations").First(&text, textID).Error; err != nil {
		c.JSON(404, gin.H{"error": "Text not found"})
		return
	}

	c.JSON(200, gin.H{"data": text})
}

// UpdateText updates text details
func (tc *TextController) UpdateText(c *gin.Context) {
	textID := c.Param("textID")

	var text models.Text
	if err := config.DB.First(&text, textID).Error; err != nil {
		c.JSON(404, gin.H{"error": "Text not found"})
		return
	}

	if err := c.BindJSON(&text); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	config.DB.Save(&text)

	c.JSON(200, gin.H{"data": text})
}

// DeleteText deletes a text from the system
func (tc *TextController) DeleteText(c *gin.Context) {
	textID := c.Param("textID")

	config.DB.Delete(&models.Text{}, textID)

	c.JSON(204, nil)
}

// CreateAnnotationInText creates a new annotation (note, highlight, underline) within a
func (tc *TextController) CreateAnnotationInText(c *gin.Context) {
	textID := c.Param("textID")
	var newAnnotation models.Annotation

	if err := c.BindJSON(&newAnnotation); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	newAnnotation.TextID = getTextIDFromParam(c, textID)

	config.DB.Create(&newAnnotation)

	c.JSON(201, gin.H{"data": newAnnotation})
}

// Utility function to get text ID from URL parameter
func getTextIDFromParam(c *gin.Context, param string) uint {
	textID, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid text ID"})
		c.Abort()
		return 0
	}
	return uint(textID)
}