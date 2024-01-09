// migrations/20240122000000_create_texts_table.go
package migrations

import (
	"github.com/jinzhu/gorm"
)

// CreateTextTable migration
func CreateTextTable(db *gorm.DB) error {
	return db.AutoMigrate(&models.Text{}).Error
}