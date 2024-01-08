package models

import "github.com/jinzhu/gorm"

// Book model
type Book struct {
   gorm.Model
   Title  string `gorm:"not null"`
   Author string `gorm:"not null"`
   ISBN   string `gorm:"unique;not null"`
   UserID uint
   User   User // Belongs to a user
   Texts  []Text // Book has many texts
}
