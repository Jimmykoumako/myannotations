package models

import "github.com/jinzhu/gorm"

// Text model
type Text struct {
   gorm.Model
   Content string `gorm:"not null"`
   Page    int    `gorm:"not null"`
   BookID  uint
   Book    Book // Belongs to a book
   Annotations []Annotation // Text has many annotations
}
