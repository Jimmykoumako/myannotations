// migrations/20240124000000_create_connections_table.go
package migrations

import (
	"github.com/jinzhu/gorm"
	"mas/models"
)

// CreateConnectionTable migration
func CreateConnectionTable(db *gorm.DB) error {
	return db.AutoMigrate(&models.Connection{}).Error
}