// migrations/20240120000000_create_users_table.go
package migrations

import (
	"github.com/jinzhu/gorm"
)

// CreateUserTable migration
func CreateUserTable(db *gorm.DB) error {
	return db.AutoMigrate(&models.User{}).Error
}