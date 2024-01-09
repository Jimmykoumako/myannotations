// migrations/20240123000000_create_annotations_table.go
package migrations

import (
	"github.com/jinzhu/gorm"
	"mas/models"
)

// CreateAnnotationTable migration
func CreateAnnotationTable(db *gorm.DB) error {
	return db.AutoMigrate(&models.Annotation{}).Error
}