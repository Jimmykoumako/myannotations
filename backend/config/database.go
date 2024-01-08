// config/database.go

package config

import (
    "fmt"
    "os"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

// GetDB initializes and returns a connection to the PostgreSQL database.
func GetDB() (*gorm.DB, error) {
    url := os.Getenv("POSTGRES_URL")
    if url == "" {
        return nil, fmt.Errorf("POSTGRES_URL environment variable not set")
    }

    db, err := gorm.Open("postgres", url)
    if err != nil {
        return nil, err
    }

    return db, nil
}
