// models/book.go

package models

import "github.com/jinzhu/gorm"

// Book represents the Book model in the database.
type Book struct {
    gorm.Model
    Title  string `json:"title"`
    Author string `json:"author"`
}
