// models/connection_feedback.go
package models

import (
	"github.com/jinzhu/gorm"
)

// ConnectionFeedback model
type ConnectionFeedback struct {
	gorm.Model
	ConnectionID uint   `gorm:"not null"`
	UserID       uint   // You may associate feedback with a user
	Feedback     string `gorm:"not null"`
}
