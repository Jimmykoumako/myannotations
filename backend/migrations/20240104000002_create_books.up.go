// migrations/20240121000000_create_books_table.go
package migrations

import (
	"github.com/jinzhu/gorm"
)

// CreateBookTable migration
func CreateBookTable(db *gorm.DB) error {
	return db.AutoMigrate(&models.Book{}).Error
}