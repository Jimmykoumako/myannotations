package models

import "github.com/jinzhu/gorm"

// Connection model
type Connection struct {
   gorm.Model
   RelationshipType string `gorm:"not null"`
   Timestamp        int64
   UserID           uint
   SourceAnnotationID uint // Connecting two annotations
   TargetAnnotationID uint
   SourceAnnotation Annotation `gorm:"foreignkey:SourceAnnotationID"` // Belongs to a source annotation
   TargetAnnotation Annotation `gorm:"foreignkey:TargetAnnotationID"` // Belongs to a target annotation
}
