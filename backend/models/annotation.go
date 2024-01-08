package models

import "github.com/jinzhu/gorm"

// Annotation model
type Annotation struct {
   gorm.Model
   Content     string `gorm:"not null"`
   Type        string `gorm:"not null"` // 'Note', 'Highlight', 'Underline', etc.
   Timestamp   int64
   Color       string // Only if applicable (e.g., for highlights)
   TextID      uint
   Text        Text // Belongs to a text
   User        User // Belongs to a user
   Connections []Connection // Annotation has many connections
}
