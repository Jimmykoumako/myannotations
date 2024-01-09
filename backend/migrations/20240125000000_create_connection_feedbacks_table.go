// migrations/20240125000000_create_connection_feedbacks_table.go
package migrations

import (
	"github.com/jinzhu/gorm"
	"mas/models"
)

// CreateConnectionFeedbackTable migration
func CreateConnectionFeedbackTable(db *gorm.DB) error {
	return db.AutoMigrate(&models.ConnectionFeedback{}).Error
}