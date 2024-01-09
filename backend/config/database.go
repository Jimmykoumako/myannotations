// config/config.go
package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

// InitDB initializes the database connection
func InitDB() {
	postgresURL := os.Getenv("POSTGRES_URL")

	// If running in a Docker container, use the provided environment variables
	if inDocker := os.Getenv("IN_DOCKER"); inDocker == "true" {
		postgresURL = "postgres://postgres:postgres@database:5432/madb"
	}

	conn, err := gorm.Open("postgres", postgresURL)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	DB = conn

	// AutoMigrate your models here if needed
	// DB.AutoMigrate(&models.User{}, &models.Book{}, ...)

	fmt.Println("Connected to the database")
}

// CloseDB closes the database connection
func CloseDB() {
	if DB != nil {
		DB.Close()
		fmt.Println("Closed the database connection")
	}
}
