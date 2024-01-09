package models

import "github.com/jinzhu/gorm"

// User model
type User struct {
   gorm.Model
   Username string `gorm:"unique;not null"`
   Email    string `gorm:"unique;not null"`
   Password string `gorm:"not null"`
   Token    string `json:"-"`
   Books    []Book   // User has many books
   Annotations []Annotation // User has many annotations
}



// SetPassword sets the hashed password for the user
func (u *User) SetPassword(password string) error {
   hashedPassword, err := HashPassword(password)
   if err != nil {
      return err
   }
   u.Password = hashedPassword
   return nil
}

// CheckPassword checks if the provided password matches the hashed password
func (u *User) CheckPassword(password string) bool {
   return CheckPasswordHash(password, u.Password)
}
