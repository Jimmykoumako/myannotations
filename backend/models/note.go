package models

import (
	"github.com/jinzhu/gorm"
)

// Note model
type Note struct {
	gorm.Model
	Content   string `gorm:"not null"`
	Timestamp int64
	TextID    uint // Foreign key referencing the Text entity
	UserID    uint // Foreign key referencing the User entity
}

// CreateNote creates a new note in the database
func CreateNote(note *Note) error {
	return DB.Create(note).Error
}

// GetNoteByID retrieves a note by its ID from the database
func GetNoteByID(noteID uint) (*Note, error) {
	var note Note
	err := DB.First(&note, noteID).Error
	return &note, err
}

// GetNotesByUserID retrieves all notes for a specific user from the database
func GetNotesByUserID(userID uint) ([]Note, error) {
	var notes []Note
	err := DB.Where("user_id = ?", userID).Find(&notes).Error
	return notes, err
}

// GetNotesByTextID retrieves all notes for a specific text from the database
func GetNotesByTextID(textID uint) ([]Note, error) {
	var notes []Note
	err := DB.Where("text_id = ?", textID).Find(&notes).Error
	return notes, err
}

// UpdateNote updates a note in the database
func UpdateNote(note *Note) error {
	return DB.Save(note).Error
}

// DeleteNote deletes a note from the database
func DeleteNote(note *Note) error {
	return DB.Delete(note).Error
}
