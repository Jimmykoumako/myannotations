package controllers

import (
   "mas/models"
   "github.com/jinzhu/gorm"
)

var DB *gorm.DB

// SetDB sets the database instance for controllers
func SetDB(db *gorm.DB) {
   DB = db
}
